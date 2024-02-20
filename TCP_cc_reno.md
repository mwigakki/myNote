## 一些需要了解的概念

### 内核skb/sk_buff 

Linux 内核 skb，是 struct sk_buff 数据结构的简称，skb可以说是内核网络子系统使用最多，也是最重要的数据结构。

![img](img/v2-983934d8b576772d7323033ae3001f48_720w.webp)

从图中可以看到，skb被用在设备驱动与TCP/IP协议栈之间的数据传输。Packet 流动方向 分为两个，一个是发送（TX or Transmit）方向；另一个是接受（RX or Receive）方向。**对于TX来说，传输层申请SKB，设备驱动释放SKB；对于RX来说，设备驱动申请SKB，网络协议栈TCP/IP Stack 负责释放之。**

#### **TX 方向**

1. 用户态应用程序基于socket系统调用接口，传送需要TX 的HTTP应用层数据（如 Ngnix产生）
2. 内核态socket 层读取用户态数据，并按照应用层的协议类型，将数据发送到对应的传输层（如HTTP对应TCP）
3. *传输层申请 skb 数据结构，并填充数据到skb，然后skb 会向下传送到网络层和链路层（链路层在内核对应网络设备层），继续添加 IP 和 MAC header到skb中*
4. *最后skb 到达网络设备驱动，其从skb 中获取到packet data 数据的虚拟地址（skb->data)，并映射出dma总线地址给 网卡进行DMA读取*

#### **RX 方向**

1. 如高速网卡收到从光纤上传入的数据包后，基于网卡内DMA控制器 和 PCIe 总线，将数据包送入设备驱动层
2. *设备驱动（Device Driver）收到数据包后，申请SKB，并将数据放入SKB 结构指向的数据空间（位于 skb结构 的head 和 end 指针之间）注：驱动放数据到skb 有两种方式，一种方式是dev_alloc_skb + memcpy；第二种是 page + build_skb 实现零拷贝，效率更高*
3. *驱动调用 NAPI Schedule 将RX数据送入TCP/IP Stack, 到达 IP 层则去除 SKB 指向数据的 IP Header，到达TCP 层去除 TCP Head*
4. *最后在内核基于IP等路由信息，查找到对应的socket，将数据基于socket 送入用户态应用程序*

#### SKB 常用成员

``` c
struct sk_buff {
    struct sk_buff		*next;     // frag_list 通过next串联起来（下文会介绍frag_list）

    unsigned int	len,       // 线性区和非线性区数据总长（下文会详细介绍）
				    data_len;  // 非线性区的数据总长，包括frags和fraglist

    __u16			queue_mapping;//网络设备层（L2）选择该packet的发送Queue ID
                                              // 该Queue ID通常会作为网络设备驱动发包的Ring ID

	__u8			ip_summed:2; /*如果协议栈没有算checksum，其被赋值为CHECKSUM_PARTIAL，
                                              * 指示硬件做checksum的无状态卸载（即硬件计算checksum）
                                              */
    __u8			encapsulation:1;/* 指示该skb指向的数据包是否为一个隧道Tunnel 包*/

	__be16			vlan_proto;  /* 当设备支持VLAN卸载时，packet L2层会没有VLAN，
                                             * VLAN protocol（1Q， 1AD)会填充在该字段，
                                             * cos 和 vlan id 填充在下一个字段
                                             */
	__u16			vlan_tci;
    __be16			protocol; // 协议类型，VLAN封包为 ETH_P_8021Q 或 ETH_P_8021AD，普通包为L3协议类型

	/* These elements must be at the end, see alloc_skb() for details.  */
	sk_buff_data_t		tail;  // skb 线性区有效数据的结尾位置  
	sk_buff_data_t		end;    // skb 线性区结束位置，从tail到end之间称为tailroom
	unsigned char		*head,  // skb 线性区开头指针，从head到data之间称为headroom
				*data;  // skb 线性区有效数据的起始地址

	refcount_t		users;  // skb 被引用次数，当引用计数为零时，即可释放
};
```



### inet_connection_sock 和 sock

`sock` 是 Socket（套接字）的缩写，它是 Linux 内核中表示网络通信端点的抽象概念。`sock` 结构体定义在 `include/net/sock.h` 头文件中，包含了通用的套接字信息，如协议族、类型、状态等。`inet_connection_sock` 是 `sock` 结构体的派生结构体，用于表示 TCP 或 UDP 连接的套接字的状态。它定义在 `include/net/inet_connection_sock.h` 头文件中。

`inet_connection_sock` 结构体在 `sock` 结构体的基础上扩展了与网络连接相关的信息，如本地和远程 IP 地址、端口号、套接字选项、拥塞控制算法等。它与 `sock` 结构体之间是继承关系，即 `inet_connection_sock` 是 `sock` 的子类。因此，`inet_connection_sock` 中包含了 `sock` 结构体的所有成员变量，同时还扩展了自己特有的成员变量来表示与网络连接相关的信息。

总结起来，**`sock` 是通用的套接字结构体**，用于表示所有类型的套接字的状态；**而 `inet_connection_sock` 是基于 `sock` 的派生结构体**，用于表示 TCP 或 UDP 连接的套接字的状态，包含了额外的网络连接相关信息。`inet_connection_sock` 扩展了 `sock` 结构体，提供了更多与网络连接相关的功能。

### tcp_sock

`struct tcp_sock`从`struct inet_connection_sock`结构体的基础上继承而来，在`struct inet_connection_sock`基础上增加了一些tcp协议相关的字段,如滑动窗口协议，拥塞算法等一些TCP专有的属性。

位置如下：`linux/include/linux/tcp.h`

``` c
struct tcp_sock {//在 inet_connection_sock  基础上增加了 滑动窗口 拥塞控制算法等tcp 专有 属性
    __be32    pred_flags;/*首部预测标志 在接收到 syn 跟新窗口 等时设置此标志 ，
    此标志和时间戳 序号等 用于判断执行 快速还是慢速路径*/

    u64    bytes_received;    /* RFC4898 tcpEStatsAppHCThruOctetsReceived
                 * sum(delta(rcv_nxt)), or how many bytes
                 * were acked.
                 */
    u32    segs_in;    /* RFC4898 tcpEStatsPerfSegsIn
                 * total number of segments in.
                 */
     u32    rcv_nxt;    /* What we want to receive next  等待接收的下一个序列号    */
    u32    copied_seq;    /* Head of yet unread data        */

/* rcv_nxt on last window update sent最早接收但没有确认的序号， 也就是接收窗口的左端，
        在发送ack的时候， rcv_nxt更新 因此rcv_wup 更新比rcv_nxt 滞后一些  */
    u32    rcv_wup;    

    u32    snd_nxt;    /* Next sequence we send 等待发送的下一个序列号        */
    u32    segs_out;    /* RFC4898 tcpEStatsPerfSegsOut
                 * The total number of segments sent.
                 */
    u64    bytes_acked;    /* RFC4898 tcpEStatsAppHCThruOctetsAcked
                 * sum(delta(snd_una)), or how many bytes
                 * were acked.
                 */
    struct u64_stats_sync syncp; /* protects 64bit vars (cf tcp_get_info()) */

     u32    snd_una;    /* First byte we want an ack for  最早一个未被确认的序号    */
     u32    snd_sml;    /* Last byte of the most recently transmitted small packet  最近发送一个小于mss的最后 一个字节序列号
    在成功发送， 如果报文小于mss，跟新这个字段 主要用来判断是否启用 nagle 算法*/
    u32    rcv_tstamp;    /* timestamp of last received ACK (for keepalives)  最近一次收到ack的时间 用于 tcp 保活*/
    u32    lsndtime;    /* timestamp of last sent data packet (for restart window) 最近一次发送 数据包时间*/
    u32    last_oow_ack_time;  /* timestamp of last out-of-window ACK */

    u32    tsoffset;    /* timestamp offset */

    struct list_head tsq_node; /* anchor in tsq_tasklet.head list */
    unsigned long    tsq_flags;

    /* Data for direct copy to user cp 数据到用户进程的控制块 有用户缓存以及其长度 prequeue 队列 其内存*/
    struct {
        struct sk_buff_head    prequeue // tcp 段 缓冲到此队列 知道进程主动读取才真正的处理;
        struct task_struct    *task;
        struct msghdr        *msg;
        int            memory;// prequeue 当前消耗的内存
        int            len;// 用户缓存中 当前可以使用的缓存大小 
    } ucopy;

    u32    snd_wl1;    /* Sequence for window update记录跟新发送窗口的那个ack 段号 用来判断是否 需要跟新窗口
    如果后续收到ack大于snd_wll 则表示需要更新 窗口*/
    u32    snd_wnd;    /* The window we expect to receive 接收方 提供的窗口大小 也就是发送方窗口大小********************/
    u32    max_window;    /* Maximal window ever seen from peer 接收方通告的最大窗口    */
    u32    mss_cache;    /* Cached effective mss, not including SACKS  发送方当前有效的mss*/

    u32    window_clamp;    /* Maximal window to advertise 滑动窗口最大值        */
    u32    rcv_ssthresh;    /* Current window clamp  当前接收窗口的阈值            */
    ......
     u32    snd_ssthresh;    /* Slow start size threshold 拥塞控制 满启动阈值        */
     u32    snd_cwnd;    /* Sending congestion window    当前拥塞窗口大小  ---发送的拥塞窗口   ***********************/
    u32    snd_cwnd_cnt;    /* Linear increase counter    自从上次调整拥塞窗口后 到目前为止接收到的
    总ack段数 如果该字段为0  表示调整拥塞窗口但是没有收到ack，调整拥塞窗口之后 收到ack段就会让
    snd_cwnd_cnt 加1 */
    u32    snd_cwnd_clamp; /* Do not allow snd_cwnd to grow above this  snd_cwnd  的最大值*/
    u32    snd_cwnd_used;//记录已经从队列发送而没有被ack的段数
    u32    snd_cwnd_stamp;//记录最近一次检验cwnd 的时间；     拥塞期间 每次会检验cwnd而调节拥塞窗口 ，
    //在非拥塞期间,为了防止应用层序造成拥塞窗口失效  因此在发送后 有必要检测cwnd
    u32    prior_cwnd;    /* Congestion window at start of Recovery.在进入 Recovery 状态时的拥塞窗口 */
    u32    prr_delivered;    /* Number of newly delivered packets to在恢复阶段给接收者新发送包的数量
                 * receiver in Recovery. */
    u32    prr_out;    /* Total number of pkts sent during Recovery.在恢复阶段一共发送的包的数量 */

     u32    rcv_wnd;    /* Current receiver window 当前接收窗口的大小        */
    u32    write_seq;    /* Tail(+1) of data held in tcp send buffer   已加入发送队列中的最后一个字节序号*/
    u32    notsent_lowat;    /* TCP_NOTSENT_LOWAT */
    u32    pushed_seq;    /* Last pushed seq, required to talk to windows */
    u32    lost_out;    /* Lost packets丢失的数据报            */
    u32    sacked_out;    /* SACK'd packets启用 SACK 时，通过 SACK 的 TCP 选项标识已接收到的段的数量。
                 不启用 SACK 时，标识接收到的重复确认的次数，该值在接收到确认新数据段时被清除。            */
    u32    fackets_out;    /* FACK'd packets    FACK'd packets 记录 SND.UNA 与 (SACK 选项中目前接收方收到的段中最高序号段) 之间的段数。FACK
            用 SACK 选项来计算丢失在网络中上的段数  lost_out=fackets_out-sacked_out  left_out=fackets_out        */

    /* from STCP, retrans queue hinting */
    struct sk_buff* lost_skb_hint; /*在重传队列中， 缓存下次要标志的段*/
    struct sk_buff *retransmit_skb_hint;/* 表示将要重传的起始包*/

    /* OOO segments go in this list. Note that socket lock must be held,
     * as we do not use sk_buff_head lock.
     */
    struct sk_buff_head    out_of_order_queue;

    /* SACKs data, these 2 need to be together (see tcp_options_write) */
    struct tcp_sack_block duplicate_sack[1]; /* D-SACK block */
    struct tcp_sack_block selective_acks[4]; /* The SACKS themselves*/

    struct tcp_sack_block recv_sack_cache[4];

    struct sk_buff *highest_sack;   /* skb just after the highest
                     * skb with SACKed bit set
                     * (validity guaranteed only if
                     * sacked_out > 0)
                     */

    int     lost_cnt_hint;/* 已经标志了多少个段 */
    u32     retransmit_high;    /* L-bits may be on up to this seqno  表示将要重传的起始包 */

    u32    prior_ssthresh; /* ssthresh saved at recovery start表示前一个snd_ssthresh得大小    */
    u32    high_seq;    /* snd_nxt at onset of congestion拥塞开始时，snd_nxt的大----开始拥塞的时候下一个要发送的序号字节*/

    u32    retrans_stamp;    /* Timestamp of the last retransmit,
                 * also used in SYN-SENT to remember stamp of
                 * the first SYN. */
    u32    undo_marker;    /* snd_una upon a new recovery episode. 在使用 F-RTO 算法进行发送超时处理，或进入 Recovery 进行重传，
                    或进入 Loss 开始慢启动时，记录当时 SND.UNA, 标记重传起始点。它是检测是否可以进行拥塞控制撤销的条件之一，一般在完成
                    拥塞撤销操作或进入拥塞控制 Loss 状态后会清零。*/
    int    undo_retrans;    /* number of undoable retransmissions. 在恢复拥塞控制之前可进行撤销的重传段数。
                    在进入 FTRO 算法或 拥塞状态 Loss 时，清零，在重传时计数，是检测是否可以进行拥塞撤销的条件之一。*/
    u32    total_retrans;    /* Total retransmits for entire connection */

    u32    urg_seq;    /* Seq of received urgent pointer  紧急数据的序号 所在段的序号和紧急指针相加获得*/
    unsigned int        keepalive_time;      /* time before keep alive takes place */
    unsigned int        keepalive_intvl;  /* time interval between keep alive probes */

    int            linger2;

/* Receiver side RTT estimation */
    struct {
        u32    rtt;
        u32    seq;
        u32    time;
    } rcv_rtt_est;

/* Receiver queue space */
    struct {
        int    space;
        u32    seq;
        u32    time;
    } rcvq_space;

/* TCP-specific MTU probe information. */
    struct {
        u32          probe_seq_start;
        u32          probe_seq_end;
    } mtu_probe;
    u32    mtu_info; /* We received an ICMP_FRAG_NEEDED / ICMPV6_PKT_TOOBIG
               * while socket was owned by user.
               */

#ifdef CONFIG_TCP_MD5SIG
    const struct tcp_sock_af_ops    *af_specific;
    struct tcp_md5sig_info    __rcu *md5sig_info;
#endif

    struct tcp_fastopen_request *fastopen_req;

    struct request_sock *fastopen_rsk;
    u32    *saved_syn;
};
```



### tcp_congestion_ops

` tcp_congestion_ops `是用于表示 TCP 拥塞控制算法的操作函数和参数的集合。实现的所有cc算法都是这个结构体类型

``` c
/*
结构体tcp_congestion_ops的定义位于linux/include/net/tcp.h:1072或/usr/src/linux-5.4.224/include/net/tcp.h:1072
*/

/*
tcp_congestion_ops 是一个结构体类型，用于描述 TCP 拥塞控制算法的操作和属性。它包含了一系列函数指针和成员变量，用于实现拥塞控制算法的各种功能。ops 是 "operations" 的缩写。
tcp_congestion_ops 结构体的具体定义可能会根据不同的内核版本而有所不同。但通常包含以下常见的成员：
name：拥塞控制算法的名称。
init：初始化拥塞控制算法的函数。
ssthresh：根据拥塞状态设置拥塞窗口阈值的函数。
cong_avoid：根据网络拥塞情况调整拥塞窗口大小的函数。
set_state：设置拥塞控制算法的状态的函数。
cwnd_event：在特定事件发生时调用的函数，用于更新拥塞窗口大小。
in_ack_event：接收到 ACK 时调用的函数，用于更新拥塞窗口大小。
min_tso_segs：最小的 TSO 分段大小。
owner：拥塞控制算法的所有者。
通过定义这些函数指针和成员变量，可以将不同的拥塞控制算法实现封装到相同的数据结构中，方便在内核中进行统一管理和调用。在实际使用中，可以根据需求选择合适的拥塞控制算法，并通过 tcp_ca_find 函数进行查找和使用。
*/
struct tcp_congestion_ops {
	struct list_head	list; // 用于将结构体链接到一个链表中，以便于管理和遍历不同的拥塞控制算法。
	u32 key; // key 和 flags：用于标识和配置拥塞控制算法的相关信息。
	u32 flags;

	/* initialize private data (optional) */
	void (*init)(struct sock *sk); // init函数，接收struct sock 类型的参数，表示 TCP 的套接字。会在 TCP 连接建立时被调用，用于初始化拥塞控制算法的状态和参数。
	/* cleanup private data  (optional) */
	void (*release)(struct sock *sk);  // 释放拥塞控制算法的私有数据。

	/* return slow start threshold (required) */
	u32 (*ssthresh)(struct sock *sk); // 这个函数会在每次拥塞发生时被调用，用于更新慢启动门限。
	/* do new cwnd calculation (required) */
	void (*cong_avoid)(struct sock *sk, u32 ack, u32 acked);  // 三个参数分别表示 TCP 的套接字、收到的 ACK 包的序号和确认的数据量。这个函数会在每次 ACK 接收时被调用，用于根据当前的网络状况调整拥塞窗口大小。
	/* call before changing ca_state (optional) */
	void (*set_state)(struct sock *sk, u8 new_state); // 用于设置 TCP 的状态。new_state表示 新的状态。这个函数会在 TCP 的状态发生变化时被调用，用于根据新的状态执行相应的操作。
	/* call when cwnd event occurs (optional) */
	void (*cwnd_event)(struct sock *sk, enum tcp_ca_event ev); // enum tcp_ca_event ev，表示拥塞窗口事件类型。功能：根据拥塞窗口事件的类型进行相应的操作。
	/* call when ack arrives (optional) */
	void (*in_ack_event)(struct sock *sk, u32 flags); // flags，表示 ACK 标志。根据 ACK 标志进行相应的操作。
	/* new value of cwnd after loss (required) */
	u32  (*undo_cwnd)(struct sock *sk); // 在发生数据包丢失后调用。功能：返回丢包后的新拥塞窗口大小。
	/* hook for packet ack accounting (optional) */
	void (*pkts_acked)(struct sock *sk, const struct ack_sample *sample); // cnt 表示被确认的数据包数，rtt 表示往返时间的估计值。这个函数会在每次 ACK 接收时被调用，用于根据被确认的数据包和往返时间的估计值更新拥塞窗口大小。
	/* override sysctl_tcp_min_tso_segs */
	u32 (*min_tso_segs)(struct sock *sk); // 设置最小的tso分段。TSO：TCP Segmentation Offload，TCP分段卸载
	/* returns the multiplier used in tcp_sndbuf_expand (optional) */
	u32 (*sndbuf_expand)(struct sock *sk);
	/* call when packets are delivered to update cwnd and pacing rate,
	 * after all the ca_state processing. (optional)
	 */
	void (*cong_control)(struct sock *sk, const struct rate_sample *rs); // const struct rate_sample *rs，表示速率样本。：根据速率样本更新拥塞窗口和传输速率。

	/* get info for inet_diag (optional) */
	size_t (*get_info)(struct sock *sk, u32 ext, int *attr,
			   union tcp_cc_info *info); // 获取用于 inet_diag 的信息。

	char 		name[TCP_CA_NAME_MAX];
	struct module 	*owner;
};
```



## TCP拥塞控制算法reno版本解读

所在版本代码所在位置是：`linux-5.4.224/net/ipv4/tcp_cong.c`，该文件是tcp拥塞控制算法的起点文件。定义了cc算法的注册、查找、分配、初始化、

``` c
// SPDX-License-Identifier: GPL-2.0-only
/*
 * Pluggable TCP congestion control support and newReno
 * congestion control.
 * Based on ideas from I/O scheduler support and Web100.
 *
 * Copyright (C) 2005 Stephen Hemminger <shemminger@osdl.org>
 */



#define pr_fmt(fmt) "TCP: " fmt // 一个预处理指令，用于定义输出格式。在这里，它定义了输出格式为 "TCP: " 加上具体的格式。

#include <linux/module.h>
#include <linux/mm.h>
#include <linux/types.h>
#include <linux/list.h>
#include <linux/gfp.h>
#include <linux/jhash.h>
#include <net/tcp.h>

static DEFINE_SPINLOCK(tcp_cong_list_lock); // 定义了一个自旋锁，用于保护 tcp_cong_list 变量的访问。自旋锁是一种轻量级的锁，用于在多个CPU之间同步访问共享资源。
static LIST_HEAD(tcp_cong_list); // 定义了一个链表头，用于存储TCP拥塞控制算法的拥塞控制模块。链表是用于存储一系列元素的数据结构，链表头是链表的起点。
// 通过name查找tcp cc算法
/* Simple linear search, don't expect many entries! */
static struct tcp_congestion_ops *tcp_ca_find(const char *name)
{
	struct tcp_congestion_ops *e;

	list_for_each_entry_rcu(e, &tcp_cong_list, list) {
		if (strcmp(e->name, name) == 0)
			return e;
	}

	return NULL;
}

// 这部分代码的作用是实现了拥塞控制算法列表的查找功能，并提供了首次使用算法时自动加载内核模块的支持。
/* Must be called with rcu lock held */
static struct tcp_congestion_ops *tcp_ca_find_autoload(struct net *net,
						       const char *name)
{
	struct tcp_congestion_ops *ca = tcp_ca_find(name);

#ifdef CONFIG_MODULES
	if (!ca && capable(CAP_NET_ADMIN)) {
		rcu_read_unlock();
		request_module("tcp_%s", name);
		rcu_read_lock();
		ca = tcp_ca_find(name);
	}
#endif
	return ca;
}

// 根据cc算法的key来查找
/* Simple linear search, not much in here. */
struct tcp_congestion_ops *tcp_ca_find_key(u32 key)
{
	struct tcp_congestion_ops *e;

	list_for_each_entry_rcu(e, &tcp_cong_list, list) {
		if (e->key == key)
			return e;
	}

	return NULL;
}

// 注册新的拥塞控制算法到可用的选项列表里
/*
 * Attach new congestion control algorithm to the list
 * of available options.
 */
int tcp_register_congestion_control(struct tcp_congestion_ops *ca)
{
	int ret = 0;
	
    // 显然，这4个ca的属性是必需的
	/* all algorithms must implement these */
	if (!ca->ssthresh || !ca->undo_cwnd ||
	    !(ca->cong_avoid || ca->cong_control)) {
		pr_err("%s does not implement required ops\n", ca->name);
		return -EINVAL;
	}
	// 生成ca唯一的 key
	ca->key = jhash(ca->name, sizeof(ca->name), strlen(ca->name));
	// 给cc算法列表加上锁
	spin_lock(&tcp_cong_list_lock); 
	if (ca->key == TCP_CA_UNSPEC || tcp_ca_find_key(ca->key)) {
		pr_notice("%s already registered or non-unique key\n",
			  ca->name);
		ret = -EEXIST;
	} else {
		list_add_tail_rcu(&ca->list, &tcp_cong_list);
		pr_debug("%s registered\n", ca->name);
	}
	spin_unlock(&tcp_cong_list_lock);

	return ret;
}
EXPORT_SYMBOL_GPL(tcp_register_congestion_control);
/*
EXPORT_SYMBOL_GPL(tcp_register_congestion_control) 是一个宏，用于将 tcp_register_congestion_control 函数导出为一个可供其他模块使用的符号。
在 Linux 内核中，为了允许不同的模块之间共享函数和变量，需要将它们导出为符号。导出的符号可以被其他模块引用和调用。EXPORT_SYMBOL_GPL 宏用于导出一个符号，并指定该符号只能由 GNU General Public License (GPL) 许可证下的模块使用。
具体到 tcp_register_congestion_control 函数，它是用于注册 TCP 拥塞控制算法的函数。通过导出该符号，其他模块可以使用该函数来注册自定义的 TCP 拥塞控制算法，以扩展内核中的拥塞控制功能。
*/

// 取消注册
/*
 * Remove congestion control algorithm, called from
 * the module's remove function.  Module ref counts are used
 * to ensure that this can't be done till all sockets using
 * that method are closed.
 */
void tcp_unregister_congestion_control(struct tcp_congestion_ops *ca)
{
	spin_lock(&tcp_cong_list_lock);
	list_del_rcu(&ca->list);
	spin_unlock(&tcp_cong_list_lock);

	/* Wait for outstanding readers to complete before the
	 * module gets removed entirely.
	 *
	 * A try_module_get() should fail by now as our module is
	 * in "going" state since no refs are held anymore and
	 * module_exit() handler being called.
	 */
	synchronize_rcu();
}
EXPORT_SYMBOL_GPL(tcp_unregister_congestion_control);

// 根据ca的name得到其key
u32 tcp_ca_get_key_by_name(struct net *net, const char *name, bool *ecn_ca)
{
	const struct tcp_congestion_ops *ca;
	u32 key = TCP_CA_UNSPEC;

	might_sleep();

	rcu_read_lock();
	ca = tcp_ca_find_autoload(net, name);
	if (ca) {
		key = ca->key;
		*ecn_ca = ca->flags & TCP_CONG_NEEDS_ECN;
	}
	rcu_read_unlock();

	return key;
}
EXPORT_SYMBOL_GPL(tcp_ca_get_key_by_name);

// 该函数的作用是通过给定的键值获取对应拥塞控制算法的名称并存储在指定的缓冲区中。
char *tcp_ca_get_name_by_key(u32 key, char *buffer)
{
	const struct tcp_congestion_ops *ca;
	char *ret = NULL;

	rcu_read_lock();
	ca = tcp_ca_find_key(key);
    // 使用 strncpy 函数将该算法的名称拷贝到指定的缓冲区 buffer 中，拷贝的长度为 TCP_CA_NAME_MAX，即拥塞控制算法名称的最大长度。
	if (ca)
		ret = strncpy(buffer, ca->name,
			      TCP_CA_NAME_MAX);
	rcu_read_unlock();

	return ret;
}
EXPORT_SYMBOL_GPL(tcp_ca_get_name_by_key);

// 为给定的socket套接字分配一个cc算法
/* Assign choice of congestion control. */
void tcp_assign_congestion_control(struct sock *sk)
{
	struct net *net = sock_net(sk);	// 获取套接字所属的网络命名空间net，
    // inet_connection_sock 是一个与 TCP/IP 协议栈中连接相关的数据结构，用于存储与连接有关的信息。在这个函数中，通过 struct inet_connection_sock *icsk 定义了一个指向该结构体的指针。
    // 简单来说，net_connection_sock 是 sock 结构体的派生结构体，用于表示 TCP 或 UDP 连接的套接字的状态。sock 是通用的套接字结构体，用于表示所有类型的套接字的状态；而 inet_connection_sock 是基于 sock 的派生结构体，用于表示 TCP 或 UDP 连接的套接字的状态，包含了额外的网络连接相关信息。inet_connection_sock 扩展了 sock 结构体，提供了更多与网络连接相关的功能。
	struct inet_connection_sock *icsk = inet_csk(sk); // 获取与套接字关联的 icsk
	const struct tcp_congestion_ops *ca;

	rcu_read_lock();
	ca = rcu_dereference(net->ipv4.tcp_congestion_control); // 获取当前网络空间的ca，使用rcu_dereference 函数用于读取 RCU 保护的数据。
	// 如果获取拥塞控制算法的所有者模块失败（可能是因为该拥塞控制算法的模块已经被卸载），则将拥塞控制算法设置为 tcp_reno，这是 TCP Reno 算法的默认拥塞控制算法。
    if (unlikely(!try_module_get(ca->owner)))
		ca = &tcp_reno;
	icsk->icsk_ca_ops = ca; // 为套接字设置拥塞控制算法。
	rcu_read_unlock();
	// 使用 memset 函数将 icsk->icsk_ca_priv 的内容清零，该变量用于存储拥塞控制算法的私有数据。
	memset(icsk->icsk_ca_priv, 0, sizeof(icsk->icsk_ca_priv));
    // 如果拥塞控制算法设置了 TCP_CONG_NEEDS_ECN 标志，说明该算法需要使用 ECN（Explicit Congestion Notification）机制，调用 INET_ECN_xmit 函数启用套接字的 ECN 发送，否则调用 INET_ECN_dontxmit 函数禁用套接字的 ECN 发送。
	if (ca->flags & TCP_CONG_NEEDS_ECN)
		INET_ECN_xmit(sk);
	else
		INET_ECN_dontxmit(sk);
}

void tcp_init_congestion_control(struct sock *sk)
{
	const struct inet_connection_sock *icsk = inet_csk(sk);

	tcp_sk(sk)->prior_ssthresh = 0;
	if (icsk->icsk_ca_ops->init)
		icsk->icsk_ca_ops->init(sk); // 如果此ca有init函数就调用
	if (tcp_ca_needs_ecn(sk))
		INET_ECN_xmit(sk);
	else
		INET_ECN_dontxmit(sk);
}

static void tcp_reinit_congestion_control(struct sock *sk,
					  const struct tcp_congestion_ops *ca)
{
	struct inet_connection_sock *icsk = inet_csk(sk);

	tcp_cleanup_congestion_control(sk);  // 清理之前的拥塞控制算法
	icsk->icsk_ca_ops = ca;// 重新赋 ca
	icsk->icsk_ca_setsockopt = 1; // 表示需要重新设置拥塞控制算法的参数。
	memset(icsk->icsk_ca_priv, 0, sizeof(icsk->icsk_ca_priv));

	if (ca->flags & TCP_CONG_NEEDS_ECN)
		INET_ECN_xmit(sk);
	else
		INET_ECN_dontxmit(sk);
	// 最后，如果当前的 TCP 状态不是 CLOSE 或 LISTEN，则调用 tcp_init_congestion_control 函数进行初始化。
	if (!((1 << sk->sk_state) & (TCPF_CLOSE | TCPF_LISTEN)))
		tcp_init_congestion_control(sk);
}

/* Manage refcounts on socket close. */
void tcp_cleanup_congestion_control(struct sock *sk)
{
	struct inet_connection_sock *icsk = inet_csk(sk);

	if (icsk->icsk_ca_ops->release)
		icsk->icsk_ca_ops->release(sk);
	module_put(icsk->icsk_ca_ops->owner);
}

// 这个函数是用于在运行时动态更改默认的 TCP 拥塞控制算法。
/* Used by sysctl to change default congestion control */
int tcp_set_default_congestion_control(struct net *net, const char *name)
{
	struct tcp_congestion_ops *ca;
	const struct tcp_congestion_ops *prev;
	int ret;

	rcu_read_lock();
	ca = tcp_ca_find_autoload(net, name);
	if (!ca) {
		ret = -ENOENT;
	} else if (!try_module_get(ca->owner)) {
		ret = -EBUSY;
	} else if (!net_eq(net, &init_net) &&
			!(ca->flags & TCP_CONG_NON_RESTRICTED)) {
		/* Only init netns can set default to a restricted algorithm */
		ret = -EPERM;
	} else { // 前三个if都是返回错误
		prev = xchg(&net->ipv4.tcp_congestion_control, ca); // 使用原子操作 xchg 将给定的算法结构体设置为网络的默认拥塞控制算法，并返回之前的默认算法。
		if (prev)
			module_put(prev->owner);

		ca->flags |= TCP_CONG_NON_RESTRICTED;
		ret = 0;
	}
	rcu_read_unlock();

	return ret;
}

// 在引导时从内核配置中设置默认值的函数。它调用 tcp_set_default_congestion_control 函数，并传递初始网络命名空间和内核配置中设置的默认拥塞控制算法。它在内核的 late_initcall 阶段被调用，以确保在系统初始化的最后阶段设置默认值。
/* Set default value from kernel configuration at bootup */
static int __init tcp_congestion_default(void)
{
	return tcp_set_default_congestion_control(&init_net,
						  CONFIG_DEFAULT_TCP_CONG);
}
late_initcall(tcp_congestion_default);

// 用于构建一个字符串，其中包含可用的拥塞控制算法值的列表。
/* Build string with list of available congestion control values */
void tcp_get_available_congestion_control(char *buf, size_t maxlen)
{
	struct tcp_congestion_ops *ca;
	size_t offs = 0;

	rcu_read_lock();
	list_for_each_entry_rcu(ca, &tcp_cong_list, list) {
		offs += snprintf(buf + offs, maxlen - offs,
				 "%s%s",
				 offs == 0 ? "" : " ", ca->name);
	}
	rcu_read_unlock();
}

// 获取当前默认的拥塞控制算法。它接受一个网络结构体指针和一个字符数组指针作为参数。通过 RCU 读取锁来保护对 net->ipv4.tcp_congestion_control 的访问，然后将该值的名称复制到给定的字符数组中。
/* Get current default congestion control */
void tcp_get_default_congestion_control(struct net *net, char *name)
{
	const struct tcp_congestion_ops *ca;

	rcu_read_lock();
	ca = rcu_dereference(net->ipv4.tcp_congestion_control);
	strncpy(name, ca->name, TCP_CA_NAME_MAX);
	rcu_read_unlock();
}

// 它遍历存储在全局变量 tcp_cong_list 中的拥塞控制算法结构体，并跳过标记为非限制的算法。然后将每个算法的名称追加到给定的字符数组中。这个函数同样使用 RCU 读取锁来保护对全局变量的访问。
/* Built list of non-restricted congestion control values */
void tcp_get_allowed_congestion_control(char *buf, size_t maxlen)
{
	struct tcp_congestion_ops *ca;
	size_t offs = 0;

	*buf = '\0';
	rcu_read_lock();
	list_for_each_entry_rcu(ca, &tcp_cong_list, list) {
		if (!(ca->flags & TCP_CONG_NON_RESTRICTED))
			continue;
		offs += snprintf(buf + offs, maxlen - offs,
				 "%s%s",
				 offs == 0 ? "" : " ", ca->name);
	}
	rcu_read_unlock();
}

// 用于更改拥塞控制算法的非限制列表
/* Change list of non-restricted congestion control */
int tcp_set_allowed_congestion_control(char *val)
{
	struct tcp_congestion_ops *ca;
	char *saved_clone, *clone, *name;
	int ret = 0;

	saved_clone = clone = kstrdup(val, GFP_USER);
	if (!clone)
		return -ENOMEM; // 内存不足

	spin_lock(&tcp_cong_list_lock);
	/* pass 1 check for bad entries */
	while ((name = strsep(&clone, " ")) && *name) {
		ca = tcp_ca_find(name);
		if (!ca) {
			ret = -ENOENT; // 表示未找到指定的拥塞控制算法。
			goto out;
		}
	}

	/* pass 2 clear old values */
	list_for_each_entry_rcu(ca, &tcp_cong_list, list)
		ca->flags &= ~TCP_CONG_NON_RESTRICTED;

	/* pass 3 mark as allowed */
	while ((name = strsep(&val, " ")) && *name) {
		ca = tcp_ca_find(name);
		WARN_ON(!ca);
		if (ca)
			ca->flags |= TCP_CONG_NON_RESTRICTED;
	}
out:
	spin_unlock(&tcp_cong_list_lock);
	kfree(saved_clone);

	return ret;
}

/* Change congestion control for socket. If load is false, then it is the
 * responsibility of the caller to call tcp_init_congestion_control or
 * tcp_reinit_congestion_control (if the current congestion control was
 * already initialized.
 */
int tcp_set_congestion_control(struct sock *sk, const char *name, bool load,
			       bool reinit, bool cap_net_admin)
{
	struct inet_connection_sock *icsk = inet_csk(sk);
	const struct tcp_congestion_ops *ca;
	int err = 0;

	if (icsk->icsk_ca_dst_locked)
		return -EPERM;

	rcu_read_lock();
	if (!load)
		ca = tcp_ca_find(name); // 如果已经加载就根据名字找到并赋值，
	else
		ca = tcp_ca_find_autoload(sock_net(sk), name); // 如果没有加载就在指定网络空间中自动加载指定名字的ca，

    // 如果找到的拥塞控制算法与套接字的当前拥塞控制算法相同，就准备跳出
	/* No change asking for existing value */
	if (ca == icsk->icsk_ca_ops) {
		icsk->icsk_ca_setsockopt = 1;
		goto out;
	}
 	
	if (!ca) { // 如果ca没找到，返回个err
		err = -ENOENT;
	} else if (!load) {// 如果ca找到了，但此ca与当前套接字的ca不同，且没有load，就尝试试获取新拥塞控制算法的所有权
		const struct tcp_congestion_ops *old_ca = icsk->icsk_ca_ops;

		if (try_module_get(ca->owner)) {
			if (reinit) {
				tcp_reinit_congestion_control(sk, ca);
			} else {
				icsk->icsk_ca_ops = ca;
				module_put(old_ca->owner);
			}
		} else {
			err = -EBUSY;
		}
	} else if (!((ca->flags & TCP_CONG_NON_RESTRICTED) || cap_net_admin)) {
		err = -EPERM;
	} else if (!try_module_get(ca->owner)) {
		err = -EBUSY;
	} else {
		tcp_reinit_congestion_control(sk, ca);
	}
 out:
	rcu_read_unlock();
	return err;
}

/* Slow start is used when congestion window is no greater than the slow start
 * threshold. We base on RFC2581 and also handle stretch ACKs properly.
 * We do not implement RFC3465 Appropriate Byte Counting (ABC) per se but
 * something better;) a packet is only considered (s)acked in its entirety to
 * defend the ACK attacks described in the RFC. Slow start processes a stretch
 * ACK of degree N as if N acks of degree 1 are received back to back except
 * ABC caps N to 2. Slow start exits when cwnd grows over ssthresh and
 * returns the leftover acks to adjust cwnd in congestion avoidance mode.
 */
u32 tcp_slow_start(struct tcp_sock *tp, u32 acked) // acked 参数，表示已经确认的字节数
{
	u32 cwnd = min(tp->snd_cwnd + acked, tp->snd_ssthresh); // 计算新的cwnd

	acked -= cwnd - tp->snd_cwnd; // 更新acked的值
    // 如果acked=0说明还处于慢开启阶段，如果acked>0，说明进入了拥塞避免阶段，多出的确认的字节数用于更新拥塞避免阶段的窗口值
	tp->snd_cwnd = min(cwnd, tp->snd_cwnd_clamp); // 
	// tp->snd_cwnd_clamp 是一个用于限制拥塞窗口大小的变量。它表示拥塞窗口在增长过程中的最大值。
	return acked; // 函数返回一个无符号32位整数，表示剩余未确认的字节数。如果多于0就把多余的acked拿去用于更新拥塞避免阶段的窗口值
}
EXPORT_SYMBOL_GPL(tcp_slow_start);

// 拥塞避免阶段的加性增, AI 阶段是在网络没有出现拥塞情况时，逐步增加拥塞窗口大小。
/* In theory this is tp->snd_cwnd += 1 / tp->snd_cwnd (or alternative w),
 * for every packet that was ACKed.
 */ 
void tcp_cong_avoid_ai(struct tcp_sock *tp, u32 w, u32 acked)
    // w：表示一个拥塞窗口的大小，acked：表示最近一次收到的确认数。
{
    // 这个函数说明，在拥塞避免阶段，并不是每收到一个ack，就把窗口加1，而是需要累积确认的包数到一定程度w，才会将窗口加1。
    // tp->snd_cwnd_cnt：表示拥塞窗口中已经发送但尚未确认的数据大小。在这个函数中，它用于累积新的确认数，以便判断是否可以增加拥塞窗口大小。
	// tp->snd_cwnd：表示当前的拥塞窗口大小。在这个函数中，它会根据之前累积的信用和新收到的确认情况进行调整。
	/* If credits accumulated at a higher w, apply them gently now. */
	if (tp->snd_cwnd_cnt >= w) { // 说明之前累积的信用已经达到了一个更高的拥塞窗口大小。在这种情况下，将 tp->snd_cwnd_cnt 重置为0，并将 tp->snd_cwnd 增加1，以逐渐应用之前累积的信用。
		tp->snd_cwnd_cnt = 0;
		tp->snd_cwnd++;
	}

	tp->snd_cwnd_cnt += acked;
    // 上一个判断if可能没过，加上acked，可能就过了，那么再来判断并准备增加窗口
	if (tp->snd_cwnd_cnt >= w) {
		u32 delta = tp->snd_cwnd_cnt / w; // 计算 delta 表示可以增加的整倍数，

		tp->snd_cwnd_cnt -= delta * w;
		tp->snd_cwnd += delta;
	}
	tp->snd_cwnd = min(tp->snd_cwnd, tp->snd_cwnd_clamp);
}
EXPORT_SYMBOL_GPL(tcp_cong_avoid_ai);

/*
 * TCP Reno congestion control
 * This is special case used for fallback as well.
 */
/* This is Jacobson's slow start and congestion avoidance.
 * SIGCOMM '88, p. 328.
 */
// 拥塞避免阶段，没有涉及丢包等触发拥塞的代码；u32 ack：表示最近一次收到的确认号。u32 acked：表示最近一次收到的确认数。
// 即使TCP流处于慢开启阶段，收到ack后也会触发这个拥塞避免的代码
void tcp_reno_cong_avoid(struct sock *sk, u32 ack, u32 acked)
{
	struct tcp_sock *tp = tcp_sk(sk);

	if (!tcp_is_cwnd_limited(sk)) // 如果拥塞窗口不受限，则直接返回，
		return;
    
	/* In "safe" area, increase. */
	if (tcp_in_slow_start(tp)) { // 如果处于慢开启阶段
		acked = tcp_slow_start(tp, acked);
		if (!acked) // 如果返回的acked不为0 说明离开了慢开启阶段需要进入拥塞避免阶段了
			return;
	}
	/* In dangerous area, increase slowly. */
	tcp_cong_avoid_ai(tp, tp->snd_cwnd, acked);
}
EXPORT_SYMBOL_GPL(tcp_reno_cong_avoid);

/* Slow start threshold is half the congestion window (min 2) */
u32 tcp_reno_ssthresh(struct sock *sk)
    // 该函数用于计算慢开启门限
{
	const struct tcp_sock *tp = tcp_sk(sk);

	return max(tp->snd_cwnd >> 1U, 2U); // 确保慢启动阈值至少为2。
}
EXPORT_SYMBOL_GPL(tcp_reno_ssthresh);

// 在拥塞避免阶段，当发生拥塞事件时，发送方会将拥塞窗口大小减小为先前的拥塞窗口大小（prior_cwnd）。
// tcp_reno_undo_cwnd 函数被用来恢复拥塞窗口大小，即将拥塞窗口大小恢复到先前的大小，以便发送方可以重新开始以较大的窗口发送数据。
u32 tcp_reno_undo_cwnd(struct sock *sk)
{
	const struct tcp_sock *tp = tcp_sk(sk);
	// 通过将当前的发送拥塞窗口（snd_cwnd）与之前的拥塞窗口（prior_cwnd）中的较大值比较，得到最大的拥塞窗口大小作为恢复后的窗口大小。
	return max(tp->snd_cwnd, tp->prior_cwnd); 
}
EXPORT_SYMBOL_GPL(tcp_reno_undo_cwnd);


struct tcp_congestion_ops tcp_reno = {
	.flags		= TCP_CONG_NON_RESTRICTED,
	.name		= "reno",
	.owner		= THIS_MODULE,
	.ssthresh	= tcp_reno_ssthresh,
	.cong_avoid	= tcp_reno_cong_avoid,
	.undo_cwnd	= tcp_reno_undo_cwnd,
};
```



## 注册自己的cc算法

``` c
#include <linux/module.h>
#include <linux/kernel.h>
#include <net/tcp.h>

// 自定义的拥塞控制算法结构体
struct tcp_congestion_ops my_congestion_ops = {
    .name = "my_congestion",
    // 设置拥塞控制算法的各种函数指针和属性
    // ...
};

static int __init my_module_init(void)
{
    int ret = tcp_register_congestion_control(&my_congestion_ops);
    if (ret != 0) {
        printk(KERN_ERR "Failed to register congestion control\n");
        return ret;
    }

    printk(KERN_INFO "Congestion control registered successfully\n");
    return 0;
}

static void __exit my_module_exit(void)
{
    // 可选：如果需要，在模块退出时取消注册拥塞控制算法
    // tcp_unregister_congestion_control(&my_congestion_ops);

    printk(KERN_INFO "Congestion control unregistered\n");
}

module_init(my_module_init);
module_exit(my_module_exit);
MODULE_LICENSE("GPL");
MODULE_AUTHOR("Your Name");
MODULE_DESCRIPTION("Sample module using tcp_register_congestion_control");

```

