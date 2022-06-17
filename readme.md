`Git is a version control system.`
`Git is free software.`
> 在vs中文件后面:
 U 应该是表示这个文件Unattached
 M 表示Modified
>                
就好比玩RPG游戏一样，每通过一关(完成一次的工作)都要保存一下游戏(add 和commit一下)。如果某一关没过去(某一次编辑出现了问题)，我们可以选择读取前一关的状态(回退上一个版本)。并且在你随时想暂停的时候都可以暂停(add,commit保存一个版本)。

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

### 从仓库中删除和恢复文件
使用命令：`git rm <file>`删除
使用命令：`git restore <file>` 恢复


### 查看上次的修改
使用命令：`git diff`，只有在文件修改后还没有add 的情况才能看的difference