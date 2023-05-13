# PyTorch与深度学习



## 1. numpy

#### 1.1 产生numpy.ndarray的几种方式

```python
import numpy as np

np.array(object, dtype=None, copy=True, order=None, subok=False, ndmin=0)
np.array([1,2,3,4,5,6])

# 初始化为0,1或空 
np.zeros(shape, dtype=float, order='C')
np.ones(shape, dtype=None, order='C')
np.empty(shape, dtype=float, order='C')	# 填充的数据无意义
# 例 
np.zeros((2,5))	# 产生一个 2 * 5 的矩阵		用() 和 [] 都行
np.ones([5,]) 	# 产生一个 包含5个元素的向量

np.arange([start,] stop[, step,], dtype=None)
# 例
np.arange(20, 36)	# [20,21,22,...,34,35]
np.linspace(start, stop, num)	# linspace和arange的不同在于linspace的第三个参数是生成的数字总数

# 随机生成
a = np.random.rand(d0, d1, ..., dn)  # 输入的就是第一维，第二维，第三维....
a = np.random.rand(3,5)		# 其中的数据是[0 , 1) 中随机
# [[0.44108417 0.13205678 0.19765883 0.9100034  0.37837717]
# [0.50288821 0.54574884 0.60314067 0.85237972 0.79760849]
# [0.60066761 0.42078304 0.57672113 0.67113562 0.44895375]]
a.fill(2)		# 将a全部填充为2，fill没有返回值
```

#### 1.2 ndarray的形状

```python
w = np.ones([3,5])
print(w.shape)	# （3，5）
print(w.shape[0]) # = len(shape)  	 3 
print(w.shape[1]) # = len(shape[0])  5 
# ndarray的shape属性是一个元组，(d1,d2,d3)，依次保存着第一维，第二维，...

# 修改形状 reshape() 可以
w = np.arange(15).reshape(3, 5) # reshape((3, 5)) 和 reshape([3,5]) 也可以
print(w.reshape(-1))	# reshape(-1)总是把矩阵转换为一个向量，[ 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14]

```

#### 1.3 ndarray的截取

```python
# ndarray的截取和list一样
w = np.arange(15).reshape(3, 5)
# [[ 0  1  2  3  4]
# [ 5  6  7  8  9]
# [10 11 12 13 14]]
w = w[1:, 2:4]
# [[ 7  8]
# [12 13]]
# 当然也可以取一行
w = w[:,]
# [ 0  5 10]	这里进行了行列转换
```

#### 1.4 ndarray的连接

```python
# 对比list的连接 append, extend 见之前的文档
w = np.arange(10).reshape([2, 5])
b = np.ones((3,2))
w = np.append(w, b)	# append的两个参数矩阵不需要行列约束，都是全部拆成一维向量再把所有元素加在最后，再返回这个一维向量
# 当然append 有参数axis可以增加行或列，但太麻烦
print(w)
# [0. 1. 2. 3. 4. 5. 6. 7. 8. 9. 1. 1. 1. 1. 1. 1.]

# numpy提供了numpy.concatenate((a1,a2,...), axis=0)函数。能够一次完成多个数组的拼接。其中a1,a2,...是数组类型的参数
a=np.array([[1,2,3],[4,5,6]])
b=np.array([[11,21,31],[7,8,9]])
c = np.concatenate((a,b),axis=0)	# axis=0 以行为基准
# array([[ 1,  2,  3],
#        [ 4,  5,  6],
#        [11, 21, 31],
#        [ 7,  8,  9]])
d = np.concatenate((a,b),axis=1)  # axis=1 以列为基准
# array([[ 1,  2,  3, 11, 21, 31],
#        [ 4,  5,  6,  7,  8,  9]])
```

#### 1.5 ndarray的计算

```python
# --------四则运算--------
# 即使不同型矩阵仍然可以四则运算
w = np.arange(15).reshape((3, 5))
print(w)
# w = [[ 0  1  2  3  4]
#  	  [ 5  6  7  8  9]
# 	  [10 11 12 13 14]]
a = np.ones([5,])  # b = [1,1,1,1,1] / [5,] 不指定列数时就认为这是行向量，一维矩阵
print(w - a)		#  w 的每一行都去减 b 这个行向量  w + b， w * b , w / b 同理 
#	[[-1.  0.  1.  2.  3.]
#	 [ 4.  5.  6.  7.  8.]
# 	[ 9. 10. 11. 12. 13.]]
# 综上，i*j 型的矩阵可以与同型矩阵进行四则运算，也可以和长度为j的向量进行运算，运算规则是每行都与此向量进行运算

# ----------点乘-----------
# 矩阵点乘
w = np.arange(6).reshape([2, 3])
print(w)
# [[0 1 2]
# [3 4 5]]
b = np.ones((3,2))
b.fill(2)
print(np.dot(w, b))
# [[ 6.  6.]
#  [24. 24.]]

# 矩阵和向量点乘
w = np.arange(12).reshape([2, 6])
# w [[ 0  1  2  3  4  5]
#    [ 6  7  8  9 10 11]]
b = np.ones((3,2))
b.fill(2)
b = b.reshape(-1)
# b [2. 2. 2. 2. 2. 2.]
print(np.dot(w, b))
# [ 30. 102.]
# 只有向量和矩阵的行向量等长时才能点乘，运算规则就是矩阵每个行向量和向量做内积
```

#### 1.6 矩阵的相关操作

```python
w = np.arange(6).reshape(2, 3)
print(w.T , w.transpose())		# 这两个都是求矩阵的转置
# [[0 3]
#  [1 4]
#  [2 5]]
```



 ## 2. PyTorch

<div style="padding:15px 20px;border-radius:10px;background:#252525;color:#dedede;font-size:23px;box-shadow:10px 10px 15px #888888">PyTorch是一个基于Python的库，用来提供一个具有灵活性的深度学习平台，PyTorch的工作流程非常接近Python的科学计算库--numpy，PyTorch的一些优点包括：<span style="color:#ee0000">运行时构建计算图，多gpu支持，自定义数据加载器，简化的预处理器</span></div>

#### 2.1 tensor

- `Tensor `，又名张量，可以简单地认为它是一个数组，且支持高效的科学计算。它可以是一个数（标量），一维数组（向量），二维数组（矩阵）或更高维的数组（高阶数据）。tensor和`numpy`的`ndarrays`类似。

	- **有多种方式生成一个tensor**

	- ``` python
		import torch
		
		x = torch.FloatTensor([2, 3])  # torch.tensor类似
		# tensor([2., 3.])
		y = torch.arange(0,4)
		# tensor([0, 1, 2, 3])
		a = torch.zeros(6).reshape(2,3)
		b = torch.ones_like(a)	# 让b 跟a 有一样的尺寸
		# tensor([[0., 0., 0.],
		#         [0., 0., 0.]])
		b = torch.randn(2,2)	# 在标准正态分布中取值
		b = torch.rand(2,2)	# 在(0, 1]的均匀分布中取值
		# tensor([[0.9402, 0.9672],
		#         [0.5175, 0.6964]])
		
		c = torch.arange(3).float() 	# 再后面直接点就可以修改数据类型
		print(c.data.type())
		```

	- **tensor 的运算**

	- ``` python
		a = torch.zeros(6).reshape(2,3)
		c = torch.tensor([1,2,3], dtype=float)
		print(c+a)	#+-*/ 都是两个张量对应位置相互运算
		print(torch.add(c,a))
		print(c.add_(a))			# inplace  原地操作
		# tensor([[1., 2., 3.],
		#         [1., 2., 3.]], dtype=torch.float64)
		# 还是和numpy一样，即使运算的两个张量不同型，也会通过广播机制将两者运算
		
		c = torch.tensor([1,2,3])
		d = torch.tensor([2,3,2])
		print(torch.dot(c,d))					# dot() 只能用于一维张量！！
		# tensor(14)
		
		c = torch.tensor([1,2,3,4,5,6]).reshape(2,3)
		d = torch.tensor([2,3,2,3,2,3]).reshape(3,2)
		print(torch.matmul(c,d))				# matmul）才能用于多维张量
		# tensor([[12, 18],
		#         [30, 45]]
		c = torch.tensor([1,2,3,4,5,6]).reshape(2,3)
		d = torch.tensor([2,3,1])
		print(torch.matmul(c,d))
		# tensor([11, 29])   n列的二维张量和长度为n的一维张量可以直接叉乘，不用reshape
		c = torch.tensor([1,2,3,4,5,6]).reshape(2,3)
		d = torch.tensor([2,3])
		print(torch.matmul(d,c))
		# tensor([14, 19, 24])
		
		# 矩阵相乘还有 torch.mm 与 torch.bmm
		# 前者实现矩阵相乘，后者实现批量矩阵相乘。
		# 还有einsum，更强大，后面再说
		```
	
	- 其他操作
	
	 ```python
	# -----索引--------
	# 注意：索引出来的结果和原数据共享内存，改一个，另一个也改了
	a = [i for i in range(1,10) if i%2 == 0]
	a = torch.Tensor(a).reshape(2,2)
	b = a[0,:]		# 如果不希望共享数据，可以 b = a[0,:].clone()
	print(b)	# tensor([2., 4.])
	b[0] = 22
	print(a)
	# tensor([[22.,  4.],
	#         [ 6.,  8.]])
	
	# ------------------------改变形状----------------------- 
	# view 或 reshape
	# 还是共享数据的，还是通过clone来解决
	a = [i for i in range(1,10) if i%2 == 0]
	a = torch.Tensor(a)
	b = a.view(2,2) # reshape(2,-1) 也行
	print(b)
	# tensor([[2., 4.],
	#        [6., 8.]])
	b += 1
	print(a)
	# tensor([3., 5., 7., 9.])
	print(b.t())	# 求转置
	print(b.inverse())	# 求逆
	print(b.diag())	# 返回对角线的元素
	print(b.trace())	# 求迹
	
	#------------------------------item()----------------------
	#只有1个数字的tensor可通过 .item() 转换为python的数据类型
	
	#-----------------------cat()-----------------------
	#连接多个张量
	a = torch.cat((torch.ones(2000), torch.zeros(2000)), dim=0)
	a.shape # 4000,  如果是dim=1 就是 2000，2
	 ```




  - Tensor 和 NumPy 相互转换


```python
# 使用numpy() 和 from_numpy() 来相互转换
# 但这两个函数产生的结果还是共享内存的
a = torch.ones(3)
b = a.numpy()
print(a, b)	# tensor([1., 1., 1.]) [1. 1. 1.]
a += 1
print(a, b)	# tensor([2., 2., 2.]) [2. 2. 2.]

c = torch.tensor(b)	# 通过tensor等类似的函数不会共享内存
c += 1
print(b, c)	# [2. 2. 2.] tensor([3., 3., 3.])
```

- [pytorch.gather 详解]([pytorch函数（1）：torch.gather()的正确理解方法_今天不用睡觉了、的博客-CSDN博客](https://blog.csdn.net/weixin_42200930/article/details/108995776?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_title~default-0.control&spm=1001.2101.3001.4242))

- [plt.scatter() 绘制散点图详解]([plt.scatter()_coder-CSDN博客_plt.scatter](https://blog.csdn.net/m0_37393514/article/details/81298503))

#### 2.2 autograd

- torch.autograd 是为了方便用户使用，专门开发的一套**自动求导引擎**，他能够根据**输入和前向传播过程自动构建计算图**，并执行**反向传播**

- Variable
	- PyTorch 在autograd 模块中实现了计算图的相关功能，autograd中的核心数据结构是Variable，Variable 封装了tensor，并*记录对tensor的操作记录用来构建计算图*。Variable的数据结构中包含三个属性：
	- 1. **data**：保存variable 所包含的tensor
		2. **grad**：保存data对应的梯度，grad也是**variable**，而非tensor
		3. **grad_fn**：指向一个Function，记录variable的历史操作
			<img src="img\自动求梯度.png" alt="自动求梯度" style="zoom:85%;" />

``` python
x = Variable(torch.ones(2,2), requires_grad=False)	
temp = Variable(torch.zeros(2,2), requires_grad=True)
y = torch.mean(x + temp + 2)
y.backward()
print(x.grad)       # 因为requires_grad=False 所以输出None
print(temp.grad)
# None
# tensor([[0.2500, 0.2500],
#         [0.2500, 0.2500]])


x = torch.ones(2,2,requires_grad=True)#默认值是False
#x.requires_grad_(False) 可改变这个参数的值
print(x)
print(x.grad_fn)	#输入None, 因为x是直接创建的，没有graf_fn

y = x + 2
print(y)
#tensor([[3., 3.],
#        [3., 3.]], grad_fn=<AddBackward0>)
print(y.grad_fn)	#<AddBackward0 object at 0x000001A9B2661448> 一个加法操作

#---------------自动求梯度---------------
x = torch.ones(2,2,requires_grad=True)
y = x + 2
z = y * y * 3
out = z.mean()
out.backward()
print(x.grad)
#tensor([[4.5000, 4.5000],
#        [4.5000, 4.5000]])
#再来方向传播一次，注意：grad是累加的
out2 = x.sum()
out2.backward()
print(x.grad)
#tensor([[5.5000, 5.5000],
#        [5.5000, 5.5000]])
out3 = x.sum()
x.grad.data.zero_()	#梯度清零
out3.backward()
print(x.grad)
#tensor([[1., 1.],
#        [1., 1.]])
```

- 注意：在使用`y.backward()`时，如果y是**标量**，则不需要传入参数；否则，需要传入一个与y同型的Tensor。这样的原因简单来说就是为了避免向量（甚至更高维张量）对张量求导，因此转换维标量对张量求导。

	所以必要时我们要把张量通过所有张量的元素加权求和的方式转换维标量。

	``` python
	x = torch.tensor([1.0,2.0,3.0,4.0], requires_grad=True)
	y = 2 * x
	z = y.view(2,2)
	v = torch.tensor([[1.0,0.1],[0.01, 0.001]],dtype=float)
	z.backward(v)
	#这里的原理是先计算a = torch.sum(z * v)，a此时是标量，再计算a对x的导数
	print(x.grad)
	#tensor([2.0000, 0.2000, 0.0200, 0.0020])
	```

	

- 计算图

	**计算图是一种特殊的有向无环图**，用于记录算子和变量之间的关系。一般用矩阵表示算子，椭圆形表示变量。

	在pytorch中计算图特点总结如下

	- autograd 根据用户对 variable 的操作构建计算图。对variable的操作抽象为Function
	- Variable 默认是不需要求导的，即 requires_grad 属性默认为False
	- 多次方向传播时，梯度是**累加**的。方向传播的中间缓存会被清空
	- pytorch采用动态图设计，可以方便地查看中间层的输出，动态地设计计算图结构

	> 动态图采用搭建和计算同时进行，灵活，易调节。pytorch采用动态图

	> 静态图先搭建，后运算，高效但不灵活。tensorflow采用静态图

#### 2.3 nn

- nn.Module

	- torch.nn 的核心数据结构是Module，它是一个抽象的概念，既可以表示神经网路中的某一层，也可以表示一个包含很多曾的**神经网络**。最常见的做法是继承nn.Module，编写自己的网络。PyTorch实现了神经网络的绝大多数的layer，这些layer都继承了nn.Module

		> 能够自动检测到自己的parameter，并将其作为学习参数。

		> 主Module能够递归查找子Module中的parameter

``` python
import torch
import torch.nn.functional as F

loss_func = F.cross_entropy		# 使用F的交叉熵损失函数
##CrossEntropyLoss()=log_softmax() + NLLLoss()

# 一个三层卷积的例子
class ConvModule(nn.Module):
    def __init__(self):
        super(ConvModule, self).__init__()
        self.conv = nn.Sequential(
            nn.Conv2d(in_channels=3, out_channels=32, kernel_size=3),
            nn.BatchNorm2d(32),
            nn.ReLU(inplace=True),
            nn.Conv2d(32, 64, 3),
            nn.BatchNorm2d(64),
            nn.ReLU(inplace=True),
            nn.Conv2d(in_channels=64, out_channels=128, kernel_size=3),
            nn.BatchNorm2d(128),
            nn.ReLU(inplace=True),
            # 不能把全连接层放在里面，因为还需要对数据降维
        )
        self.fc = nn.Linear(128, num_classes)

    def forward(self, x):
        out = self.conv(x)
        # 通过平均池化来将图像变为1×1
        out = F.avg_pool2d(out, 26)
        out = out.squeeze()
        return self.fc(out)

```


  	

- 优化器

	PyTorch 将深度学习中常见的优化方法全部封装在**torch.optim**中，其设计非常灵活，能够很方便地扩展成自定义的优化方法。

	所有的优化方法都继承基类 **optim.Optimizer**，并实现了自己的优化步骤，最基本的优化方法是随机梯度下降法 SGD。需要重点注意的是优化时如何调整学习率。

	> SGD: 每次更新随机选择并计算一个样本进行迭代
	>
	> opt = optim.学习率优化算法(model.parameters(), lr=学习率)
	>
	> opt.step()  给参数更新一次参数
	
	```python
	import torch.optim as optim
	
	
	class Mnist_Logistic(nn.Module):
	    def __init__(self):
	        super(Mnist_Logistic, self).__init__()
	        self.lin = nn.Linear(784, 10)  # 输入784维，输出10维
	
	    def forward(self, x):
	        return self.lin(x)
	
	def accuracy(pred, label):
	    preds = torch.argmax(pred, dim=1)
	    return (preds == label).float().mean()
	
	model = Mnist_Logistic()    # 此线性模型只是简单的计算 x @ w + b
	loss_func = F.cross_entropy  # 相当于给这个函数取别名
	lr = 0.5
	epochs = 5
	opt = optim.SGD(model.parameters(), lr=lr)
	
	for epoch in range(epochs):
	    for x, y in train_iter:
	        x = x.reshape(-1, 784)
	        pred = model(x)
	        loss = loss_func(pred, y)
	        loss.backward()
	        # 使用optimizer来自动地更新参数
	        opt.step()	# 给模型更新一次参数
	        opt.zero_grad()
	
	```

	
	
	*调整学习率*
	
	- 修改optimizer.param_groups 中对应的学习率。
	- 新建优化器，由于optimizer十分轻量级，构建开销小，故可以构建新的optimizer



#### 2.4 nn.Conv2d

`torch.nn.Conv2d(in_channels, out_channels, kernel_size, stride=1, padding=0, dilation=1, groups=1, bias=True)`

<table><thead><tr><th align="left">参数</th><th align="left">参数类型</th><th align="left"></th><th align="left">描述</th></tr></thead><tbody><tr><td align="left"><code>in_channels</code></td><td align="left">int</td><td align="left">Number of channels in the input image</td><td align="left">输入图像通道数</td></tr><tr><td align="left"><code>out_channels</code></td><td align="left">int</td><td align="left">Number of channels produced by the convolution</td><td align="left">卷积产生的通道数</td></tr><tr><td align="left"><code>kernel_size</code></td><td align="left">(int or tuple)</td><td align="left">Size of the convolving kernel</td><td align="left">卷积核尺寸，可以设为1个int型数或者一个(int, int)型的元组。例如(2,3)是高2宽3卷积核</td></tr><tr><td align="left"><code>stride</code></td><td align="left">(int or tuple, optional)</td><td align="left">Stride of the convolution. Default: 1</td><td align="left">卷积步长，默认为1。可以设为1个int型数或者一个(int, int)型的元组。</td></tr><tr><td align="left"><code>padding</code></td><td align="left">(int or tuple, optional)</td><td align="left">Zero-padding added to both sides of the input. Default: 0</td><td align="left">填充操作，控制<code>padding_mode</code>的数目。</td></tr><tr><td align="left"><code>padding_mode</code></td><td align="left">(string, optional)</td><td align="left">‘zeros’, ‘reflect’, ‘replicate’ or ‘circular’. Default: ‘zeros’</td><td align="left"><code>padding</code>模式，默认为Zero-padding 。</td></tr><tr><td align="left"><code>dilation</code></td><td align="left">(int or tuple, optional)</td><td align="left">Spacing between kernel elements. Default: 1</td><td align="left">扩张操作：控制kernel点（卷积核点）的间距，默认值:1。</td></tr><tr><td align="left"><code>groups</code></td><td align="left">(int, optional)</td><td align="left">Number of blocked connections from input channels to output channels. Default: 1</td><td align="left">group参数的作用是控制分组卷积，默认不分组，为1组。</td></tr><tr><td align="left"><code>bias</code></td><td align="left">(bool, optional)</td><td align="left">If True, adds a learnable bias to the output. Default: True</td><td align="left">为真，则在输出中添加一个可学习的偏差。默认：True。</td></tr></tbody></table>






#### 2.5 使用GPU加速

以下数据结构分为CPU和GPU两个版本

- Tensor
- Variable （包括parameter）
- nn.Module (包括常用的layer，loss function等)

``` python
use_gpu = torch.cuda.is_available()	# 检查是否可使用gpu

##先指定一个设备
device = torch.device('cuda:0')  # 指定device为0号GPU，如使用CPU填'cpu'

net = ConvModule().to(device)	# 将模型放在指定的gpu上去
data = data.to(device).float()	# 数据这样放进gpu
target = target.to(device).long()
```





#### 2.6 保存和加载

有时需要将Tensor, Variable 等对象保存到硬盘，以便之后再用。

- 这些信息都是保存为Tensor。
- 使用torch.save(obj, filename)函数和torch.load(filename)即可完成保存和加载
- 在save/load 时可指定使用的pickle模块，在load时还可以将GPU tensor映射到CPU或其他GPU

``` python
torch.save(data_tensor, 'data_tensor.pt')

data_tensor_load = torch.load('data_tensor.pt')
```





## 3. torchvision

#### 3.1 torchvision.transforms

transforms 包含了常用的图像预处理方法。

**transforms.Compose()**：将一系列的transforms有序组合，实现时按照这些方法依次对图像操作。

``` python
from torchvision import transforms

train_transform = transforms.Compose([
    transforms.Resize((32, 32)),  # 缩放
    transforms.RandomCrop(32, padding=4),  # 随机裁剪
    transforms.ToTensor(),  # 图片转张量，同时归一化0-255 ---》 0-1
    transforms.Normalize(norm_mean, norm_std),  # 标准化均值为0标准差为1
])

#使用
# 构建MyDataset实例
train_data = RMBDataset(data_dir=train_dir, transform=train_transform)
# 构建DataLoder
train_loader = DataLoader(dataset=train_data, batch_size=BATCH_SIZE, shuffle=True)
```

#### 3.2 torchvision.datasets

**torchvision.datasets** : 设计上都是继承torch.utils.data.Dataset，常用数据集的dataset实现，MNIST，CIFAR-10，ImageNet等

``` python
import torch
import torchvision
import torchvision.transforms as transforms

# 函数原型
torchvision.datasets.MNIST(root, train=True, transform=None, target_transform=None, downloan=False)
# root（string）–存在MNIST/processed/training.pt和MNIST/processed/test.pt的数据集的根目录。
# train（bool，可选）–如果为True，则从training.pt创建数据集，否则从test.pt创建数据集。
# transform（callable，可选）–接受PIL图像并返回已转换的版本。
# target_transform（callable，可选）–接受目标并对其进行转换的函数/转换。
# downloan（bool，可选）–如果为true，则从internet下载数据集并将其放在根目录中。如果数据集已下载，则不会再次下载

#*******************举例说明***************************
mnist_train = torchvision.datasets.FashionMNIST(root='', train=True, download=False, transform=transforms.ToTensor())
mnist_test = torchvision.datasets.FashionMNIST(root='', train=False, download=False, transform=transforms.ToTensor())

batch_size = 256
train_iter = torch.utils.data.DataLoader(mnist_train, batch_size, shuffle=True, num_workers=0)
test_iter = torch.utils.data.DataLoader(mnist_test, batch_size, shuffle=False, num_workers=0)

for x, y in train_iter:
    #print(x.shape, y.shape) # torch.Size([256, 1, 28, 28]) torch.Size([256])
    x = x.squeeze(1).reshape(batch_size, -1)    #将第1维压缩掉（该维度的值的个数为1才可以以压缩），再把一个图片的数据reshape到一行
    print(x.shape)	#torch.Size([256, 784])

    
# 获得cifar10的数据集
cifar_train = torchvision.datasets.CIFAR10(root=path, train=True, download=True, transform=transform)
cifar_test = torchvision.datasets.CIFAR10(root=path, train=False, download=True, transform=transform)
batch_size = 512
cifar_trainloader = torch.utils.data.Dataloadr(cifar_train, batch_size=batch_size, shuffle=True, num_workers=0)
cifar_testloader = torch.utils.data.Dataloadr(cifar_test, batch_size=batch_size, shuffle=False, num_workers=0)
```

还可以用来读取本地的数据

``` python
from torchvision import datasets

transform = transforms.Compose([
    transforms.Resize((200, 200)),
    transforms.ToTensor()
])
dataset = datasets.ImageFolder('../../../dataset/去雾数据集', transform=transform)
# 这里的路径只需要指明到数据文件夹的位置，它会自动读取路径下的所有文件夹其中的图片。
print(dataset)
# 但dataset一般会把所有的所以读进去，训练数据和测试数据全部读到一起了，这是用Subset来拆分
test_set = torch.utils.data.Subset(dataset, range(520))
train_set = torch.utils.data.Subset(dataset, range(520, 1040))
```

`ImageFolder`详解

``` python
dataset=torchvision.datasets.ImageFolder(
                       root, transform=None, 
                       target_transform=None, 
                       loader=<function default_loader>, 
                       is_valid_file=None)

# 这个ImageFloder总是把数据和标签放在一起，很麻烦，于是写一个它的子类
from typing import Optional, Callable, Any
from torchvision.datasets import ImageFolder
from torchvision.datasets.folder import default_loader, IMG_EXTENSIONS


class CustomImageFloder(datasets.ImageFolder):
    def __init__(
            self,
            root: str,
            transform: Optional[Callable] = None,
            target_transform: Optional[Callable] = None,
            loader: Callable[[str], Any] = default_loader,
            is_valid_file: Optional[Callable[[str], bool]] = None,
    ):
        super(ImageFolder, self).__init__(root, loader, IMG_EXTENSIONS if is_valid_file is None else None,
                                          transform=transform,
                                          target_transform=target_transform,
                                          is_valid_file=is_valid_file)
        self.imgs = self.samples

    def __getitem__(self, idx):
        # 直接用ImageFloder的话标签去掉太麻烦了
        # 于是只返回数据就行，不用返回类别了
        return super().__getitem__(idx)[0]
```

参数详解：

root：图片存储的根目录，即各类别文件夹所在目录的上一级目录。
transform：对图片进行预处理的操作（函数），原始图片作为输入，返回一个转换后的图片。
target_transform：对图片类别进行预处理的操作，输入为 target，输出对其的转换。如果不传该参数，即对 target 不做任何转换，返回的顺序索引 0,1, 2…
loader：表示数据集加载方式，通常默认加载方式即可。
is_valid_file：获取图像文件的路径并检查该文件是否为有效文件的函数(用于检查损坏文件)

**我们得到的`dataset`,它的结构就是[(img_data,class_id),(img_data,class_id),…]**

`class_id`是读取的第几个文件夹的标识

返回的dataset都有以下三种属性：

- `self.classes`：用一个 list 保存类别名称
- `self.class_to_idx`：类别对应的索引，与不做任何转换返回的 target 对应
- `self.imgs`：保存(img-path, class) tuple的 list



自定义的数据集装进loader

```python
from torch.utils.data import TensorDataset
from torch.utils.data import DataLoader

train_ds = TensorDataset(x_train, y_train)
test_ds = TensorDataset(x_test, y_test)
train_loader = DataLoader(train_ds, batch_size=bs, shuffle=True)    # shuffle=True 每次迭代都会打乱
test_loader = DataLoader(test_ds, batch_size=bs)
```

下面介绍以numpy格式读取mnist或fashion-mnist

``` python
import os
import gzip
import numpy as np
def load_mnist(path, kind='train'):
    """Load MNIST data from `path`"""
    labels_path = os.path.join(path,
                               '%s-labels-idx1-ubyte.gz'
                               % kind)
    images_path = os.path.join(path,
                               '%s-images-idx3-ubyte.gz'
                               % kind)
    with gzip.open(labels_path, 'rb') as lbpath:
        labels = np.frombuffer(lbpath.read(), dtype=np.uint8,
                               offset=8)
    with gzip.open(images_path, 'rb') as imgpath:
        images = np.frombuffer(imgpath.read(), dtype=np.uint8,
                               offset=16).reshape(len(labels), 784)
    return images, labels
#这样的读取的images.shape为(60000,782), labels.shape为(60000,)
#数据为0~256之间的灰度值
```

#### 3.3 torchvision.model

**torchvision.model **: 常用的模型预训练，AlexNet，VGG， ResNet，GoogLeNet等



## 4. Tensor Attributes

#### 4.1 torch.dtype

**torch.dtype** 表示`torch.Tensor`的数据类型，`PyTorch`有八种不同的数据类型

|        Data type         |             dtype             | Tensor types         |
| :----------------------: | :---------------------------: | :------------------- |
|  32-bit floating point   | torch.float32 or torch.float  | torch.*.FloatTensor  |
|  64-bit floating point   | torch.float64 or torch.double | torch.*.DoubleTensor |
|  16-bit floating point   |  torch.float16 or torch.half  | torch.*.HalfTensor   |
| 8-bit integer (unsigned) |          torch.uint8          | torch.*.ByteTensor   |
|  8-bit integer (signed)  |          torch.int8           | torch.*.CharTensor   |
| 16-bit integer (signed)  |  torch.int16 or torch.short   | torch.*.ShortTensor  |
| 32-bit integer (signed)  |   torch.int32 or torch.int    | torch.*.IntTensor    |
| 64-bit integer (signed)  |   torch.int64 or torch.long   | torch.*.LongTensor   |

``` python
x = torch.Tensor([[1, 2, 3, 4, 5], [6, 7, 8, 9, 10]])
print x.type()
#torch.FloatTensor
```



#### 4.2 torch.device

**torch.device** 代表`torch.Tensor` 分配到的设备的对象

`torch.device`包含一个设备类型（`'cpu'`或`'cuda'`设备类型）和可选的设备的序号。如果设备序号不存在，则为当前设备; 例如，`torch.Tensor`用设备构建`'cuda'`的结果等同于`'cuda:X'`,其中`X`是`torch.cuda.current_device()`的结果。

``` python
print(torch.cuda.current_device())
# 使用cpu运行这句就会报错	AssertionError: Torch not compiled with CUDA enabled
# 使用gpu就会返回设备号0,1等等
```

构造`torch.device`可以通过字符串/字符串和设备编号

``` python
device = torch.device('cuda:0')
# device(type='cuda', index=0)

torch.device('cpu')
# device(type='cpu')

torch.device('cuda')  # current cuda device
# device(type='cuda')

torch.device('cuda', 0)
# device(type='cuda', index=0)

torch.device('cpu', 0)
# device(type='cpu', index=0)
```

构造`tensor`时指定设备

``` python
torch.randn((2,3), device=torch.device('cuda:0'))
torch.randn((2,3), device='cuda:0')
torch.randn((2,3), device=1)

# 将数据放到指定的设备上
data = data.to(device).float()
target = target.to(device).long()

# 将模型放到指定的设备上
net = ConvModule().to(device)
```

#### 4.3 torch.layout

**torch.layout** 表示`torch.Tensor`内存布局的对象。



## 5. PIL 

`PIL` 中有两个进行数据增强相关模块：

- `ImageEnhance`：增强（或减弱）图像的亮度、对比度、色度和锐度
- `ImageFilter`：对图像添加一些滤镜效果

#### 5.1 ImageEnhance

`ImageEnhance` 是 PIL 中的图像增强模块，通过 `	` 导入该模块，其主要包含四个功能：

1. `ImageEnhance.Brightness`：调整图像的亮度

```python
from PIL import Image
from PIL import ImageEnhance

img = Image.open(img_path)
enh_bri = ImageEnhance.Brightness(img)
new_img = enh_bri.enhance(factor=1.5)
Image._show(enh_bri)
```

`factor` 参数控制图像的明亮程度，范围为0到无穷。`factor` 为 0 时生成纯黑图像，为 1 时还是原始图像，该值小于1为亮度减弱，大于1为亮度增强，值越大，图像越亮。

2. `ImageEnhance.Color`：调整图像的色彩平衡

用法如下：

```python
img = Image.open(img_path)
enh_col = ImageEnhance.Color(img)
new_img = enh_col.enhance(factor=1.5)
```

`factor` 参数控制图像的色彩平衡，即色彩的鲜艳程度，范围为0到无穷 。`factor` 为 0 时生成灰度图像，为 1 时还是原始图像，该值越大，图像颜色越饱和。

3. `ImageEnhance.Contrast`：调整图像的对比度

用法如下：

```python
img = Image.open(img_path)
enh_con = ImageEnhance.Contrast(img)
new_img = enh_con.enhance(factor=1.5)
```

`factor` 参数控制图像的对比度，范围为0到无穷。`factor` 为 0 时生成全灰图像，为 1 时还是原始图像，该值越大，图像颜色对比越明显。

4. `ImageEnhance.Sharpness`：调整图像的锐化程度

用法如下：

```python
img = Image.open(img_path)
enh_sha = ImageEnhance.Sharpness(img)
new_img = enh_sha.enhance(factor=1.5)
```

`factor` 参数控制图像的锐度，范围为 ![[公式]](https://www.zhihu.com/equation?tex=%280%2C+2%29) 。`factor` 为 0 时生成模糊的图像，为 1 时还是原始图像，为 2 时生成完全锐化的图像。

#### 5.2 ImageFilter

`ImageFilter` 是 PIL 中的图像滤镜模块，通过 `from PIL import ImageFilter` 导入该模块。

ImageFilter 模块提供了很多滤镜操作（10种），包含模糊、平滑、锐化、边界增强等滤镜效果的处理。这些滤镜主要通过 Image 类的`filter()`方法来对图片进行操作。

| filter name                   |        含义        |
| :---------------------------- | :----------------: |
| ImageFilter.BLUR              |        模糊        |
| ImageFilter.CONTOUR           |      轮廓滤镜      |
| ImageFilter.DETAIL            |      细节增强      |
| ImageFilter.EDGE_ENHANCE      |      边缘增强      |
| ImageFilter.EDGE_ENHANCE_MORE | 边缘增强(阈值更大) |
| ImageFilter.EMBOSS            |      浮雕滤镜      |
| ImageFilter.FIND_EDGES        |      边缘滤镜      |
| ImageFilter.SMOOTH            |      平滑滤镜      |
| ImageFilter.SMOOTH_MORE       | 平滑滤镜(阈值更大) |
| ImageFilter.SHARPEN           |        锐化        |

这个用法很简单，比如下面实现轮廓滤镜和浮雕滤镜：

```python
from PIL import Image
from PIL.ImageFilter import EMBOSS, CONTOUR

img = Image.open(img_path)
img_1 = img.filter(EMBOSS)
img_2 = img.filter(CONTOUR)
```



#### 5.3 image和tensor相互转化

```python
import torch
from torchvision import datasets
from torchvision.transforms import transforms
from PIL import Image
from PIL import ImageEnhance

transform = transforms.Compose([
    transforms.Resize((200, 200)),
    transforms.ToTensor(),
    transforms.Normalize((0.5, 0.5, 0.5), (0.5, 0.5, 0.5))
])

img = Image.open(path)
enh_sha = ImageEnhance.Sharpness(img)
new_img = enh_sha.enhance(factor=1.5)
########## image to  tensor
img_tensor = transform(new_img)
########## tensor  to image
img = transforms.ToPILImage()(img_tensor)
```



## 6. pyplot

python 中的画图库

使用示例：

``` python
import matplotlib.pyplot as plt


x = [2, 3, 4, 5]
y = [i**2 for i in range(1, 5)]
z = [2, 4, 5]
plt.plot(x, '*')	# 给plot传入数组
plt.plot(y, 'o')	# 定义y 的形状，还有's', '*', '-','--'等
plt.plot(z, 'r')	# 定义z 的颜色，还有'b', 'g'等 颜色和形状可以组合使用
plt.xlabel('xxx')
plt.ylabel('yyy')
plt.title('xyz')
# plt.plot(x, y)则表示以输入的点的坐标为(x, y)了
plt.legend(['x', 'y', 'z'])		# 标识线段的名称
plt.axis([0, 10, 0, 20])
plt.show()
```

点击了解 [pyplot的更多用法]([Python可视化库Matplotlib绘图入门详解_pyplot (sohu.com)](https://www.sohu.com/a/343708772_120104204))



## 7. upsample

一个卷积层举例：

``` python
self.conv = nn.Sequential(
            # 定义单层卷积的输入输出通道数以及卷积核的大小, 得到图片  16 * 196 * 196
            nn.Conv2d(in_channels=3, out_channels=16, kernel_size=(5, 5)),
            # 批归一化
            nn.BatchNorm2d(16),
            # 激活函数
            nn.ReLU(inplace=True),
            # 得到图片  16 * 98 * 98
            nn.MaxPool2d(kernel_size=(2, 2)),    # stride 默认是kernelsize的大小
            # 再来一层 , 得到 64 * 96 * 96
            nn.Conv2d(in_channels=16, out_channels=64, kernel_size=(3, 3)),
            nn.BatchNorm2d(64),
            nn.ReLU(inplace=True),

            # 逆卷积用来升采样， 得到图片 16 * 98 * 98，
            nn.ConvTranspose2d(in_channels=64, out_channels=16, kernel_size=(3, 3), bias=False),
    		# 逆卷积按照卷积的计算过程反过来推导即可
            nn.BatchNorm2d(16),
            nn.ReLU(inplace=True),
            # 升采样 得到 16 * 196 * 196， 采用双线性插值算法
            nn.Upsample(scale_factor=2, mode='bilinear', align_corners=True),
            # 逆卷积用来升采样， 得到图片 3 * 200 * 200
            nn.ConvTranspose2d(in_channels=16, out_channels=3, kernel_size=(5, 5), bias=False),
            nn.BatchNorm2d(3),
            nn.ReLU(inplace=True),
        )
```

其中：

> ```python
> CLASS torch.nn.Upsample(size=None, scale_factor=None, mode='nearest', align_corners=None)
> ```

**参数：**

- **size** ([*int*](https://docs.python.org/3/library/functions.html#int) *or* *Tuple [*[*int*](https://docs.python.org/3/library/functions.html#int)*] or* *Tuple [*[*int*](https://docs.python.org/3/library/functions.html#int)*,* [*int*](https://docs.python.org/3/library/functions.html#int)*] or* *Tuple [*[*int*](https://docs.python.org/3/library/functions.html#int)*,* [*int*](https://docs.python.org/3/library/functions.html#int)*,* [*int*](https://docs.python.org/3/library/functions.html#int)*] ,* *optional*) – 根据不同的输入类型制定的输出大小
- **scale_factor** ([*float*](https://docs.python.org/3/library/functions.html#float) *or* *Tuple [*[*float*](https://docs.python.org/3/library/functions.html#float)*] or* *Tuple [*[*float*](https://docs.python.org/3/library/functions.html#float)*,* [*float*](https://docs.python.org/3/library/functions.html#float)*] or* *Tuple [*[*float*](https://docs.python.org/3/library/functions.html#float)*,* [*float*](https://docs.python.org/3/library/functions.html#float)*,* [*float*](https://docs.python.org/3/library/functions.html#float)*] ,* *optional*) – 指定输出为输入的多少倍数。如果输入为tuple，其也要制定为tuple类型
- **mode** ([*str*](https://docs.python.org/3/library/stdtypes.html#str)*,* *optional*) – 可使用的上采样算法，有`'nearest'`, `'linear'`, `'bilinear'`, `'bicubic'` and `'trilinear'`. 默认使用`'nearest'`
- **align_corners** ([*bool*](https://docs.python.org/3/library/functions.html#bool)*,* *optional*) – 如果为True，输入的角像素将与输出张量对齐，因此将保存下来这些像素的值。仅当使用的算法为`'linear'`, `'bilinear'`or `'trilinear'`时可以使用。默认设置为`False`

<hr/>

举例：

```python
import torch
from torch import nn
input = torch.arange(1, 5, dtype=torch.float32).view(1, 1, 2, 2)
input
```

返回：

```python
tensor([[[[1., 2.],
          [3., 4.]]]])
```

 

```python
m = nn.Upsample(scale_factor=2, mode='nearest')
m(input)
```

返回：

```python
tensor([[[[1., 1., 2., 2.],
          [1., 1., 2., 2.],
          [3., 3., 4., 4.],
          [3., 3., 4., 4.]]]])
```

 

```python
m = nn.Upsample(scale_factor=2, mode='bilinear',align_corners=False)	# 采用双线性插值算法
m(input)
```

返回：

```python
tensor([[[[1.0000, 1.2500, 1.7500, 2.0000],
          [1.5000, 1.7500, 2.2500, 2.5000],
          [2.5000, 2.7500, 3.2500, 3.5000],
          [3.0000, 3.2500, 3.7500, 4.0000]]]])
```

 

```python
m = nn.Upsample(scale_factor=2, mode='bilinear',align_corners=True)
m(input)
```

返回：

```python
tensor([[[[1.0000, 1.3333, 1.6667, 2.0000],
          [1.6667, 2.0000, 2.3333, 2.6667],
          [2.3333, 2.6667, 3.0000, 3.3333],
          [3.0000, 3.3333, 3.6667, 4.0000]]]])
```

 

```python
m = nn.Upsample(size=(3,5), mode='bilinear',align_corners=True)
m(input)
```

返回：

```python
tensor([[[[1.0000, 1.2500, 1.5000, 1.7500, 2.0000],
          [2.0000, 2.2500, 2.5000, 2.7500, 3.0000],
          [3.0000, 3.2500, 3.5000, 3.7500, 4.0000]]]])
```

 

专门用于2D数据的线性插值算法: `UpsamplingNearest2d`，参数等跟上面的差不多，省略

```
CLASS torch.nn.UpsamplingNearest2d(size=None, scale_factor=None)
```



专门用于2D数据的双线性插值算法: `UpsamplingBilinear2d`，参数等跟上面的差不多，省略

```
CLASS torch.nn.UpsamplingBilinear2d(size=None, scale_factor=None)
```



## 8. pandas

#### 8.1 Series

- `pandas` 中一维数据结构称为 `Series`，它类似于一维数组的对象，它由一组数据（各种Numpy数据类型）以及一组与之相关的数据标签（即索引）组成。（索引可以是任何类型），Series 可以直接用 [index] 来访问

- ```python
	pandas.Series( data, index, dtype, name, copy)
	```

- - **data**：一组数据(ndarray 类型)。
	- **index**：数据索引标签，如果不指定，默认从 0 开始。
	- **dtype**：数据类型，默认会自己判断。
	- **name**：设置名称。
	- **copy**：拷贝数据，默认为 False。

- ```python
	a = [1, 2, 3, 4]
	var = pd.Series(a, index=[1, 'index2', False, False])
	print(var, '\n********')
	print(var[False])	
	#输出###########################
	'''
	1         1
	index2    2
	False     3
	False     4
	dtype: int64 
	********
	False    3
	False    4
	dtype: int64
	'''
	##########################
	
	#也可以用字典生成
	sites = {1: "Google", 2: "Runoob", 3: "Wiki"}
	myvar = pd.Series(sites)
	# myvar = pd.Series(sites, index = [1, 2]) 就表示只取了前两个
	print(myvar)
	
	```



#### 8.2 DataFrame

- `pandas` 中二维数据结构称为 `DataFrame`，它类似于一维数组的对象，它是一个表格型的数据结构，它含有一组有序的列，每列可以是不同的值类型（数值、字符串、布尔型值）。<code>DataFrame</code> 既有行索引也有列索引，它可以被看做由 Series 组成的字典（共同用一个索引）。DataFrame 不能直接用 [ ] 访问行。但可以用 [ 列名 ] 来访问列

![DataFrame结构](https://www.runoob.com/wp-content/uploads/2021/04/pandas-DataStructure.png)

- ```python
	pandas.DataFrame( data, index, columns, dtype, copy)
	```

	参数说明：

	- **data**：一组数据(ndarray、series, map, lists, dict 等类型)。
	- **index**：索引值，或者可以称为行标签。
	- **columns**：列标签，默认为 RangeIndex (0, 1, 2, …, n) 。
	- **dtype**：数据类型。
	- **copy**：拷贝数据，默认为 False。

data[] , []里只接收header， 即只接收列

data.loc[]	, 这里才接收行

``` python
data = [['google', 10], ['runoob', 12], [13, 'iop']]
df = pd.DataFrame(data)
print(df)
'''
        0    1
0  google   10
1  runoob   12
2      13  iop
'''

data = [['google', 10], ['runoob', 12], [13, 'iop']]
df = pd.DataFrame(data, index=[2, 'str', 2], columns=['column1', 2.2])
print(df)
'''
	column1  2.2
2    google   10
str  runoob   12
2        13  iop
'''

dt = [{'a':1, 'b':2}, {'c':3, 'b':4, 'd':5}]
df = pd.DataFrame(dt)	# 同样可以指定column来指定取哪些列
print(df)
'''
     a  b    c    d
0  1.0  2  NaN  NaN
1  NaN  4  3.0  5.0
'''
```

df 的常用方法：

- `tail( n )` 方法用于读取尾部的 n 行，如果不填参数 n ，默认返回 5 行，空行各个字段的值返回 **NaN**。

- `head( n )` 方法用于读取前面的 n 行，如果不填参数 n ，默认返回 5 行。

- `info() `方法返回表格的一些基本信息。

***DateFame转为numpy的ndarray以及list***

``` python
dt = [
    [33,44,55],
    [66,77,99],
    [888,665,5554],
    [np.NAN,np.NAN,np.NAN],
    [82348,6235,52354],
]

df = pd.DataFrame(dt)

print(df.values.tolist())
print(np.array(df))
# 其实df.values的本质就返回numpy.ndarray类型

'''
[[33.0, 44.0, 55.0], [66.0, 77.0, 99.0], [888.0, 665.0, 5554.0], [nan, nan, nan], [82348.0, 6235.0, 52354.0]]
[[3.3000e+01 4.4000e+01 5.5000e+01]
 [6.6000e+01 7.7000e+01 9.9000e+01]
 [8.8800e+02 6.6500e+02 5.5540e+03]
 [       nan        nan        nan]
 [8.2348e+04 6.2350e+03 5.2354e+04]]
'''

# 如果直接用[df]转成list的话， 会把index和columns带上
a = [df]
print(a)

'''
[         0       1        2
0     33.0    44.0     55.0
1     66.0    77.0     99.0
2    888.0   665.0   5554.0
3      NaN     NaN      NaN
4  82348.0  6235.0  52354.0]
'''
```





#### 8.3 locate

-  `loc[]`：按照索引取
- `iloc[]`：硬按数据下标取 [0,1,2]

``` python
weapons = {
    'name': ['bow of dusk', 'dagger of thief', 'the wind waker'],
    'price': [300, 210, 420],
    'duration': [100, 100, 100],
    'attack': [64, 45, 20],
    'magic': [float('NaN'), 0, 50]
}
weapons_df = pd.DataFrame(weapons, columns=['name', 'duration', 'price', 'attack', 'magic'])
print(weapons_df)

'''
              name  duration  price  attack  magic
0      bow of dusk       100    300      64    NaN
1  dagger of thief       100    210      45    0.0
2   the wind waker       100    420      20   50.0
'''
# print(weapons_df[0]) 直接这样会报错

weapons_df.index = ['1', 2, 3]	# 可以直接修改对象的属性
print(weapons_df.loc['1'])
print(type(weapons_df.loc['1']))	# 直接用[]取出了Series

print(weapons_df.loc[['1']])
#print(weapons_df.loc[[0]])
print(type(weapons_df.loc[['1']]))	# 用[[]]取出来还是DataFrame

'''
name        bow of dusk
duration            100
price               300
attack               64
magic               NaN
Name: 1, dtype: object
<class 'pandas.core.series.Series'>
          name  duration  price  attack  magic
1  bow of dusk       100    300      64    NaN
<class 'pandas.core.frame.DataFrame'>
'''
print(weapons_df.loc[['1', 2]]) # 要取多行就只能用[[]]
print(weapons_df.iloc[[1, 2]])	# int locate
#print(weapons_df.iloc[[1, 1]])	# 取1行 两遍
'''
              name  duration  price  attack  magic
1      bow of dusk       100    300      64    NaN
2  dagger of thief       100    210      45    0.0
              name  duration  price  attack  magic
2  dagger of thief       100    210      45    0.0
3   the wind waker       100    420      20   50.0
'''
```



#### 8.4 CSV

- CSV（Comma-Separated Values，逗号分隔值，有时也称为字符分隔值，因为分隔字符也可以不是逗号），其文件以纯文本形式存储表格数据（数字和文本）。

``` python
weapons = {
    'name': ['bow of dusk', 'dagger of thief', 'the wind waker'],
    'price': [300, 210, 420],
    'duration': [100, 100, 100],
    'attack': [64, 45, 20],
    'magic': [float('NaN'), 0, 50],
}

weapons_df = pd.DataFrame(weapons, columns=['name', 'duration', 'price', 'attack', 'magic'])

weapons_df.to_csv('weapon.csv', index=False)	# 保存到csv文件，不要保存index

df = pd.read_csv('weapon.csv')		# 读取csv文件，第一行默认解释为header，不作为数据
print(df)


df = pd.read_csv('weapon.csv', header=None, names=[78,88,79,1111,5])
#header默认为0，指将第一行作为列名，=None就表示不这样，names可以指定列名
print(df)
'''
              78        88     79      1111   5   
0             name  duration  price  attack  magic
1      bow of dusk       100    300      64    NaN
2  dagger of thief       100    210      45    0.0
3   the wind waker       100    420      20   50.0
'''
```

`read_csv()`的参数：

- **header:** 指定第几行作为列名(忽略注解行)，如果没有指定列名，默认header=0; 如果指定了列名header=None
- **names** 指定列名，如果文件中不包含header的行，应该显性表示header=None
- **index_col:** 默认为None 用列名作为DataFrame的行标签，如果给出序列，则使用MultiIndex。如果读取某文件,该文件每行末尾都有带分隔符，考虑使用index_col=False使panadas不用第一列作为行的名称
- **squeeze:** 默认为False, True的情况下返回的类型为Series
- **prefix:**默认为none, 当header =None 或者没有header的时候有效，例如’x’ 列名效果 X0, X1, …
- **dtype:** E.g. {‘a’: np.float64, ‘b’: np.int32} 指定数据类型
	- **skiprows:** list-like or integer or callable, default None 忽略某几行或者从开始算起的几行
- **nrows：** int 读取的行数



#### 8.5 数据清洗

- 数据清洗是对一些没有用的数据进行处理的过程。很多数据集存在数据缺失、数据格式错误、错误数据或重复数据的情况，如果要对使数据分析更加准确，就需要对这些没有用的数据进行处理。

***将 NaN 丢弃：***

- 如果我们要删除包含空字段的行，可以使用 **dropna()** 方法，语法格式如下：

```python
DataFrame.dropna(axis=0, how='any', thresh=None, subset=None, inplace=False)
```

**参数说明：**

- axis：默认为 **0**，表示逢空值剔除整行，如果设置参数 **axis＝1** 表示逢空值去掉整列。
- how：默认为 **'any'** 如果一行（或一列）里任何一个数据有出现 NaN 就去掉整行，如果设置 **how='all'** 一行（或列）都是 NaN 才去掉这整行。
- thresh：设置需要多少非空值的数据才可以保留下来的。
- subset：设置想要检查的列。如果是多个列，可以使用列名的 list 作为参数。
- inplace：如果设置 True，将计算得到的值直接覆盖之前的值并返回 None，修改的是源数据。

我们可以通过 **isnull()** 判断各个单元格是否为空。

``` python
df = pd.read_csv('property-data.csv')
print(df)
# print(df['NUM_BEDROOMS'])
# print(df['NUM_BEDROOMS'].isnull())
new_df = df.dropna()
print(new_df)

'''
           PID  ST_NUM     ST_NAME OWN_OCCUPIED NUM_BEDROOMS NUM_BATH SQ_FT
0  100001000.0   104.0      PUTNAM            Y            3        1  1000
1  100002000.0   197.0   LEXINGTON            N            3      1.5    --
2  100003000.0     NaN   LEXINGTON            N          NaN        1   850
3  100004000.0   201.0    BERKELEY           12            1      NaN   700
4          NaN   203.0    BERKELEY            Y            3        2  1600
5  100006000.0   207.0    BERKELEY            Y          NaN        1   800
6  100007000.0     NaN  WASHINGTON          NaN            2   HURLEY   950
7  100008000.0   213.0     TREMONT            Y            1        1   NaN
8  100009000.0   215.0     TREMONT            Y           na        2  1800
           PID  ST_NUM    ST_NAME OWN_OCCUPIED NUM_BEDROOMS NUM_BATH SQ_FT
0  100001000.0   104.0     PUTNAM            Y            3        1  1000
1  100002000.0   197.0  LEXINGTON            N            3      1.5    --
8  100009000.0   215.0    TREMONT            Y           na        2  1800

'''
```

***将 NaN 替换：***

- 我们使用 **fillna() **来替换数据（其参数与`dropna` 差不多）
- Pandas使用 **mean()**、**median()** 和 **mode()** 方法计算列的均值、中位数值和众数。一般用它们来替换

``` python
x = df['ST_NUM'].mean()  
df['ST_NUM'].fillna(x, inplace=True)	# 取一列操作时，为了是原df也改变，就必须使inplace为True

print(df)

'''
           PID      ST_NUM     ST_NAME OWN_OCCUPIED NUM_BEDROOMS NUM_BATH SQ_FT
0  100001000.0  104.000000      PUTNAM            Y            3        1  1000
1  100002000.0  197.000000   LEXINGTON            N            3      1.5    --
2  100003000.0  191.428571   LEXINGTON            N          NaN        1   850
3  100004000.0  201.000000    BERKELEY           12            1      NaN   700
4          NaN  203.000000    BERKELEY            Y            3        2  1600
5  100006000.0  207.000000    BERKELEY            Y          NaN        1   800
6  100007000.0  191.428571  WASHINGTON          NaN            2   HURLEY   950
7  100008000.0  213.000000     TREMONT            Y            1        1   NaN
8  100009000.0  215.000000     TREMONT            Y           na        2  1800
'''
```

- `interpolate`将 NaN 按某种方法替换
- DataFrame.interpolate(method=‘linear’, axis=0, limit=None, inplace=False, limit_direction=‘forward’, limit_area=None, downcast=None, **kwargs)

``` python
df = pd.DataFrame([(0.0,  np.nan, -1.0, 1.0),
                    (np.nan, 2.0, np.nan, np.nan),
                   (2.0, 3.0, np.nan, 9.0),
                    (np.nan, 4.0, -4.0, 16.0)],
                   columns=list('abcd'))
print(df)
df.interpolate(method='linear', inplace=True, limit_direction='both', axis=1)
print(df)

'''
     a    b    c     d
0  0.0  NaN -1.0   1.0
1  NaN  2.0  NaN   NaN
2  2.0  3.0  NaN   9.0
3  NaN  4.0 -4.0  16.0
     a    b    c     d
0  0.0 -0.5 -1.0   1.0
1  2.0  2.0  2.0   2.0
2  2.0  3.0  6.0   9.0
3  4.0  4.0 -4.0  16.0
'''
```

***数据格式化：***

``` python
#日期格式化
date = {
    'date': ['2020/12/20', '2020.12.18', '20201216'],
    'duration': [41, 43, 4500]
}
df = pd.DataFrame(date, index=['day1', 'day2', 'day3'])
df['date'] = pd.to_datetime(df['date'])

print(df)

'''
           date  duration
day1 2020-12-20        41
day2 2020-12-18        43
day3 2020-12-16      4500
'''
```

***替换错误数据：***

``` python
df.iloc[2, 1] = 45
# df.loc['day3', 'duration'] = 45
print(df)

'''
           date  duration
day1 2020-12-20        41
day2 2020-12-18        43
day3 2020-12-16        45
'''
```

- `replace()` 将某值用其他值替换
- DataFrame.replace(to_replace=None, value=None, inplace=False, limit=None, regex=False, method=‘pad’)

``` python
df = pd.DataFrame({'A': [0, 1, 2, 3, 4],
                 'B': [0, 6, 7, 8, 9],
                   'C': ['a', 'b', 'c', 'd', 'e']})
print(df)
# df.replace(0, 5, inplace=True)
df.replace([6, 7], [67, 76], inplace=True)#可以是列表与数的组合
# df.replace({0: 10, 1: 100, 2: 200}, inplace=True)#也可以用字典的形式来进行替换
df.replace({'A': [0, 1], 'B': [0, 9]}, 100, inplace=True)#可以是字典和数值等 的组合
print(df)

'''
   A  B  C
0  0  0  a
1  1  6  b
2  2  7  c
3  3  8  d
4  4  9  e
     A    B  C
0  100  100  a
1  100   67  b
2    2   76  c
3    3    8  d
4    4  100  e
'''
```

***清洗重复数据：***

如果我们要清洗重复数据，可以使用 **duplicated()** 和 **drop_duplicates()** 方法。

``` python
person = {
  "name": ['Google', 'Runoob', 'Runoob', 'Taobao'],
  "age": [50, 40, 40, 23]
}
df = pd.DataFrame(person)

print(df.duplicated())
'''
0    False
1    False
2     True
3    False
dtype: bool
'''

df.drop_duplicates(inplace=True)
print(df)
'''
     name  age
0  Google   50
1  Runoob   40
3  Taobao   23
'''

```

#### 8.6 取得唯一值

- `unique()` 是以 数组形式（numpy.ndarray）返回Series的所有唯一值（特征的所有唯一值）

- `nunique() ` Return number of unique elements in the object. 即返回的是唯一值的个数

``` python
dt = {
    'f1': ['a', 1, 2, 11],
    'f2': ['b', 3, 4, 22],
    'f3': ['c', 3, 5, 33],
    'cls': [1, 2, 4, 1]
}
dt = pd.DataFrame(dt)
print(dt)
print(dt['cls'].unique())
print(dt['cls'].nunique())

'''
   f1  f2  f3  cls
0   a   b   c    1
1   1   3   3    2
2   2   4   5    4
3  11  22  33    1
[1 2 4]
3
'''
```

#### 8.7 按某列分组

`groupby`(by=None, axis=0, level=None, as_index: bool = True, sort: bool = True, group_keys: bool = True, squeeze: bool = False, observed: bool = False)

- 得到`pandas.core.groupby.generic.DataFrameGroupBy `对象，在此基础上可以进行很多操作，比如求分组的数量，分组内的最大值最小值平均值等。在sql中，就是大名鼎鼎的groupby操作
- 返回的结果同样是一个generater

``` python
data = [
    ['tom', 'c1', 68, 99],
    ['jerry', 'c1', 88, 51],
    ['lucy', 'c2', 66, 77],
    ['jack', 'c2', 85, 74],
]

df = pd.DataFrame(data, columns=['name', 'class', 'math', 'english'], )
print(df, '\n**********')
res = df.groupby(by='class')
print(res.mean())

'''
    name  class  math  english
0    tom     c1    68       99
1  jerry     c1    88       51
2   lucy     c2    66       77
3   jack     c2    85       74 
**********
       math  english
class               
c1     78.0     75.0
c2     75.5     75.5
'''
# 可直接将其转化为list
# list中每个元素是一个元组，元组第一个位置保存类别，第二个位置就是相应的DataFrame
print(list(res))
'''
[('c1',     name class  math  english
0    tom    c1    68       99
1  jerry    c1    88       51), ('c2',    name class  math  english
2  lucy    c2    66       77
3  jack    c2    85       74)]
'''
```

***agg聚合操作***

- 聚合操作是`groupby`后非常常见的操作。聚合操作可以用来min, max, sum, mean, median, std, var(方差), count等。

``` python
print(res.agg({'math': 'median', 'english': 'mean'}))

'''
       math  english
class               
c1     78.0     75.0
c2     75.5     75.5
'''
```

***transform***

``` python
# 如果此时想要将班级数学中位数成绩 和英语平均成绩加到总的成绩单df中

df['math_median'] = res['math'].transform('median')
df['english_mean'] = res['english'].transform('mean')
print(df)

'''
    name class  math  english  math_median  english_mean
0    tom    c1    68       99         78.0          75.0
1  jerry    c1    88       51         78.0          75.0
2   lucy    c2    66       77         75.5          75.5
3   jack    c2    85       74         75.5          75.5
'''
```

***apply***

- apply 可以实现更加复杂的操作，可以将函数作为参数传进去

``` python
# 比如我们想要得到数学成绩最高的学生
def get_highest_math_grade(x):
    # 这里自动传入的x是groupby之后的所有DataFrame的集合体，
    # 对x进行DataFrame的操作对作用在所有的DataFrame上
    y = x.sort_values(by='math', ascending=True)
    return y.iloc[-1, :]

highest_math_grade = res.apply(get_highest_math_grade)
print(highest_math_grade)

'''
        name class  math  english
class                            
c1     jerry    c1    88       51
c2      jack    c2    85       74
'''

```

#### 8.8 判断是否为空 

`isnull()` 判断dataframe是否为空

``` python
dt = [
    [33,44,55],
    [66,77,99],
    [888,665,5554],
    [np.NAN,np.NAN,np.NAN],
    [82348,6235,52354],
]

df = pd.DataFrame(dt)

print(df.isnull()) 

'''
       0      1      2
0  False  False  False
1  False  False  False
2  False  False  False
3   True   True   True
4  False  False  False
'''

print(df.isnull().any()) # 判断每列是否任意值为空	.all() 就是判断是否全为空
'''
0    True
1    True
2    True
'''
print(df.isnull().any().any())
# True
```



