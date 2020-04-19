## 1. Builder设计模式

LoadingCache是通过newBuilder的方式建立的，这里用到的是Builder设计模式，一个bean中有个静态内部类Builder，其中的属性就是外部类对应的属性，它对外暴露属性，并把返回Buildler自身，直到使用bulid()，new出外部类，如下demo所示：

```java
/**
 * 构建模式demo
 *
 * @author Luwan
 * @date 2020/4/19
 */
public class PersonBuilderDemo {

    private String name;
    private int age;
    private boolean sex;

    public String getName() {
        return name;
    }

    public int getAge() {
        return age;
    }

    public boolean isSex() {
        return sex;
    }

    public static class Builder {
        private String name;
        private int age;
        private boolean sex;

        public Builder name(String n) {
            name = n;
            return this;
        }

        public Builder age(int a) {
            age = a;
            return this;
        }

        public Builder sex(boolean s) {
            sex = s;
            return this;
        }

        public PersonBuilderDemo build() {
            return new PersonBuilderDemo(this);
        }
    }

    private PersonBuilderDemo(Builder builder) {
        name = builder.name;
        age = builder.age;
        sex = builder.sex;
    }
}
```

来看下guava cache的Builder代码的实现，首先是创建一个CacheBuilder:

```java
/**
   * 静态公共方法暴露出来进行构建
   */
  public static CacheBuilder<Object, Object> newBuilder() {
    return new CacheBuilder<Object, Object>();
  }
```

剩下的就是利用各种属性进行赋值，这里使用其中一个作为例子:

```java
/**
   * 1、对参数进行校验
   * 2、赋值之后返回CacheBuilder对象本身，和上边的Person本身是很想的
   */
  public CacheBuilder<K, V> maximumSize(long size) {
    checkState(this.maximumSize == UNSET_INT, "maximum size was already set to %s",
        this.maximumSize);
    checkState(this.maximumWeight == UNSET_INT, "maximum weight was already set to %s",
        this.maximumWeight);
    checkState(this.weigher == null, "maximum size can not be combined with weigher");
    checkArgument(size >= 0, "maximum size must not be negative");
    this.maximumSize = size;
    return this;
  }
```

第三步就是调用build()方法进行构建LocalCache类：

```java
/**
   * 1、LocalCache调用其静态内部类LocalLoadingCache的构造方法new了一个LocalCache，这一点和上边的demo很像
   */
  public <K1 extends K, V1 extends V> LoadingCache<K1, V1> build(
      CacheLoader<? super K1, V1> loader) {
    checkWeightWithWeigher();
    return new LocalCache.LocalLoadingCache<K1, V1>(this, loader);
  }
 
  static class LocalLoadingCache<K, V>
      extends LocalManualCache<K, V> implements LoadingCache<K, V> {
 
    LocalLoadingCache(CacheBuilder<? super K, ? super V> builder,
        CacheLoader<? super K, V> loader) {
      super(new LocalCache<K, V>(builder, checkNotNull(loader)));
    }
  }
 
  /**
   * new LocalCache的过程就是把builder中的属性赋值到LocalCache中属性的过程
   */
  LocalCache(
      CacheBuilder<? super K, ? super V> builder, @Nullable CacheLoader<? super K, V> loader) {
    concurrencyLevel = Math.min(builder.getConcurrencyLevel(), MAX_SEGMENTS);
 
    keyStrength = builder.getKeyStrength();
    valueStrength = builder.getValueStrength();
 
    keyEquivalence = builder.getKeyEquivalence();
    valueEquivalence = builder.getValueEquivalence();
 
    maxWeight = builder.getMaximumWeight();
    weigher = builder.getWeigher();
    expireAfterAccessNanos = builder.getExpireAfterAccessNanos();
    expireAfterWriteNanos = builder.getExpireAfterWriteNanos();
    refreshNanos = builder.getRefreshNanos();
 
    removalListener = builder.getRemovalListener();
    removalNotificationQueue = (removalListener == NullListener.INSTANCE)
        ? LocalCache.<RemovalNotification<K, V>>discardingQueue()
        : new ConcurrentLinkedQueue<RemovalNotification<K, V>>();
 
    ticker = builder.getTicker(recordsTime());
    entryFactory = EntryFactory.getFactory(keyStrength, usesAccessEntries(), usesWriteEntries());
    globalStatsCounter = builder.getStatsCounterSupplier().get();
    defaultLoader = loader;
 
    int initialCapacity = Math.min(builder.getInitialCapacity(), MAXIMUM_CAPACITY);
    if (evictsBySize() && !customWeigher()) {
      initialCapacity = Math.min(initialCapacity, (int) maxWeight);
    }
```

CacheLoader是一个抽象类，需要返回一个实现类，一般直接new一个重写load方法即可：

```java
/**
 * CacheLoader是一个抽象类，返回的是CacheLoader的一个实现，至于其中的方法什么时候调用，完全看localCache的需要
 */
 new CacheLoader<String, String>() { 
   @Override 
    public String load(String key) { 
     return service.query(key); 
   } 
 }
```

至此guava localcache对象便生成了。