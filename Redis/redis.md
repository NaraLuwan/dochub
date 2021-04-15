3. Redis
[ ] redis 数据类型
- String set zset hash list hyperloglog bitmap geo
[ ] 底层数据结构
- redis中所有场景中出现的字符串，基本都是由动态字符串来实现的，有三个属性
  - free：还剩多少空间
  - len：字符串长度
  - buf：存放的字符数组
优点：空间预分配+惰性空间释放 -> 减少修改字符串内存重分配次数
- int：redis中存放的各种数字
- 双向链表：
- ziplist：
- 哈希表：
- 跳表：
[ ] 过期策略

[ ] 高可用方案：支持持久化-RDB/AOF 
RDB(redis database backup file)数据快照：压缩二进制
  - 配置方式：save 60 1  # 最近1分钟内 至少产生10000次写入
  - 优点：RDB文件压缩写入，体积较小；加载RDB文件速度很快，短时间启动
  - 缺点：数据不全；生成RDB文件消耗大量CPU和内存资源
  - 适用场景：数据备份、主从全量复制，对丢数据不敏感的业务
AOF(append only file)追加日志：原始操作命令，参数等
  - 配置方式：appendonly yes #开启  appendfsync everysec  # 文件刷盘方式：每秒
  - 优点：更新及时，数据更完整，降低丢失数据风险
  - 缺点：随着时间增长AOF文件会越来越大；增加磁盘负担，比如开启每秒刷盘
      - 解决：AOF文件很大时，触发AOF重写，Redis会扫描整个实例的数据，重新生成一个AOF文件。
  - 适用场景：对丢失数据敏感的场景
通常我们会选择AOF+每秒刷盘这种方式，既能保证良好的写入性能，在实例宕机时最多丢失1秒的数据，做到性能和安全的平衡。
[ ] .redis 是如何 rehash 的? 