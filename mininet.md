从mininet中退回到命令行：`ctrl+d`

开启客户端：`xterm h1 h2 h3`，可以开很多个



### 手动加流表

##### 绑端口：

`/home/sinet21/P4/behavioral-model/targets/simple_switch/simple_switch --log-console --thrift-port 9091 -i 1@enp59s0f2 -i 2@enp59s0f3 l3fwd-wo-chksm-gen-bmv2.json`

>  交换机路径  --log-console(显示日志的,可选)  -thrift-port  9091(一般从9091开始)  -i(绑哪个网卡的)  端口@以太网网口名(ifconfig中的) (绑多个一样格式)    p4文件名.json

#####  导入流表：

`sudo python runtime_CLI.py --thrift-port 9092 < s1-runtime.txt`

`sudo ./runtime_CLI.py < s1-runtime.txt --thrift-port 9092`

