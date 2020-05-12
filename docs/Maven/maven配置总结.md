## maven 仓库分类

Maven根据坐标寻找构件的时候，它先会查看本地仓库，如果本地仓库存在构件，则直接使用；如果没有，则从远程仓库查找，找到后，下载到本地。

### 本地仓库
默认情况下，每个用户在自己的用户目录下都有一个路径名为.m2/repository/的仓库目录。我们也可以在 settings.xml 文件配置本地仓库的地址
```xml
<localRepository>D:\MavenHouse\repository</localRepository>
```
### 远程仓库
本地仓库好比书房，而远程仓库就像是书店。对于Maven来说，每个用户只有一个本地仓库，但是可以配置多个远程仓库。

### 中央仓库
Maven必须要知道至少一个可用的远程仓库，中央仓库就是这样一个默认的远程仓库，Maven 默认有一个 super pom 文件。
如：D:\apache-maven-3.0.4\lib 下的 maven-model-builder-3.0.4.jar 中的 org/apache/maven/model/pom-4.0.0.xml
```xml
  <repositories>
    <repository>
      <id>central</id>
      <name>Central Repository</name>
      <url>https://repo.maven.apache.org/maven2</url>
      <layout>default</layout>
      <snapshots>
        <enabled>false</enabled>
      </snapshots>
    </repository>
  </repositories>
```

maven是通过id来标识远程仓库的，我们可以通过配置镜像的mirrorOf为central来覆盖默认的中央仓库。

## 镜像
镜像（Mirroring）是冗余的一种类型，一个磁盘上的数据在另一个磁盘上存在一个完全相同的副本即为镜像。

为什么配置镜像?
```text
1.一句话，你有的我也有，你没有的我也有。（拥有远程仓库的所有 jar，包括远程仓库没有的 jar）
2.还是一句话，我跑的比你快。（有时候远程仓库获取 jar 的速度可能比镜像慢，这也是为什么我们一般要配置中央仓库的原因，外国的 maven 仓库一般获取速度比较慢）
```

**注意:当远程仓库被镜像匹配到的，则在获取 jar 包将从镜像仓库获取，而不是我们配置的 repository 仓库, repository 将失去作用**

### mirrorOf 标签

mirrorOf 标签里面放置的是 repository 配置的 id,为了满足一些复杂的需求，Maven还支持更高级的镜像配置：
```text
external:* = 不在本地仓库的文件才从该镜像获取
repo,repo1 = 远程仓库 repo 和 repo1 从该镜像获取
*,!repo1 =  所有远程仓库都从该镜像获取，除 repo1 远程仓库以外
* = 所用远程仓库都从该镜像获取
```

### 仓库之间的优先级

```text
本地仓库 > 私服 （profile）> 远程仓库（repository）和 镜像 （mirror） > 中央仓库 （central）
```

镜像等同于远程仓库，中央仓库可以理解为默认的远程仓库，所以真正的优先级为：

```text
本地仓库 > 私服（profile）> 远程仓库（repository）
```


