## 安装和使用

0. 直接官网下载安装。

1. 以管理员身份打开 cmd 命令行工具，切换到mysql的bin目录
2. 初始化数据库：mysqld --initialize --console
3. 执行完成后，会输出 root 用户的初始默认密码，如： CVCHW3(Pjke!
4. 输入以下安装命令：`mysqld install`
5. 启动输入以下命令即可：​`net start mysql`

我们希望使用用户：root，密码：123456。于是改密码：登录进mysql后：`alter user 'root'@'localhost' identified by '123456';`

**启动服务**：

- `net start mysql`	（如果拒绝访问就用管理员权限打开cmd，或者已经启动了）

**登录：**

- `mysql -u root -p` 回车， 然后输密码

**退出mysql** :

- `exit`

## 数据库的操作
- 显示所有数据库： `show databases;`
- 使用某个数据库：` use [name];`
- 创建数据库：`CREATE DATABASE 数据库名;`

``` mysql
create database if not exists mydb;  # 查询前判断是否存在 
```

- 删除数据库： `drop database 数据库名;`

## 数据表的操作

### 增加表和数据

- 创建数据表：`CREATE TABLE table_name (column_name column_type);`

``` mysql
例：
mysql> create table student(
    -> id int unsigned primary key AUTO_INCREMENT ,
    -> name varchar(20) not null,
    -> gender enum('男', '女'),
    -> birthday date,
    -> class varchar(30),
    -> is_active BOOLEAN DEFAULT TRUE);
```

- 向表中添加数据：`INSERT INTO table_name ( field1, field2,...fieldN ) VALUES ( value1, value2,...valueN );`

``` mysql
insert into student values  # 这里插入所有列的数据，所以省略了( field1, field2,...fieldN ) 
    -> (2, 'tom', '男', '2000-8-9', '一年级2班'),
    -> (3, 'lily', '女', '2000/11/19', '一年级2班');
```

### 删除表和数据

- 删除某张表： `drop table  [IF EXISTS]  表名;`

- 删除数据：`delete from 表名 where 条件`

``` mysql
DELETE FROM customers WHERE customer_id IN ( SELECT customer_id FROM orders WHERE order_date < '2023-01-01');
```

### 查询表和数据

- 显示此数据库中的所有表：`show tables;`
- 显示表的表头：`SHOW COLUMNS FROM 数据表;` 或者`desc 表名`
- 显示数据表的详细索引信息，包括 PRIMARY KEY（主键）: `SHOW INDEX FROM 数据表:`

- 查询数据：`select * from 表名 [where 条件 order by 字段1 asc, 字段2 desc]` ： 表示按照条件和字段1升序排序，如果字段1相同，就按字段2降序排列。

- 统计总共数据量：`select count(列序号，一般写1就行) from 表名;`
- 限制查询数量，常用于分页：`select * from 表名 limit n1,n2;` n1表示跳过前n1条（可省略）；n2表示一共查n2条数据。
- 连接两个SELECT 语句的结果到一个结果集合中去，并去除重复的行。

``` mysql
SELECT column1, column2, ... FROM table1 WHERE condition1
UNION
SELECT column1, column2, ... FROM table2 WHERE condition2
[ORDER BY column1, column2, ...];

SELECT city FROM customers
UNION
SELECT city FROM suppliers
ORDER BY city; # 把两个表中的city一起查询出来

SELECT product_name FROM products WHERE category = 'Electronics'
UNION
SELECT product_name FROM products WHERE category = 'Clothing'
ORDER BY product_name;  # 选择电子产品和服装类别的产品名称，并按产品名称升序排序。
```

### 修改表和数据

- 修改表的结构：MySQL 的 **ALTER** 命令用于修改数据库、表和索引等对象的结构。

``` mysql
ALTER TABLE old_table_name RENAME TO new_table_name;  # 修改表名
ALTER TABLE table_name ADD COLUMN new_column_name datatype; # 增加列
ALTER TABLE table_name CHANGE COLUMN old_column_name new_column_name datatype; # 修改列名
ALTER TABLE TABLE_NAME MODIFY COLUMN column_name new_datatype; # 修改列的数据结构
ALTER TABLE table_name DROP COLUMN column_name; # 删除表中一列
```

- 修改数据：`update 表名 set 字段名1=xxx,字段名2=xxx [where 条件]`

### where 子句

如需有条件地从表中选取数据，可将 WHERE 子句添加到 SELECT 语句中。WHERE 子句用于在 MySQL 中过滤查询结果，只返回满足特定条件的行。

在where子句中可以使用的操作：

- `=, !=, >, <, >=, <=, not, is null, is not null`

- 使用`and` 和`or` 连接多个条件

``` mysql
SELECT * FROM products WHERE category = 'Electronics' AND price > 100.00;
SELECT * FROM orders WHERE order_date >= '2023-01-01' OR total_amount > 1000.00;
```

- 模糊查询 like：

``` mysql
SELECT * FROM customers WHERE first_name LIKE 'J%'; # 查询名以J开头的人
```

- MYSQL不仅支持 like，还支持正则表达式。 使用 **REGEXP** 或**RLIKE**操作符来（两者等效）进行正则表达式匹配。

``` mysql
SELECT name FROM person_tbl WHERE name REGEXP '^st'; # 查找 name 字段中以 'st' 为开头的所有数据：
```

- 在取值空间中选择 in；也可以 not in；使用 in 可以连接select 语句

``` mysql
SELECT * FROM countries WHERE country_code IN ('US', 'CA', 'MX');
select * from student where id in (select id from student where gender="男"); # 此语句冗余，仅作展示
```

- 取值范围之间 between ，一般用在数字或时间

``` mysql
SELECT * FROM orders WHERE order_date BETWEEN '2023-01-01' AND '2023-12-31';
```

- 排序 order by，升序 asc，降序 desc

``` mysql
SELECT column1, column2, ... FROM table_name ORDER BY column1 [ASC | DESC], column2 [ASC | DESC], ...;
```

- 分组 group by；在分组的列上我们可以使用 COUNT, SUM, AVG, MAX, MIN等函数。

``` mysql
SELECT customer_id, SUM(order_amount) AS total_amount FROM orders GROUP BY customer_id;
```

### 连接多个表

我们可以在 SELECT, UPDATE 和 DELETE 语句中使用 MySQL 的 JOIN 来联合多表查询。

不常用的一种连接方式是笛卡尔积连接，它的速度最慢，尽量不要用，下面是一个示例：

``` mysql
select * from student, lession_student where student.id = lession_student.student_id ;
```

项目中常用的是内连接，JOIN 按照功能大致分为如下三类：

- `INNER JOIN`（内连接或等值连接）：获取两个表字段匹配关系的记录。

![img](img/img_innerjoin.gif)

``` mysql
SELECT column1, column2, ...
FROM table1
INNER JOIN table2 ON table1.column_name = table2.column_name;
# 如
SELECT order.order_id, customers.customer_name
FROM orders
INNER JOIN customers on orders.customer_id = customers.customer_id;
# 如
SELECT orders.order_id, customers.customer_name
FROM orders
INNER JOIN customers ON orders.customer_id = customers.customer_id
WHERE orders.order_date >= '2023-01-01';
```

- `LEFT JOIN`（左连接）：返回左表的所有行，并包括右表中匹配的行，如果右表中没有匹配的行，将返回 NULL 值，以下是 LEFT JOIN 语句的基本语法：

![img](img/img_leftjoin.gif)

``` mysql
SELECT column1, column2, ...
FROM table1
LEFT JOIN table2 ON table1.column_name = table2.column_name;
# 
SELECT customers.customer_id, customers.customer_name, orders.order_id
FROM customers
LEFT JOIN orders ON customers.customer_id = orders.customer_id;
```

以上 SQL 语句将选择客户表中的客户 ID 和客户名称，并包括左表 customers 中的所有行，以及匹配的订单 ID（如果有的话）。

- `RIGHT JOIN`（右连接）：获取右表所有的记录，即使左表没有对应匹配的记录。

![img](img/img_rightjoin.gif)

``` mysql
SELECT column1, column2, ...
FROM table1
RIGHT JOIN table2 ON table1.column_name = table2.column_name;
#
SELECT customers.customer_id, orders.order_id
FROM customers
RIGHT JOIN orders ON customers.customer_id = orders.customer_id;
```

以上 SQL 语句将选择右表 orders 中的所有订单 ID，并包括左表 customers 中匹配的客户 ID。如果在 customers 表中没有匹配的客户 ID，相关列将显示为 NULL。



## mysql数据类型

MySQL 支持多种类型，大致可以分为三类：**数值，时间日期，字符串**

### 数值类型

MySQL 支持所有标准 SQL 数值数据类型。

| 类型         | 大小                                     | 范围（有符号）                                               | 范围（无符号）                                               | 用途            |
| :----------- | :--------------------------------------- | :----------------------------------------------------------- | :----------------------------------------------------------- | :-------------- |
| TINYINT      | 1 Bytes                                  | (-128，127)                                                  | (0，255)                                                     | 小整数值        |
| SMALLINT     | 2 Bytes                                  | (-32 768，32 767)                                            | (0，65 535)                                                  | 大整数值        |
| MEDIUMINT    | 3 Bytes                                  | (-8 388 608，8 388 607)                                      | (0，16 777 215)                                              | 大整数值        |
| INT或INTEGER | 4 Bytes                                  | (-2 147 483 648，2 147 483 647)                              | (0，4 294 967 295)                                           | 大整数值        |
| BIGINT       | 8 Bytes                                  | (-9,223,372,036,854,775,808，9 223 372 036 854 775 807)      | (0，18 446 744 073 709 551 615)                              | 极大整数值      |
| FLOAT        | 4 Bytes                                  | (-3.402 823 466 E+38，-1.175 494 351 E-38)，0，(1.175 494 351 E-38，3.402 823 466 351 E+38) | 0，(1.175 494 351 E-38，3.402 823 466 E+38)                  | 单精度 浮点数值 |
| DOUBLE       | 8 Bytes                                  | (-1.797 693 134 862 315 7 E+308，-2.225 073 858 507 201 4 E-308)，0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308) | 0，(2.225 073 858 507 201 4 E-308，1.797 693 134 862 315 7 E+308) | 双精度 浮点数值 |
| DECIMAL      | 对DECIMAL(M,D) ，如果M>D，为M+2否则为D+2 | 依赖于M和D的值                                               | 依赖于M和D的值                                               | 小数值          |

### 时期时间

每个时间类型有一个有效值范围和一个"零"值，当指定不合法的MySQL不能表示的值时使用"零"值。

| 类型      | 大小 ( bytes) | 范围                                                         | 格式                | 用途                     |
| :-------- | :------------ | :----------------------------------------------------------- | :------------------ | :----------------------- |
| DATE      | 3             | 1000-01-01/9999-12-31                                        | YYYY-MM-DD          | 日期值                   |
| TIME      | 3             | '-838:59:59'/'838:59:59'                                     | HH:MM:SS            | 时间值或持续时间         |
| YEAR      | 1             | 1901/2155                                                    | YYYY                | 年份值                   |
| DATETIME  | 8             | '1000-01-01 00:00:00' 到 '9999-12-31 23:59:59'               | YYYY-MM-DD hh:mm:ss | 混合日期和时间值         |
| TIMESTAMP | 4             | '1970-01-01 00:00:01' UTC 到 '2038-01-19 03:14:07' UTC结束时间是第 **2147483647** 秒，北京时间 **2038-1-19 11:14:07**，格林尼治时间 2038年1月19日 凌晨 03:14:07 | YYYY-MM-DD hh:mm:ss | 混合日期和时间值，时间戳 |

### 字符串类型

| 类型       | 大小                  | 用途                            |
| :--------- | :-------------------- | :------------------------------ |
| CHAR       | 0-255 bytes           | 定长字符串                      |
| VARCHAR    | 0-65535 bytes         | 变长字符串                      |
| TINYBLOB   | 0-255 bytes           | 不超过 255 个字符的二进制字符串 |
| TINYTEXT   | 0-255 bytes           | 短文本字符串                    |
| BLOB       | 0-65 535 bytes        | 二进制形式的长文本数据          |
| TEXT       | 0-65 535 bytes        | 长文本数据                      |
| MEDIUMBLOB | 0-16 777 215 bytes    | 二进制形式的中等长度文本数据    |
| MEDIUMTEXT | 0-16 777 215 bytes    | 中等长度文本数据                |
| LONGBLOB   | 0-4 294 967 295 bytes | 二进制形式的极大文本数据        |
| LONGTEXT   | 0-4 294 967 295 bytes | 极大文本数据                    |

**注意**：char(n) 和 varchar(n) 中括号中 n 代表字符的个数，并不代表字节个数，比如 CHAR(30) 就可以存储 30 个字符。

CHAR 和 VARCHAR 类型类似，但它们保存和检索的方式不同。它们的最大长度和是否尾部空格被保留等方面也不同。在存储或检索过程中不进行大小写转换。

BINARY 和 VARBINARY 类似于 CHAR 和 VARCHAR，不同的是它们包含二进制字符串而不要非二进制字符串。也就是说，它们包含字节字符串而不是字符字符串。这说明它们没有字符集，并且排序和比较基于列值字节的数值值。

BLOB 是一个二进制大对象，可以容纳可变数量的数据。有 4 种 BLOB 类型：TINYBLOB、BLOB、MEDIUMBLOB 和 LONGBLOB。它们区别在于可容纳存储范围不同。

有 4 种 TEXT 类型：TINYTEXT、TEXT、MEDIUMTEXT 和 LONGTEXT。对应的这 4 种 BLOB 类型，可存储的最大长度不同，可根据实际情况选择。

### 枚举与集合类型（Enumeration and Set Types）

- **ENUM**: 枚举类型，用于存储单一值，可以选择一个预定义的集合。
- **SET**: 集合类型，用于存储多个值，可以选择多个预定义的集合。



## 可能的出错

使用navicat连接mysql数据库时 报错：Client does not support authentication protocol requested by server 

解决方法：[MySql错误 1251 - Client does not support authentication protocol requested by server 解决方案-CSDN博客](https://blog.csdn.net/OCEAN_C/article/details/89719578)

``` mysql
mysql> alter user 'root'@'localhost' identified with mysql_native_password by '123456';
mysql> flush privileges;
```







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
