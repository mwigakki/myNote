# nginx

[Nginx超详细入门教程，收藏慢慢看 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/348921580)

[Nginx 教程 (w3schools.cn)](https://www.w3schools.cn/nginx/)

[如何在 Ubuntu 20.04 中安装和配置 Nginx - 卡拉云 (kalacloud.com)](https://kalacloud.com/blog/how-to-install-nginx-on-ubuntu-20-04/)

## 一、nginx介绍

### 1. 什么是Nginx？

Nginx (engine x) 是一个高性能的HTTP和反向代理web服务器，同时也提供了IMAP/POP3/SMTP服务。

其特点是占有内存少，并发能力强，事实上nginx的并发能力在同类型的网页服务器中表现较好。

### 2. 为什么要使用Nginx？

随着当今互联网的迅速发展，单点服务器早已无法承载上万个乃至数十万个用户的持续访问。比如一台Tomcat服务器在理想状态下只能够可以承受住2000个左右的并发量，为了解决这个问题，就需要多台Tomcat服务器来进行负载均衡。

那么，应该如何实现负载均衡？Nginx就是其中的一种解决方案，当用户访问网站时，Nginx拦截到这个访问请求，并将其通过轮询的方式均匀地分配到不同的服务器上。

![](https://pic4.zhimg.com/80/v2-19ad1edacec5c87d918b7fcc6574de0f_720w.jpg)

### 3. 什么是正向代理？

正向代理，就是客户端将自己的请求率先发给代理服务器，通过代理服务器将请求转发给服务器。我们常用的VPN就是一种代理服务器，为了可以连上国外的网站，客户端需要使用一个可以连接外网的服务器作为代理，并且客户端能够连接上该代理服务器。

![正向代理](https://pic2.zhimg.com/80/v2-d1da364c1b6c2f0024867c3f8ce49c59_720w.jpg)



### 4. 什么是反向代理？



## 二、Nginx在Linux下的安装

这里演示在Ubuntu20.04 LTS下安装和配置 nginx

### 1. 安装

由于 Nginx 可以从 ubuntu 软件源中获得，因此我们可以使用 apt 来安装 Nginx。

我们可以使用以下命令安装 Nginx 到 Ubuntu 中。

```bash
sudo apt update
sudo apt install nginx
```

选择 Y 来开始安装，apt 会帮你把 Nginx 和它所必备的依赖安装到我们的服务器中。

此时运行命令`curl 本机地址` 就可以看到Nginx的欢迎界面的HTML了，说明我们安装成功。

### 2. 



## 三、Nginx配置入门

### 配置文件



### 反向代理单个服务器



### 反向代理多台服务器





## 四、运行测试

### 启动服务



### 关闭服务



## 五、负载均衡

### 什么是负载均衡？



### 配置Nginx负载均衡



