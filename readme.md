`Git is a version control system.`
`Git is free software.`
> 在vs中文件后面:
 U 应该是表示这个文件Unattached
 M 表示Modified
>                

### 安装git（windows）
官网下载后，一直下一步就行。只有把默认编辑器从vim改成VScode
安装git完成后，需要在命令行设置：
> `git config --global user.name "Your Name"`
> `git config --global user.email "email@example.com"`
相当于自报家门

### 初始化
初始化一个Git仓库：
- 新建一个文件夹，在此文件夹中打开CMD
- 使用`git init`命令。
- 就可以在此目录中看到一个`.git`隐藏文件夹了，这就说明此目录已经是一个git仓库了

### 添加文件到Git仓库
分两步：
- 使用命令`git add <file>`，注意，可反复多次使用，添加多个文件；
    - 使用`git add` 后的文件就进入`Staged Changes` 的状态了，修改后还没 add 的话就只是`Changes` 的状态
- 使用命令`git commit -m <message>`，完成。

### 查看当前仓库的状态
使用命令`git status` 查看当前仓库的状态，

### 从仓库中删除文件
使用命令：`git rm <file>`