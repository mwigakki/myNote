# Nginx

[Nginx 教程 (w3schools.cn)](https://www.w3schools.cn/nginx/)

[如何在 Ubuntu 20.04 中安装和配置 Nginx - 卡拉云 (kalacloud.com)](https://kalacloud.com/blog/how-to-install-nginx-on-ubuntu-20-04/)

## 一、Nginx介绍

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

反向代理与正向代理不同，正向代理是代理了客户端，而反向代理则是代理服务器端。在有多台服务器分布的情况下，为了能让客户端访问到的IP地址都为同一个网站，就需要使用反向代理。

![img](http://pic1.zhimg.com/80/v2-55476d2362c5706329c5693326b303bc_720w.jpg)

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

### 2. 检查防火墙

在测试 Nginx 之前，我们需要调整防火墙，让他允许 Nginx 服务通过。Nginx `ufw` 在安装时会把他自身注册成为服务。

```bash
sudo ufw app list
```

输出结果：

```bash
kalacloud@chuan-server:~$ sudo ufw app list
Available applications:
  Nginx Full
  Nginx HTTP
  Nginx HTTPS
  OpenSSH
```

你可以看到 Nginx 提供了三个配置文件：

- **Nginx Full**

    开端口80 正常，未加密的网络流量

    端口443 TLS / SSL加密的流量

- **Nginx HTTP**

    仅打开端口80 正常，未加密

- **Nginx HTTPS**

    仅打开端口443 TLS / SSL加密

### 3. 检查web服务器

在安装结束后，Ubuntu 会启动 Nginx 。 Web 服务器应该已经在运行了。

我们可以通过 `systemd` 来检查 init 系统状态，确保它正在运行：

```bash
$ systemctl status nginx
kalacloud@chuan-server:~$ systemctl status nginx
● nginx.service - A high performance web server and a reverse proxy server
     Loaded: loaded (/lib/systemd/system/nginx.service; enabled; vendor preset:>
     Active: active (running) since Wed 2020-07-29 07:21:53 UTC; 23min ago
       Docs: man:nginx(8)
   Main PID: 3340 (nginx)
      Tasks: 2 (limit: 2248)
     Memory: 4.3M
     CGroup: /system.slice/nginx.service
             ├─3340 nginx: master process /usr/sbin/nginx -g daemon on; master_>
             └─3341 nginx: worker process

Jul 29 07:21:52 chuan-server systemd[1]: Starting A high performance web server>
Jul 29 07:21:53 chuan-server systemd[1]: Started A high performance web server >
lines 1-13/13 (END)
```

如上所示，这个服务已经成功启动。接下来我们要直接来测试 Nginx 是否可以通过浏览器访问。

在浏览器中输入本机IP即可访问。也可以使用`curl http://内网地址` 来查看是否显示HTML

``` bash
root@VM-8-17-ubuntu:/home/ubuntu# curl 10.0.8.17
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
.....
```

## 三、管理Nginx

我们已经启动起来 Web 服务器了。接下来，让我们来学习一下 Nginx 的基本命令。

要停止Web服务器，输入：(systemctl: system control)

```bash
sudo systemctl stop nginx
```

要在停止时，启动Web服务器，键入：

```bash
sudo systemctl start nginx
```

要停止，然后再次启动该服务，键入：

```bash
sudo systemctl restart nginx
```

如果我们只是修改配置，Nginx 可以在不中断的情况下热加载。我们可以键入：

```bash
sudo systemctl reload nginx
```

默认情况下，Nginx 会在服务器启动时，跟随系统启动，如果我们不想这样，我们可以用这个命令来禁止：

```bash
sudo systemctl disable nginx
```

要重新让系统启动时引导 Nginx 启动，那么我们可以输入：

```bash
sudo systemctl enable nginx
```



## 四、Nginx文件及目录结构

### 内容

- `/var/www/html` 默认的 Web 页面。默认打开可以看到 Nginx 页面。
- `/var/www/html` 实际的 Web 内容。默认请看下只有 Nginx 自己的原生页面。我们可以通过更改 Nginx 配置来更改文件。

### 服务器配置

- `/etc/nginx` Nginx 配置目录。所有 Nginx 的配置文件都在这里。
- `/etc/nginx/nginx.conf` Nginx 的配置文件。大多数全局配置可以通过这个文件来修改。
- `/etc/nginx/sites-available/sites-enabled` 用来存储服务器下每个站点服务器块的目录。 默认情况下 Nginx 不会直接使用目录下的配置文件，需要我们更改配置来告诉 Nginx 来去读。
- `/etc/nginx/sites-enabled/sites-available` 这里是存储已经启用站点服务器块的目录。
- `/etc/nginx/snippets` 这个目录包含一些 Nginx 的配置文件。可打开详细查看这些配置文件到文档进行学习。

### 服务器日志

- `/var/log/nginx/access.log` 这里是 Nginx 到日志文件，对 Web 服务器的每个请求都会记录在这个日志中。
- `/var/log/nginx/error.log` 记录 Nginx 运行过程中发生的错误日志。







