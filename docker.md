# docker

https://www.w3cschool.cn/docker/docker-tutorial.html

[Docker镜像（image）详解 (biancheng.net)](http://c.biancheng.net/view/3143.html)

## 1 容器简介

**Linux容器**是与系统其他部分隔离开的一系列进程，从另一个镜像运行，并由该镜像提供支持进程所需的全部文件。容器提供的镜像包含了应用的所有依赖项，因而在从开发到测试再到生产的整个过程中，它都具有可移植性和一致性。

#### 容器不就是虚拟化吗

是，但也不竟然。我们用一种简单方式来思考一下：

- 虚拟化使得许多操作系统可同时在单个系统上运行。

- 容器则可共享同一个操作系统内核，将应用进程与系统其他部分隔离开

## 2. Docker 的优点

- **1、简化程序：**
  Docker 让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 机器上，便可以实现虚拟化。Docker改变了虚拟化的方式，使开发者可以直接将自己的成果放入Docker中进行管理。方便快捷已经是 Docker的最大优势，过去需要用数天乃至数周的 任务，在Docker容器的处理下，只需要数秒就能完成。

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

Docker 使用**客户端-服务器 (C/S) 架构**模式，使用远程API来管理和创建Docker容器。

Docker 容器 通过 Docker 镜像 来创建。

![docker示例图](D:\record\learnGit\docker.assets\576507-docker1.png)

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

查看docker版本

> docker --version

查看docker系统

> docker system info

## 5. Docker 镜像使用

检查 Docker 主机的本地仓库中是否包含镜像。

> docker image 

各个选项说明:

- **REPOSITORY：**表示镜像的仓库源
- **TAG：**镜像的标签
- **IMAGE ID：**镜像ID
- **CREATED：**镜像创建时间
- **SIZE：**镜像大小

同一仓库源可以有多个 TAG，代表这个仓库源的不同个版本，如ubuntu仓库源里，有15.10、14.04等多个不同的版本，我们使用 **REPOSITORY:TAG** 来定义不同的镜像。

将镜像取到 Docker 主机本地的操作是拉取。

Docker 主机安装之后，本地并没有镜像。

`docker image pull` 是下载镜像的命令。镜像从远程镜像仓库服务的仓库中下载。

默认情况下，镜像会从 Docker Hub 的仓库中拉取。docker image pull alpine:latest 命令会从 Docker Hub 的 alpine 仓库中拉取标签为 latest 的镜像。
