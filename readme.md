`Git is a version control system.`
`Git is free software.`

 cmd 进入git怎么退出: 按q 回车

> 在vs中文件后面: 
 A 表示 Added。
 D 表示 Deleted
 M 表示 Modified.
 U 表示 Untracked。

>                就好比玩RPG游戏一样，每通过一关(完成一次的工作)都要保存一下游戏(add 和commit一下)。如果某一关没过去(某一次编辑出现了问题)，我们可以选择读取前一关的状态(回退上一个版本)。并且在你随时想暂停的时候都可以暂停(add,commit保存一个版本)。

# 安装与常用操作

### 安装git（windows）

官网下载后，一直下一步就行。只有把默认编辑器从vim改成VScode需要注意
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
使用命令`git status` 查看当前仓库的状态，(是否有修改没有add或commit)

### 查看上次的修改
`git diff`查看工作区和暂存区差异，
`git diff --cached`查看暂存区和仓库差异，
`git diff HEAD `查看工作区和仓库的差异，

# 版本回退

### 查看提交日志
`git log`命令显示从最近到最远的提交日志 
如果嫌输出信息太多，看得眼花缭乱的，可以试试加上`--pretty=oneline`参数：即 `git log --pretty=oneline`
该命令最主要是用来看版本号的，下面是一些示例：
``` 
76c1fafea484761fc13ad96062e5baa5751863f8 (HEAD -> master) add log
f680593ab67427a0ecc9feb2e06296dd39d09bde add some analogy rhetoric
4e14ce03b86b0418967b0aff285e6005d8ad21d4 commmmmmit
13cdea478b95d596a5e431340453da6873d4b8d4 commmmmmit
6480669c9130f720bd633c11629acc7bbe7a812b delect .txt and add .md
23739862302591dbf1cd0652df14723b851318c8 second commit
d9677798cd9c97b0ba6a3ba6d695c3e06cb78c3a first time commit
```

### 回退
```
首先，Git必须知道当前版本是哪个版本，
在Git中，用HEAD表示当前版本，上一个版本就是HEAD^，上上一个版本就是HEAD^^，
当然往上10个版本写10个^比较容易数不过来，所以写成HEAD~10。
```
使用命令：`git reset --hard HEAD^` 回退到上一个版本 ,(在cmd下要写两个^, 因为cmd中^是转义符号，相当于linux的\)

也可使用命令：`git reset --hard <版本号>` , 版本号不必输全，输入前几个能够唯一分辨就行

Git的版本回退速度非常快，因为Git在内部有个指向当前版本的HEAD指针，当你回退版本的时候，Git仅仅是把HEAD从指向了指定的版本
```
┌────┐
│HEAD│
└────┘
   │
   └──> ○ third commit
        │
        ○ second commit
        │
        ○ first commit
改为指向add distributed：

┌────┐
│HEAD│
└────┘
   │
   │    ○ third commit
   │    │
   └──> ○ second commit
        │
        ○ first commit
```

### 重返未来
如果版本回退之后后悔了想回到之前最新的版本：
重新使用`git reset --hard <版本号>`就行，只要知道想去的版本的版本号就行
但如果之前的命令行关掉了，不记得之前的版本号了，也没有关系
使用命令：`git reflog` 查看所有版本历史

## 工作区和暂存区
- **工作区（Working Directory）**: 就是你在电脑里能看到的目录
- **版本库（Repository）**: 工作区有一个隐藏目录.git，这个不算工作区，而是Git的版本库。

Git的版本库里存了很多东西，其中最重要的就是称为stage（或者叫index）的暂存区，还有Git为我们自动创建的第一个分支master，以及指向master的一个指针叫HEAD。
![工作区与版本库示例图1](https://cdn.liaoxuefeng.com/files/attachments/919020037470528/0)

当修改、删除或添加一些文件进去时(使用了`add或rm`)，工作区里的更改会记录到`stage`中
现在，暂存区的状态就变成这样了：
![工作区与版本库示例图2](https://cdn.liaoxuefeng.com/files/attachments/919020074026336/0)

一旦提交(`commit`)后，如果你又没有对工作区做任何修改，那么工作区就是“干净”的：
```
git status
On branch master
nothing to commit, working tree clean
```
现在版本库变成了这样，暂存区就没有任何内容了：
![工作区与版本库示例图3](https://cdn.liaoxuefeng.com/files/attachments/919020100829536/0)


# 管理修改
- Git比其他版本控制系统设计得优秀，因为Git跟踪并管理的是修改，而非文件。每次修改，如果不用`git add`到暂存区，那就不会加入到`commit`中。
### 丢弃工作区的修改
新版命令`git restore file` 更容易记忆，记住这个就好
命令`git checkout -- file`也可立马丢弃 **工作区** 的修改：
   - `git checkout -- file`命令中的`--`很重要，没有`--`，就变成了“切换到另一个分支”的命令，我们在后面的分支管理中会再次遇到`git checkout`命令。
### 丢弃暂存区的修改
新版命令`git restore --staged <file>` 更容易理解，记住这个就好
命令`git reset HEAD <file>`可以把暂存区的修改撤销掉（unstage），重新放回工作区：
`git reset`命令既可以回退版本，也可以把暂存区的修改回退到工作区。当我们用HEAD时，表示最新的版本。

### 删除和恢复文件
使用命令：`git rm <file>`，相当于是删除工作目录中的test.txt文件,并把此次删除操作提交到了暂存区

**误删文件后的恢复**
1. 手动将文件删除：此时相当与将工作区删除了文件，直接使用使用命令：`git restore <file>` 恢复
2. 使用命令：`git rm <file>`删除了文件，此时删除的操作已经被提交到暂存区staged了。此时需要先使用命令`git restore --staged <file>` 将暂存区的删除恢复到工作区，再使用命令：`git restore <file>` 恢复文件


