Git is a version control system.
Git is free software.
== 在vs中文件后面的 U 应该是表示这个文件Unattached
==                M 表示Modified

### 安装git
安装git完成后，需要在命令行设置：
> git config --global user.name "Your Name"
> git config --global user.email "email@example.com"

### 初始化
初始化一个Git仓库，使用`git init`命令。

添加文件到Git仓库，分两步：
- 使用命令git add <file>，注意，可反复多次使用，添加多个文件；
    - 使用git add 后的文件就进入`Staged Changes` 的状态了，修改后还没add 的话就只是`Changes` 的状态
- 使用命令git commit -m <message>，完成。

