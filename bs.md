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



**网络流定义**为符合某种规范以及超时约束的一系列数据包的集合[35]。这里的流规范主要 有：五元组（源 IP 地址、目的 IP 地址、源端口、目的端口、协议）、二元组（源 IP 地址、目 的 IP 地址）、目的 IP 地址等，这里的流规范也称为流标识；超时约束是指超过一定时间不再 通信的流就认为是已结束，是按照时间划分流的一种方式。流表示的是一条端到端的通信状 态，按照方向性又可以分为有向流和无向流，无向流不区分源 IP 地址（源端口）和目的 IP 地 址（目的端口），比如对于二元组无向流，源 IP 地址为 A，目的 IP 地址为 B 的二元组流与源 IP 地址为 B，目的 IP 地址为 A 的流定义为同一条流



**流量矩阵**：

`互联网流量矩阵：入门` ； `结合流内相关性和流间相关性进行流量矩阵预测`

流量矩阵可以完整地描述网络中所有流量的分布，是网络规划、网络管理以及流量工程 （Traffic Engineering）中的关键参数



**流量数据集**：美国 Abilene 骨干网络数据集[48]和无线 Mesh 网络 UCSB（University of California, Santa Barbara）数据集 [49]。

本实验采用 CAIDA passive 2016 数据集[55]，该数据集是来自 CAIDA 的 Equinix-Chicago 监视器的匿名被动流量 trace，是美国西雅图和芝加哥之间的高速主干链路流量数据，其数据 格式为 pcap 文件，此数据集可用于研究互联网流量的特征，包括流量的分布、流长以及流持 续时间等。

[48] The Abilene Dataset 2004[EB/OL].http://www.cs.utexas.edu/~yzhang/research/AbileneTM/. 

[49] The UCSB Dataset 2007[EB/OL].http://moment.cs.ucsb.edu/meshnet/datasets.

[55] The CAIDA Anonymized Internet Traces 2016[EB/OL].http://www.caida.org.





首先在各自交换机上运行程序手机流量信息，当发现大变化流是才通知控制器进行预测等行动

**在单个虚拟或物理网络设备上收集负载和流量统计信息**已经得到了广泛的研究[52,53]，在 设备上收集的数据用于在交换机内做出本地决策。

[52] Basat R B, Einziger G, Friedman R, et al. Heavy hitters in streams and sliding windows[C]//IEEE INFOCOM 2016The 35th Annual IEEE International Conference on Computer Communications. IEEE, 2016: 1-9. 

[53] Basat R B, Einziger G, Friedman R, et al. Optimal elephant flow detection[C]//IEEE INFOCOM 2017-IEEE Conference on Computer Communications. IEEE, 2017: 1-9.

**全局测量**是在数据平面的一些网络设备中收 集测量数据，然后集中式的控制器合并从网络测量点收集的数据，并创建整个流量的网络范 围视图[54]。