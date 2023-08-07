# linux内核CC算法学习

源码下载：[Index of /pub/linux/kernel/](https://mirrors.edge.kernel.org/pub/linux/kernel/) ，本文以 5.4.224 版本为例讲解。

在新下载的内核中，网络相关代码的位置：`usr/src/linux-x.x.xxx/net/`

在`tcp_ipv4.c`中，`tcp_prot`定义了tcp的各个接口。(如下所示)

``` c
struct proto tcp_prot = {
  .name = "TCP",
  .owner = THIS_MODULE,
  .close = tcp_close,
  .connect = tcp_v4_connect,
  .disconnect = tcp_disconnect,
  .accept = inet_csk_accept,
  .destroy = tcp_v4_destroy_sock,
  .shutdown = tcp_shutdown,
  .setsockopt = tcp_setsockopt,
  .getsockopt = tcp_getsockopt,
  .recvmsg = tcp_recvmsg,
  .sendmsg = tcp_sendmsg,
  .sendpage = tcp_sendpage,
  .backlog_rcv = tcp_v4_do_rcv,
  .get_port = inet_csk_get_port,
  .twsk_prot = &tcp_timewait_sock_ops,
  .rsk_prot = &tcp_request_sock_ops,
};
```

所有TCP拥塞控制算法都需要实现`tcp_congestion_ops`这样的结构，其定义位于 `/usr/src/linux-x.x.xxx/include/net/tcp.h`中。每一种拥塞控制算法实现这样的结构，然后向系统注册。系统将注册的拥塞控制算法组织成一个**单链表**。

可以在已有的若干拥塞控制算法中看到它们各自的实现，如：(`/usr/src/linux-x.x.xxx/include/ipv4/xxx.c`)

``` c
static struct tcp_congestion_ops bictcp __read_mostly = {	// tcp_bic.c 
	.init		= bictcp_init,
	.ssthresh	= bictcp_recalc_ssthresh,
	.cong_avoid	= bictcp_cong_avoid,
	.set_state	= bictcp_state,
	.undo_cwnd	= tcp_reno_undo_cwnd,
	.pkts_acked     = bictcp_acked,
	.owner		= THIS_MODULE,
	.name		= "bic",
};

struct tcp_congestion_ops tcp_reno = {  // tcp_cong.c
	.flags		= TCP_CONG_NON_RESTRICTED,
	.name		= "reno",
	.owner		= THIS_MODULE,
	.ssthresh	= tcp_reno_ssthresh,
	.cong_avoid	= tcp_reno_cong_avoid,
	.undo_cwnd	= tcp_reno_undo_cwnd,
};

static struct tcp_congestion_ops tcp_bbr_cong_ops __read_mostly = { // tcp_bbr.c
	.flags		= TCP_CONG_NON_RESTRICTED,
	.name		= "bbr",
	.owner		= THIS_MODULE,
	.init		= bbr_init,
	.cong_control	= bbr_main,
	.sndbuf_expand	= bbr_sndbuf_expand,
	.undo_cwnd	= bbr_undo_cwnd,
	.cwnd_event	= bbr_cwnd_event,
	.ssthresh	= bbr_ssthresh,
	.min_tso_segs	= bbr_min_tso_segs,
	.get_info	= bbr_get_info,
	.set_state	= bbr_set_state,
};
```

TCP头部字段`tcphdr`定义位置：`usr/src/linux-x.x.xxx/include/uapi/tcp.h`

机器收到数据包后的处理过程的主程序在`usr/src/linux-x.x.xxx/net/ipv4/ip_input.c/*ip_rcv_core`函数中

在  `usr/src/linux-x.x.xxx/net/ipv4/tcp_input.c/tcp_rcv_state_process()`中，6306行，将发送窗口改为接收到的包的cwnd。（原本就有的）



## 网络核心数据结构

### sock 底层数据结构

#### `sock`

- 位置：`/include/net/sock.h`

- sock 结构是比较通用的网络层描述块，构成传输控制块的基础，与具体的协议族无关。它描述了各协议族的公共信息，因此不能直接作为传输层控制块来使用。不同协议族的传输层在使用该结构的时候都会对其进行拓展，来适合各自的传输特性。例如，**`inet_sock`结构由sock 结构及其它特性组成，构成了IPV4 协议族传输控制块的基础。**

#### `sk_buff`

- 位置：`/include/linux/skbuff.h`

- `struct sk_buff - socket buffer`:这一结构体在各层协议中都会被用到。该结构体存储了**网络数据报的所有信息**。包括各层的头部以及payload，以及必要的各层实现相关的信息。在文件中对skb的属性进行了详细解释。

### inet 层相关数据结构

#### `ip_options`

- 位置：`/include/net/inet_sock.h`

- 定义ip头的可选项

#### `inet_connection_sock`

- 位置：`/include/net/inet_connection_sock.h`

- 它是所有面向传输控制块的表示。其在`inet_sock`的基础上，增加了有关连接，确认，重传等成员

### TCP 层相关数据结构

#### `tcphdr`

- 位置：`/include/uapi/tcp.h`

- 这一结构正是TCP 首部及各个字段定义。

#### `tcp_options_received`

- 位置：`/include/linux/tcp.h`

- TCP 头部的选项字段。

#### `tcp_sock` 

- 位置：`/include/linux/tcp.h`

- 该数据结构是TCP 协议的控制块，它在`inet_connection_sock`结构的基础上扩展了滑动窗口协议、拥塞控制算法等一些TCP 的专有属性



## TCP Option

TCP Option 解释：[常用的TCP Option_blakegao的博客-CSDN博客](https://blog.csdn.net/blakegao/article/details/19419237)

TCP选项字段构造：[TCP/IP OPTION字段 - 莫扎特的代码 - 博客园 (cnblogs.com)](https://www.cnblogs.com/codingMozart/p/9067552.html)

解析收到的数据包的option字段的位置在：`net/ipv4/tcp_input.c/tcp_parse_options`

一般Option的格式为TLV结构，如下所示：

| Kind / Type（1 Byte） | Length（1 Byte）[整个这个选项段的总长度] | Value |
| --------------------- | ---------------------------------------- | ----- |

当前，TCP常用的Option如下所示：

| Kind(Type) | Length | Name                                                         | Reference | 描述 & 用途                                |
| ---------- | ------ | ------------------------------------------------------------ | --------- | ------------------------------------------ |
| 0          | 1      | EOL                                                          | RFC 793   | 选项列表结束                               |
| 1          | 1      | NOP                                                          | RFC 793   | 无操作（用于补位填充）                     |
| 2          | 4      | MSS                                                          | RFC 793   | 最大Segment长度                            |
| 3          | 3      | WSOPT                                                        | RFC 1323  | 窗口扩大系数（Window Scaling Factor）      |
| 4          | 2      | SACK-Premitted                                               | RFC 2018  | 表明支持SACK                               |
| 5          | 可变   | SACK                                                         | RFC 2018  | SACK Block（收到乱序数据）                 |
| 8          | 10     | TSPOT                                                        | RFC 1323  | Timestamps                                 |
| 19         | 18     | TCP-MD5                                                      | RFC 2385  | MD5认证                                    |
| 28         | 4      | UTO（RFC 5482: TCP option to announce/request UTO Not yet implemented in Linux） | RFC 5482  | User Timeout（超过一定闲置时间后拆除连接） |
| 29         | 可变   | TCP-AO                                                       | RFC 5925  | 认证（可选用各种算法）                     |
| 253/254    | 可变   | Experimental                                                 | RFC 4727  | 保留，用于科研实验                         |

注：

> 1.    EOL和NOP Option（Kind 0、Kind 1）只占1 Byte，没有Length和Value字段；
> 2.    NOP用于 将TCP Header的长度补齐至32bit的倍数（由于Header Length字段以32bit为单位，因此TCP Header的长度一定是32bit的倍数）；
>3.    SACK-Premitted Option占2 Byte，没有Value字段；
> 4.    其余Option都以1 Byte的“Kind”开头，指明Option的类型；Length指明Option的总长度（包括Kind和Length）
>5.    对于收到“不能理解”的Option，TCP会无视掉，并不影响该TCP Segment的其它内容；
> 5.    自己新增可选项字段时，需要修改IP头中的总长度字段，TCP头中的offset字段，并且修改TCP的校验和。如果知识修改字段的话就需要重新计算TCP的校验和。



## 简称查询

| **简称**    | **含义**                                                     |
| ----------- | ------------------------------------------------------------ |
| **ca**      | congestion control algorithm 拥塞控制算法                    |
| tp          | tcp_sock                                                     |
| th          | tcphdr  , tcp 头部字段                                       |
| icsk        | inet_connection_sock                                         |
| icsk_ca_ops | inet_connection_sock congestion control algorithm options，当前网络连接所采用的拥塞控制算法。 |
| snd_cwnd    | 拥塞窗口                                                     |
| snd_wnd     | 当前发送窗口                                                 |
| snd_nxt     | 下一个发送序号                                               |
| snd_una     | 最早的未确认序号                                             |
| ihl         | ip header length  ： ip头部长度                              |
| TCPOLEN     | TCP OPTION LENGTH                                            |
|             |                                                              |

上面的snd_wnd、snd_una、snd_nxt三个字段组成了滑动窗口（已发送还未确认的包in_flight，以及即将发送的包）。 如下图所示：

![image-20221115204043714](img/image-20221115204043714.png)



拥塞窗口`cwnd`：防止**过多的数据注入到网络中**导致网络发生拥塞。

发送窗口（滑动窗口）`swnd`：防止源端发送过去，收端却接收不过来。

接收窗口`rwnd`：放在ACK中的窗口字段的值，收端告诉源端自己还能接收多少。

发送窗口的值 **`swnd = min(cwnd, rwnd)`**



## lkseagle代码修改的部分

`/usr/src/linux-5.4.224/net/ipv4/tcp_cong.c/tcp_reno_cong_avoid()` 函数， // 都是我自己注释的， /\*官方自己的 注释\*/

下面这些都是新增的

``` c
	if(tp->rx_opt.my_wnd){
		tp->snd_cwnd = tp->rx_opt.my_wnd << tp->rx_opt.snd_wscale;
		printk("my_wnd is ok!");
	}
	else{
		tp->snd_cwnd=10;
	}
	printk("cwnd:%d,my_wnd:%d,snd_wscale:%d \n", tp->snd_cwnd, tp->rx_opt.my_wnd, tp->rx_opt.snd_wscale);
```



`/usr/src/linux-5.4.224/include/linux/tcp.h/tcp_options_received`  结构体，在最后增加了自定的my_wnd

`usr/src/linux-x.x.xxx/net/ipv4/tcp_input.c/tcp_parse_options()` 解析可选字段的函数中，为了使我们新增的my_wnd字段能够解析，3980-3983 增加了 if 里面的解析。同时还有这个函数中新增了7个 printk 打印输出。（ctrl + F 找到可以删除）
