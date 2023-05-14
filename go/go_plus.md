# go_plus

[【置顶】Go语言学习之路/Go语言教程 | 李文周的博客 (liwenzhou.com)](https://www.liwenzhou.com/posts/Go/golang-menu/)

[Go语言标准库文档](https://studygolang.com/pkgdoc) 

https://www.topgoer.com/

[数据结构和算法（Golang实现） - 《数据结构和算法（Golang实现）》 - 书栈网 · BookStack](https://www.bookstack.cn/read/hunterhug-goa.c/README.md)

 ## 1. Go实现TCP/UDP通信

### socket

`Socket`是应用层与TCP/IP协议族通信的中间软件抽象层。在设计模式中，`Socket`其实就是一个门面模式，它把复杂的TCP/IP协议族隐藏在`Socket`后面，对用户来说只需要调用Socket规定的相关函数，让`Socket`去组织符合指定的协议数据然后进行通信。

<img src="img\image-20220508194119681.png" alt="image-20220508194119681" style="zoom:67%;" />

### Go语言实现TCP通信

#### TCP服务端

一个TCP服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。因为Go语言中创建多个goroutine实现并发非常方便和高效，所以我们可以每建立一次链接就创建一个goroutine去处理。

TCP服务端程序的处理流程：

1. 监听端口
2. 接收客户端请求建立链接
3. 创建goroutine处理链接。

我们使用Go语言的net包实现的TCP服务端代码如下：

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {

	defer conn.Close()
	// 3.与client通信
	reader := bufio.NewReader(os.Stdin)
	for {
		var tmp = [128]byte{}
		n, e3 := conn.Read(tmp[:])
		if e3 != nil {
			fmt.Println("read from connection failed, err:", e3)
			return
		}
		fmt.Println(string(tmp[:n]))
		fmt.Print("[server] 请说：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
	}
}

func main() {
	// 1.本地端口启动服务
	listener, e1 := net.Listen("tcp", "127.0.0.1:20000")
	if e1 != nil {
		fmt.Println("start tcp server on 127.0.0.1:20000 failed, err:", e1)
		return
	}
	// 2.等待client建立连接
	for {
		conn, e2 := listener.Accept()
		if e2 != nil {
			fmt.Println("accept failed, err:", e2)
			return
		}
		go process(conn)
	}
}
```

将上面的代码保存之后编译成`server`或`server.exe`可执行文件。

#### TCP客户端

一个TCP客户端进行TCP通信的流程如下：

1. 建立与服务端的链接
2. 进行数据收发
3. 关闭链接

使用Go语言的net包实现的TCP客户端代码如下：

``` go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1.与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("connect server on 127.0.0.1:20000 failed, err:", err)
		return
	}
	// 2.发送数据
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("[client] 请说：")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "exit" {
			break
		}
		conn.Write([]byte(text))
		var tmp = [128]byte{}
		n, e3 := conn.Read(tmp[:])
		if e3 != nil {
			fmt.Println("read from connection failed, err:", e3)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
	conn.Close()
}
```

将上面的代码编译成`client`或`client.exe`可执行文件，先启动server端再启动client端，在client端输入任意内容回车之后就能够在server端看到client端发送的数据，从而实现TCP通信。

#### TCP黏包

**为什么会出现粘包**

主要原因就是tcp数据传递模式是流模式，在保持长连接的时候可以进行多次的收和发。

“粘包”可发生在发送端也可发生在接收端：

1. 由Nagle算法造成的发送端的粘包：Nagle算法是一种改善网络传输效率的算法。简单来说就是当我们提交一段数据给TCP发送时，TCP并不立刻发送此段数据，而是等待一小段时间看看在等待期间是否还有要发送的数据，若有则会一次把这两段数据发送出去。
2. 接收端接收不及时造成的接收端粘包：TCP会把接收到的数据存在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时的把TCP的数据取出来，就会造成TCP缓冲区中存放了几段数据。

**黏包示例**

服务端代码如下：

```go
// socket_stick/server/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [1024]byte
	for {
		n, err := reader.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
```

客户端代码如下：

```go
// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
		// time.Sleep(time.Millisecond) 加上等待时间试试
	}
}
```

将上面的代码保存后，分别编译。先启动服务端再启动客户端，可以看到服务端输出结果如下：

```bash
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?Hello, Hello. How are you?
收到client发来的数据： Hello, Hello. How are you?Hello, Hello. How are you?
```

客户端分10次发送的数据，在服务端并没有成功的输出10次，而是多条数据“粘”到了一起。

**解决办法**

出现”粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。

封包：封包就是给一段数据加上包头，这样一来数据包就分为包头和包体两部分内容了(过滤非法包时封包会加入”包尾”内容)。包头部分的长度是固定的，并且它存储了包体的长度，根据包头长度固定以及包头中含有包体长度的变量就能正确的拆分出一个完整的数据包。

我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

```go
// socket_stick/proto/proto.go
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
```

接下来在服务端和客户端分别使用上面定义的`proto`包的`Decode`和`Encode`函数处理数据。

服务端代码如下：

```go
// socket_stick/server2/main.go

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
```

客户端代码如下：

```go
// socket_stick/client2/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
```

### Go语言实现UDP通信

UDP协议（User Datagram Protocol）中文名称是用户数据报协议，是OSI（Open System Interconnection，开放式系统互联）参考模型中一种**无连接**的传输层协议，不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信，但是UDP协议的实时性比较好，通常用于视频直播相关领域。

#### UDP服务端

使用Go语言的`net`包实现的UDP服务端代码如下：

```go
// UDP/server/main.go

// UDP server端
func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}
}
```

#### UDP客户端

使用Go语言的`net`包实现的UDP客户端代码如下：

```go
// UDP 客户端
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	sendData := []byte("Hello server")
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return
	}
	fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
```

## 2. 递归

#### 递归实现斐波那契数列

``` go
func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
```

#### 递归实现二分查找

``` go
func BinarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return BinarySearch(nums, target, left, mid-1)
	} else {
		return BinarySearch(nums, target, mid+1, right)
	}
}
```

#### 非递归实现二分查找

``` go
func BinarySearch2(nums []int, target int) int {
	left, mid, right := 0, -1, len(nums)-1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
```

## 3. 链表

#### 单链表

``` go
type LinkNode struct {
	val      int
	nextNode *LinkNode
}


func main() {
	var n4 LinkNode = LinkNode{4, nil}
	var n3 LinkNode = LinkNode{3, &n4}
	var n2 LinkNode = LinkNode{2, &n3}
	var n1 LinkNode = LinkNode{1, &n2}
	// 这里tmp是LinkNode指针类型，因为struct是值类型不能用来==nil，
	for tmp := &n1; tmp != nil; tmp = tmp.nextNode {
		fmt.Println(tmp.val)
	}
}
```

#### 双向链表

go自带的包container/list就是双向链表

导入list包：`import  list "container/list"`

``` go
// list中的元素
type Element struct {
	next, prev *Element
	list *List
	Value any
}
// 由Element来维护前后指针
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int     // current list length excluding (this) sentinel element
}
// 初始化或清空list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}
// 可以看出，List一直有一个不可读取的root节点，
```

需要注意的是使用List时，不会直接用List对象，因为list里的函数的caller几乎都是Element，不是List，所以一般用List.Back()。其他插值，删除等等方法参见[Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国 (studygolang.com)](https://studygolang.com/pkgdoc)

#### 环链表

go自带的包container/ring就是环链表

导入ring包：`import  ring "container/ring"`

``` go
// 只是环的一个元素，因为只要拿一个环的一个元素就拿到了整个环
type Ring struct {
	next, prev *Ring
	Value      any // 可以直接改
}
// New creates a ring of n elements.
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
```

## 4. 栈

go没有自带的栈类型，我们可以用数组(切片)或者列表来实现

- **数组实现**：能快速随机访问存储的元素，通过下标 `index` 访问，支持随机访问，查询速度快，但存在元素在数组空间中大量移动的操作，增删效率低。

- **链表实现**：只支持顺序访问，在某些遍历操作中查询速度慢，但增删元素快

#### 数组实现

``` go
// 实现数组栈 ArrayStack1
type ArrayStack struct {
	array []int // 底层切片
	size  int
	lock  sync.Mutex
}

// 入栈
func (stack *ArrayStack) Push1(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.array = append(stack.array, v)
	stack.size++
}

// 出栈
func (stack *ArrayStack) Pop1() int {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	if stack.size <= 0 {
		panic("stack is empty.")
	}
	v := stack.array[stack.size-1]
	stack.size--
	stack.array = stack.array[:stack.size]
	return v
	/*
		如果切片偏移量向前移动 stack.array[0 : stack.size-1]，
		表明最后的元素已经不属于该数组了，数组变相的缩容了。
		此时，切片被缩容的部分并不会被回收，仍然占用着空间，
		所以空间复杂度较高，但操作的时间复杂度为：O(1)。

		如果我们创建新的数组 newArray，然后把老数组的元素复制到新数组，
		就不会占用多余的空间，但移动次数过多，时间复杂度为：O(n)。
	*/
}

func (stack *ArrayStack) Peek1() int {
	if stack.size == 0 {
		panic("stack is empty.")
	}
	return stack.array[stack.size-1]
}
func (stack *ArrayStack) Size1() int {
	return stack.size
}
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
```

#### 链表实现

``` go
// 实现链表栈 LinkStack
type LinkStack struct {
	root *LinkNode // 链表起点
	size int
	lock sync.Mutex
}
type LinkNode struct {
	Next  *LinkNode
	value int
}

// 入栈
// 如果把root当作栈底的话，那么入栈出栈等等都要一直把整个栈找一遍找到栈顶，时间复杂度为O(n)，不合适
// 故把root当作栈顶
func (stack *LinkStack) Push2(v int) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	node := LinkNode{value: v}
	if stack.root == nil {
		stack.root = &node
	} else {
		node.Next = stack.root
		stack.root = &node
	}
	stack.size++
}

// 出栈
func (stack *LinkStack) Pop2() int {
	if stack.size <= 0 {
		panic("empty")
	}
	stack.lock.Lock()
	defer stack.lock.Unlock()
	result := stack.root.value
	stack.root = stack.root.Next
	stack.size--
	return result
}

// 获得栈顶
func (stack *LinkStack) Peek2() int {
	if stack.size <= 0 {
		panic("empty")
	}
	return stack.root.value
}

// 栈大小
func (stack *LinkStack) Size2() int {
	return stack.size
}

// 栈是否为空
func (stack *LinkStack) IsEmpty2() bool {
	return stack.size == 0
}
```

## 5. 队列

#### 实现数组队列 ArrayQueue

``` go
// 实现数组队列 ArrayQueue
type ArrayQueue struct {
	array []int
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Add(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.array = append(queue.array, v)
	queue.size++
}

func (queue *ArrayQueue) Remove() int {
	if queue.size <= 0 {
		panic("empty")
	}
	queue.lock.Lock()
	defer queue.lock.Unlock()
	v := queue.array[0]
	queue.array = queue.array[1:]
	queue.size--
	return v
}

```

#### 实现链表队列 LinkQueue

``` go
// 实现链表队列 LinkQueue
type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}
type LinkNode struct {
	Next  *LinkNode
	Value int
}

// 入队
func (queue *LinkQueue) Add2(v int) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	node := LinkNode{Value: v}
	// 当然加在链头链尾都可以，只是时间复杂度不一样
	// 但入队出队的的时间复杂度总是一个O(n),一个O(1)
	if queue.root == nil {
		queue.root = &node
	} else {
		tmp := queue.root
		for ; tmp.Next != nil; tmp = tmp.Next {
		}
		tmp.Next = &node
	}
	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove2() int {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	if queue.size <= 0 {
		panic("empty")
	}
	tmp := queue.root
	v := tmp.Value
	queue.root = tmp.Next
	queue.size--
	return v
}
```

## 6. 集合

`Set` 可以没有顺序关系，也可以按值排序，算一种特殊的不可重复列表。一般实现无序不重复的集合。

因为我们知道字典的键是不重复的，所以只要我们不考虑字典的值，就可以实现集合，我们来实现存整数的集合 `Set`：

``` go
type set struct {
	m    map[int]struct{} // 值使用空结构体不会占空间
	size int
	lock sync.RWMutex // 读写锁
}

func NewSet(cap int) *set {
	tmp := make(map[int]struct{}, cap)
	return &set{m: tmp}
}

func (s *set) Add(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[item] = struct{}{}
	s.size = len(s.m)
}

func (s *set) Remove(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, has := s.m[item]
	if has {
		delete(s.m, item)
		s.size = len(s.m)
	}
}
func (s *set) Has(item int) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, has := s.m[item]
	return has
}

// 查看集合大小
func (s *set) Len() int {
	return s.size
}

// 集合是够为空
func (s *set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

// 清除集合所有元素
func (s *set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = make(map[int]struct{})
	s.size = 0
}

// 将集合转化为列表
func (s *set) List() []int {
	s.lock.Lock()
	defer s.lock.Unlock()
	ans := make([]int,0)
	fmt.Println(s.size)
	for k := range s.m {
		ans = append(ans, k)
	}
	return ans
}
func main() {
	s := NewSet(0)
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(3)
	fmt.Print(s.Has(4))
	s.Remove(4)
	fmt.Println(s.m)
	fmt.Println(s.List())

}
```

## 7. 树

树是一种比较高级的基础数据结构，由 `n` 个有限节点组成的具有层次关系的集合。（算树的高度或深度时不算根节点这一层）

树的定义：

1. 有节点间的层次关系，分为父节点和子节点。
2. 有唯一一个根节点，该根节点没有父节点。
3. 除了根节点，每个节点有且只有一个父节点。
4. 每一个节点本身以及它的后代也是一棵树，是一个递归的结构。
5. 没有后代的节点称为叶子节点，没有节点的树称为空树

### 二叉树

- **二叉树**：每个节点最多只有两个儿子节点的树。

- **满二叉树**：叶子节点与叶子节点之间的高度差为 `0` 的二叉树，即整棵树是满的，树呈满三角形结构。

- **完全二叉树**：完全二叉树是由满二叉树而引出来的，设二叉树的深度为 `k`，除第 `k` 层外，其他各层的节点数都达到最大值，且第 `k` 层所有的节点都连续集中在最左边。数组也可以用来表示二叉树，一般用来表示完全二叉树。

树根据子节点的多寡，有二叉树，三叉树，四叉树等，我们这里主要介绍二叉树。

**二叉树的数学特征**

1. 高度为 `h≥0` 的二叉树至少有 `h+1` 个结点，比如最不平衡的二叉树就是退化的线性链表结构，所有的节点都只有左儿子节点，或者所有的节点都只有右儿子节点。
2. 高度为 `h≥0` 的二叉树至多有 `2^h+1` 个节点，比如这棵树是满二叉树。
3. 含有 `n≥1` 个结点的二叉树的高度至多为 `n-1`，由 `1` 退化的线性链表可以反推。
4. 含有 `n≥1` 个结点的二叉树的高度至少为 `logn`，由 `2` 满二叉树可以反推。
5. 在二叉树的第 `i` 层，至多有 `2^(i-1)` 个节点，比如该层是满的。

``` go
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}
type ArrayQueue struct {
	array []*TreeNode
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()
	queue.array = append(queue.array, v)
	queue.size++
}

func (queue *ArrayQueue) Remove() *TreeNode {
	if queue.size <= 0 {
		panic("empty")
	}
	queue.lock.Lock()
	defer queue.lock.Unlock()
	v := queue.array[0]
	queue.array = queue.array[1:]
	queue.size--
	return v
}

/*
构建一棵树后，我们希望遍历它，有四种遍历方法：
	先序遍历：先访问根节点，再访问左子树，最后访问右子树。
	后序遍历：先访问左子树，再访问右子树，最后访问根节点。
	中序遍历：先访问左子树，再访问根节点，最后访问右子树。(这三种都是深度优先遍历)
	层次遍历：每一层从左到右访问每一个节点。（此即为广度优先遍历）
*/

// 先序遍历
func PreOrder(tree *TreeNode) {
	fmt.Print(tree.Data, " ")
	if tree.Left != nil {
		PreOrder(tree.Left)
	}
	if tree.Right != nil {
		PreOrder(tree.Right)
	}
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	PreOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	PreOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	// 先打印左子树
	PostOrder(tree.Left)
	// 再打印右字树
	PostOrder(tree.Right)
	// 再打印根节点
	fmt.Print(tree.Data, " ")
}

// 层次遍历
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	queue := new(ArrayQueue)
	queue.Add(tree)
	for queue.size > 0 {
		element := queue.Remove()
		fmt.Print(element.Data, " ")
		if element.Left != nil {
			queue.Add(element.Left)
		}
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}
```

## 8. 排序

### 交换排序

#### 冒泡排序

比如现在有一堆乱序的数，比如：`5 9 1 6 8 14 6 49 25 4 6 3`。

第一轮迭代：从第一个数开始，**依次比较相邻的两个数（一直这样做）**，如果前面一个数比后面一个数大，那么交换位置，直到处理到最后一个数，最后的这个数是最大的。

第二轮迭代：因为最后一个数已经是最大了，现在重复第一轮迭代的操作，但是只处理到倒数第二个数。

第三轮迭代：因为最后一个数已经是最大了，最后第二个数是次大的，现在重复第一轮迭代的操作，但是只处理到倒数第三个数。

``` go
// 交换排序之冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
```

冒泡排序的平均时间复杂度为：`O(n^2)`。

冒泡排序交换和比较的次数相加是一个和 `N` 有关的平方数，所以冒泡排序的最好和最差时间复杂度都是：`O(n^2)`。

冒泡排序算法是**稳定**的，因为如果两个相邻元素相等，是不会交换的，保证了稳定性的要求。

#### 快速排序

快速排序通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。

步骤如下：

1. 先从数列中取出一个数作为基准数。一般取第一个数。
2. 分区过程，将比这个数大的数全放到它的右边，小于或等于它的数全放到它的左边。
3. 再对左右区间重复第二步，直到各区间只有一个数。

``` go
// 交换排序之快速排序
func QuickSort(nums []int, left, right int) {
	if right > left {
		mid := partition(nums, left, right)
		QuickSort(nums, left, mid-1)
		QuickSort(nums, mid+1, right)
	}
}

// 快排的分割函数(nums[right]是有值的)
func partition(nums []int, left, right int) int {
	for i := left + 1; i <= right; i++ {
		if nums[i] <= nums[left] {
			nums[i], nums[left] = nums[left], nums[i]
			left = i
		} else { // 把较大的数往后放
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
	return left
}
```

快速排序的平均时间复杂度为：`O(nlogn)`。

最差的情况下，每次都不能平均地切分，每次切分都因为基准数是最大的或者最小的，不能分成两个数列，这样时间复杂度变为了 `T(n) = T(n-1) + O(n)`，按照主定理计算可以知道时间复杂度为：`O(n^2)`

快速排序是不稳定的，因为切分过程中进行了交换，相同值的元素可能发生位置变化。

**改进**

- **切分**的结果极大地影响快速排序的性能，为了避免切分不均匀情况的发生，有几种方法改进：

1. 每次进行快速排序切分时，先将数列随机打乱，再进行切分，这样随机加了个震荡，减少不均匀的情况。当然，也可以随机选择一个基准数，而不是选第一个数。
2. 每次取数列头部，中部，尾部三个数，取三个数的中位数为基准数进行切分。

>  方法 1 相对好，而方法 2 引入了额外的比较操作，一般情况下我们可以随机选择一个基准数。

快速排序可以继续进行算法改进。

1. 在**小规模数组的情况下，直接插入排序**的效率最好，当快速排序递归部分进入小数组范围，可以切换成直接插入排序。
2. 排序数列可能存在大量重复值，使用三向切分快速排序，将数组分成三部分，大于基准数，等于基准数，小于基准数，这个时候需要维护三个下标。
3. 使用伪尾递归减少程序栈空间占用，使得栈空间复杂度从 `O(logn)~log(n)` 变为：`O(logn)`。

#### 归并排序

我们先介绍两个有序的数组合并成一个有序数组的操作。

1. 先申请一个辅助数组，长度等于两个有序数组长度的和。
2. 从两个有序数组的第一位开始，比较两个元素，哪个数组的元素更小，那么该元素添加进辅助数组，然后该数组的元素变更为下一位，继续重复这个操作，直至数组没有元素。
3. 返回辅助数组。

归并排序就是不端地进行归并操作，

![image-20220521204838520](\img\image-20220521204838520.png)

``` go
// 自顶而下归并排序，排序范围是[begin, end)
func MergeSort(nums []int, begin, end int) {
	if end-begin > 1 { // 至少要两个数才需要归并
		mid := begin + (end-begin)/2
		MergeSort(nums, begin, mid)
		MergeSort(nums, mid, end)
		merge(nums, begin, mid, end)
	}
}

// 完成一次归并，例如把[1,4,2,3]mid=2, 归并为[1,2,3,4]
func merge(nums []int, begin, mid, end int) {
	// 申请额外的空间来合并两个有序数组，这两个数组是 nums[begin,mid),array[mid,end)
	leftSize := mid - begin
	rightSize := end - mid
	newNums := make([]int, 0, leftSize+rightSize)
	l, r := begin, mid
	for l < mid && r < end {
		if nums[l] < nums[r] {
			newNums = append(newNums, nums[l])
			l++
		} else {
			newNums = append(newNums, nums[r])
			r++
		}
	}
	newNums = append(newNums, nums[l:mid]...)
	newNums = append(newNums, nums[r:end]...)
	for i := 0; i < len(newNums); i++ {
		nums[begin+i] = newNums[i]
	}
}
```

时间复杂度是 `O(nlogn)`。

### 选择排序

#### 简单选择排序

选择排序，一般我们指的是简单选择排序，也可以叫直接选择排序，它不像冒泡排序一样相邻地交换元素，而是**通过选择最小的元素，放到最前面去，每轮迭代只需交换一次**。虽然交换次数比冒泡少很多，但效率和冒泡排序一样的糟糕。

现在有一堆乱序的数，比如：`5 9 1 6 8 14 6 49 25 4 6 3`。

第一轮迭代，从第一个数开始，左边到右边进行扫描，找到最小的数 1，与数列里的第一个数交换位置。

第二轮迭代，从第二个数开始，左边到右边进行扫描，找到第二小的数 3，与数列里的第二个数交换位置。

第三轮迭代，从第三个数开始，左边到右边进行扫描，找到第三小的数 4，与数列里的第三个数交换位置。

第N轮迭代：….

``` go 
// 选择排序之简单选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := nums[i]
		idx := -1
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				idx = j
			}
		}
		nums[i], nums[idx] = nums[idx], nums[i]
	}
}
```

最佳和最坏时间复杂度仍然是：`O(n^2)`。

**改进**

上面的算法需要从某个数开始，一直扫描到尾部，我们可以优化算法，使得复杂度减少一半。

我们每一轮，除了找最小数之外，还找最大数，然后分别和前面和后面的元素交换，这样循环次数减少一半

#### 堆排序





### 插入排序

#### 直接插入排序

每次把一个数插到已经排好序的数列里面形成新的排好序的数列，以此反复。

``` go
// 插入排序之直接插入排序
func InsertSort(nums []int) {
	tmp := 0
	for i := 1; i < len(nums); i++ {
		tmp = nums[i]
		for k := i - 1; k >= 0; k-- {
			if nums[k] > tmp {
				nums[k+1], nums[k] = nums[k], nums[k+1]
			} else {
				break
			}
		}
	}
}
```

直接插入排序的平均复杂度是`O(n^2)`

数组规模 `n` 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。

数列中的**有序性越高，插入排序的性能越高**，因为待排序数组有序性越高，插入排序比较的次数越少。

#### 希尔排序





