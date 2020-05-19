## 前言

出于安全方面考虑MySql-Server 只允许本机(localhost, 127.0.0.1)来连接访问. 这对于 Web-Server 与 MySql-Server 都在同一台服务器上的网站架构来说是没有问题的. 但随着网站流量的增加, 后期服务器架构可能会将 Web-Server 与 MySql-Server 分别放在独立的服务器上, 以便得到更大性能的提升, 此时 MySql-Server 就要修改成允许 Web-Server 进行远程连接。

### 登录 Mysql-Server 连接本地 mysql (默认只允许本地连接)
```bash
mysql -u root -p
```
```text
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 163
Server version: 5.7.29-0ubuntu0.16.04.1 (Ubuntu)

Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
```

### 修改 Mysql-Server 用户配置

```sql
select User,Host from mysql.user;
```
```text
+------+----------+-----------+
  | User | Password | Host      |
  +------+----------+-----------+
  | root |          | localhost |
  +------+----------+-----------+
  1 row in set (0.00 sec)
```

默认有一个 root 用户, 密码为空, 只允许 localhost 连接。

可以添加一个新的用户允许所有ip访问：

```sql
GRANT ALL PRIVILEGES ON  *.*  TO test@"%" IDENTIFIED BY 'test1234';flush privileges;
```
其中，test为用户名，test1234为密码，%表示允许所有ip，也可以指定为具体的ip。

### 修改 Mysql 配置文件 my.ini

```bash
sudo vim /etc/mysql/mysql.conf.d/mysqld.cnf
```

将 bind-address = 127.0.0.1 这一行注释掉即可。

### 重启 MySQL

```bash
sudo /etc/init.d/mysql restart
```

```text
[ ok ] Restarting mysql (via systemctl): mysql.service.
```