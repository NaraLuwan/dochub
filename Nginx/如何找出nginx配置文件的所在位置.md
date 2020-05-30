前言：对于一台陌生的服务器或安装太久忘了位置，怎么才能简单快速的找到配置文件的位置呢？

## 第一步：先找出nginx可执行文件路径

### 情况一：如果程序在运行中
```bash
ps -ef | grep nginx
```
```text
root      6155     1  0 19:53 ?        00:00:00 nginx: master process /usr/sbin/nginx -g daemon on; master_process on;
root      6301  6155  0 20:04 ?        00:00:00 nginx: worker process
root      6302  6155  0 20:04 ?        00:00:00 nginx: worker process
root      6303  6155  0 20:04 ?        00:00:00 nginx: worker process
root      6304  6155  0 20:04 ?        00:00:00 nginx: worker process
appops    6408  2448  0 20:24 pts/8    00:00:00 grep nginx
```
通常是 /usr/sbin/nginx

### 情况二：程序并没有运行

#### 查看软件安装路径
```bash
whereis nginx
```

#### 查询运行文件所在路径
```bash
which nginx
```

#### 当然还有另外的查询方法
rpm包安装的
```bash
rpm -qa | grep "nginx"
```
yum方法安装的
```bash
yum list installed | grep "nginx"
```

## 第二步

通过-t参数查看生效配置

```text
-t : test configuration and exit
```

如：
```bash
/usr/sbin/nginx -t
```
