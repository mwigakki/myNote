## mysql

#### 初次安装:

​	以管理员身份打开 cmd 命令行工具，切换到mysql的bin目录
​	初始化数据库：mysqld --initialize --console
​	执行完成后，会输出 root 用户的初始默认密码，如： CVCHW3(Pjke!
​	输入以下安装命令：
​		mysqld install
​	启动输入以下命令即可：
​		net start mysql
​	

用户：root
密码：123456
改密码：登录进mysql后：alter user 'root'@'localhost' identified by '123456';

#### 启动服务

net start mysql	（如果拒绝访问就用管理员权限打开cmd，或者已经启动了）

#### 登录

  ---> mysql -u root -p 然后输密码

#### 进入mysql后：
显示所有数据库： show databases;
​使用某个数据库： use [name];
​	创建数据库：CREATE DATABASE 数据库名;
​	删除数据库： drop database <数据库名>;
​	创建数据表：CREATE TABLE table_name (column_name column_type);

``` mysql
例：
mysql> create table student(
    -> id int unsigned primary key,
    -> name varchar(20) not null,
    -> gender enum('男', '女'),
    -> birthday date,
    -> class varchar(30));
```

#### 进入某个数据库后：
​	显示此数据库中的所有表：show tables;
​	显示表的表头：SHOW COLUMNS FROM 数据表;
​	显示数据表的详细索引信息，包括PRIMARY KEY（主键）:SHOW INDEX FROM 数据表:
​	删除某张表： drop table [name];
​	添加数据：INSERT INTO table_name ( field1, field2,...fieldN )
​                       VALUES
​                       ( value1, value2,...valueN );

例：

``` mysql
insert into student values
    -> (2, 'tom', '男', '2000-8-9', '一年级2班'),
    -> (3, 'lily', '女', '2000/11/19', '一年级2班');
```

#### 退出mysql : exit





1、MySQL 中有哪些常见日志？
重做日志（redo log）：物理日志
作用是确保事务的持久性。 redo 日志记录事务执行后的状态，用来恢复未写入 data file 的已提交事务数据。

回滚日志（undo log）：逻辑日志
作用是保证数据的原子性。 保存了事务发生之前的数据的一个版本，可以用于回滚，同时可以提供多版本并发控制下的读（MVCC），也即非锁定读。

二进制日志（binlog）：逻辑日志
常用于主从同步或数据同步中，也可用于数据库基于时间点的还原。

错误日志（errorlog）
记录着 MySQL 启动和停止，以及服务器在运行过程中发生的错误的相关信息。在默认情况下，系统记录错误日志的功能是关闭的，错误信息被输出到标准错误输出。

普通查询日志（general query log）
记录了服务器接收到的每一个命令，无论命令语句是否正确，因此会带来不小开销，所以也是默认关闭的。

慢查询日志（slow query log）
记录执行时间过长和没有使用索引的查询语句（默认 10s），同时只会记录执行成功的语句。

中继日志（relay log）
在从节点中存储接收到的 binlog 日志内容，用于主从同步。　



第一范式 第二范式 第三范式
