# QUIC



- **QUIC 核心特性连接建立延时低**

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

    Karn提出的算法：TCP在计算加权平均往返时延RTTs时，只要报文重传了，就不采用此次的往返时间样本了。这样子RTTs和超时重传时间RTO不会有太大波动。

    但这样了导致的新问题是，当链路状态恶化时无法及时更新RTO导致经常发生重传，因此Karn算法进行修正，每发生一次重传就让RTO增大一些，典型做法是将RTO翻倍。

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

















TCP为什么需要三个握手四次挥手？

**对于建链接的3次握手，**主要是要初始化Sequence Number 的初始值。通信的双方要互相通知对方自己的初始化的Sequence Number，（双方用的sequence number当然是不一样的，比如C端用x，S端用y）。——所以叫SYN，全称Synchronize Sequence Numbers。这个号要作为以后的数据通信的序号，以保证应用层接收到的数据不会因为网络上的传输的问题而乱序（TCP会用这个序号来拼接数据）。

**对于4次挥手，**因为TCP是全双工的，所以，发送方和接收方都需要Fin和ACK。