## 工作原理
Filebeat 是使用 Golang 实现的轻量型日志采集器，也是 Elasticsearch stack 里面的一员。本质上是一个 agent ，可以安装在各个节点上，根据配置读取对应位置的日志，并上报到相应的地方去。其工作流程如下图所示：

![](https://github.com/NaraLuwan/dochub/tree/master/images/2021110301.png)

当启动 Filebeat 程序时，它会启动一个或多个查找器去检测指定的日志目录或文件。对于查找器 prospector 所在的每个日志文件，FIlebeat 会启动收集进程 harvester。 每个 harvester 都会为新内容读取单个日志文件，并将新日志数据发送到后台处理程序，后台处理程序会集合这些事件，最后发送集合的数据到 output 指定的目的地。

除了图中提到的各个组件，整个 filebeat 主要包含以下重要组件：

Crawler：负责管理和启动各个 Input
Input：负责管理和解析输入源的信息，以及为每个文件启动 Harvester。可由配置文件指定输入源信息。
Harvester: Harvester 负责读取一个文件的信息。
Pipeline: 负责管理缓存、Harvester 的信息写入以及 Output 的消费等，是 Filebeat 最核心的组件。
Output: 输出源，可由配置文件指定输出源信息。
Registrar：管理记录每个文件处理状态，包括偏移量、文件名等信息。当 Filebeat 启动时，会从 Registrar 恢复文件处理状态。
filebeat 的整个生命周期，几个组件共同协作，完成了日志从采集到上报的整个过程。