## 流量预测

**ResNet-GCN-LSTM模型**

将网络流量的变化作为重要的触发流量预测的条件或直接作为预测的入参

边缘交换机在遇到大流时向控制器发送消息让其进行预测

对SDN网络中**转发节点进行重要度排序**：基于LSTM的SDN网络流量预测研究.pdf

SDN控制器

- **集中式SDN控制器**：使得数据面与控制面间的通信管道 记忆 自身的计算能力 成为网络瓶颈
- **分布式SDN控制器**：更适应于大规模网络，另一方面往往带来了负载均衡问题，在SDN网络中负载均衡问题体现在两方面：
    - 数据面的负载均衡问题，使得部分链路负载高而部分链路负载低，网络整体性能差
    - 控制面的负载均衡问题，主要讨论的是如何充分利用各个控制器的资源。某些控制器负载高而某些控制器负载低影响全网的性能。

流测度：

网络流量预测增强的高效分布式大流检测研究

TCM 图概要算法压缩流量矩阵

预测的大流的剩余长度也是一个重要指标



**网络流定义**为符合某种规范以及超时约束的一系列数据包的集合[35]。这里的流规范主要 有：五元组（源 IP 地址、目的 IP 地址、源端口、目的端口、协议）、二元组（源 IP 地址、目 的 IP 地址）、目的 IP 地址等，这里的流规范也称为流标识；超时约束是指超过一定时间不再 通信的流就认为是已结束，是按照时间划分流的一种方式。流表示的是一条端到端的通信状 态，按照方向性又可以分为有向流和无向流，无向流不区分源 IP 地址（源端口）和目的 IP 地 址（目的端口），比如对于二元组无向流，源 IP 地址为 A，目的 IP 地址为 B 的二元组流与源 IP 地址为 B，目的 IP 地址为 A 的流定义为同一条流

在数据包中获取时间戳来得到时间信息





**流量矩阵**：

`互联网流量矩阵：入门` ； `结合流内相关性和流间相关性进行流量矩阵预测`

流量矩阵可以完整地描述网络中所有流量的分布，是网络规划、网络管理以及流量工程 （Traffic Engineering）中的关键参数

服务器端口、数据包字节中位数、数据包总数、ACK 数据包总数等



### **流量数据集**：

美国 Abilene 骨干网络数据集[48]和无线 Mesh 网络 UCSB（University of California, Santa Barbara）数据集 [49]。

本实验采用 CAIDA passive 2016 数据集[55]，该数据集是来自 CAIDA 的 Equinix-Chicago 监视器的匿名被动流量 trace，是美国西雅图和芝加哥之间的高速主干链路流量数据，其数据 格式为 pcap 文件，此数据集可用于研究互联网流量的特征，包括流量的分布、流长以及流持 续时间等。

[48] The Abilene Dataset 2004[EB/OL].http://www.cs.utexas.edu/~yzhang/research/AbileneTM/. 

[49] The UCSB Dataset 2007[EB/OL].http://moment.cs.ucsb.edu/meshnet/datasets.

[55] The CAIDA Anonymized Internet Traces 2016[EB/OL].http://www.caida.org.

[50] Moore A. ， Zuev D., Crogan M., et al. Discriminators for use inflow-based classification [M]. London: Queen Mary Universityof London, 2005.

本文采用网络公开数据集“ICSX VPN-nonVPN”，该数据集由捕获 到的不同应用程序流量的大小为 1.4GB 的 PCAP 格式文件[60]

[60].Draper-Gil G, Lashkari A H, Mamun M S I, et al. Characterization of encrypted and vpn traffic using time-related[C]//Proceedings of the 2nd international conference on information systems security and privacy (ICISSP). 2016: 407-414.



图 3.5 显 示了经过以上步骤处理后的数据包长度概率密度分布图，如柱状图所示，数据集 中的数据包长度变化很大。

![image-20230221112723512](img/image-20230221112723512.png)



首先在各自交换机上运行程序手机流量信息，当发现大变化流是才通知控制器进行预测等行动

**在单个虚拟或物理网络设备上收集负载和流量统计信息**已经得到了广泛的研究[52,53]，在 设备上收集的数据用于在交换机内做出本地决策。

[52] Basat R B, Einziger G, Friedman R, et al. Heavy hitters in streams and sliding windows[C]//IEEE INFOCOM 2016The 35th Annual IEEE International Conference on Computer Communications. IEEE, 2016: 1-9. 

[53] Basat R B, Einziger G, Friedman R, et al. Optimal elephant flow detection[C]//IEEE INFOCOM 2017-IEEE Conference on Computer Communications. IEEE, 2017: 1-9.

**全局测量**是在数据平面的一些网络设备中收 集测量数据，然后集中式的控制器合并从网络测量点收集的数据，并创建整个流量的网络范 围视图[54]。



本实验基于 P4 的可编程网络运行环境需要用到的工具有：protobuf、grpc、PI、behavioralmodel[58]、p4c 和 mininet。protobuf 是一种二进制的数据传输格式，效率和兼容性都很好。grpc 是一个通用、高性能的开源 RPC 框架。behavioral-model 简称 BMv2，是一款支持 P4 编程 的软件交换机。p4c 是 P4 语言推荐的编译器。mininet 是一个网络模拟器，用于搭建包含虚拟 主机、交换机、控制器和链路的网络拓扑。

**流量重放工具 tcpreplay** 对 CAIDA 流量的 pcap 文件进行重放。

git clone https://gitee.com/hahahawt /p4-guide.git

**流量预测对比实验**：与经典LSTM，ARIMA模型对比

 

## 路径迁移

方法：强化学习

选择最优路径时，要考虑到链路时延、可用带宽、丢包率、链路利用率 4 个因素

在 SDN 网络中，将 SDN 路由优化问题看成决策问题，使用强化学习技术来优 化路由。主要的思路为：首先将实际网络抽象为一个学习环境，网络结构、流量矩 阵等视为状态，路径权重视为一项动作，QoS 服务质量、链路利用率之类的维护策 略被用作奖励；然后代理不断接收最新的状态来训练模型以优化网络性能；最后数 据中心在新流量到达时可以快速计算适当的路由路径。



传统多路径调度存在的问题

**可变路径 MTU**。由于每个冗余路径可能具有不同的 MTU，因此这意味着整个 路径 MTU 可以在逐个数据包的基础上进行更改，从而忽略了路径 MTU 发现的有用 性。

**可变延迟**。由于每个冗余路径可能涉及不同的等待时间，因此使数据包采用单 独的路径可能会导致数据包始终无序到达，从而增加了交付延迟和缓冲要求。数据 包重新排序使 TCP 认为，当具有较高序列号的数据包在较早的数据包之前到达时， 就发生了丢失。当在“最新”数据包之前接收到三个或更多数据包时，TCP 会进入“快速重传”的模式，该模式会试图不必要地重传延迟的数据包，消耗额外的带宽。 因此，重新排序可能会损害网络性能

**调试**。在常见的调试实用程序，例如 ping 和 traceroute，如果存在多个路径时， 其可靠性会较低很多，甚至可能呈现完全错误的结果。

> 2）解决方案
>
>  在 RFC2991 中提供了三种方法来提高多路径算法的效率。 
>
> 模 N 哈希。为了从 N 个下一跳的列表中选择下一跳，路由器对标识流的数据包 头字段执行模 N 哈希。这具有快速的优点，但是每当添加或删除下一跳时，都会以 更改路径的所有流的 (N-1) / N 为代价。 
>
> 哈希阈值。路由器首先通过对标识流的数据包头字段进行哈希运算来选择密钥。 在哈希函数的输出空间中，已为 N 个下一跳分配了唯一的区域。通过将哈希和同区 域边界进行比较，路由器可以确定哈希值值属于哪个区域，从而确定要使用哪个下 一跳。该方法的优点是，在添加或删除下一跳时，仅影响区域边界附近的流量。对 于 ECMP 哈希阈值，可以通过简单的除法进行查找。添加或删除下一跳时，仅有所 有流的 1/4 和 1/2 之间会更改路径。 
>
> 最高随机权重（HRW）。路由器通过对标识流的数据包头字段以及下一跳的地 址执行哈希操作，为每个下一跳计算密钥。然后，路由器选择具有最高结果密钥值 的下一跳。这具有使受下一跳添加或删除影响的流的数量最小化（仅为它们的 1 / N） 的优点，但是它的成本大约是模 N 哈希的 N 倍

当前数据中心网络中普遍采用基于哈希的等价多路径路由算法（ECMP）进行 路由转发[52]。

然而该算法是数据流到路径的静态映射，不考虑链路利用率和流大小，从而会不可避免的出现数据流碰撞问题，容易造成网络拥塞[53]。

[52] Hopps C. Analysis of an equal-cost multi-path algorithm[M]. United States: RFC Editor, 2000.

[53] Dixit A., Prakash P., Charlie H Y., et al. On the impact of packet spraying in data center networks[C]// Proc of IEEE Infocom. Piscataway, NJ: IEEE Press, 2013: 2130-2138.

本文对比的网络性能指标以及部分对比实验借 鉴了文献[57]中的评价方法，网络性能指标主要包括网络的平均吞吐量、链路利用 率、链路带宽利用率、平均往返时延、整体丢包率。

[57] 黄马驰. 基于 SDN 的数据中心网络流量调度策略研究[D]. 重庆:重庆邮电大学, 2017.



具体来说，控制器计算从源节点到目的节点的最优网络状态 t S ，并从动作集 合 A 中**选取一条最优路径 a，进行数据转发**，此时的网络状态则更新为 t S  + 1，控 制器**根据转发后的网络状态给予该条转发路径一个奖励值**，每轮转发后，均对奖 励值进行更新。如果后续某个状态下通过这条路径转发数据，则该状态获得的奖 励值增大，后续数据通过该条路径的概率也会增大。同理，通过该条路径转发数据过多时，其奖励值会逐步减少，直到有另一条奖励值更大的状态出现代替该条 路径转发。



SDN中数据包转发是基于流表的[33]！

[33].Kuźniar M, Perešíni P, Kostić D, et al. Methodology, measurement and analysis of flow table update characteristics in hardware openflow switches[J]. Computer Networks, 2018, 136: 22-36.
