# docker

[前言 - Docker — 从入门到实践 (gitbook.io)](https://yeasy.gitbook.io/docker_practice/)

## 1. 容器简介

#### 为什么要用容器

- 传统的应用部署方式是通过插件或脚本来安装应用。这样做的缺点是应用的运行、配置、管理、所有生存周期将与当前操作系统绑定，这样做并不利于应用的升级更新/回滚等操作，当然也可以通过创建虚机的方式来实现某些功能，但是虚拟机非常重，并不利于移植。

- 软件开发的一大目的就是隔离，应用程序在运行时相互独立互不干扰，这种隔离实现起来是很不容易的，其中一种解决方案就是虚拟机技术，通过将应用程序部署在不同的虚拟机中从而实现隔离，另一种就是**容器**技术。

#### 容器相对虚拟机的优势

- **每个虚拟机一套OS，而容器共享一个操作系统/内核**
  
  > 虚拟机模型将底层硬件资源划分到虚拟机当中。每个虚拟机都是包含了虚拟 CPU、虚拟 RAM、虚拟磁盘等资源的一种软件结构。因此，每个虚拟机都需要有自己的操作系统来声明、初始化并管理这些虚拟资源。
  >
  > 但是，**操作系统本身是有其额外开销的**。例如，每个操作系统都消耗一点 CPU、一点 RAM、一点存储空间等。每个操作系统都需要独立的许可证，并且都需要打补丁升级，每个操作系统也都面临被攻击的风险。
  >
  > 通常将这种现象称作 **OS Tax** 或者 **VM Tax**，每个操作系统都占用一定的资源。
  
  > 容器模型具有在宿主机操作系统中运行的单个内核。在一台主机上运行数十个甚至数百个容器都是可能的——**容器共享一个操作系统/内核**。
  >
  > 这意味着只有一个操作系统消耗 CPU、RAM 和存储资源，只有一个操作系统需要授权，只有一个操作系统需要升级和打补丁。同时，只有一个操作系统面临被攻击的风险。简言之，就是只有一份 OS 损耗。
  >
  > 在上述单台机器上只需要运行 4 个业务应用的场景中，也许问题尚不明显。但当需要运行成百上千应用的时候，就会引起质的变化。
  >
  > 同时，容器也保证了不同应用程序的相互隔离。
  >
  > 总之，**容器只隔离应用程序的运行时环境但容器之间可以共享同一个操作系统**
  
- **docker容器的启动时间比虚拟机快很多**
  
  > 因为**容器并不是完整的操作系统**，所以其启动要远比虚拟机快。
  >
  > 在容器内部并不需要内核，也就没有定位、解压以及初始化的过程——更不用提在内核启动过程中对硬件的遍历和初始化了。这些在容器启动的过程中统统都不需要！唯一需要的是位于下层操作系统的共享内核是启动了的！最终结果就是，容器可以在 1s 内启动。
  >
  > 唯一对容器启动时间有影响的就是容器内应用启动所花费的时间。

同时，容器还具有如下优势

- 持续开发、集成和部署：提供可靠且频繁的容器镜像构建/部署，并使用快速和简单的回滚(由于镜像不可变性)。
- 开发和运行相分离：在build或者release阶段创建容器镜像，使得应用和基础设施解耦。
- 开发，测试和生产环境一致性：在本地或外网（生产环境）运行的一致性。
- 资源隔离，且利用高效

#### 容器和虚拟机模型

容器和虚拟机都依赖于宿主机才能运行。宿主机可以是笔记本，是数据中心的物理服务器，也可以是公有云的某个实例。

在下面的示例中，假设宿主机是一台需要运行 4 个业务应用的物理服务器。

在**虚拟机模型**中，首先要开启物理机并启动 Hypervisor 引导程序。一旦 Hypervisor（虚拟机监控程序） 启动，就会占有机器上的全部物理资源，如 CPU、RAM、存储和 NIC。

Hypervisor 接下来就会将这些物理资源划分为虚拟资源，并且看起来与真实物理资源完全一致。

然后 Hypervisor 会将这些资源打包进一个叫作虚拟机（VM）的软件结构当中。这样用户就可以使用这些虚拟机，并在其中安装操作系统和应用。

前面提到需要在物理机上运行 4 个应用，所以在 Hypervisor 之上需要创建 4 个虚拟机并安装 4 个操作系统，然后安装 4 个应用。当操作完成后，结构如下图所示。

<img src="http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01336346.gif" style="zoom: 67%;" />


而**容器模型**则略有不同。

服务器启动之后，所选择的操作系统会启动。在 Docker 世界中可以选择 Linux，或者内核支持内核中的容器原语的新版本 Windows。

与虚拟机模型相同，OS 也占用了全部硬件资源。在 OS 层之上，需要安装容器引擎（如 Docker）。

容器引擎可以获取系统资源，比如进程树、文件系统以及网络栈，接着将资源分割为安全的互相隔离的资源结构，称之为容器。

每个容器看起来就像一个真实的操作系统，在其内部可以运行应用。按照前面的假设，需要在物理机上运行 4 个应用。

因此，需要划分出 4 个容器并在每个容器中运行一个应用，如下图所示。

<img src="http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01424234.gif"  />

从更高层面上来讲：

**Hypervisor 是硬件虚拟化（Hardware Virtualization）——Hypervisor 将硬件物理资源划分为虚拟资源**。

**容器是操作系统虚拟化（OS Virtualization）——容器将系统资源划分为虚拟资源**。



## 2. Docker 架构

**容器是一种通用技术，docker只是其中的一种实现**

Docker 包括三个基本概念:

- 镜像（Image）：Docker 镜像（Image），就相当于是一个 root 文件系统。比如官方镜像 ubuntu:16.04 就包含了完整的一套 Ubuntu16.04 最小系统的 root 文件系统。
- 容器（Container）：镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中的类和实例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。Docker Container 通过 Docker Image来创建。
- 仓库（Repository）：仓库可看成一个代码控制中心，用来保存镜像。

镜像可以理解为一种构建时（build-time）结构，而容器可以理解为一种运行时（run-time）结构，

即可以简单的把**image**理解为**可执行程序**，**container**就是**运行起来的进程**。

还有一个概念：`dockerfile`

- 写程序需要源代码，那么“写”`image`就需要`dockerfile`，`dockerfile`就是`image`的源代码，docker就是"编译器"。

    因此我们只需要在`dockerfile`中指定需要哪些程序、依赖什么样的配置，之后把`dockerfile`交给“编译器”docker进行“编译”，也就是`docker build`命令，生成的可执行程序就是image，之后就可以运行这个image了，这就是`docker run`命令，image运行起来后就是docker container。



Docker 使用**客户端-服务器 (C/S) 架构**模式。

其中docker服务端是一个服务进程，管理着所有的容器。docker客户端则扮演着docker服务端的远程控制器，可以用来控制docker的服务端进程。大部分情况下，docker服务端和客户端运行在一台机器上。



<img src="https://raw.githubusercontent.com/mwigakki/learnGit/master/docker.assets/576507-docker1.png" alt="docker示例图" style="zoom: 150%;" />

| 概念                   | 说明                                                         |
| :------------------------------------ | :-------------------------------------------- |
| Docker 镜像(Images)    | Docker 镜像是用于创建 Docker 容器的模板，比如 Ubuntu 系统。  |
| Docker 容器(Container) | 容器是独立运行的一个或一组应用，是镜像运行时的实体。         |
| Docker 客户端(Client)  | Docker 客户端通过命令行或者其他工具使用 Docker API (https://docs.docker.com/reference/api/docker_remote_api) 与 Docker 的守护进程通信。 |
| docker daemon(守护进程) | daemon的主要功能包括镜像管理、镜像构建、REST API、身份验证、安全、核心网络以及编排。 |
| Docker 主机(Host)      | 一个物理或者虚拟的机器用于执行 Docker 守护进程和容器。       |
| Docker 仓库(Registry)  | Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库。Docker Hub([https://hub.docker.com](https://hub.docker.com/)) 提供了庞大的镜像集合供使用。 |
| Docker Machine         | Docker Machine是一个简化Docker安装的命令行工具，通过一个简单的命令行即可在相应的平台上安装Docker，比如VirtualBox、 Digital Ocean、Microsoft Azure。 |

**docker daemon 工作机制**

Docker Daemon可以认为是通过Docker Server模块接受Docker Client的请求，并在Engine中处理请求，然后根据请求类型，创建出指定的Job并运行，运行过程的作用有以下几种可能：向Docker Registry获取镜像，通过graphdriver执行容器镜像的本地化操作，通过networkdriver执行容器网络环境的配置，通过execdriver执行容器内部运行的执行工作等。



## 3. Ubuntu Docker 安装

**前提条件**

Docker 要求 Ubuntu 系统的内核版本高于 3.10 ，查看本页面的前提条件来验证你的 Ubuntu 版本是否支持 Docker。

通过` uname -r`命令查看你当前的内核版本

**使用脚本安装 Docker**

- 1、获取最新版本的 Docker 安装包

```bash
sudo apt-get update
```

```bash
sudo apt-get install -y docker.io
```

- 2、启动docker 后台服务

```bash
sudo service docker start
```

> 提示：如果使用 Linux，并且还没有将当前用户加入到本地 Docker UNIX 组中，则可能需要在下面的命令前面添加 sudo。

查看docker版本

``` bash
docker --version
```

查看docker系统（可能需要sudo）

``` bash
docker system info
```

如果在 Linux 中遇到无权限访问的问题，需要确认当前用户是否属于本地 Docker UNIX 组。如果不是，可以通过`usermod -aG docker <user>`来添加，然后退出并重新登录 Shell，改动即可生效。

## 4. Docker 镜像 

[Docker镜像（image）详解 (biancheng.net)](http://c.biancheng.net/view/3143.html)

检查 Docker 主机的本地仓库中是否包含镜像。

``` bash
docker image ls docker
docker images
```

各个选项说明:

- **REPOSITORY：**表示镜像的仓库源
- **TAG：**镜像的标签
- **IMAGE ID：**镜像ID
- **CREATED：**镜像创建时间
- **SIZE：**镜像大小

> 通过 docker image 命令的提示可以看到更多的命令

同一仓库源可以有多个 TAG，代表这个仓库源的不同个版本，如ubuntu仓库源里，有15.10、14.04等多个不同的版本，我们使用 **REPOSITORY:TAG** 来定义不同的镜像。

将镜像取到 Docker 主机本地的操作是拉取。

Docker 主机安装之后，本地并没有镜像。

#### 搜索镜像

使用命令：

``` bash
docker search <imageName>
```

通过 CLI 的方式搜索 Docker Hub。

简单模式下，该命令会搜索所有“NAME”字段中包含特定字符串的仓库。

例：

``` bash
root@VM-8-17-ubuntu:/home/ubuntu# docker search mysql
NAME                           DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                          MySQL is a widely used, open-source relation…   12761     [OK]       
mariadb                        MariaDB Server is a high performing open sou…   4899      [OK]       
percona                        Percona Server is a fork of the MySQL relati…   579       [OK]       
phpmyadmin                     phpMyAdmin - A web interface for MySQL and M…   556       [OK]       
bitnami/mysql                  Bitnami MySQL Docker Image                      71                   [OK]
linuxserver/mysql-workbench                                                    37                   
linuxserver/mysql              A Mysql container, brought to you by LinuxSe…   35               
......
```

需要注意，上面返回的镜像中既有官方的也有非官方的（如后三个），可以使用 --filter "is-official=true"，使命令返回内容只显示官方镜像。

#### 拉取镜像

``` bash
docker image pull <repository>:<tag>
```

 是下载镜像的命令。镜像从远程镜像仓库服务的仓库中下载。不写tag时默认是 latest。

其次，标签为 latest 的镜像没有什么特殊魔力！标有 latest 标签的镜像不保证这是仓库中最新的镜像！例如，Alpine 仓库中最新的镜像通常标签是 edge。通常来讲，使用 latest 标签时需要谨慎！

从非官方仓库拉取镜像也是类似的，只需要在仓库名称面前加上 Docker Hub 的用户名或者组织名称。

默认情况下，镜像会从` Docker Hub `的仓库中拉取。`docker image pull alpine:latest` 命令会从 `Docker Hub `的 `alpine` 仓库中拉取标签为 latest 的镜像。再将最新的`ubuntu`也拉取下来看看`docker image pull ubuntu:latest` ，然后使用`docker image ls`查看拉取结果：

```bash
root@VM-8-17-ubuntu:~# docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
ubuntu       latest    27941809078c   13 days ago   77.8MB
alpine       latest    e66264b98777   3 weeks ago   5.53MB
```

可见`Alpine`的大小是非常小的。目前 Docker 官方已开始推荐使用 `Alpine` 替代之前的 `Ubuntu` 做为基础镜像环境。这样会带来多个好处。包括镜像下载速度加快，镜像安全性提高，主机之间的切换更方便，占用更少磁盘空间等。

#### 删除镜像

当读者不再需要某个镜像的时候，可以通过命令从 Docker 主机删除该镜像。

```bash
docker image rm <镜像ID>或<repository>:<tag>
```

#### 推送镜像

首先需要登录docker的站点，在https://hub.docker.com/注册个账号，然后docker login，成功后使用docker tag修改测试运行过的image，然后再docker push 【账号】/【docker】：【tag】，如果没给tag，默认就是latest，如果制定了，以后pull镜像就需要指定tag

#### 利用 commit 理解镜像构

> 注意： `docker commit` 命令除了学习之外，还有一些特殊的应用场合，比如被入侵后保存现场等。但是，**不要使用 `docker commit` 定制镜像**，定制镜像应该使用 `Dockerfile` 来完成。

镜像是容器的基础，每次执行 `docker run` 的时候都会指定哪个镜像作为容器运行的基础。在之前的例子中，我们所使用的都是来自于 Docker Hub 的镜像。直接使用这些镜像是可以满足一定的需求，而当这些镜像无法直接满足需求时，我们就需要定制这些镜像。接下来的几节就将讲解如何定制镜像。

回顾一下之前我们学到的知识，镜像是多层存储，每一层是在前一层的基础上进行的修改；而容器同样也是多层存储，是在以镜像为基础层，在其基础上加一层作为容器运行时的存储层。

现在让我们以定制一个 Web 服务器为例子，来讲解镜像是如何构建的。

``` bash
docker run --name webserver -d -p 80:80 nginx
```

> -d 表示后台运行容器，-p 指定映射到宿主机上的端口，80是http的端口，这样我们可以用浏览器去访问这个 `nginx` 服务器。

> 如果是在本机运行的 Docker，那么可以直接访问：`http://localhost` ，如果是在虚拟机、云服务器上安装的 Docker，则需要将 `localhost` 换为虚拟机地址或者实际云服务器地址。直接用浏览器访问的话，我们会看到默认的 Nginx 欢迎页面。

<img src="D:\record\learnGit\docker.assets\image-20220623105826877.png" alt="nginx欢迎页面" style="zoom:80%;" />

现在，假设我们非常不喜欢这个欢迎页面，我们希望改成欢迎 Docker 的文字，我们可以使用 `docker exec` 命令进入容器，修改其内容。

``` bash
$ docker exec -it webserver bash
root@3729b97e8226:/# echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
root@3729b97e8226:/# exit
exit
```

我们以交互式终端方式进入 `webserver` 容器，并执行了 `bash` 命令，也就是获得一个可操作的 Shell。

然后，我们用 `<h1>Hello, Docker!</h1>` 覆盖了 `/usr/share/nginx/html/index.html` 的内容。

现在我们再刷新浏览器的话，会发现内容被改变了。

<hr/>

我们修改了容器的文件，也就是改动了容器的**存储层**。我们可以通过 `docker diff` 命令看到具体的改动。

现在我们定制好了变化，我们希望能将其保存下来形成镜像。

要知道，当我们运行一个容器的时候（如果不使用卷的话），我们做的任何文件修改都会被记录于容器存储层里。而 Docker 提供了一个 `docker commit` 命令，可以将容器的存储层保存下来成为镜像。换句话说，就是**在原有镜像的基础上，再叠加上容器的存储层，并构成新的镜像**。以后我们运行这个新镜像的时候，就会拥有原有容器最后的文件变化。

`docker commit` 的语法格式为：

``` bash
docker commit [选项] <容器ID或容器名> [<仓库名>[:<标签>]]
```

我们可以用下面的命令将容器保存为镜像：

``` bash
ubuntu@VM-8-17-ubuntu:~$ docker commit --message "modified nginx default index page" webserver nginx:v2
sha256:5098f2c1dd57b81ebeca0467725ea1306b63b32200d52e2a814b0f4abf36779b
```

我们可以在 `docker image ls` 中看到这个新定制的镜像：

我们还可以用 `docker history` 具体查看镜像内的历史记录，如果比较 `nginx:latest` 的历史记录，我们会发现新增了我们刚刚提交的这一层。

``` bash
ubuntu@VM-8-17-ubuntu:~$ docker history nginx:v2
IMAGE          CREATED              CREATED BY                                      SIZE      COMMENT
5098f2c1dd57   About a minute ago   nginx -g daemon off;                            1.19kB    modified nginx default index page
3c58beffe63c   8 hours ago          /bin/sh -c #(nop)  CMD ["nginx" "-g" "daemon…   0B        
<missing>      8 hours ago          /bin/sh -c #(nop)  STOPSIGNAL SIGQUIT           0B        
<missing>      8 hours ago          /bin/sh -c #(nop)  EXPOSE 80                    0B        
<missing>      8 hours ago          /bin/sh -c #(nop)  ENTRYPOINT ["/docker-entr…   0B        
<missing>      8 hours ago          /bin/sh -c #(nop) COPY file:09a214a3e07c919a…   4.61kB
...
```

新的镜像定制好后，我们可以来运行这个镜像。

``` bash
docker run --name web2 -d -p 81:80 nginx:v2
```

这里我们命名为新的服务为 `web2`，并且映射到 `81` 端口。访问 `http://localhost:81` 看到结果，其内容应该和之前修改后的 `webserver` 一样。

至此，我们第一次完成了定制镜像，使用的是 `docker commit` 命令，手动操作给旧的镜像添加了新的一层，形成新的镜像，对镜像多层存储应该有了更直观的感觉。

这里我们命名为新的服务为 `web2`，并且映射到 `81` 端口。访问 `http://localhost:81` 看到结果，其内容应该和之前修改后的 `webserver` 一样。

至此，我们第一次完成了定制镜像，使用的是 `docker commit` 命令，手动操作给旧的镜像添加了新的一层，形成新的镜像，对镜像多层存储应该有了更直观的感觉。



**另一个例子**

使用命令创建一个新的容器，并安装`ping`命令：

``` bash
docker container run -it ubuntu /bin/bash
apt-get update
apt-get upgrade
apt install iputils-ping # ping
```

此时该容器即可使用`ping`了。

然后使用 commit

``` bash
docker commit <容器ID> <保存的名字>
```

此时使用`docker image ls` 可以看到多了一个image

``` bash
root@VM-8-17-ubuntu:/home/ubuntu# docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED         SIZE
learn/ping   latest    bab13b3b508a   4 seconds ago   122MB
ubuntu       latest    27941809078c   2 weeks ago     77.8MB
alpine       latest    e66264b98777   4 weeks ago     5.53MB
```

此时就可以使用新的image，`learn/ping` 来使用ping命令了。

```bash
root@VM-8-17-ubuntu:/home/ubuntu# docker run learn/ping ping www.baidu.com
PING www.a.shifen.com (110.242.68.3) 56(84) bytes of data.
64 bytes from 110.242.68.3 (110.242.68.3): icmp_seq=1 ttl=250 time=10.6 ms
64 bytes from 110.242.68.3 (110.242.68.3): icmp_seq=2 ttl=250 time=10.6 ms
```





## 5. Docker 容器 

容器是镜像的运行时实例。正如从虚拟机模板上启动 VM 一样，用户也同样可以从单个镜像上启动一个或多个容器。

虚拟机和容器最大的区别是容器更快并且更轻量级——与虚拟机运行在完整的操作系统之上相比，容器会共享其所在主机的操作系统/内核。

<img src="http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01016212.gif" alt="单个镜像启动多个容器的示意图" style="zoom: 80%;" />

容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于自己的独立的 [命名空间](https://en.wikipedia.org/wiki/Linux_namespaces)。



#### 创建并使用容器

创建一个新的容器并运行一个命令:

``` bash
docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]			# contianer可以省略
```

OPTIONS说明：(常用)

- -i: 以交互模式运行容器，通常与 -t 同时使用；
- --name="xxx": 为容器指定一个名称；
- -t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；

更多见：[Docker run 命令 | 菜鸟教程 (runoob.com)](https://www.runoob.com/docker/docker-run-command.html)

该命令可以携带很多参数，在其基础的格式中，指定了启动所需的镜像以及要运行的应用。

``` bash
docker container run <image> <app>
```

例：

``` bash
docker container run -it ubuntu /bin/bash
```

则会启动某个 Ubuntu Linux 容器，并运行 Bash Shell 作为其应用。会发现 Shell 提示符发生了变化，说明目前已经位于容器内部了。若尝试在容器内执行一些基础命令，可能会发现某些指令无法正常工作。这是因为大部分容器镜像都是经过高度优化的。这意味着某些命令或者包可能没有安装。

> -it 参数可以将当前终端连接到容器的 Shell 终端之上。

如果想启动 PowerShell 并运行一个应用，则可以使用命令`docker container run -it microsoft- /powershell:nanoserver pwsh.exe`。

容器随着其中运行应用的退出而终止。其中 Linux 容器会在 Bash Shell 退出后终止，而 Windows 容器会在 PowerShell 进程终止后退出。

一个简单的验证方法就是启动新的容器，并运行 sleep 命令休眠 10s。容器会启动，然后运行休眠命令，在 10s 后退出。

如果在 Linux 主机（或者在 Linux 容器模式下的 Windows 主机上）运行`docker container run alpine:latest sleep 10`命令，Shell 会连接到容器 Shell 10s 的时间，然后退出。

下面是常用命令：

| 命令  (参数`<container>`可以是容器 ID 或名称)  | 描述                                             |
| :--------------------------------------------- | :----------------------------------------------- |
| `docker ps` 或 `docker container ls`           | 列出正在运行的容器                               |
| `docker ps -a   ` 或 `docker container ls -a ` | 列出所有容器                                     |
| `docker ps -s` 或 `docker container ls -s `    | 列出正在运行的容器 *（带 CPU/内存）*             |
| `docker images` 或 `docker image ls`           | 列出所有镜像                                     |
| `docker exec -it <container> bash`             | 连接到容器                                       |
| `docker logs <container>`                      | 显示容器的控制台日志                             |
| `docker stop <container>`                      | 停止一个容器                                     |
| `docker start <container>`                     | 启动一个容器                                     |
| `docker restart <container>`                   | 重启一个容器                                     |
| `docker rm <container>`                        | 移除一个stoped容器，在后面加上`-f`可强制销毁容器 |
| `docker port <container>`                      | 显示容器的端口映射                               |
| `docker top <container>`                       | 列出进程                                         |
| `docker kill <container>`                      | 杀死一个running中的容器                          |
| `docker container prune`                       | 清除所有处于终止状态的容器                       |

***docker stop 与 docker kill的区别***

- **docker stop**: *Stop a running container (**send SIGTERM, and then SIGKILL after grace period**) [...] The main process inside the container will receive SIGTERM, and after a grace period, SIGKILL. [emphasis mine]*
- **docker kill**: *Kill a running container (**send SIGKILL, or specified signal**) [...] The main process inside the container will be sent SIGKILL, or any signal specified with option --signal. [emphasis mine]*

- docker stop：支持“**优雅退出**”。先发送SIGTERM信号，在一段时间之后（10s）再发送SIGKILL信号。Docker内部的应用程序可以接收SIGTERM信号，然后做一些“退出前工作”，比如保存状态、处理当前请求等。
- docker kill：发送SIGKILL信号，应用程序直接退出。

#### docker进入、退出容器

查看某一个进程的详细信息：

``` bash
docker inspect <容器ID>
```

**进入容器伪终端命令：**

```bash
docker attach <容器ID>
```

另一条命令也能达到类似的效果：

```bash
docker exec -it <容器ID/容器name> /bin/bash 
```

但此命令是让容器运行最后的那个命令，（当然可以换成其他命令），所以用它进入伪终端再exit不会kill容器

**退出容器命令**

```bash
exit	# exit表示退出并停止容器
```

或者

```bash
Ctrl+P+Q	# 这样退出就不会删除容器
```

容器如果不运行任何进程则无法存在

#### 导入和导出

如果要导出本地某个容器，可以使用 `docker export` 命令。

``` bash
root@VM-8-17-ubuntu:/home/ubuntu# docker container ls -a
CONTAINER ID   IMAGE        COMMAND       CREATED        STATUS        PORTS     NAMES
a547f74dea73   learn/ping   "bash"        13 hours ago   Up 12 hours             silly_jennings
5f48864bc0ba   ubuntu       "/bin/bash"   16 hours ago   Up 13 hours             myname2
583b4f91da24   ubuntu       "bash"        17 hours ago   Up 17 hours             boring_herschel
root@VM-8-17-ubuntu:/home/ubuntu# docker export a54 > myubuntu.tar
root@VM-8-17-ubuntu:/home/ubuntu# ls
myubuntu.tar  node_modules  package-lock.json  samples-server-master
```

这样将导出容器快照到本地文件。

<hr/>

可以使用 `docker import` 从容器快照文件中再导入为镜像，例如:

``` bash
root@VM-8-17-ubuntu:/home/ubuntu# cat myubuntu.tar | docker import - test/myubuntu:v1.0
sha256:d5de5ef0bb18fd0576c25115987aa31c7867d7a44c7c9f75b0a55d20df58c1d0
root@VM-8-17-ubuntu:/home/ubuntu# docker images
REPOSITORY      TAG       IMAGE ID       CREATED          SIZE
test/myubuntu   v1.0      d5de5ef0bb18   12 seconds ago   113MB
```

此外，也可以通过指定 URL 或者某个目录来导入，例如

``` bash
$ docker import http://example.com/exampleimage.tgz example/imagerepo
```

*注：用户既可以使用* `*docker load*` *来导入镜像存储文件到本地镜像库，也可以使用* `*docker import*` *来导入一个容器快照到本地镜像库。这两者的区别在于容器快照文件将丢弃所有的历史记录和元数据信息（即仅保存容器当时的快照状态），而镜像存储文件将保存完整记录，体积也要大。此外，从容器快照文件导入时可以重新指定标签等元数据信息。*

#### 容器生命周期

容器的生命周期，从**创建、运行、休眠，直至销毁**的整个过程。

使用`docker container run ...` 创建一个容器。

可以根据需要多次停止、启动、暂停以及重启容器，并且这些操作执行得很快。

停止容器就像停止虚拟机一样。尽管已经停止运行，容器的全部配置和内容仍然保存在 Docker 主机的文件系统之中，并且随时可以重新启动。

使用`docker rm <container>` 销毁一个容器。



## 6. Docker 仓库

仓库（`Repository`）是集中存放镜像的地方。

一个容易混淆的概念是注册服务器（`Registry`）。实际上注册服务器是管理仓库的具体服务器，每个服务器上可以有多个仓库，而每个仓库下面有多个镜像。从这方面来说，仓库可以被认为是一个具体的项目或目录。例如对于仓库地址 `docker.io/ubuntu` 来说，`docker.io` 是注册服务器地址，`ubuntu` 是仓库名。

大部分时候，并不需要严格区分这两者的概念。

[访问仓库 - Docker — 从入门到实践 (gitbook.io)](https://yeasy.gitbook.io/docker_practice/repository)

## 7. 数据管理

![数据管理](https://3503645665-files.gitbook.io/~/files/v0/b/gitbook-x-prod.appspot.com/o/spaces%2F-M5xTVjmK7ax94c8ZQcm%2Fuploads%2Fgit-blob-5950036bba1c30c0b1ab52a73a94b59bbdd5f57c%2Ftypes-of-mounts.png?alt=media)

#### 数据卷

`数据卷` 是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性：

- `数据卷` 可以在容器之间共享和重用
- 对 `数据卷` 的修改会立马生效
- 对 `数据卷` 的更新，不会影响镜像
- `数据卷` 默认会一直存在，即使容器被删除

>  注意：`数据卷` 的使用，类似于 Linux 下对目录或文件进行 mount，镜像中的被指定为挂载点的目录中的文件会复制到数据卷中（仅数据卷为空时会复制）。

**创建一个数据卷**

``` bash
$ docker volume create my-vol
```

查看所有的 `数据卷`

``` bash
$ docker volume ls

DRIVER              VOLUME NAME
local               my-vol
```

在主机里使用以下命令可以查看指定 `数据卷` 的信息

``` bash
$ docker volume inspect my-vol
[
    {
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/my-vol/_data",
        "Name": "my-vol",
        "Options": {},
        "Scope": "local"
    }
]
```

**启动一个挂载数据卷的容器**

在用 `docker run` 命令的时候，使用 `--mount` 标记来将 `数据卷` 挂载到容器里。在一次 `docker run` 中可以挂载多个 `数据卷`。

下面创建一个名为 `web` 的容器，并加载一个 `数据卷` 到容器的 `/usr/share/nginx/html` 目录。

[数据卷 - Docker — 从入门到实践 (gitbook.io)](https://yeasy.gitbook.io/docker_practice/data_management/volume)

**删除数据卷**

 ``` bash
$ docker volume rm my-vol
 ```

`数据卷` 是被设计用来持久化数据的，它的生命周期独立于容器，Docker 不会在容器被删除后自动删除 `数据卷`，并且也不存在垃圾回收这样的机制来处理没有任何容器引用的 `数据卷`。如果需要在删除容器的同时移除数据卷。可以在删除容器的时候使用 `docker rm -v` 这个命令。

无主的数据卷可能会占据很多空间，要清理请使用以下命令

 ``` bash
 $ docker volume prune
 ```

#### 挂载主机目录

[挂载主机目录 - Docker — 从入门到实践 (gitbook.io)](https://yeasy.gitbook.io/docker_practice/data_management/bind-mounts)

## 8. 网络管理

Network Namespace 是 Linux 内核提供的功能，是实现网络虚拟化的重要功能，它能创建多个隔离的网络空间，它们有独自网络栈信息。不管是虚拟机还是容器，运行的时候仿佛自己都在独立的网络中。而且不同Network Namespace的资源相互不可见，彼此之间无法直接通信。


每个新创建的 Network Namespace 默认有一个本地环回接口 lo，除此之外，所有的其他网络设备(物理/虚拟网络接口，网桥等)只能属于一个 Network Namespace。每个 socket 也只能属于一个 Network Namespace。

可以为每个联网容器创建一个命名空间，当然这需要有足够的物理网卡。当物理网卡不足而需要多个网络容器时，就需要虚拟网卡设备。



Docker 允许通过外部访问容器或容器互联的方式来提供网络服务。

#### 外部访问容器

容器中可以运行一些网络应用，要让外部也可以访问这些应用，可以通过 `-P` 或 `-p` 参数来指定端口映射。

当使用 `-P` 标记时，Docker 会随机映射一个端口到内部容器开放的网络端口。

下载一个nginx做服务器，再开一个它的容器。

``` bash
docker pull nginx
docker run -d -P nginx
```

同样的，可以通过 `docker logs` 命令来查看访问记录。

``` bash
$ docker logs fa

172.17.0.1 - - [25/Aug/2020:08:34:04 +0000] "GET / HTTP/1.1" 200 612 "-" "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0" "-"
```

`-p` 则可以指定要映射的端口，并且，在一个指定端口上只可以绑定一个容器。

支持的格式有 `ip:hostPort:containerPort | ip::containerPort | hostPort:containerPort`。



#### Docker网络类型

Docker提供了5种网络类型，其中常见的两种：bridge及host

- **Bridge**

    **Bridge（网桥）是Docker默认使用的网络类型**，用于**同一主机上**的docker容器相互通信，**网络中的所有容器可以通过IP互相访问，连接到同一个网桥的docker容器可以相互通信**。Bridge网络通过网络接口`docker0` 与主机桥接，启动docke时就会自动创建，新创建的容器默认都会自动连接到这个网络，可以在主机上通过`ifconfig docker0`查看到该网络接口的信息

    ![](https://upload-images.jianshu.io/upload_images/22206660-0386dd7d38199d27.png?imageMogr2/auto-orient/strip)

```  bash
root@VM-8-17-ubuntu:/home/ubuntu# ifconfig docker0
docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        inet6 fe80::42:88ff:fed9:a67d  prefixlen 64  scopeid 0x20<link>
```

- Host

  Host模式下，容器的网络接口不与宿主机网络隔离。**在容器中监听相应端口的应用能够直接被从宿主机访问，即不需要做端口映射**。host网络仅支持Linux，host网络没有与宿主机网络隔离，可能引发**安全隐患或端口冲突**

#### 宿主机访问容器应用

这种是最常见的场景，使用容器部署应用，给其他机器提供服务，此时**容器和宿主机IP是通的，可以直接使用容器的虚拟IP+端口进行访问容器服务**，如果想**使用宿主机的IP访问到容器服务，需要将宿主机的某端口映射到容器的服务端口**，即对外暴露端口供宿主机和其他机器访问，**如果不发布端口，外界将无法访问这些容器**

 
