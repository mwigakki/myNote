# TQUIC

[QUIC 协议详解 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/405387352)

[科普：QUIC协议原理分析 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/32553477)

[深入 HTTP/3（一）｜从 QUIC 链接的建立与关闭看协议的演进 · SOFAStack](https://www.sofastack.tech/blog/deeper-into-http/3-evolution-of-the-protocol-from-the-creation-and-closing-of-quic-links/)

## **1、QUIC 简介**

QUIC 全称：Quick UDP Internet Connections，是一种基于 UDP 的传输层协议。由 Google 自研，2012 年部署上线，2013 年提交 IETF，2021 年 5 月，IETF 推出标准版 RFC9000。

![img](https://pic4.zhimg.com/80/v2-d61a62fdfb08ed3882e1018136ce6b2f_720w.jpg)

从协议栈可以看出：QUIC = HTTP/2 + TLS + UDP

## 2、QUIC的传输机制

### 2.1、回顾TCP

和 TCP 一样，QUIC 的首要目标也是提供一个可靠、有序的流式传输协议。不仅如此，QUIC 还要保证原生的数据安全以及传输的高效。

**QUIC 就是在以一种更简洁高效的机制去对标 TCP+TLS**。当然，和 TCP+TLS 一样，QUIC 建联流程的本质都是在为上述特性服务，由于 QUIC 是基于 UDP 重新设计的协议，便也就没那么多的历史包袱。

首先看一下TCP+TLS的连接建立：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*lIq1RrDVzaoAAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:67%;" />

可以看到，在一个 client 正式开始发送应用层数据之前，需要 3 个 RTT 的交互，这算是一个非常大的开销。而从流程上来看，TCP 握手和 TLS 的握手似乎比较相似，有融合在一起的可能。的确有相关的文献探讨过在 SYN 报文里融合 ClientHello 的可行性，不过由于以下原因，这部分的探索也慢慢不了了之。

> <i style="color:black; font-size:20px;">TCP 的设计问题</i>
>
> ***<u>重传歧义问题</u>***：设想这样的一个场景，发送端发送了一个 TCP 报文，由于通信的中间设备发生了阻塞，导致该报文被延迟转发了，发送端迟迟未收到 ACK，便重新发送了一个 TCP 报文，在新的 TCP 报文达到接收端时，被延迟转发的报文也达到了接收端，接收端只会响应一个 ACK。而客户端收到 ACK 时，并不清楚这个 ACK 是对延迟转发的报文的 ACK，还是新的报文的 ACK，带来的影响也就是 RTT 的估计会不准确，从而影响拥塞控制算法的行为，降低网络效率。
>
> *<u>**难用的 TCP keepalive**</u>*：比如 TCP 连接中的另一方因为停电突然断网，我们并不知道连接断开，此时发送数据失败会进行重传，由于重传包的优先级要高于 keepalive 的数据包，因此 keepalive 的数据包无法发送出去。只有在长时间的重传失败之后我们才能判断此连接断开了。
>
> *<u>**队头阻塞问题**</u>*：严格来说这并不算 TCP 自身的问题，因为 TCP 本身是一个面向链接的协议，它保证了一个链接上的数据可靠传输，也算完成了任务。然而随着互联网的普及，人们利用网络传输的数据越来越多，如果将所有数据都放在一个 TCP 链接上传输，其中某一个数据发生丢包，后面的数据的传输都会被 block 住，严重影响效率。当然，使用多个 TCP 链接传输数据是一种解决方案，但多个链接又会带来新的开销问题及链接管理问题。



QUIC 因此提出一个新的建立连接机制，**<u>*把传输和加密握手合并成一个*</u>**，以最小化延迟（1-RTT）建立连接。**后续的连接（repeat connections）还可以通过上一次连接时缓存的信息**（TLS 1.3 Diffie-Hellman 公钥、传输参数、`NEW_TOKEN` 帧生成的令牌，等）**直接在一个经过认证且加密的通道传输数据**（0-RTT），这极大减小了建立连接的时延。

### 2.2、QUIC的加密和传输握手

QUIC 加密握手提供以下属性：

- 认证密钥交换，其中

  - 服务端总是经过身份验证

  - 客户端可以选择性进行身份验证
  - 每个连接都会产生不同并且不相关的密钥
  - 密钥材料（keying material）可用于 0-RTT 和 1-RTT 数据包的保护

- 两个端点（both endpoints）传输参数的认证值，以及服务端传输参数的保密保护

- 应用协议的认证协商（TLS 使用 [ALPN](https://link.zhihu.com/?target=https%3A//www.rfc-editor.org/info/rfc7301)）

简单来说：

> 客户端向服务器发送**连接ID**，然后服务器向客户端**返回令牌**和服务器的公共Diffie-Hellman值，这允许服务器和客户端就**初始密钥达成一致**。客户端和服务器可以立即开始交换数据，同时**建立最终的会话密钥**。
>
>  此外，与TLS 1.3类似，在服务器和客户端第一次“会面”后，它们**缓存会话密钥**，并且在新请求时，不需要握手。这就是所谓的0-RTT特性。



**<u>*密钥交换过程*</u>**

**QUIC** **首次连接需要1RTT，具体过程如下：**

![img](https://pic4.zhimg.com/80/v2-6e56fae6180c9791665e4798876f1347_720w.jpg)

**step1：** 客户端发送Inchoate Client Hello消息（CHLO）请求建立连接。

**step2：** 服务器根据一组质数p以及其原根g和a（长期私钥）算出A（长期公钥），将Apg（通过CA证书私钥加密后）放在serverConfig里面，发到Rejection消息（REJ）到客户端；（DH密钥算法）

> 服务器一开始不直接使用随机生成的短期密钥的原因就是因为客户端可以0缓存下服务端的长期公钥，这样在下一次连接的时候客户端就可以直接使用这个长期公钥实现0-RTT握手并直接发送加密数据

**setp3&4：** 客户端在接收到REJ消息后，会随机选择一个数b（短期密钥），并用CA证书获取的公钥解密出**serverConfig**里面的p、A和b就可以算出初始密钥K，并将B（Complete client hello消息）和用初始密钥K加密的Data数据发到服务器。

**step5：** 服务器收到客户端发来的公开数B，再利用p、g计算得到同样的初始秘钥K，来解密客户端发来的数据。这时会利用其他加密算法**随机生成此次会话密钥K'** ，再通过初始密钥K加密K'发送给客户端(SHLO)（每次会话都是用随机密钥，并且服务器会定期更新a和A，实际上这就是为了保证前向安全性）

> 在密码学中，前向保密（Forward Secrecy）是密码学中通讯协议的安全属性，指的是当前使用的主密钥泄漏不会导致过去的会话密钥泄漏。

**step6：** 客户端收到SHLO后利用初始密钥K解出会话密钥K'，二者后续的会话都使用K'加密。

简要的的 QUIC 建联流程：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*AcjQRb4tXS4AAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:80%;" />

(图中client对handshake的回复时就可以带上数据包了，图中并未标出) 

整个流程实际上还是和 TCP 一样是一个请求-响应的模式，然而相比于 TCP+TLS，我们还看到了一些不一样的地方：

1. 图中多了”`init packet`”、”`handshake packet`”、”`short-header packet`”的概念；

2. 图中多了 `pkt_number`的概念以及`stream`+`offset`的概念；

3. `pkt_number` 的下标变化似乎有些奇怪。

而这些不同机制就是 QUIC 实现相比于 TCP 来说更高效的点，让我们来逐一分析。

#### 2.3.1 pkt_number 的设计

pkt_number 从流程图看起来，和 TCP 的 seq 字段比较类似，然而实际上还是有不少差别，可以说，pkt_number 的设计就是为了解决前面提到的 TCP 的问题的，我们来看看 pkt_number 的设计：

1. **从 0 开始的下标**：保证数据的有序
2. **永远自增的 pkt_number**：这里的永远自增指的是 pkt_number 的明文随每个 QUIC packet 发送，都会自增 1。pkt_number 的自增解决的是二义性问题，接收端收到 pkt_number 对应的 ACK 之后，可以明确的知道到底是重传的报文被 ACK 了，还是新的报文被 ACK 了，这样 RTT 的估计及丢包检测，就可以更加精细。同时quic还是用offset保证了数据的有序。

3. **加密 pkt_number 以保障安全**：当然 pkt_number 从 0 开始技术便也就遇上了和 TCP一样的安全问题，解决方案也很简单，就是用为 pkt_number 加密，pkt_number 加密后，中间人便无法获取到加密的 pkt_number 的 key，便也无法获取到真实的 pkt_number，也就无法通过观测 pkt_number 来预测后续的数据发送。

TLS 严格来说并不是一个状态严格递进的协议，每进入一个新的状态，还是有可能会收到上一个状态的数据，这么说有点抽象。

举个例子，TLS1.3 引入了一个 0-RTT 的技术，该技术允许 client 在通过 clientHello 发起 TLS 请求时，同时发送一些应用层数。我们当然期望这个应用层数据的过程相对于握手过程来说是异步且互不干扰的，而如果他们都是用同一个 pkt_number 来标示，那么应用层数据的丢包势必会导致对握手过程的影响。所以，QUIC 针对握手的状态设计了三种不同的 pkt_number space：（它们是三种不同类型的quic包，但可以放在一个UDP报文中传输）

(1) init;

(2) Handshake;

(3) Application Data。

分别对应：

(1) TLS 握手中的明文数据传输，即图中的 init packet;

(2) TLS 中通过 traffic secret 加密的握手数据传输，即图中 handshake packet;

(3) 握手完成后的应用层数据传输及 0-RTT 数据传输，即图中的 short header packet 以及图中暂未画出的 0-RTT packet。

三种不同的 space 保证了三个流程的丢包检测互不影响。

#### 2.3.2 基于 stream 的有序传输

我们知道 QUIC 是基于 UDP 实现的协议，而 UDP 是不可靠的面向报文的协议，这和 TCP 基于 IP 层的实现并没有什么本质上的不同，都是：

(1) 底层只负责尽力而为的，以 packet 为单位的传输;

(2) 上层协议实现更关键的特性，如可靠，有序，安全等。

从前面我们知道 TCP 的设计导致了链接维度的队头阻塞问题，而仅仅依靠 pkt_number 也无法实现数据的有序，所以 QUIC 必须要一种更细粒度的机制来解决这些问题：

1. **流既是一种抽象，也是一种单位**

TCP 队头阻塞的根因来自于其一条链接只有一个发送流，流上任意一个 packet 的阻塞都会导致其他数据的失败。当然解决方案也不复杂，我们只需要在一条 QUIC 链接上抽象出多个流来即可，整体的思路如下图：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*Zwt4QYKc0OgAAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:50%;" />

只要能保证各个 stream 的发送独立，那么我们实际上就避免了 QUIC 链接本身的队头阻塞，即一个流阻塞我们也可以在其他流上发送数据。

有了单链接多流的抽象，我们再来看 QUIC 的传输有序性设计，实际上 QUIC 在 stream 层面之上，还有更细粒度的单位，称作 frame。一个承载数据的 frame 携带有一个 offset 字段，表明自己相对于初始数据的偏移量。而初始偏移量为 0，这个思路等价于 pkt_number 等于 0，即不需要握手即可开始数据的发送。熟悉 HTTP/2、GRPC 的同学应该比较清楚这个 offset 字段的设计，和流式数据的传输是一样的。

2. 一个 TLS 握手也是一个流

虽然 TLS 数据并没有一个固定的 stream 来标示，但其可以被看作为一个特定的 stream，或者说是所有其他 stream 能建立起来的初始 stream，因为它其实也是基于 offset 字段和固定的 frame 来承载的，这也就是 TLS 数据有序的保障所在。

3. 基于 frame 的控制

有了 frame 这一层的抽象之后，我们当然可以做更多的事情。除了承载实际数据之外，也可以承载一些控制数据，QUIC 的设计汲取了 TCP 的经验教训，对于 keepalive、ACK、stream 的操作等行为，都设置了专门的控制 frame，这也就实现了数据和控制的完全解耦，同时保证了基于 stream 的细粒度控制，让 QUIC 成为更精细化的传输层协议。



讲到这里，其实可以看到我们从 QUIC 建联流程的探讨中，已经明确了 QUIC 设计的目标，正如文章中一直在强调的概念：

“无论是建联还是什么流程，都是在为实现 QUIC 的特性而服务”。

我们现在对 QUIC 的特性及实现已经有了一些认知，来小结一下：

![图片](https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*cKsLTJsZXOAAAAAAAAAAAAAAARQnAQ)

### 2.3、连接关闭

#### 2.3.1 从 TCP 看 QUIC 链接的优雅关闭

链接关闭是一个简单的诉求，可以简单梳理为两个目标：

1. 用户可以**主动优雅关闭**链接，并能通知对方，释放资源。

2. 当通信一端无法建立的链接时，有一个通知对方的机制。

然而诉求很简单，TCP 的实现却很不简单，我们来看看这个 TCP 链接关闭流程状态机的转移图：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*JPQGQJmhX50AAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:80%;" />

这个流程看起来就够复杂了，牵扯出来的问题就更不少，比如经典面试题：

- “为什么需要 TIME_WAIT？”

- “TIME_WAIT 的连接数过多需要怎么处理？”

- “tcp_tw_reuse 和 tcp_tw_recycle 等内核参数的作用和区别？”

而这**一切问题的根因都来源于 TCP 的链接和流的绑定**，或者说是控制信令和数据通道的耦合。

我们不禁要提出一个灵魂拷问“我们需要的是全双工的数据传输模式，但我们真的需要在链接维度做这个事情吗？”这么说似乎有一些抽象，还是以 TCP 的 TIME_WAIT 设计为例子来说明，我们来自底向上的看看 TCP 的问题：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*vbBnRZjSfw0AAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:67%;" />

再回到我们的问题上，如果我们将流和链接区分开，在链接维度保证流的控制指令可靠传输，链接本身实现一个简单的单工关闭过程，即通信一端主动关闭链接，则整个链接关闭，是否一切就简单起来了呢？

当然这就是 QUIC 的解法，有了这一层思路之后，我们来整理下 QUIC 链接关闭的诉求：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*0ExjRIEVUwQAAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:80%;" />

抛开流的关闭流程设计，在链接维度我们就可以得到一个清爽的状态机：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*R2iKTbcGnUoAAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:67%;" />

可以看到，得益于单工的关闭模式，在整个 QUIC 链接关闭的流程里，关闭指令只有一个，即图中的` CONNECTION_CLOSE`，而关闭的状态也只有两个，即图中的 closing 和 draing。我们先来看看两种状态下终端的行为：

- `closing `状态：当用户主动关闭链接时，即进入该状态，该状态下，终端收到所有的应用层数据都将只会回复 `	CONNECTION_CLOSE`
- `draining `状态：当终端收到`CONNECTION_CLOSE` 时，即进入该状态，该状态下终端不再回复任何数据

更简单的是，**`CONNECTION_CLOSE` 是一个不需要被 ACK 的指令**，也就意味着不需要重传。因为从链接维度而言，我们只需要保证最后能成功关闭的链接，并且新的链接不被老的关闭指令影响即可，这种简单的 `CONNECTION_CLOSE` 指令就能实现所有的诉求。

#### 2.3.2 更安全的 reset 的方式

当然，链接关闭也分为多种情况，和 TCP 一样，除了上一节提到的 QUIC 一端主动关闭链接的模式，QUIC还需要具有提供无法响应时，直接reset对端连接的能力。

而 QUIC reset 对端链接的方式相比于 TCP 来说更加安全，该机制被称作 stateless reset，这并不是一个十分复杂的机制。在 QUIC 链接建立好之后，QUIC 双方会同步一个 token，而后续的链接关闭将通过校验这个 token 来判断该对端是否有权限来 reset 这个链接，这种机制从根本上规避了前面提到的 TCP 被恶意 reset 的这种攻击模式，整个流程如下图：

<img src="https://gw.alipayobjects.com/mdn/rms_1c90e8/afts/img/A*81IfQa8sW50AAAAAAAAAAAAAARQnAQ" alt="图片" style="zoom:67%;" />

当然，stateless reset 的方案并不是万金油，安全的代价是更窄的使用范围。因为为了保障安全，token 必须通过安全的数据通道进行传输(在 QUIC 中被限定为 NEW_TOKEN frame)，并且接收端需要维持一个记录这个 token 的状态，也只有通过这个状态才可以保证 token 的有效性。

因此 stateless reset 被限定为 QUIC 链接关闭的最后手段，并且也只能在只能使用在客户端和服务端均处于一个相对正常的情况下正常工作，比如这样的情况 stateless reset 就不适用，服务端并没有监听在 443 端口，但客户端发送数据到 443 端口，而这种情况在 TCP 协议栈下是可以 RST 掉的。

#### 3.3 工程考量的超时断链

keepalive 机制本身并没有什么花样，都是一个计时器加探测报文即可搞定，而 QUIC 得益于链接和数据流的拆分，关闭链接变成了一个非常简单的事情，keepalive 也就变得更简单易用，QUIC 在链接维度上提供了名为 PING frame 的控制指令，用于主动探测保活。

更简单的是，QUIC 在超时后关闭链接的方式是 silently close，即不通知对端，本机直接释放掉链接所有的资源。silently close 的好处是资源可以立即得到释放，特别是对于 QUIC 这种单链接上既要维护 TLS 状态，也要维护流状态的协议来说，有很大的收益。

但其劣势在于后续如果有之前链接数据到来，则只能通过 stateless reset 的方式通知对端关闭链接，stateless reset 的关闭相对来说 CONNECTION_CLOSE 开销更大。因此这部分可以说完全是一个 tradeoff，而这部分的设计方案的最终敲定，更多来自于大量工程实践的经验与结果。



## **3、QUIC 实现原理**

### **3.1、数据格式**

一个 QUIC 数据包的格式如下：

![img](https://pic2.zhimg.com/80/v2-60231adb6c7014c7f043712839f77ab5_720w.jpg)

由` header` 和` data `两部分组成。

`header` 是明文的，包含 4 个字段：`Flags`，`Connection ID`，`QUIC Version`，`Packet Number`。

其中`Connection ID` 在C和S处是不一样的。`Packet Number`是严格递增的。

`data` 是加密的，可以包含 1 个或多个` frame`，每个` frame` 又分为` type` 和` payload`，其中 `payload` 就是应用数据(以及控制信息)；



数据帧有很多类型：`Stream、ACK、Padding、Window_Update、Blocked` 等，这里重点介绍下用于传输应用数据的 `Stream` 帧。

![img](https://pic2.zhimg.com/80/v2-f1cb88ac186e851a724a85dbd4f3de01_720w.jpg)

**Frame Type：** 帧类型，占用 1 个字节

![img](https://pic3.zhimg.com/80/v2-eb52f6d7f4d12a2599085f62f9240162_720w.jpg)

（1）*<u>Bit7</u>*：必须设置为 1，表示` Stream `帧

（2）<u>*Bit6*</u>：如果设置为 1，表示发送端在这个` stream `上已经结束发送数据，流将处于半关闭状态

（3）<u>*Bit5*</u>：如果设置为 1，表示 `Stream `头中包含` Data length `字段

（4）<u>*Bit432*</u>：表示 offset 的长度。000 表示 0 字节，001 表示 2 字节，010 表示 3 字节，以此类推

（5）<u>*Bit10*</u>：表示 `Stream ID` 的长度。00 表示 1 字节，01 表示 2 字节，10 表示 3 字节，11 表示 4 字节

**Stream ID：** 流 ID，用于标识数据包所属的流。后面的流量控制和多路复用会涉及到

**Offset：**偏移量，表示该数据包在整个数据中的偏移量，用于数据排序。重传包的Offset和之前包的Offset是一样的，解决了重传歧义问题。 

**Data Length：** 数据长度，占用 2 个字节，表示实际应用数据的长度，16个bit，最长为2^16-1

**Data：** 实际的应用数据

### 3.2、建立连接

先分析下 HTTPS 的握手过程，包含 TCP 握手和 TLS 握手，

TCP 握手：

<img src="https://pic2.zhimg.com/80/v2-1fb94488942494ec9425ecf6682ed6e1_720w.jpg" alt="img" style="zoom:50%;" />

从图中可以看出，TCP 握手需要 2 个 RTT。(TCP的第三次握手时就可以带上数据了)

> ***<u>TCP为什么需要三个握手四次挥手？</u>***
>
> **对于建链接的3次握手，**主要是要<span style="color:#ee1122;font-size:16px">
> 初始化Sequence Number 的值</span>。通信的双方要互相通知对方自己的初始化的Sequence Number，（双方用的sequence number当然是不一样的，比如C端用x，S端用y）。——所以叫`SYN`，全称`Synchronize Sequence Numbers`。这个号要作为以后的数据通信的序号，以保证应用层接收到的数据不会因为网络上的传输的问题而乱序（TCP会用这个序号来拼接数据）。这样就保证的TCP的可靠传输。
>
> **对于4次挥手，**因为TCP是全双工的，所以，发送方和接收方都需要`FIN`和`ACK`。C端发送`FIN`给S端，S端恢复`ACK`，此时连接处于“半关闭”状态。只有C端到S端的连接关闭，S端还是可以给C端发信息。C端可以接收。然后S端给C端发`FIN`，C端收到后回复`ACK`后要等待两个`MSL(最长报文段交付时间)`后连接才算全部关闭。等这些个时间是因为C端最后回复的`ACK`可能丢了，此时S端超时重传`FIN`，C端必须能接收此重传的`FIN`，不然S端就无法正常关闭。

TLS 握手：密钥协商（1.3 版本）

<img src="https://pic4.zhimg.com/80/v2-86cc5e2cf1b509083759b9c22800e2a7_720w.jpg" alt="img" style="zoom:50%;" />

TLS 握手是启动 TLS 加密通信会话的过程的描述。在 TLS 握手过程中(使用Diffie-Hellman密钥交换算法)，通信双方要交换消息以相互确认，彼此验证，**确立它们将使用的 TLS 版本和密码套件(一组加密算法)，并生成一致的会话密钥**。

QUIC可以使用标准的TLS1.3加密协议



## QUIC AND TLS1.3

![image-20220713172415854](C:\Users\a\AppData\Roaming\Typora\typora-user-images\image-20220713172415854.png)

![image-20220713172426370](C:\Users\a\AppData\Roaming\Typora\typora-user-images\image-20220713172426370.png)



![image-20220713172436225](C:\Users\a\AppData\Roaming\Typora\typora-user-images\image-20220713172436225.png)

![image-20220713172507006](C:\Users\a\AppData\Roaming\Typora\typora-user-images\image-20220713172507006.png)

> STK : Source Token



## QUIC不被看好？

它没有真正重新审视传输协议的本质：<u>***仍然高度依赖于确认***</u>，这不适合无线世界的性能挑战。

QUIC专注于握手优化（非常重要！）但未能真正重新审视传输协议的本质（传输！）。为什么这很重要？因为传统的传输技术是在有线用户世界中定义的。这个世界早已不复存在，不可靠和不可预测的无线链路主宰着用户体验。还有其他选择吗？是：**擦除编码（erasure coding），而不是ARQ。**

简而言之，QUIC取代了TCP和TLS的组合，采用了跨层的传输和安全方法。在QUIC下，UDP用作“传输”。为什么选择UDP？一个非常实用的决定：使用UDP可以在用户空间实现非常快速的部署，而修改TCP需要很长时间才能被采用。

[QUIC vs TCP+TLS and why QUIC is not the next big thing (codavel.com)](https://blog.codavel.com/quic-vs-tcptls-and-why-quic-is-not-the-next-big-thing)







- **QUIC 核心特性连接建立延时低**

QUIC建立连接时，加密算法的版本协商与传输层握手合并完成，以减小延迟。

在连接上实际传输数据时需要建立并使用一个或多个数据流。

- **改进的拥塞控制** 

    TCP 的拥塞控制实际上包含了四个算法：<u>慢启动，拥塞避免，快速重传，快速恢复</u>。

    QUIC 协议当前默认使用了 TCP 协议的 *<u>**Cubic 拥塞控制算法**</u>* [6]，同时也支持 *<u>CubicBytes, Reno, RenoBytes, BBR, PCC</u>* 等拥塞控制算法。

- **可插拔**

    能够非常灵活地生效，变更和停止。体现在如下方面：

    1. 应用程序层面就能实现不同的拥塞控制算法，不需要操作系统，不需要内核支持。这是一个飞跃，因为传统的 TCP 拥塞控制，必须要端到端的网络协议栈支持，才能实现控制效果。而内核和操作系统的部署成本非常高，升级周期很长，这在产品快速迭代，网络爆炸式增长的今天，显然有点满足不了需求。
    2. 即使是单个应用程序的不同连接也能支持配置不同的拥塞控制。就算是一台服务器，接入的用户网络环境也千差万别，结合大数据及人工智能处理，我们能为各个用户提供不同的但又更加精准更加有效的拥塞控制。比如 BBR 适合，Cubic 适合。
    3. 应用程序不需要停机和升级就能实现拥塞控制的变更，我们在服务端只需要修改一下配置，reload 一下，完全不需要停止服务就能实现拥塞控制的切换。

- **单调递增的 Packet Number**

    TCP 为了保证可靠性，使用了基于字节序号的 Sequence Number 及 Ack 来确认消息的有序到达。

    QUIC 同样是一个可靠的协议，它使用 Packet Number 代替了 TCP 的 sequence number，并且每个 Packet Number 都严格递增，也就是说就算 Packet N 丢失了，重传的 Packet N 的 Packet Number 已经不是 N，而是一个比 N 大的值。而 TCP 呢，重传 segment 的 sequence number 和原始的 segment 的 Sequence Number 保持不变，也正是由于这个特性，引入了 Tcp 重传的歧义问题。

    ![img](https://pic2.zhimg.com/80/v2-8db4c3c378edaac0060b4238e3554091_720w.jpg)

    如上图所示，超时事件 RTO 发生后，客户端发起重传，然后接收到了 Ack 数据。由于序列号一样，这个 Ack 数据到底是原始请求的响应还是重传请求的响应呢？不好判断。

    如果算成原始请求的响应，但实际上是重传请求的响应（上图左），会导致采样 RTT 变大。如果算成重传请求的响应，但实际上是原始请求的响应，又很容易导致采样 RTT 过小。

    **<u>*那么TCP是如何解决重传歧义问题的呢？*</u>**

    > Karn提出的算法：TCP在计算加权平均往返时延RTTs时，只要报文重传了，就不采用此次的往返时间样本了。这样子RTTs和超时重传时间RTO不会有太大波动。
    >
    > 但这样了导致的新问题是，当链路状态恶化时无法及时更新RTO导致经常发生重传，因此Karn算法进行修正，每发生一次重传就让RTO增大一些，典型做法是将RTO翻倍。

    **<u>*QUIC是如何解决的*</u>**

    由于 Quic 重传的 Packet 和原始 Packet 的 Pakcet Number 是严格递增的，所以很容易就解决了这个问题。

    ![img](https://pic2.zhimg.com/80/v2-086cc0ac3b95eb5bfe84a2d87bb2a645_720w.jpg)

    如上图所示，RTO 发生后，根据重传的 Packet Number 就能确定精确的 RTT 计算。如果 Ack 的 Packet Number 是 N+M，就根据重传请求计算采样 RTT。如果 Ack 的 Pakcet Number 是 N，就根据原始请求的时间计算采样 RTT，没有歧义性。

    但是单纯依靠严格递增的 Packet Number 肯定是无法保证数据的顺序性和可靠性。QUIC 又引入了一个 Stream Offset 的概念。

    即一个 Stream 可以经过多个 Packet 传输，Packet Number 严格递增，没有依赖。但是 Packet 里的 Payload 如果是 Stream 的话，就需要依靠 Stream 的 Offset 来保证应用数据的顺序。如图所示，发送端先后发送了 Pakcet N 和 Pakcet N+1，Stream 的 Offset 分别是 x 和 x+y。

    假设 Packet N 丢失了，发起重传，重传的 Packet Number 是 N+2，但是它的 Stream 的 Offset 依然是 x，这样就算 Packet N + 2 是后到的，依然可以将 Stream x 和 Stream x+y 按照顺序组织起来，交给应用程序处理。

    ![img](https://pic2.zhimg.com/80/v2-60985053d9de4e8e74042c33587ec35d_720w.jpg)

- **不允许 Reneging **？

    什么叫 Reneging 呢？就是接收方丢弃已经接收并且上报给 SACK 选项的内容 [8]。TCP 协议不鼓励这种行为，但是协议层面允许这样的行为。主要是考虑到服务器资源有限，比如 Buffer 溢出，内存不够等情况。

    Reneging 对数据重传会产生很大的干扰。因为 Sack 都已经表明接收到了，但是接收端事实上丢弃了该数据。

    QUIC 在协议层面禁止 Reneging，一个 Packet 只要被 Ack，就认为它一定被正确接收，减少了这种干扰。

- **更多的 Ack 块**

    TCP 的 Sack（Selective ack） 选项能够告诉发送方已经接收到的连续 Segment  的范围，方便发送方进行选择性重传。

    由于 TCP 头部最大只有 60 个字节，标准头部占用了 20 字节，所以 Tcp Option 最大长度只有 40 字节，再加上 Tcp Timestamp option 占用了 10 个字节 [25]，所以留给 Sack 选项的只有 30 个字节。

    每一个 Sack Block 的长度是 8 个，加上 Sack Option 头部 2 个字节，也就意味着 Tcp Sack Option 最大只能提供 3 个 Block。

    但是 Quic Ack Frame 可以同时提供 256 个 Ack Block，在丢包率比较高的网络下，更多的 Sack Block 可以提升网络的恢复速度，减少重传量。

- **Ack Delay 时间**

    Tcp 的 Timestamp 选项存在一个问题 [25]，它只是回显了发送方的时间戳，但是没有计算接收端接收到 segment 到发送 Ack 该 segment 的时间。这个时间可以简称为 Ack Delay。

    这样就会导致 RTT 计算误差。如下图：

    ![img](https://pic3.zhimg.com/80/v2-5466d84601e2fe87de92f06daba2f88e_720w.jpg)

    可以认为 TCP 的 RTT 计算：

    ![img](https://pic1.zhimg.com/80/v2-49198a21c42de7615246ab7452cf51ac_720w.jpg)

    而 Quic 计算如下：

    ![img](https://pic3.zhimg.com/80/v2-8cba385d551d72b867f8c437a34b9aba_720w.jpg)

    当然 RTT 的具体计算没有这么简单，需要采样，参考历史数值进行平滑计算，参考如下公式 [9]。

    ![img](https://pic4.zhimg.com/80/v2-04c5bd39d745430fb858ec3458e81003_720w.jpg)

- **基于 stream 和 connecton 级别的流量控制（双级别流控）**

QUIC 的流量控制 [22] 类似 HTTP2，即在 Connection 和 Stream 级别提供了两种流量控制。为什么需要两类流量控制呢？主要是因为 QUIC 支持多路复用。

1. Stream 可以认为就是一条 HTTP 请求。
2. Connection 可以类比一条 TCP 连接。多路复用意味着在一条 Connetion 上会同时存在多条 Stream。既需要对单个 Stream 进行控制，又需要针对所有 Stream 进行总体控制。

QUIC 实现流量控制的原理比较简单：

通过` window_update` 帧告诉对端自己可以接收的字节数，这样发送方就不会发送超过这个数量的数据。

通过 `BlockFrame `告诉对端由于流量控制被阻塞了，无法发送数据。

QUIC 的流量控制和 TCP 有点区别，TCP 为了保证可靠性，窗口左边沿向右滑动时的长度取决于已经确认的字节数。如果中间出现丢包，就算接收到了更大序号的 Segment，窗口也无法超过这个序列号。

但 QUIC 不同，就算此前有些 packet 没有接收到，它的滑动只取决于接收到的最大偏移字节数。

![图 5 Quic Flow Control](https://pic2.zhimg.com/80/v2-5f2ad22131453b3fd4c6351094b6b581_720w.jpg)

针对 Stream：

![img](https://pic1.zhimg.com/80/v2-536b207b20e3ddc1bbeac1942fae0214_720w.jpg)

针对 Connection：

![img](https://pic3.zhimg.com/80/v2-7465703d69fbcc169c8f397b377e67a2_720w.jpg)

同样地，STGW 也在连接和 Stream 级别设置了不同的窗口数。

最重要的是，我们可以在内存不足或者上游处理性能出现问题时，通过流量控制来限制传输速率，保障服务可用性

- **没有队头阻塞的多路复用**

QUIC 的多路复用和 HTTP2 类似。在一条 QUIC 连接上可以并发发送多个 HTTP 请求 (stream)。但是 QUIC 的多路复用相比 HTTP2 有一个很大的优势。

QUIC 一个连接上的多个 stream 之间没有依赖。这样假如 stream2 丢了一个 udp packet，也只会影响 stream2 的处理。不会影响 stream2 之前及之后的 stream 的处理。

这也就在很大程度上缓解甚至消除了队头阻塞的影响。

多路复用是 HTTP2 最强大的特性 [7]，能够将多条请求在一条 TCP 连接上同时发出去。但也恶化了 TCP 的一个问题，队头阻塞 [11]，如下图示：

![图 6 HTTP2 队头阻塞](https://pic2.zhimg.com/80/v2-2dd2a9fb8693489b9a0b24771c8a40a1_720w.jpg)

HTTP2 在一个 TCP 连接上同时发送 4 个 Stream。其中 Stream1 已经正确到达，并被应用层读取。但是 Stream2 的第三个 tcp segment 丢失了，TCP 为了保证数据的可靠性，需要发送端重传第 3 个 segment 才能通知应用层读取接下去的数据，虽然这个时候 Stream3 和 Stream4 的全部数据已经到达了接收端，但都被阻塞住了

不仅如此，由于 HTTP2 强制使用 TLS，还存在一个 TLS 协议层面的队头阻塞 [12]。

![图 7 TLS 队头阻塞](https://pic3.zhimg.com/80/v2-f1c2dcdb8f3cb56c260f408420cea502_720w.jpg)

Record 是 TLS 协议处理的最小单位，最大不能超过 16K，一些服务器比如 Nginx 默认的大小就是 16K。由于一个 record 必须经过数据一致性校验才能进行加解密，所以一个 16K 的 record，就算丢了一个字节，也会导致已经接收到的 15.99K 数据无法处理，因为它不完整。

那 QUIC 多路复用为什么能避免上述问题呢？

1. QUIC 最基本的传输单元是 Packet，不会超过 MTU 的大小，整个加密和认证过程都是基于 Packet 的，不会跨越多个 Packet。这样就能避免 TLS 协议存在的队头阻塞。
2. Stream 之间相互独立，比如 Stream2 丢了一个 Pakcet，不会影响 Stream3 和 Stream4。不存在 TCP 队头阻塞。

![图 8 QUIC 多路复用时没有队头阻](https://pic4.zhimg.com/80/v2-9e649330ab729b6438a8586c8b4f1bd3_720w.jpg)塞的问题

当然，并不是所有的 QUIC 数据都不会受到队头阻塞的影响，比如 QUIC 当前也是使用 Hpack 压缩算法 [10]，由于算法的限制，丢失一个头部数据时，可能遇到队头阻塞。

总体来说，QUIC 在传输大量数据时，比如视频，受到队头阻塞的影响很小，因为一个视频内容会被分为多个stream传输。

- **加密认证的报文**

TCP 协议头部没有经过任何加密和认证，所以在传输过程中很容易被中间网络设备篡改，注入和窃听。比如修改序列号、滑动窗口。这些行为有可能是出于性能优化，也有可能是主动攻击。

但是 QUIC 的 packet 可以说是武装到了牙齿。除了个别报文比如 PUBLIC_RESET 和 CHLO，所有报文头部都是经过认证的，报文 Body 都是经过加密的。

这样只要对 QUIC 报文任何修改，接收端都能够及时发现，有效地降低了安全风险。

如下图所示，红色部分是 Stream Frame 的报文头部，有认证。绿色部分是报文内容，全部经过加密。

![img](https://pic1.zhimg.com/80/v2-04f12b295aae3fb44b490b852e5c1e44_720w.jpg)

- **连接迁移**

一条 TCP 连接 [17] 是由四元组标识的（源 IP，源端口，目的 IP，目的端口）。什么叫连接迁移呢？就是当其中任何一个元素发生变化时，这条连接依然维持着，能够保持业务逻辑不中断。当然这里面主要关注的是客户端的变化，因为客户端不可控并且网络环境经常发生变化，而服务端的 IP 和端口一般都是固定的。

比如大家使用手机在 WIFI 和 4G 移动网络切换时，客户端的 IP 肯定会发生变化，需要重新建立和服务端的 TCP 连接。

又比如大家使用公共 NAT 出口时，有些连接竞争时需要重新绑定端口，导致客户端的端口发生变化，同样需要重新建立 TCP 连接。

针对 TCP 的连接变化，MPTCP[5] 其实已经有了解决方案，但是由于 MPTCP 需要操作系统及网络协议栈支持，部署阻力非常大，目前并不适用。

所以从 TCP 连接的角度来讲，这个问题是无解的。

那 QUIC 是如何做到连接迁移呢？很简单，任何一条 QUIC 连接不再以 IP 及端口四元组标识，而是以一个 64 位的随机数作为 ID 来标识，这样就算 IP 或者端口发生变化时，只要 ID 不变，这条连接依然维持着，上层业务逻辑感知不到变化，不会中断，也就不需要重连。

由于这个 ID 是客户端随机产生的，并且长度有 64 位，所以冲突概率非常低。









