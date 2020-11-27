## 有哪些场景需要进行选举
1. 集群在启动的过程中,需要选举Leader
2. 集群正常启动后,leader因故障挂掉了,需要选举Leader
3. 集群中的Follower数量不足以通过半数检验,Leader会挂掉自己,选举新leader
4. 集群正常运行,新增加1个Follower

选举过程参考：[深入理解 ZK集群的Leader选举](https://www.cnblogs.com/ZhuChangwu/p/11622763.html)

## zookeeper脑裂会出现多个leader吗？

有很小的概率会出现多个leader的情况，但是会进入崩溃恢复或者服务不可用，不会出现数据不一致的情况。

具体场景1：比如集群里有A、B两个机房，A机房有s1、s2、s3三台机器，B机房有s4一台机器，集群启动时s4被选举为leader，s1、s2、s3为follower，如果两个机房网络断掉了，A机房的三台机器连接不上leader会进行选举，由于3台机器满足过半机制是可以选举出一个leader的，此时就会出现两个leader。

其实这种情况s4发现连接不上足够数量的follower（三台），会主动shut down自己然后进入选举状态，如果网络一直不通，则该节点不可用，整个集群还是可用的，如果网络通了，s4最终会认领到新的leader。
