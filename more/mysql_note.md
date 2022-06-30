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

