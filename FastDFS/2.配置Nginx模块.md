## 环境&版本
- ubuntu 16.04 LTS
- FastDFS V5.05
- Nginx 1.12.1

## 背景
使用 FastDFS 如果需要支持 http 方式访问文件，需要通过 Nginx 来实现。另外，FastDFS 通过 Tracker 服务器，将文件放在 Storage 服务器存储， 但是同组存储服务器之间需要进行文件复制， 有同步延迟的问题。
                                          
假设 Tracker 服务器将文件上传到了 192.168.51.128，上传成功后文件 ID已经返回给客户端，此时 FastDFS 存储集群机制会将这个文件同步到同组存储 192.168.51.129，在文件还没有复制完成的情况下，客户端如果用这个文件 ID 在 192.168.51.129 上取文件,就会出现文件无法访问的错误。
                                          
针对这种问题FastDFS作者提供了 fastdfs-nginx-module 可以重定向文件链接到源服务器取文件，避免客户端由于复制延迟导致的文件无法访问错误，因此需要先安装 Nginx 再集成 fastdfs-nginx-module 模块。

## 安装&配置 Nginx
### 前提
如果本机已经安装了Nginx需要先卸载或者找到对应版本的fastdfs-nginx-module，这里假设已安装过了且选择卸载重新安装。
1. 先备份已安装的模块和配置（如果有的话）
```bash
/usr/local/nginx/sbin/nginx -V
```
备份 config arguments参数里的内容

2. 卸载已安装安装Nginx
```bash
# 卸载删除除了配置文件以外的所有文件
apt-get remove nginx nginx-common

# 卸载所有东东，包括删除配置文件
apt-get purge nginx nginx-common
 
# 在上面命令结束后执行，主要是卸载删除Nginx的不再被使用的依赖包
apt-get autoremove

# 卸载删除两个主要的包
apt-get remove nginx-full nginx-common
```
### 依赖环境安装
1.  gcc 安装
```bash
apt-get install build-essential
apt-get install g++
```
2. PCRE pcre-devel 安装
```bash
apt-get install libpcre3 libpcre3-dev
```
3. zlib 安装
```bash
apt-get install zlib1g-dev
```
4. OpenSSL 安装
```bash
apt-get install openssl libssl-dev
```

### 安装 Nginx
1. 下载nginx
```bash
wget -c https://nginx.org/download/nginx-1.12.1.tar.gz
```
2. 解压
```bash
tar -zxvf nginx-1.12.1.tar.gz
cd nginx-1.12.1
```
3. 使用默认配置
```bash
./configure
```
4. 编译、安装
```bash
make
make install
```
5. 启动nginx
```bash
cd /usr/local/nginx/sbin/
./nginx
``` 
```text
其它命令
# ./nginx -s stop
# ./nginx -s quit
# ./nginx -s reload
```

6. 设置开机启动
```bash
vim /etc/rc.local
```
添加一行：
```text
/usr/local/nginx/sbin/nginx
```

7. 设置执行权限
```bash
# chmod 755 rc.local
```

8. 查看nginx的版本及模块
```bash
/usr/local/nginx/sbin/nginx -V
```
![github](https://github.com/NaraLuwan/spring-boot-learning/blob/master/img/2020011404.png)

9. 防火墙中打开Nginx端口（默认的 80） 
```bash
vim /etc/sysconfig/iptables
```

添加以下内容：
```text
-A INPUT -m state --state NEW -m tcp -p tcp --dport 80 -j ACCEPT
```

重启防火墙：
```bash
service iptables restart
```

### 测试访问文件
1. 修改nginx.conf
```bash
vim /usr/local/nginx/conf/nginx.conf
```

添加以下内容，将 /group1/M00 映射到 /ljzsg/fastdfs/file/data
```text
location /group1/M00 {
    alias /ljzsg/fastdfs/file/data;
}
```
2. 重启nginx
```bash
/usr/local/nginx/sbin/nginx -s reload
```
![github](https://github.com/NaraLuwan/spring-boot-learning/blob/master/img/2020011405.png)

3. 在浏览器访问之前上传的图片
```text
http://file.server.com/group1/M00/00/00/CvKxKF4dN0iARCm2AAAADzqXUB8903.txt
```

## 配置 Nginx 模块
### 安装 fastdfs-nginx-module
1. 下载 fastdfs-nginx-module
```bash
wget https://github.com/happyfish100/fastdfs-nginx-module/archive/5e5f3566bbfa57418b5506aaefbe107a42c9fcb1.zip
```

2. 解压
```bash
unzip 5e5f3566bbfa57418b5506aaefbe107a42c9fcb1.zip
```
重命名
```bash
mv fastdfs-nginx-module-5e5f3566bbfa57418b5506aaefbe107a42c9fcb1  fastdfs-nginx-module-master
```

3. 在nginx中添加模块
先停掉nginx服务
```bash
/usr/local/nginx/sbin/nginx -s stop
```

进入解压包目录
```bash
cd /home/luwan/software/nginx-1.12.1/
```

添加模块
```bash
./configure --add-module=../fastdfs-nginx-module-master/src
```

重新编译、安装
```bash
./make 
./make install
```

4. 查看Nginx的模块
```bash
/usr/local/nginx/sbin/nginx -V
```
有下面这个就说明添加模块成功
![github](https://github.com/NaraLuwan/spring-boot-learning/blob/master/img/2020011406.png)

5. 复制 fastdfs-nginx-module 源码中的配置文件到/etc/fdfs 目录并修改
```bash
cd /softpackages/fastdfs-nginx-module-master/src
cp mod_fastdfs.conf /etc/fdfs/
```

修改以下内容：
```text
# 连接超时时间
connect_timeout=10

# Tracker Server
tracker_server=file.server.com:22122

# StorageServer 默认端口
storage_server_port=23000

# 如果文件ID的uri中包含/group**，则要设置为true
url_have_group_name = true

# Storage 配置的store_path0路径，必须和storage.conf中的一致
store_path0=/home/luwan/fastdfs/file
```

6. 复制 FastDFS 的部分配置文件到/etc/fdfs 目录
```bash
cd /softpackages/fastdfs-5.05/conf/
cp anti-steal.jpg http.conf mime.types /etc/fdfs/
```

7. 配置nginx，修改nginx.conf
```bash
vim /usr/local/nginx/conf/nginx.conf
```
在80端口下添加fastdfs-nginx模块
```text
location ~/group([0-9])/M00 {
    ngx_fastdfs_module;
}
```
注意：

　　listen 80 端口值是要与 /etc/fdfs/storage.conf 中的 http.server_port=80 （前面改成80了）相对应。如果改成其它端口，则需要统一，同时在防火墙中打开该端口。

　　location 的配置，如果有多个group则配置location ~/group([0-9])/M00 ，没有则不用配group。

8. 在/ljzsg/fastdfs/file 文件存储目录下创建软连接，将其链接到实际存放数据的目录（这一步可以省略）

```bash
ln -s /ljzsg/fastdfs/file/data/ /ljzsg/fastdfs/file/data/M00
```
9. 启动nginx
```bash
/usr/local/nginx/sbin/nginx
```
10. 在地址栏访问
```text
http://file.server.com/group1/M00/00/00/CvKxKF4dN0iARCm2AAAADzqXUB8903.txt
```
能下载文件就算安装成功。注意和第三点中直接使用nginx路由访问不同的是，这里配置 fastdfs-nginx-module 模块，可以重定向文件链接到源服务器取文件。

最终部署图：
![github](https://github.com/NaraLuwan/spring-boot-learning/blob/master/img/2020011407.png)