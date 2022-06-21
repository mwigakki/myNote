# docker

https://www.w3cschool.cn/docker/docker-tutorial.html

[Docker镜像（image）详解 (biancheng.net)](http://c.biancheng.net/view/3143.html)

[什么是Docker？看这一篇干货文章就够了！ - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/187505981)

## 1. 容器简介

软件开发的一大目的就是隔离，应用程序在运行时相互独立互不干扰，这种隔离实现起来是很不容易的，其中一种解决方案就是虚拟机技术，通过将应用程序部署在不同的虚拟机中从而实现隔离，另一种就是**容器**技术。

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

<img src="http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01424234.gif" style="zoom:67%;" />

从更高层面上来讲：

**Hypervisor 是硬件虚拟化（Hardware Virtualization）——Hypervisor 将硬件物理资源划分为虚拟资源**。

**容器是操作系统虚拟化（OS Virtualization）——容器将系统资源划分为虚拟资源**。



> 容器是一种通用技术，docker只是其中的一种实现。

**docker的优势**

- **1、简化程序：**
  Docker 让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 机器上，便可以实现虚拟化。Docker改变了虚拟化的方式，使开发者可以直接将自己的成果放入Docker中进行管理。方便快捷已经是 Docker的最大优势，过去需要用数天乃至数周的任务，在Docker容器的处理下，只需要数秒就能完成。

- **2、避免选择恐惧症：**
  如果你有选择恐惧症，还是资深患者。Docker 帮你 打包你的纠结！比如 Docker 镜像；Docker 镜像中包含了运行环境和配置，所以 Docker 可以简化部署多种应用实例工作。比如 Web 应用、后台应用、数据库应用、大数据应用比如 Hadoop 集群、消息队列等等都可以打包成一个镜像部署。

- **3、节省开支：**
  一方面，云计算时代到来，使开发者不必为了追求效果而配置高额的硬件，Docker 改变了高性能必然高价格的思维定势。Docker 与云的结合，让云空间得到更充分的利用。不仅解决了硬件管理的问题，也改变了虚拟化的方式。

- **4、一致性：**

  容器非常适合持续集成和持续交付（CI / CD）工作流程，请考虑以下示例方案：

  - 您的开发人员在本地编写代码，并使用 Docker 容器与同事共享他们的工作。
  - 他们使用 Docker 将其应用程序推送到测试环境中，并执行自动或手动测试。
  - 当开发人员发现错误时，他们可以在开发环境中对其进行修复，然后将其重新部署到测试环境中，以进行测试和验证。
  - 测试完成后，将修补程序推送给生产环境，就像将更新的镜像推送到生产环境一样简单。

​    Docker 允许开发人员使用您提供的应用程序或服务的本地容器在标准化环境中工作，从而简化了开发的生命周期。

## 3. Docker 架构

Docker 包括三个基本概念:

- 镜像（Image）：Docker 镜像（Image），就相当于是一个 root 文件系统。比如官方镜像 ubuntu:16.04 就包含了完整的一套 Ubuntu16.04 最小系统的 root 文件系统。
- 容器（Container）：镜像（Image）和容器（Container）的关系，就像是面向对象程序设计中的类和实例一样，镜像是静态的定义，容器是镜像运行时的实体。容器可以被创建、启动、停止、删除、暂停等。
- 仓库（Repository）：仓库可看成一个代码控制中心，用来保存镜像。

镜像可以理解为一种构建时（build-time）结构，而容器可以理解为一种运行时（run-time）结构，

Docker 使用**客户端-服务器 (C/S) 架构**模式，使用远程API来管理和创建Docker容器。

Docker 容器 通过 Docker 镜像 来创建。

![docker示例图](https://raw.githubusercontent.com/mwigakki/learnGit/master/docker.assets/576507-docker1.png)

| 概念                   | 说明                                                         |
| :------------------------------------ | :-------------------------------------------- |
| Docker 镜像(Images)    | Docker 镜像是用于创建 Docker 容器的模板，比如 Ubuntu 系统。  |
| Docker 容器(Container) | 容器是独立运行的一个或一组应用，是镜像运行时的实体。         |
| Docker 客户端(Client)  | Docker 客户端通过命令行或者其他工具使用 Docker API (https://docs.docker.com/reference/api/docker_remote_api) 与 Docker 的守护进程通信。 |
| Docker 主机(Host)      | 一个物理或者虚拟的机器用于执行 Docker 守护进程和容器。       |
| Docker 仓库(Registry)  | Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库。Docker Hub([https://hub.docker.com](https://hub.docker.com/)) 提供了庞大的镜像集合供使用。 |
| Docker Machine         | Docker Machine是一个简化Docker安装的命令行工具，通过一个简单的命令行即可在相应的平台上安装Docker，比如VirtualBox、 Digital Ocean、Microsoft Azure。 |

## 4. Ubuntu Docker 安装

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

> docker --version

查看docker系统（可能需要sudo）

> docker system info

## 5. Docker 镜像使用

检查 Docker 主机的本地仓库中是否包含镜像。

> docker image ls

各个选项说明:

- **REPOSITORY：**表示镜像的仓库源
- **TAG：**镜像的标签
- **IMAGE ID：**镜像ID
- **CREATED：**镜像创建时间
- **SIZE：**镜像大小

> 通过docker image 命令的提示可以看到更多的命令

同一仓库源可以有多个 TAG，代表这个仓库源的不同个版本，如ubuntu仓库源里，有15.10、14.04等多个不同的版本，我们使用 **REPOSITORY:TAG** 来定义不同的镜像。

将镜像取到 Docker 主机本地的操作是拉取。

Docker 主机安装之后，本地并没有镜像。

**拉取镜像**

`docker image pull <repository>:<tag>`  是下载镜像的命令。镜像从远程镜像仓库服务的仓库中下载。不屑tag时默认时latest。

其次，标签为 latest 的镜像没有什么特殊魔力！标有 latest 标签的镜像不保证这是仓库中最新的镜像！例如，Alpine 仓库中最新的镜像通常标签是 edge。通常来讲，使用 latest 标签时需要谨慎！

从非官方仓库拉取镜像也是类似的，读者只需要在仓库名称面前加上 Docker Hub 的用户名或者组织名称。

默认情况下，镜像会从` Docker Hub `的仓库中拉取。`docker image pull alpine:latest` 命令会从 `Docker Hub `的 `alpine` 仓库中拉取标签为 latest 的镜像。再将最新的`ubuntu`也拉取下来看看`docker image pull ubuntu:latest` ，然后我们使用`docker image ls`查看拉取结果：

```bash
root@VM-8-17-ubuntu:~# docker image ls
REPOSITORY   TAG       IMAGE ID       CREATED       SIZE
ubuntu       latest    27941809078c   13 days ago   77.8MB
alpine       latest    e66264b98777   3 weeks ago   5.53MB
```

可见`Alpine`的大小是非常小的。目前 Docker 官方已开始推荐使用 `Alpine` 替代之前的 `Ubuntu` 做为基础镜像环境。这样会带来多个好处。包括镜像下载速度加快，镜像安全性提高，主机之间的切换更方便，占用更少磁盘空间等。

**删除镜像**

当读者不再需要某个镜像的时候，可以通过 `docker image rm 镜像ID或<repository>:<tag>` 命令从 Docker 主机删除该镜像。其中，rm 是 remove 的缩写。



## 6. Docker容器（container）详解

容器是镜像的运行时实例。正如从虚拟机模板上启动 VM 一样，用户也同样可以从单个镜像上启动一个或多个容器。

虚拟机和容器最大的区别是容器更快并且更轻量级——与虚拟机运行在完整的操作系统之上相比，容器会共享其所在主机的操作系统/内核。

![单个镜像启动多个容器的示意图](http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01016212.gif)

启动容器的简便方式是使用`docker container run`命令。

该命令可以携带很多参数，在其基础的格式`docker container run <image> <app>`中，指定了启动所需的镜像以及要运行的应用。

`docker container run -it ubuntu /bin/bash`则会启动某个 Ubuntu Linux 容器，并运行 Bash Shell 作为其应用。会发现 Shell 提示符发生了变化，说明目前已经位于容器内部了。若尝试在容器内执行一些基础命令，可能会发现某些指令无法正常工作。这是因为大部分容器镜像都是经过高度优化的。这意味着某些命令或者包可能没有安装。

如果想启动 PowerShell 并运行一个应用，则可以使用命令`docker container run -it microsoft- /powershell:nanoserver pwsh.exe`。

> -it 参数可以将当前终端连接到容器的 Shell 终端之上。

容器随着其中运行应用的退出而终止。其中 Linux 容器会在 Bash Shell 退出后终止，而 Windows 容器会在 PowerShell 进程终止后退出。

一个简单的验证方法就是启动新的容器，并运行 sleep 命令休眠 10s。容器会启动，然后运行休眠命令，在 10s 后退出。

如果在 Linux 主机（或者在 Linux 容器模式下的 Windows 主机上）运行`docker container run alpine:latest sleep 10`命令，Shell 会连接到容器 Shell 10s 的时间，然后退出。

可以使用 `docker container stop` 命令**手动停止容器运行**，并且使用 `docker container start` 再次启动该容器。

如果再也不需要该容器，则使用 `docker container rm` 命令来删除容器。

`docker container ls` 列出所有容器

#### 容器进程

[Docker容器（container）详解 (biancheng.net)](http://c.biancheng.net/view/3150.html)
