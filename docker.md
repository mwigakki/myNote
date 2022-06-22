# docker

https://www.w3cschool.cn/docker/docker-tutorial.html

[什么是Docker？看这一篇干货文章就够了！ - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/187505981)

[Docker — 从入门到实践 | Docker 从入门到实践 (docker-practice.com)](https://vuepress.mirror.docker-practice.com/)

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

镜像的唯一标识是其 ID 和摘要（DIGEST），而一个镜像可以有多个标签。

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

#### 其他

##### 分层存储

因为镜像包含操作系统完整的 `root` 文件系统，其体积往往是庞大的，因此在 Docker 设计时，就充分利用 [Union FS (opens new window)](https://en.wikipedia.org/wiki/Union_mount)的技术，将其设计为分层存储的架构。所以严格来说，镜像并非是像一个 `ISO` 那样的打包文件，镜像只是一个虚拟的概念，其实际体现并非由一个文件组成，而是由一组文件系统组成，或者说，由多层文件系统联合组成。

镜像构建时，会一层层构建，前一层是后一层的基础。每一层构建完就不会再发生改变，后一层上的任何改变只发生在自己这一层。比如，删除前一层文件的操作，实际不是真的删除前一层的文件，而是仅在当前层标记为该文件已删除。在最终容器运行的时候，虽然不会看到这个文件，但是实际上该文件会一直跟随镜像。因此，在构建镜像的时候，需要额外小心，每一层尽量只包含该层需要添加的东西，任何额外的东西应该在该层构建结束前清理掉。

分层存储的特征还使得镜像的复用、定制变的更为容易。甚至可以用之前构建好的镜像作为基础层，然后进一步添加新的层，以定制自己所需的内容，构建新的镜像。

##### 镜像体积

如果仔细观察，会注意到，这里标识的所占用空间和在 Docker Hub 上看到的镜像大小不同。比如，`ubuntu:18.04` 镜像大小，在这里是 `63.3MB`，但是在 [Docker Hub (opens new window)](https://hub.docker.com/layers/ubuntu/library/ubuntu/bionic/images/sha256-32776cc92b5810ce72e77aca1d949de1f348e1d281d3f00ebcc22a3adcdc9f42?context=explore)显示的却是 `25.47 MB`。这是因为 Docker Hub 中显示的体积是压缩后的体积。在镜像下载和上传过程中镜像是保持着压缩状态的，因此 Docker Hub 所显示的大小是网络传输中更关心的流量大小。而 `docker image ls` 显示的是镜像下载到本地后，展开的大小，准确说，是展开后的各层所占空间的总和，因为镜像到本地后，查看空间的时候，更关心的是本地磁盘空间占用的大小。

另外一个需要注意的问题是，`docker image ls` 列表中的镜像体积总和并非是所有镜像实际硬盘消耗。由于 Docker 镜像是多层存储结构，并且可以继承、复用，因此不同镜像可能会因为使用相同的基础镜像，从而拥有共同的层。由于 Docker 使用 Union FS，相同的层只需要保存一份即可，因此实际镜像硬盘占用空间很可能要比这个列表镜像大小的总和要小的多。

#####  虚悬镜像

上面的镜像列表中，还可以看到一个特殊的镜像，这个镜像既没有仓库名，也没有标签，均为 `<none>`。：

```bash
<none>               <none>              00285df0df87        5 days ago          342 MB
```

这个镜像原本是有镜像名和标签的，原来为 `mongo:3.2`，随着官方镜像维护，发布了新版本后，重新 `docker pull mongo:3.2` 时，`mongo:3.2` 这个镜像名被转移到了新下载的镜像身上，而旧的镜像上的这个名称则被取消，从而成为了 `<none>`。除了 `docker pull` 可能导致这种情况，`docker build` 也同样可以导致这种现象。由于新旧镜像同名，旧镜像名称被取消，从而出现仓库名、标签均为 `<none>` 的镜像。这类无标签镜像也被称为 **虚悬镜像(dangling image)** ，可以用下面的命令专门显示这类镜像：

```bash
$ docker image ls -f dangling=true
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
<none>              <none>              00285df0df87        5 days ago          342 MB
```

一般来说，虚悬镜像已经失去了存在的价值，是可以随意删除的，可以用下面的命令删除。

```bash
$ docker image prune
```

##### 中间层镜像

为了加速镜像构建、重复利用资源，Docker 会利用 **中间层镜像**。所以在使用一段时间后，可能会看到一些依赖的中间层镜像。默认的 `docker image ls` 列表中只会显示顶层镜像，如果希望显示包括中间层镜像在内的所有镜像的话，需要加 `-a` 参数。	

```bash
$ docker image ls -a
```

这样会看到很多无标签的镜像，与之前的虚悬镜像不同，这些无标签的镜像很多都是中间层镜像，是其它镜像所依赖的镜像。这些无标签镜像不应该删除，否则会导致上层镜像因为依赖丢失而出错。实际上，这些镜像也没必要删除，因为之前说过，相同的层只会存一遍，而这些镜像是别的镜像的依赖，因此并不会因为它们被列出来而多存了一份，无论如何你也会需要它们。只要删除那些依赖它们的镜像后，这些依赖的中间层镜像也会被连带删除。

##### 以特定格式显示

默认情况下，`docker image ls` 会输出一个完整的表格，但是我们并非所有时候都会需要这些内容。比如，刚才删除虚悬镜像的时候，我们需要利用 `docker image ls` 把所有的虚悬镜像的 ID 列出来，然后才可以交给 `docker image rm` 命令作为参数来删除指定的这些镜像，这个时候就用到了 `-q` 参数。

```bash
$ docker image ls -q
5f515359c7f8
05a60462f8ba
fe9198c04d62
```

`--filter` 配合 `-q` 产生出指定范围的 ID 列表，然后送给另一个 `docker` 命令作为参数，从而针对这组实体成批的进行某种操作的做法在 Docker 命令行使用过程中非常常见，不仅仅是镜像，将来我们会在各个命令中看到这类搭配以完成很强大的功能。

另外一些时候，我们可能只是对表格的结构不满意，希望自己组织列；或者不希望有标题，这样方便其它程序解析结果等，这就用到了 [Go 的模板语法 (opens new window)](https://gohugo.io/templates/introduction/)。

比如，下面的命令会直接列出镜像结果，并且只包含镜像ID和仓库名：

```bash
$ docker image ls --format "{{.ID}}: {{.Repository}}"
5f515359c7f8: redis
05a60462f8ba: nginx
fe9198c04d62: mongo
00285df0df87: <none>
```

像其它可以承接多个实体的命令一样，可以使用 `docker image ls -q` 来配合使用 `docker image rm`，这样可以成批的删除希望删除的镜像。我们在“镜像列表”章节介绍过很多过滤镜像列表的方式都可以拿过来使用。

比如，我们需要删除所有仓库名为 `redis` 的镜像：

```bash
$ docker image rm $(docker image ls -q redis)
```

或者删除所有在 `mongo:3.2` 之前的镜像：

```bash
$ docker image rm $(docker image ls -q -f before=mongo:3.2)
```





## 5. Docker 容器 

容器是镜像的运行时实例。正如从虚拟机模板上启动 VM 一样，用户也同样可以从单个镜像上启动一个或多个容器。

虚拟机和容器最大的区别是容器更快并且更轻量级——与虚拟机运行在完整的操作系统之上相比，容器会共享其所在主机的操作系统/内核。

<img src="http://c.biancheng.net/uploads/allimg/190417/4-1Z41G01016212.gif" alt="单个镜像启动多个容器的示意图" style="zoom: 80%;" />

容器的实质是进程，但与直接在宿主执行的进程不同，容器进程运行于属于自己的独立的 [命名空间](https://en.wikipedia.org/wiki/Linux_namespaces)。

前面讲过镜像使用的是分层存储，容器也是如此。每一个容器运行时，是以镜像为基础层，在其上创建一个当前容器的存储层，我们可以称这个为容器运行时读写而准备的存储层为 **容器存储层**。

容器存储层的生存周期和容器一样，容器消亡时，容器存储层也随之消亡。因此，任何保存于容器存储层的信息都会随容器删除而丢失。

按照 Docker 最佳实践的要求，容器不应该向其存储层内写入任何数据，容器存储层要保持无状态化。所有的文件写入操作，都应该使用 [数据卷（Volume）](https://vuepress.mirror.docker-practice.com/data_management/volume.html)、或者 [绑定宿主目录](https://vuepress.mirror.docker-practice.com/data_management/bind-mounts.html)，在这些位置的读写会跳过容器存储层，直接对宿主（或网络存储）发生读写，其性能和稳定性更高。

数据卷的生存周期独立于容器，容器消亡，数据卷不会消亡。因此，使用数据卷后，容器删除或者重新运行之后，数据却不会丢失。

**当利用 `docker run` 来创建容器时，Docker 在后台运行的标准操作包括**：

- 检查本地是否存在指定的镜像，不存在就从 [registry](https://vuepress.mirror.docker-practice.com/repository/) 下载
- 利用镜像创建并启动一个容器
- 分配一个文件系统，并在只读的镜像层外面挂载一层可读写层
- 从宿主主机配置的网桥接口中桥接一个虚拟接口到容器中去
- 从地址池配置一个 ip 地址给容器
- 执行用户指定的应用程序
- 执行完毕后容器被终止

#### 创建并使用容器

创建一个新的容器并运行一个命令:

``` bash
docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]			# contianer可以省略
```

OPTIONS说明：(常用)

- -i: 以交互模式运行容器，通常与 -t 同时使用；
- --name="xxx": 为容器指定一个名称；
- -t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；
- -d: 让 Docker 在后台运行而不是直接把执行命令的结果输出在当前宿主机下

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

#### 保存对容器的修改

使用命令创建一个新的容器，并安装`ping`命令：

``` bash
docker container run -it ubuntu /bin/bash
apt-get update
apt-get upgrade
apt install iputils-ping # ping
```

此时该容器即可使用`ping`了。

当对某一个容器做了修改之后（通过在容器中运行某一个命令），可以把对容器的修改保存下来，并**制作为一个image**，这样下次可以从保存后的最新状态运行该容器。docker中保存状态的过程称之为commit，它保存的新旧状态之间的区别，从而产生一个新的版本。（但是不建议用commit制作镜像，而应该用dockerfile）

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

#### 容器生命周期

容器的生命周期，从**创建、运行、休眠，直至销毁**的整个过程。

使用`docker container run ...` 创建一个容器。

可以根据需要多次停止、启动、暂停以及重启容器，并且这些操作执行得很快。

停止容器就像停止虚拟机一样。尽管已经停止运行，容器的全部配置和内容仍然保存在 Docker 主机的文件系统之中，并且随时可以重新启动。

使用`docker rm <container>` 销毁一个容器。
