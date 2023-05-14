[markwdown语法]([GitHub - younghz/Markdown: Markdown 基本语法。](https://github.com/younghz/Markdown))

# 学python

## 1. python

​		python是一种解释型语言，没有编译这个环节，Java中有，将.java编译成.class，

​		python 中数字类型和字符串类型的值是**不可变**的，表面上改变值实际上就是指向了另一个内存空间

​		python 中，变量是以内容为基准而不是像 c 中以变量名为基准

​		‘’‘ 多行注释

​		‘’‘			其实就是没被引用的字符串



## 2. print

```python
print(self, *args, sep=' ', end='\n', file=None)  # 输出函数，可自定义分隔符和结尾符
fp = open('D:/text.txt', 'a+')
print("some words", file=fp)	
fp.close() # 输出到文件
```

​				

## 3. \r 和 \n 

​	\r 是return，回车，
​	\n是newline，换行
​	\r 就是简单的把光标移动到此行的开头，所以如果\r后继续输出内容就会覆盖掉前面的内容
​	\b backspace 向前退一个格，所以如果\b后继续输出内容就会覆盖掉前面的内容的最后一个字符

​	

## 4. 浮点数计算不精确

```python
print(1.1+2.2) # = 3.3000...00003  # 浮点数计算不精确
print(Decimal('1.1')+Decimal('2.2'))   # 用decimmal模块来解决
```



## 5. 常用运算符

​	python中// 是整除， ** 是求次方
​	python 支持链式赋值
​	a = b = c = 1
​	python 支持解包赋值
​	a, b, c = 1, 2, 3

python && 必须用 and， || 必须用or ， 异或 ^ ,  取反 ~，位运算不变，用 not 表示非



## 6. in 和 not in

简单的判断字符串是否在字符串中，数字是否在数组中（数字类型不严格检查）

```
ls1 = [11, 22, 33]
print(11 in ls1)      # True
print(11.0 in ls1)    # True
print('he' in 'hello')  # True
```



## 7. is 

```python
print('对象的数据类型：', type(num))
print('对象的存储地址：', id(num))
a = b =1 
a is b  # True
a = 1.0 
b = 1 
a is b  #  返回false 
# 	is 关键字比较的就是两个对象的id，
#  简单来说 is 比较的是地址，还可以用 is not
```



## 8. 一切皆对象，对象皆bool

​		python一切皆对象，所有对象都有bool值，用bool(obj)来得到.
以下对象的bool值为**False**：False, 0, 0.0, None, 空字符串，空列表 [], list()，空元组(), tuple()，空字典{}, dict()，空集合 set()
所有任何对象都可以直接用在条件表达式的条件部分，会隐式转换

本质是每个对象都从object那里继承了  \__bool__() 函数，通过重写这个函数可以修改对象的bool值



## 9. 杂项 if, pass, for, range，yield，sum

```python
if condition：
	x
else :
 	y
# 可简写为  x if condition1 else y

# pass 语句做为占位符 ，什么也不做

# range(start, stop, step) 创建一个整数序列，参数只有stop是必须的，只有在用到range对象时才回去计算序列中的相关元素

# for in 语句   for 是专门用来遍历可迭代对象的， while和 do while 才是循环
for item in 'hello':
	print(tiem)
    
''' else if 在python在中叫 elif
	else 可以和while或者for以及try语句搭配使用
	也是在循环最后条件不成立时执行else
	但break跳出来就不会执行else
'''
for i in range(10):
    print(i)
    if i == 8:
        break
else:
    print('qw')		# for 正常跑完就会走到else，有跳出的话就不会走这里
   
又是函数返回多个值而我们只需要接受其中的部分时可以用_
a, b, _ = fun()		# fun() 返回三个值而我们只需要前两个
next(iter)		# 跳过可迭代对象的一个元素
next(file)		# 跳过文件的一行

#----------yield-----------
#yield的作用是将一个函数变成一个迭代器
def f():
    a = 1
    for i in range(7):
        yield a
        a += i
        if a > 10:
            return
b = f()
print(type(b))
print(b)
print(list(b))
#<class 'generator'>
#<generator object f at 0x000001735C71B3C8>
#[1, 1, 2, 4, 7]

#----------------sum----------------
a = sum(i for i in range(5))
print(a) #10 可直接求得iteration的值，而不用遍历了
```



## 10. 列表 list

相当于C的数组， 由 list() 或[ ... ] 来创建 ，**没有数据类型限制**

注意，列表的id不是首地址，列表的id和列表中元素的id没有任何关系，list的对象指向了一段空间，这段空间中存在着列表个元素的id

列表可以用**负数索引**来去元素，最后一个元素索引是-1，以此类推

python根据需要动态分配内存和回收内存

#### 1. list的特点

- [] , 有序，可重复，元素可变 

#### 2.list的切片

```
		1. list[start: stop: step], 切片的结果是原列表的拷贝
		2. list.index(val) 获得索引
		3. 如果step是负数，则会从右往左计算进行切片，specially，list[::-1] 返回列表的倒序
		4. list[start: stop: step] 同样可以作为左值 ，故可以修改 ls1[1, 3] = ls2,  也可删除 ls1[1,3] = []
		5. 多维列表也是这样切片 list[ start: stop: step ,  start: stop: step]
```

#### 3.增删改

1. append(obj) 在最后添加一个元素
2. extend(obj)在list末尾添加obj中所有元素，而不是把obj作为一个元素添加
3. insert(index,obj) 在指定位置插入obj
4. remove(obj) 删除这个元素，只删一次， 没有这个元素就抛出异常
5. pop(index) 删除指定位置的元素，默认last
6. clear() 清除所有元素
7. del list 删除列表对象

#### 4. 列表生成式

```python
ls2 = [i for i in range(1, 10)]  # 1 到 9
ls2 = [i**2 for i in range(1, 10)] #  1 到 81
print([i for i in range(1,10) if i%2 == 0]) # 后面还可以直接跟条件语句 [2,4,6,8]
```



## 11. 字典

通过 {key: value, ...} 或 dict（key=value, ...） 来创建

#### 1. 特点

{key: value}  无序，key不可重复，value可重复，元素可变

> 字典中键值对在存放时会先计算key的哈希值作为存放地址，所以字典中key不能改变，不能重复，只能是不可变对象（如不能使用list作为key）

可以用in 或者 not in 判断**key**是否在字典当中

#### 2. 获取元素

通过dict[key] 或 dict.get(key) 来获得值，其中，dict[不存在的key] 系统会报错，

dict.get(不存在的key) 得到None，甚至dict.get(不存在的key，defaultvalue) 可以有默认值

dict.keys() 获得所有的key  然后可以用list(keys)转成列表
dict.values()与上同理
dict.items() 得到所有键值对，每个键值对组成一个元组，然后共同放在一个list中

for it in dict： 其中it只是key

```python
for i in dct:
    print(i, dct[i])
```

#### 3. 增删改

1. del dict[key] 删除指定键值对
2. dict.clear() , 清空字典
3. dict[key] = value   来新增或修改

#### 4. 字典生成式

``` python
dic = { item: price for item , price in zip(items, prices) }
```

> 此zip() 用于将可迭代对象作为参数，将对象中对应的元素打包成一个元组，然后将所有元组组成一个列表



## 12. 元组

不可变序列  tuple(('1', 2, True))  也可以tuple( [1,2] ) 或 (1,'2',True) 甚至可以省略这个括号 tp = 1, '2'

元组有自动装包解包过程

> ​	注意：只包含一个元素的元组声明时要多加一个逗号 tp = (10, ) 

#### 1. 特点

有序，可重复，**不可变**

>  不可变的好处：多任务环境下就不需要加锁了

#### 2. 获取元素

因元组是有序的，故直接用tuple[ index ] 就可以访问

同样可用负的index来从右往左访问

#### 3. 增删改

1. 不能删除元组中的元素，只能删除整个元组 del tuple

2. 不能给元组新增元素，除非新建一个再重新赋值

3. 不能修改元素，tuple [ index ] 不能做左值，tuple[ index ] = value 是非法的

4. 元素的截取也都是会返回一个新的元组，各种截取操作与列表的操作一致

#### 4. 元组生成式

```python
tp = tuple(i for i in range(10))
# 必须显示地加上tuple 不然会认为这只是一个generator
```



## 13. 集合

#### 1. 特点

1. 无序，不可重复，可变
2. 在集合中元素的地址是通过hash计算出来的
3.  s = { 1, 'qw', True} 或者 set( list / set/ tuple / str / range/ 可迭代对象) 各种方法都可得到集合
4. 集合中不会存储相同的元素，集合就是没有value只有key的字典

#### 2. 集合和集合的关系

```python
# 直接用==判断是否相等
s1.issubset(s2) 	# 判断s1是否是s2的子集
s1.issuperset(s2)	# 判断s1是否是s2的超集
s1.isdisjoint(s2)	# 判断s1和s2是否没有交集
s1.intersection(s2)  = s1 & s2 	# 返回两者的交集, 
s1.union(s2)  =  s1 | s2		# 返回两者的并集，  
s1.difference(s2) =  s1 - s2  # 返回s1在s2中的差集， 
s1.symmetric_difference(s2)   =  s1.union(s2) - s1.intersection(s2) = s1 ^ s2  # 返回对称差集

```

#### 3. 增删改

1. 集合中元素的判断 in 和 not in 
2. add(obj)  	增加一个元素，集合中不能添加无法hash的对象，如不能添加list，set，dict，
3. update(obj), 将各种可迭代对象**其中的值**添加进去
4. remove(obj) 删除一个元素，obj不存在就抛出异常
5. discard(obj) 删除一个元素，不抛出异常
6. pop() 随机删除一个

#### 4. 集合生成式

``` python
{i**2 for i in range(10, 15)}
```



## 15. 字符串

字符串 str ， python中的基本数据类型，

#### 1. 特点

str是一个**不可变**的字符序列，字符串也不会被修改，人为地修改实际上就是创建了新的字符串然后指向它

字符串拼接很消耗资源，所以拼接建议使用join()方法，而非+，因为join() 只new一个对象

> ----字符串的驻留机制---- 字符串在内存中只保留一份，即使由多个变量指向它
> 		字符串的长度为0或1， 符合标识符的字符串，字符串只在编译时进行驻留而非运行时，[-5, 256] 之间的整数数字
> 		（交互模式下才能体现，pycharm有字符串优化）
> 		s1 = sys.intern(s2) 强制使两个字符串指向同一个字符串常量
> 		字符串唯一地放在字符串常量池中，创建相同字符串时都是指向的同一份字符串，不会开辟开空间，

​	

#### 2. str的函数

##### 1. 查找

``` python
index() 和 rindex() ： 查找子串第一次(最后一次）出现的位置，找不到抛出异常ValueError
find() 和 rfind() ： 同上， 只是找不到不会抛出异常，而是返回-1  ， 建议多用find()
```



##### 2. 大小写相关

``` python
lower() 将所有字符转化为小写		islower()  字符串中的字母是否全为小写
upper() .............大写		 isupper() ....................大写
swapcase() 将所有大写转换为小写，所有小写转换为大写
capitalize() 将首字母大写， 其余小写
title() 将每个单词的第一个字符转换为大写，剩余字符转换为小写
```

##### 3. 字符串对齐

```python
str.center(总宽度，填充字符='')	把str 居中
str.ljust() , str.rjust() 参数同上，左对齐，右对齐
str.zfill(宽度)  右对齐，用0填充，一般对数字字符串使用，会自动识别负号等
	
```

##### 4. 字符串分割

````python
str.split(sep='',maxsplit), 以分隔符来分割，返回字符串列表
str.rsplit(sep='',maxsplit) 从右侧开始分割
````

##### 5. 字符串的判断

```python
isidentifier() 是否是合法的标识符
isspace()  是否全由空白字符组成 （回车，换行，制表符）
isalpha() 是否全是字母
isdecimal() 是否全是十进制数字
isnumeric() 是否全是数字  汉字数字，罗马数字都是
isalnum() 是否只由字母或者数字组成， 包括汉字和汉字数字
```

##### 6. 替换和连接

```python
str.replace(old, new, count=len(str)) 用old字符串替换new字符串，换count次， 
str.join(obj) 将列表，元组或字符串中的字符(串)合并为新的字符串，连接符是str
# 将多个字符串连接起来  当然直接用加号是可以的
# 用格式化工具  耗时差不多
s1, s2, s3 = '我', '爱你', '中国'
s = '%s%s%s'% (s1, s2, s3)
```

##### 7. 字符串比较

每个字符按序比较ascii码 unicode码

ord(字符) 获得字符的原始值ascii码， chr(数字) 获得数字对应的字符 

##### 8. 字符串切片

切片产生新的对象	s2 = s1[start：stop： step]

##### 9. 字符串格式化

 原字符，不要转义字符起作用 在字符串前加上r 或R

用 % 或者 {} 做为占位符

```python
%s: 字符串，%i：整数， %f：浮点数 %d：整数
例：print('qwer %f qwewqe %i' % (41.2, 0b11))

{0}, {1} 顺序占位符
例：print('asdas {0}asd {1} asdas{0}'.format(12, '___'))

用f-string
name = 'tom'	
age = 12
print(f'姓名 {name} ， 年龄 {age}')
```

##### 10. 精度和宽度

```python
print('%10d' % 66) ， 输出填充到10个字符
print('%.3f' % 3.1415926)  输出的浮点数保留三位小数	
%x.yd 小数点前表示宽度，小数点后表示精度
在{0:.3}中的3表示保留3位数字
```

##### 11. 字符串的编码转换

````python
utf_a = '天涯共此时啊'    # 默认utf-8
gbk_a = utf_a.encode('GBK')
print(gbk_a.decode('GBK'))
````



## 16. 函数

#### 1. 函数的定义 

````python
def functionName(arg1 = 1 : int , arg2....) -> returnType
````

> 可以写默认值，但注意不要引起歧义，大不了人为加上形参名就好 fun(a = 1, b =2), 默认值定义必须从右往左

#### 2. 返回类型

​		不需要返回值可以不return,  也可在一定条件下 return

​		如果返回多个值，直接 return res1, res2, res3   就相当于把它们三个组成一个**元组**，调用处可以直接用三个引用接收即可， r1, r2, r3 = fun()

​		函数的返回类型不需要明确说明，当然也可以用 def fun() -> int : 来指明需要返回int，但即使不return也不报错，而且就算返回其他类型也只是警告，不会报错

	#### 3. 变量的作用域

和 C++ JAVA 差不多， 只是给局部变量加上global关键字后就变成全局变量了，然后global的对象和函数外的全局变量指向相同的引用

``` python
a = 1
def fun():
    global a
    a += 1
    return a
print(fun())		# 2
```

#### 4. 函数的参数

##### 1. 参数传递

> 对于不可变对象，如果int，str, tuple等等，在函数中改变这些的形参值都不会影响实参

> 而对于可变对象如 list，dict, set  改变了形参就改变了实参

```python
def fun2(a, ls1, ls2):
    a = (12, 13)
    ls1 = [1, 2, 3]
    ls2.append(6)
a = (10, 11)
lst1 = [1, 2]
lst2 = [4, 5]
fun2(1, lst1, lst2)
print(a, lst1, lst2)
# (10, 11) [1, 2] [4, 5, 6]
```

##### 2. 个数可变的位置形参

​		函数定义时不知道参数时 可以使用 def fun(*args)  传入的多个实参会放在一个**元组**里给args

````python
def fun3(*args):
    print(type(args))
    for i in args:
        print(i, end=', ')
fun3(1, '2', True, [i for i in range(10, 15)])
# 输出
# <class 'tuple'>
# 1, 2, True, [10, 11, 12, 13, 14], 
````



##### 3. 个数可变的位置实参

```python
def fun6(a, b, c):
	print('a =', a, 'b = ', b, 'c=', c
ls1 = [10, 20, 30]
fun6(*ls1)   
fun6(*range(3)) 
fun6(*(2,4,6))  #  把各种可迭代对象作为位置实参传入函数
```
##### 3. 个数可变的关键字形参

函数定义时不知道参数时 可以使用 def fun(\**args)  传入的多个实参会放在一个**字典**里给args

```python
def fun3(**args):
    print(type(args))
    for i in args:
        print(i, args[i], sep=':', end=', ')
fun3(a = 10, b = '11', c = 12.0)
# 输出
# <class 'dict'>
# a:10, b:11, c:12.0, 
```
##### 4. 个数可变的关键字实参

```python
def fun6(a, b, c):
	print('a =', a, 'b = ', b, 'c=', c)
di1 = {'a': 11, 'b': 22, 'c': 33}
fun6(**di1）	 # 注意形参名要和key一致
```



## 17. 异常处理try

```python
import traceback
# 可以打印异常的信息

try:
	可能出错的代码
except someError1:
	捕获异常之后的操作
except someError2 as e:
	捕获异常之后的操作
    traceback.print_exc()
	......
except BaseException(所有异常的父类):
	捕获异常之后的操作
else:
	没由捕获到异常之后的操作
finally:
	不管捕获没有，最终都要进行的操作
```



## 18.类

#### 1. 类的组成

- 类属性
- 实例属性
- 实例方法
- 静态方法
- 类方法

```python
class Student:
    species = '人'   # 这是类属性，所有对象共享

    def __init__(self, name, age):
        print('这是构造函数')
        self.name = name  # name 和 age就是实例属性
        self.age = age

    def eat(self):      # 实例方法， 第一个参数必须是实例对象
        print(self.name, '学生在吃东西', self)

    @staticmethod
    def sfunc():       # 静态方法
        print('这是静态方法')

    @classmethod
  	def cfunc(cls):     # 类方法   cls就是Student
        print('这是类方法', cls.species)
```



#### 2. 类中属性和方法的调用

```python
print(Student.species)  # 类名调用类属性
st1 = Student('tom', 19)
st2 = Student('jerry', 11)
print(st1.name)         # 调用实例属性
print(st2.species)      # 对象调用类属性
st1.eat()               # 只能由对象来调用实例方法	Student.eat(st1) 也行
Student.sfunc()         # 静态方法可以由对象或者类名来调用
st1.sfunc()
Student.cfunc()         # 类方法可以由对象或者类名来调用
st2.cfunc()
```



#### 3. 动态绑定属性和方法

python 是动态语言，在创建对象后仍然可以动态地绑定属性和方法（即使没有这些属性和方法）

```python
def show():
    print('这是show方法，在class之外')

st2.gender = '男'    # 只给st2这个属性了
print(st2.name, st2.age, st2.gender)
st2.show = show  # 只给st2绑定这个方法
st2.show()      # 只由st2能show    
```

#### 4.访问权限

python没有专门规定访问权限的关键字，如果不希望某个属性不能再类外访问，可以给属性前加两个_

```python
class Teacher:
    def __init__(self, name, income):
        self.name = name
        self.__income = income

t = Teacher('bob', 100)
print(t.name)
print(t.__income)		# 报错 'Teacher' object has no attribute '__income'
print(t._Teacher__income) 		# 但这样还是可以强行的访问，不讲武德
```

#### 5. 继承

```python
class 子类类名(父类1, 父类2, ...):
    pass
```

python中所有类默认继承object，就像java一样

但python允许多继承

```python
class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    def show(self):
        print(self.name, self.age)

class Student(Person):
    def __init__(self, name, age, stuid):
        super().__init__(name, age)
        self.stuid = stuid
        
    def show(self):	# 方法重写
        super().show()				# 通过super()得到一个父类的对象
        print(self.stuid)

class Teacher(Person):
    def __init__(self, name, age, tchid):
        super().__init__(name, age)
        self.tchid = tchid

st1 = Student('dique', 18, 111110)
tch1 = Teacher('tom', 55, 152020)
st1.show()
tch1.show()
```



#### 6.方法重写

python只有方法重写没有方法重载，因为传入的参数是不指定类型的。

重写和C++，Java一样的



#### 7. 类方法和静态方法

都可以直接由类来调用，

- 类方法调用会传入类本身cls， 可以通过 cls 来获取到类属性和其他类方法
- 静态方法与这个类关系不大，不能获得类的任何属性和方法，



#### 8. object类

任何类都继承object类，任何类和对象都可以通过dir(类或对象名) 来查看object类的方法

object类中的方法

- \__str__()  返回对象的描述，推荐重写，相当于Java中的toString()，可以直接输入对象来获得描述了

- \__sizeof__() 返回对象的内存大小，单位byte

- 对象.\__dict__ 返回对象 { 属性名：属性值 } 的字典

- 类.\__dict__ 返回类 { 属性名：属性值 } 的字典

- 对象.\__class__返回对象的类

- 类.\__bases__ 返回类的父类

- 类.\__base__ 返回类的第一个父类

- 类.\__mor__ 返回类的层次结构

- 类.\__subclasses__() 返回类的子类们

- ```python
	class A:
	    def __init__(self, num):
	        self.num = num
	
	    def __add__(self, other):
	        n = self.num + other.num
	        res = A(n)
	        return res
	# 重写__add__ 可以使得对象获得加法运算，other没有类型限制，返回值也没有，想返回啥都行
	    def __sub__(self, other):
	        return self.num - other.num
	# 减法同理
	    def __mul__(self, other):
	        return self.num * other.num
	```

- ```python
	    def __bool__(self):
	        if self.num % 2 == 0:
	            return True
	        else:
	            return False
	     # 修改对象的bool值
	```

- ```python
	def __len__(self):
	    return self.num
	
	# 对象就可以使用内置函数 len(obj) 来获得 __len__的返回值
	```

- object的 \__new__(cls, *args, **kwargs): 函数，才是真正能实例化对象的函数，init是给对象的属性赋值而已

	```python
	class Person:
	    def __init__(self, name, age):
	        print('__init__被调用了，self的id为 %s' % id(self))
	        self.name = name
	        self.age = age
	
	    def __new__(cls, *args, **kwargs):
	        # obj = super().__new__(cls)
	        # print("调用new, cls的id为{0}".format(id(cls)))
	        # print('创建的对象的id为%s' % id(obj))
	        # return obj
	        pass
	
	
	p1 = Person('tom', 12)
	print(p1.age)
	# 连 init 都不会执行，且p1.age 还会报错'NoneType' object has no attribute 'age'，
	# p1 就是None，根本没有实例化任何东西
	
	# 取消new中的注释会发现，new中的obj就是新建的对象，就是init中的self
	```

- 



#### 9. python的多态

```python
class Animal:
    def eat(self):
        print('动物吃东西')

class Dog(Animal):
    def eat(self):
        print('狗吃骨头')

class Cat(Animal):
    def eat(self):
        print('猫吃鱼')

class Person:
    def eat(self):
        print('人吃五谷杂粮')


def fun(animal):
    animal.eat()


fun(Dog())
fun(Cat())
fun(Animal())
fun(Person())
# python是弱类型语言，无法实现java c++那样的多态。
# 静态语言实现多态的三个必要条件：继承，方法重写，父类引用指向子类对象
# 动态语言不关心数据的类型，只关系对象的行为
```



#### 10. 浅拷贝与深拷贝

- 变量的赋值： cpu1 = CPU() ; cpu2 = cpu1;  cpu2 只是指向了cpu1指向的实体，并没有创建新的对象

- 浅拷贝：python默认的拷贝都是浅拷贝，拷贝时，对象包含的子对象内容不拷贝，因此，源对象和拷贝对象会使用同一个子对象

	```python
	class CPU:
	    pass
	class Disk:
	    pass
	class Computer:
	    def __init__(self, cpu, disk):
	        self.cpu = cpu
	        self.disk = disk
	
	cpu1 = CPU()
	cpu2 = cpu1
	# print(id(cpu1), id(cpu2))		# 显然一样
	disk1 = Disk()
	pc1 = Computer(cpu1, disk1)
	
	import copy
	
	pc2 = copy.copy(pc1)
	print(id(pc1), id(pc2))			# 拷贝的对象两者不相同， 浅拷贝只拷贝源对象	
	print(id(pc1.cpu), id(pc2.cpu)) # 浅拷贝，对象的子对象两者相同
	```

- 深拷贝：使用copy 模块的 deepcopy 函数，递归拷贝对象中包含的子对象，源对象和拷贝对象所有的子对象都不相同

	```python
	pc3 = copy.deepcopy(pc1)
	print(id(pc1), id(pc3))				# 源对象不同
	print(id(pc1.cpu), id(pc3.cpu))		# 子对象也不同
	```



## 19. 模块

- ***模块***：module，一个.py的文件就是一个模块

- ***函数和模块的关系***：一个模块里包含多个函数

- ***模块的好处***：
- - 方便其他程序或脚本导入并使用
	- 避免变量名和函数名冲突
	- 提高代码的可维护性
	- 提高代码的可重用性

- ***自定义模块***： 新建一个.py 文件即可（模块名不要与python自带的模块名相同）

- ***导入模块***
	- import 模块名 [as 别名]
	- from 模块名 import 函数/ 变量 / 类

- ***模块的使用***

	- ``` python
		import demo17_module
		
		print(type(demo17_module))		# <class 'module'>
		print(id(demo17_module))
		stu = demo17_module.Student() 		# 使用 模块名.模块中内容 来使用
		```

	- ``` python
		from math import pi
		print(pi)					# 这样导入就可以直接使用
		```

	- 如果自己写的模块导入时出错，可以右键模块所在的directory，mark directory as ， source root



## 20. main

```python
def add(a, b):
    return a + b

print('code out of the main')

if __name__ == '__main__':
    print(add(1, 2))

print('code out of the main2')
# 运行此程序，这三句都会按序运行，

# 但如果将此模块导入其他程序时，那么main 里面的程序就不会运行了
```



## 21. python中的包

- 包是一个本层次的目录结构，它将一组功能相近的模块组织再一个目录下，
- 作用：代码规范，避免名称冲突
- ***包 package 和 目录 directory 的区别***
	- 包含 \__init\__.py文件的是包,  \__inti__文件就写了此包可以完成的功能等等
	- 没有的就是目录， 普通的资源目录
- 包的导入：
	- import 包名 . 模块名  [as 别名]
	- form 包名 import 模块名  [as 别名]
	- form 包名 . 模块名 import 类 / 函数 / 变量



## 22. python中常用的内置模块

- sys：与python解释器机器os相关的标准库
- time：与时间相关的模块

``` python
import time

startTime = time.perf_counter()
'''
执行代码
'''
endTime = time.perf_counter()
runTime = endTime - startTime	# 单位是秒
print(runTime)

startTime = time.perf_counter_ns()
'''
执行代码
'''
endTime = time.perf_counter_ns()    # 单位是纳秒
runTime = endTime - startTime
print('耗时: %.3f' % (runTime / 10**6), 'ms')
```

- os
- calendar
- urllib
- json
- re：正则
- math
- decimal： 用于进行精确控制运算精度、有效位数金额四舍五入的十进制运算
- logging



## 23. 第三方模块的安装

- cmd - pip install 模块名
- 不行的话就 cmd - python -m pip install 模块名  , pip 不行就一律加上 python -m 



## 24. 编码格式

- python的解释器使用的是unicode(内存)
- .py 文件再磁盘上使用 utf-8 （外存）



## 25. 文件操作

- 内置函数open() 创建文件对象
	- file = open(filename [, mode = r, encoding ])   ，默认只读， 默认文本文件中字符的编码格式是gbk

- 读取文件内容

	```python
	file = open(file='../../a.txt', mode='r')
	if file:
	    print('打开文件成功')
	    print(file.readlines())     # 读取每行的内容并放在一个列表中
	file.close()
	```

#### 常用文件打开模式

| 打开模式 |                        描述                        | 文件指针位置 |
| :------: | :------------------------------------------------: | :----------: |
|    r     |             只读模式，文件不存在就报错             |   文件开头   |
|    w     | 只写模式，文件不存在就创建，文件存在会覆盖原有内容 |   文件开头   |
|    a     |   追加模式，文件不存在就创建，在文件末尾追加内容   |   文件末尾   |
|    b     |      二进制方式打开，不能单独使用，如：rb, wb      |      -       |
|    +     |         读写方式打开，不能单独使用，如：a+         |      -       |

- - 文本文件：存储普通字符，默认为Unicode字符集
	- 二进制文件，内容以字节进行存储，如 mp3 , png

- ```python
	# 二进制文件的复制	..\ 退回到上一级目录
	file = open(file='a.ico', mode='rb')
	copyfile = open(file='copya.ico', mode='wb')
	copyfile.write(file.read())
	copyfile.close()
	file.close()
	```

#### 文件常用方法

|                 方法名 |                          描述                          |
| :---------------------: | :----------------------------------------------------: |
|           read([size]) |               读取size个字符，不然就全部               |
|             readline() |                        读取一行                        |
|            readlines() |             读取每行的内容并放在一个列表中             |
|             write(str) |                    将字符串写入文件                    |
|     writelines(s_list) |        将字符串列表写入文本文件，不会添加换行符        |
| seek(offset[, whence]) |    把文件指针移动到新的位置上<br />whence自己搜索吧    |
|                 tell() |                 返回文件指针当前的位置                 |
|                flush() |          把缓冲区的内容写入文件，而不关闭文件          |
|                close() | 把缓冲区的内容写入文件，并且关闭文件，释放文件相关资源 |

#### with语句

- with语句可以自动管理上下文资源，无论什么原因跳出来 with语句块，都能保证文件正确的关闭，以保证资源的释放

	```python
	with open('a.txt', 'r') as file:
	    print(file.read())
	```

- 上下文管理器：一个类对象实现了 \_\_enter\_\_()  和 \_\_exit\_\_() ，那么这个类就遵守上下文管理协议，这个类的实例对象就称为上下文管理器，open('a.txt', 'r') 返回的这个对象就是上下文管理器

	```python
	# 该类遵守了上下文管理协议
	class MyContentMgr:
	    def __enter__(self):
	        print('enter方法被执行了')
	        return self
	
	    def __exit__(self, exc_type, exc_val, exc_tb):
	        print('exit方法执行了')
	
	    def show(self):
	        print('show方法') # print('show方法', 1/0) 即使这里抛出异常了也会调用exit方法
	
	with MyContentMgr() as mcm:
	    mcm.show()
	    print('with语句中')
	
	print('with语句外')
	
	# 运行结果如下
	# enter方法被执行了
	# show方法
	# with语句中
	# exit方法执行了
	# with语句外
	```



## 26. 目录操作

- `os模块`是python内置的与操作系统功能和文件系统相关的模块，该模块中的语句的执行结果与操作系统相关

- os模块和`os.path `模块用于对目录或文件进行操作

	```python
	# 调用可执行文件
	# os.system('notepad.exe')
	# os.system('calc.exe')       # 上一个程序要关闭之后才能调用到下一个程序
	
	os.startfile(r'D:\Notepad++\notepad++.exe')
	```

- os模块操作目录相关函数

|              函数               |                         描述                         |
| :-----------------------------: | :--------------------------------------------------: |
|            getcwd()             |                  返回当前的工作目录                  |
|          listdir(path)          |            返回指定路径下的文件和目录信息            |
| mkdirs(path1/path2..  [, mode]) |                     创建多级目录                     |
|     makedir[path [, mode]]      |              创建目录，目录已存在就报错              |
|           rmdir(path)           |                       删除目录                       |
|  removedirs(path1/path2 ... )   |                     删除多级目录                     |
|           chdir(path)           |               将path设置为当前工作目录               |
|           walk(path)            | 遍历path及其所有子目录，返回元组加列表，自己百度用法 |

- os.path模块操作目录相关函数

	|       函数       |                  描述                  |
	| :--------------: | :------------------------------------: |
	|  abspath(path)   |        获取文件或目录的绝对路径        |
	|   exists(path)   |         判断文件或目录是否存在         |
	| join(path, name) |      将目录与目录或文件名拼接起来      |
	|   split(path)    |            分离目录和文件名            |
	|  splitext(path)  |           分离文件名和拓展名           |
	|  basename(path)  |         从一个目录中提取文件名         |
	|  dirname(path)   | 从一个路径中提取文件路径，不包括文件名 |
	|   isdir(path)    |             判断是否为路径             |

	​	

	```python
	import os.path
	
	absp = os.path.abspath('demo25_ospath.py')
	print(os.path.exists(absp), os.path.exists('demo1.py'))
	print(os.path.isdir(absp))
	print(os.path.join(absp[:absp.rfind('\\'):], 'demo25_ospath.py'))  # 又得到absp了
	
	print(os.path.split(absp))
	
	print(os.path.basename(absp))
	print(os.path.dirname(absp))		# 分开目录和文件
	```



* `pathlib模块`

	pathlib模块提供了一个面向对象API来解析、建立、测试和处理文件名和路径，而不是使用底层字符串操作。

	- 目录操作
	
	- ``` python
		from pathlib import Path
		p = Path()						# 当前目录
		p = Path('a', 'b', 'c/d')		# 当前目录下的a/b/c/d
		p = Path('/etc/')				# 根下的etc目录
		
		```

	- 获得父目录
	
	- ``` python
		import pathlib
		print(pathlib.Path.cwd().parent)		#获得父目录
		print(pathlib.Path.cwd().parent.parent)	#获得爷目录
		
		```

	- 路径拼接与分解
	
	- ``` python
		p = Path()
		print(p)
		p = p / 'a'
		print(p.absolute())		# 得到绝对路径    D:\JetBrains\pythonProject\convolution\practice1\a
		p1 = 'b' / p
		print(p1.absolute())	# 注意b所在的位置  D:\JetBrains\pythonProject\convolution\practice1\b\a
		p2 = Path('c')
		p3 = p2 / p1
		print(p3)			    # c\b\a
		print(p3.parts)			# ('c', 'b', 'a')
		print(p3.joinpath('etc', 'init.d', Path('httpd')))  # 将调用者和参数都连起来
		# c\b\a\etc\init.d\httpd
		```





## 27. 项目打包

- 安装第三方模块	pip install PyInstall

- 执行打包操作

- - pyinstaller -F  程序所在绝对路径	-F: 只生产一个exe文件
	- C:\Users\用户名\dist\程序名.exe   

	

## 28. 位运算

| 位运算符 | 说明     | 使用形式 | 举 例                            |
| -------- | -------- | -------- | -------------------------------- |
| &        | 按位与   | a & b    | 4 & 5                            |
| \|       | 按位或   | a \| b   | 4 \| 5                           |
| ^        | 按位异或 | a ^ b    | 4 ^ 5                            |
| ~        | 按位取反 | ~a       | ~4                               |
| <<       | 按位左移 | a << b   | 4 << 2，表示整数 4 按位左移 2 位 |
| >>       | 按位右移 | a >> b   | 4 >> 2，表示整数 4 按位右移 2 位 |



## 29. 随机数

``` python
import random
 
print(random.randint(0,9))
```

> ```
> random.randint(a,b)
> 函数返回数字 N ，N 为 a 到 b 之间的数字（a <= N <= b），包含 a 和 b。
> ```



## 30. 正则

[python正则]([Python 正则表达式 | 菜鸟教程 (runoob.com)](https://www.runoob.com/python/python-reg-expressions.html))

[正则语法]([正则表达式 – 语法 | 菜鸟教程 (runoob.com)](https://www.runoob.com/regexp/regexp-syntax.html))

