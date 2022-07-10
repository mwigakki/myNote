# caddy

Caddy是具有自动HTTPS的HTTP/2 web服务器。它是一个轻量级的、商业支持的web服务器，可以使用Let’s Encrypt自动获取和更新SSL/TLS证书。 

其最出色的功能包括：

- 使用Caddyfile轻松配置。
- 默认情况下自动启用HTTPS（通过“加密”）
- 默认情况下为HTTP/2。换句话说，这对于维护我们网站的安全性很重要。
- 虚拟主机，因此多个站点可以正常工作。
- 实验性QUIC支持最先进的变速箱。
- TLS会话验证密钥轮换可实现更安全的连接。
- 可扩展插件，因为便捷的Web服务器是有用的。
- 在没有外部依赖的任何地方运行。

## 在Ubuntu 20.04上安装Caddy服务器

运行以下命令以添加存储库：

`echo "deb [trusted=yes] https://apt.fury.io/caddy/ /" | sudo tee -a /etc/apt/sources.list.d/caddy-fury.list`

更新APT缓存:

```shell
sudo apt update
```

最后，使用以下命令安装Caddy：

```shell
sudo apt install caddy
```

需要一段时间安装，安装时，请记住，如果使用防火墙，则必须允许访问端口`80`和`443`。

另外，您可以使用systemctl检查Caddy的操作。

```shell
sudo systemctl status caddy
```

如果看到caddy 的active状态时 failed，那就使用下面命令来开启caddy

```shell
sudo systemctl start caddy
```

现在打开Web浏览器，然后转到服务器或域的IP地址，就可以看到caddy默认的欢迎页面了。

![caddy默认的欢迎页面](http://iplayio-cn.litchilab.com/Fo-80dKpDOTSeD-8FGf_g6BUMRE-)



由caddy的欢迎页面可以得知，caddy的配置文件位于 `/etc/caddy/Caddyfile`，默认页面是放在 `/usr/share/caddy`中的，而我们可以将网站页面放在`/var/www/html`.



- Chrome 浏览器开启quic：

Chrome 浏览器访问：`chrome://flags`

启用：`enable-quic`

测试是否已经支持 QUIC：https://quic.nginx.org/
