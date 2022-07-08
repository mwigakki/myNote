> 在vs中文件后面: 
 A 表示 Added。
 D 表示 Deleted
 M 表示 Modified.
 U 表示 Untracked。

> 就好比玩RPG游戏一样，每通过一关(完成一次的工作)都要保存一下游戏(add 和commit一下)。如果某一关没过去(某一次编辑出现了问题)，我们可以选择读取前一关的状态(回退上一个版本)。并且在你随时想暂停的时候都可以暂停(add,commit保存一个版本)。

 cmd 进入git怎么退出: 按q 回车

 查看分支合并情况： `git log --graph --pretty=oneline --abbrev-commit`

# 安装与常用操作

### 安装git（windows）

官网下载后，一直下一步就行。只有把默认编辑器从vim改成VScode需要注意

安装git完成后，需要在命令行设置：

> `git config --global user.name "Your Name"`
> 
> `git config --global user.email "email@example.com"`

相当于自报家门

然后还要给github加上一个ssh key以便方便地使用github，见后面。

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
```  bash
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
使用命令：`git reset --hard HEAD^` 回退到上一个版本 ,(在cmd下要写两个^, 因为cmd中^是转义符号，相当于linux的\ )

也可使用命令：`git reset --hard <版本号>` , 版本号不必输全，输入前几个能够唯一分辨就行

Git的版本回退速度非常快，因为Git在内部有个指向当前版本的HEAD指针，当你回退版本的时候，Git仅仅是把HEAD从指向了指定的版本
``` bash
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

# 远程仓库
自行注册GitHub账号。由于你的本地Git仓库和GitHub仓库之间的传输是通过SSH加密的，所以，需要一点设置：
- 第1步：创建SSH Key。在用户主目录（windows的.ssh目录c盘：/user(或者是用户)/你的用户名(你自己之前起过的)/.ssh）下，看看有没有.ssh目录，如果有，再看看这个目录下有没有id_rsa和id_rsa.pub这两个文件，如果已经有了，可直接跳到下一步。如果没有，打开Shell（Windows下打开Git Bash或CMD），创建SSH Key：
  > ssh-keygen -t rsa -C "youremail@example.com"

  你需要把邮件地址换成你自己的邮件地址，然后一路回车，使用默认值即可，无需密码
  如果一切顺利的话，可以在用户主目录里找到.ssh目录，里面有id_rsa和id_rsa.pub两个文件，这两个就是SSH Key的秘钥对，**id_rsa是私钥**，不能泄露出去，id_rsa.pub是公钥，可以放心地告诉任何人。

第2步：登陆[GitHub](https://github.com/)，打开“`Account settings`”，“`SSH and GPG Keys`”页面：然后，点“`New SSH Key`”，填上任意Title，在Key文本框里粘贴`id_rsa.pub`文件的内容，最后add即可

为什么GitHub需要SSH Key呢？因为GitHub需要识别出你推送的提交确实是你推送的，而不是别人冒充的，而Git支持SSH协议，所以，GitHub只要知道了你的公钥，就可以确认只有你自己才能推送。
当然，GitHub允许你添加多个Key。假定你有若干电脑，你一会儿在公司提交，一会儿在家里提交，只要把每台电脑的Key都添加到GitHub，就可以在每台电脑上往GitHub推送了。

### 添加远程仓库
登陆GitHub，然后，在右上角找到“Create a new repo”按钮，然后...，创建一个新的仓库。

目前，在GitHub上的这个learngit仓库还是空的，可以看看github给了很多的提示（包括从命令行创建一个新的仓库，或者从已有的仓库push），GitHub告诉我们，可以从这个仓库克隆出新的仓库，也可以把一个已有的本地仓库与之关联，然后，把本地仓库的内容推送到GitHub仓库。

现在，我们根据GitHub的提示，在本地的learngit仓库下运行命令：
> git remote add origin git@github.com:mwigakki/learnGit.git

添加后，远程库的名字就是`origin`，这是Git默认的叫法，也可以改成别的，但是origin这个名字一看就知道是远程库。
下一步，就可以把本地库的所有内容推送到远程库上：
> git push -u origin master

把本地库的内容推送到远程，用`git push`命令，实际上是把当前分支`master`**推送**到远程。

由于远程库是空的，我们第一次推送`master`分支时，加上了`-u`参数，Git不但会把本地的`master`分支内容推送的远程新的`master`分支，还会把本地的`master`分支和远程的`master`分支关联起来，在以后的推送或者拉取时就可以简化命令。

推送成功后，可以立刻在GitHub页面中看到远程库的内容已经和本地一模一样：

#### 推动到远程
从现在起，只要本地作了提交，就可以通过命令：
> git push <远程仓库名> <要推送的分支>

例
> git push origin master

把本地master分支的最新修改推送至GitHub。

#### SSH警告
当你第一次使用Git的clone或者push命令连接GitHub时，会得到一个警告：

``` bash
The authenticity of host 'github.com (xx.xx.xx.xx)' can't be established.
RSA key fingerprint is xx.xx.xx.xx.xx.
Are you sure you want to continue connecting (yes/no)?
```
这是因为Git使用SSH连接，而SSH连接在第一次验证GitHub服务器的Key时，需要你确认GitHub的Key的指纹信息是否真的来自GitHub的服务器，输入yes回车即可。

Git会输出一个警告，告诉你已经把GitHub的Key添加到本机的一个信任列表里了：
> Warning: Permanently added 'github.com' (RSA) to the list of known hosts.

这个警告只会出现一次，后面的操作就不会有任何警告了。
如果你实在担心有人冒充GitHub服务器，输入yes前可以对照GitHub的RSA Key的指纹信息是否与SSH连接给出的一致。

### 删除远程库
如果添加的时候地址写错了，或者就是想删除远程库，可以用git remote rm <name>命令。使用前，建议先用git remote -v查看远程库信息
> git remote -v：
``` bash
origin  git@github.com:michaelliao/learn-git.git (fetch)
origin  git@github.com:michaelliao/learn-git.git (push)
```
然后，根据名字删除，比如删除origin：
> git remote rm origin

此处的“删除”其实是解除了本地和远程的绑定关系，并不是物理上删除了远程库。远程库本身并没有任何改动。要真正删除远程库，需要登录到GitHub，在后台页面找到删除按钮再删除。

“删除”后再次使用

> git remote add origin git@github.com:mwigakki/learnGit.git

可以重新将本地和远程绑定。

### 从远程库克隆
在github项目右上角`CODE`里找到 SSH地址，然后在本地合适的位置打开CMD或git bash，使用命令：
> git clone <ssh地址>

即可将项目克隆下来了

# 分支
在版本回退里，你已经知道，每次提交，Git都把它们串成一条时间线，这条时间线就是一个分支。截止到目前，只有一条时间线，在Git里，这个分支叫主分支，即`master`分支。`HEAD`严格来说不是指向提交，而是指向`master`，`master`才是指向提交的，所以，`HEAD`指向的就是当前分支。

下面是一个简单的分支使用示例：

一开始的时候，`master`分支是一条线，Git用`master`指向最新的提交，再用HEAD指向`master`，就能确定当前分支，以及当前分支的提交点：

![master分支](https://www.liaoxuefeng.com/files/attachments/919022325462368/0)

每次提交，master分支都会向前移动一步，这样，随着你不断提交，master分支的线也越来越长。

当我们创建新的分支，例如`dev`时，Git新建了一个指针叫`dev`，指向`master`相同的提交，再把`HEAD`指向`dev`，就表示当前分支在`dev`上：

![master和dev分支1](https://www.liaoxuefeng.com/files/attachments/919022363210080/l)

Git创建一个分支很快，因为除了增加一个dev指针，改改HEAD的指向，工作区的文件都没有任何变化！

不过，从现在开始，对工作区的修改和提交就是针对dev分支了，比如新提交一次后，dev指针往前移动一步，而master指针不变：

![master和dev分支2](https://www.liaoxuefeng.com/files/attachments/919022387118368/l)

假如我们在dev上的工作完成了(commit)，就可以把dev合并到master上。Git怎么合并呢？最简单的方法，就是直接把master指向dev的当前提交，就完成了合并：

![master和dev分支3](https://www.liaoxuefeng.com/files/attachments/919022412005504/0)

当我们把`dev`分支的工作成果合并到`master`分支上：
``` bash
$ git merge dev
Updating d46f35e..b17d20e
Fast-forward
 readme.txt | 1 +
 1 file changed, 1 insertion(+)
```

`git merge`命令用于合并指定分支到当前分支。合并后，再查看`readme.txt`的内容，就可以看到，和`dev`分支的最新提交是完全一样的。

注意到上面的`Fast-forward`信息，Git告诉我们，这次合并是“快进模式”，也就是直接把`master`指向`dev`的当前提交，所以合并速度非常快。

当然，也不是每次合并都能`Fast-forward`，我们后面会讲其他方式的合并。

所以Git合并分支也很快！就改改指针，工作区内容也不变！

合并完分支后，甚至可以删除dev分支。删除dev分支就是把dev指针给删掉，删掉后，我们就剩下了一条master分支：

![master和dev分支3](https://www.liaoxuefeng.com/files/attachments/919022479428512/0)


**下面是具体命令：**

- 查看当前所有分支：`git branch` , 当前分支前面会标一个*号。
- 创建一个从HEAD处创建一个新分支：`git branch dev`
- 切换分支：`git switch dev`或`git checkout dev`
- 创建并切换分支：`git switch -c dev`或`git checkout -b dev`，切换分支后进行相应的修改和提交commit
- 将指定分支合并到**当前分支**： `git merge <分支名>`
- 删除分支：`git branch -d dev` , `-D`表示强制刪除

注：在dev分支上修改了文件，但是并没有执行git add. git commit命令，然后切换到master分支，仍然能看到dev分支的改动。只是因为所有的修改还只是在工作区，连暂存区都不是，git都还不知道有任何的修改。

### 需要解决冲突的分支合并
准备一个新的分支`feature1`
> git switch -c feature1

对文件进行修改，然后再`feature1`分支上提交。
切换到`master`分支
> git switch master

再对文件进行修改并提交。
现在，master分支和feature1分支各自都分别有新的提交，变成了这样：
![分支冲突](https://www.liaoxuefeng.com/files/attachments/919023000423040/0)

这种情况下，Git无法执行“快速合并”，只能试图把各自的修改合并起来，但这种合并就可能会有冲突，我们试试看：
``` bash
git merge feature1
Auto-merging readme.txt
CONFLICT (content): Merge conflict in readme.txt
Automatic merge failed; fix conflicts and then commit the result.
```
果然冲突了！Git告诉我们，readme.txt文件存在冲突，必须手动解决冲突后再提交。git status也可以告诉我们冲突的文件

``` bash
git status
On branch master
Your branch is ahead of 'origin/master' by 2 commits.
  (use "git push" to publish your local commits)

You have unmerged paths.
  (fix conflicts and run "git commit")
  (use "git merge --abort" to abort the merge)

Unmerged paths:
  (use "git add <file>..." to mark resolution)

	both modified:   readme.txt

no changes added to commit (use "git add" and/or "git commit -a")
```

我们可以直接查看readme.txt的内容：
``` bash
Git has a mutable index called stage.
Git tracks changes of files.
<<<<<< HEAD
Creating a new branch is quick & simple.
=======
Creating a new branch is quick AND simple.
>>>>>>> feature1
```
git已经帮我们做了标记,用`<<<<<<<`，`=======`，`>>>>>>>`标记出不同分支的内容，方便我们对冲突的地方查询修改，

接下来我们就在此基础上修改文件并重新`add`, `commit`，再次`merge`合并即可。

现在，`master`分支和`feature1`分支变成了下图所示：
![冲突分支合并](https://www.liaoxuefeng.com/files/attachments/919023031831104/0)

用带参数的`git log`也可以看到分支的合并情况：
``` bash
git log --graph --pretty=oneline --abbrev-commit
*   cf810e4 (HEAD -> master) conflict fixed
|\  
| * 14096d0 (feature1) AND simple
* | 5dc6824 & simple
|/  
* b17d20e branch test
* d46f35e (origin/master) remove test.txt
* b84166e add test.txt
```

最后，删除feature1分支

### 分支管理策略
通常，合并分支时，如果可能，Git会用`Fast forward`模式，但这种模式下，删除分支后，会丢掉分支信息。

如果要强制禁用`Fast forward`模式，Git就会在`merge`时生成一个新的`commit`，这样，从分支历史上就可以看出分支信息。
下面我们实战一下`--no-ff`方式的g`it merge`：
首先，仍然创建并切换dev分支：
> git switch -c dev
> Switched to a new branch 'dev'

修改readme.txt文件，并提交一个新的commit：
``` bash
$ git add readme.txt 
$ git commit -m "add merge"
[dev f52c633] add merge
 1 file changed, 1 insertion(+)
```

现在，我们切换回master：
``` bash
$ git switch master
Switched to branch 'master'
```

准备合并dev分支，请注意--no-ff参数，表示禁用Fast forward：
``` bash
git merge --no-ff -m "merge with no-ff" dev
Merge made by the 'recursive' strategy.
 readme.txt | 1 +
 1 file changed, 1 insertion(+)
```

因为本次合并要创建一个新的commit，所以加上-m参数，把commit描述写进去。

合并后，我们用`git log`看看分支历史：

``` bash
$ git log --graph --pretty=oneline --abbrev-commit
*   e1e9c68 (HEAD -> master) merge with no-ff
|\  
| * f52c633 (dev) add merge
|/  
*   cf810e4 conflict fixed
...

而使用fast forward方式的就是这样的
* 1496baa (HEAD -> master, dev) ff commit in dev
*   2328ae0 merge with no-ff
```
可以看到，不使用Fast forward模式，merge后就像这样：
![不使用Fast forward模式merge](https://www.liaoxuefeng.com/files/attachments/919023225142304/0)

**分支策略**
在实际开发中，我们应该按照几个基本原则进行分支管理：

首先，master分支应该是非常稳定的，也就是仅用来发布新版本，平时不能在上面干活；

那在哪干活呢？干活都在dev分支上，也就是说，dev分支是不稳定的，到某个时候，比如1.0版本发布时，再把dev分支合并到master上，在master分支发布1.0版本；

你和你的小伙伴们每个人都在dev分支上干活，每个人都有自己的分支，时不时地往dev分支上合并就可以了。

所以，团队合作的分支看起来就像这样：
![合作开发](https://www.liaoxuefeng.com/files/attachments/919023260793600/0)

dev分支的内容
dev，然后使用fast forward方式提交

### bug分支
软件开发中，bug就像家常便饭一样。有了bug就需要修复，在Git中，由于分支是如此的强大，所以，每个bug都可以通过一个新的临时分支来修复，修复后，合并分支，然后将临时分支删除。

当你接到一个修复一个代号101的bug的任务时，很自然地，你想创建一个分支issue-101来修复它，但是，等等，当前正在dev上进行的工作还没有提交：

``` bash
$ git status
On branch dev
Changes to be committed:
  (use "git reset HEAD <file>..." to unstage)

	new file:   hello.py

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   readme.txt
```

并不是你不想提交，而是工作只进行到一半，还没法提交，预计完成还需1天时间。但是，必须在两个小时内修复该bug，怎么办？

幸好，Git还提供了一个`stash`功能，可以把当前工作现场“储藏”起来，等以后恢复现场后继续工作：
``` bash
git stash
Saved working directory and index state WIP on dev: f52c633 add merge
```
<hr/>
[此句写于issue-101分支，作为bug分支的练习] : 现在修复了一个bug
<hr/>

现在，用git status查看工作区，就是干净的（除非有没有被Git管理的文件），因此可以放心地创建分支来修复bug。

首先确定要在哪个分支上修复bug，假定需要在master分支上修复，就从master创建临时分支：
``` bash
git checkout master
Switched to branch 'master'
Your branch is ahead of 'origin/master' by 6 commits.
  (use "git push" to publish your local commits)

git checkout -b issue-101
Switched to a new branch 'issue-101'
```

现在修复bug，...，然后提交：
``` bash
git add readme.txt 
git commit -m "fix bug 101"
[issue-101 4c805e2] fix bug 101
 1 file changed, 1 insertion(+), 1 deletion(-)
```

修复完成后，切换到master分支，并完成合并，最后删除issue-101分支：
``` bash
git switch master
Switched to branch 'master'
Your branch is ahead of 'origin/master' by 6 commits.
  (use "git push" to publish your local commits)

git merge --no-ff -m "merged bug fix 101" issue-101
Merge made by the 'recursive' strategy.
 readme.txt | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```

太棒了，原计划两个小时的bug修复只花了5分钟！现在，是时候接着回到dev分支！
使用`git status`工作区是干净的，刚才的工作现场存到哪去了？用`git stash list`命令看看：
``` bash
git stash list
stash@{0}: WIP on dev: f52c633 add merge
```
工作现场还在，Git把stash内容存在某个地方了，但是需要恢复一下，有两个办法：

一是用`git stash apply`恢复，但是恢复后，stash内容并不删除，你需要用`git stash drop`来删除；

另一种方式是用`git stash pop`，恢复的同时把stash内容也删了：

再用git stash list查看，就看不到任何stash内容了：

> git stash list

你可以多次stash，恢复的时候，先用git stash list查看，然后恢复指定的stash，用命令：

> git stash apply stash@{0}

在master分支上修复了bug后，我们要想一想，dev分支是早期从master分支分出来的，所以，这个bug其实在当前dev分支上也存在。

那怎么在dev分支上修复同样的bug？重复操作一次，提交不就行了？

有木有更简单的方法？

有！

同样的bug，要在dev上修复，我们只需要把4c805e2 fix bug 101这个提交所做的修改“复制”到dev分支。注意：我们只想复制4c805e2 fix bug 101这个提交所做的修改，并不是把整个master分支merge过来。

为了方便操作，Git专门提供了一个`cherry-pick`命令，让我们能复制一个特定的提交到当前分支：
``` bash
git branch
* dev
  master
$ git cherry-pick 4c805e2
[master 1d4b803] fix bug 101
 1 file changed, 1 insertion(+), 1 deletion(-)
```

`git cherry-pick`可以理解为”挑拣”提交，它会获取某一个分支的单笔提交，并作为一个新的提交引入到你当前分支上。当我们需要在本地合入其他分支的提交时，如果我们不想对整个分支进行合并，而是只想将某一次提交合入到本地当前分支上，那么就要使用`git cherry-pick`了。

Git自动给dev分支做了一次提交，注意这次提交的commit是1d4b803，它并不同于master的4c805e2，因为这两个commit只是改动相同，但确实是两个不同的commit。用git cherry-pick，我们就不需要在dev分支上手动再把修bug的过程重复一遍。

有些聪明的童鞋会想了，既然可以在master分支上修复bug后，在dev分支上可以“重放”这个修复过程，那么直接在dev分支上修复bug，然后在master分支上“重放”行不行？当然可以，不过你仍然需要git stash命令保存现场，才能从dev分支切换到master分支。


### Feature分支
软件开发中，总有无穷无尽的新的功能要不断添加进来。

添加一个新功能时，你肯定不希望因为一些实验性质的代码，把主分支搞乱了，所以，每添加一个新功能，最好新建一个feature分支，在上面开发，完成后，合并，最后，删除该feature分支。

现在，你终于接到了一个新任务：开发代号为Vulcan的新功能，该功能计划用于下一代星际飞船。

于是准备开发:

> git switch -c feature-vulcan
> Switched to a new branch 'feature-vulcan'

5分钟后，开发完毕：使用`add`， `commit`将开发提交到feature分支

切回`dev`，`git switch dev`准备合并

一切顺利的话，feature分支和bug分支是类似的，合并，然后删除。

但是！

就在此时，接到上级命令，因经费不足，新功能必须取消！

虽然白干了，但是这个包含机密资料的分支还是必须就地销毁：

``` bash
git branch -d feature-vulcan
error: The branch 'feature-vulcan' is not fully merged.
If you are sure you want to delete it, run 'git branch -D feature-vulcan'.
```

销毁失败。Git友情提醒，`feature-vulcan`分支还没有被合并，如果删除，将丢失掉修改，如果要强行删除，需要使用大写的`-D`参数。

现在我们强行删除：
``` bash
$ git branch -D feature-vulcan
Deleted branch feature-vulcan (was 287773e).
```

终于删除成功！

总结：

开发一个新feature，最好新建一个分支；

如果要丢弃一个没有被合并过的分支，可以通过git branch -D <name>强行删除。

### 抓取分支
多人协作时，大家都会往master和dev分支上推送各自的修改。

当你的小伙伴（注意要把SSH Key添加到GitHub）从远程库clone时，默认情况下，你的小伙伴只能看到本地的master分支。

现在，你的小伙伴要在dev分支上开发，就必须创建远程origin的dev分支到本地，于是他用这个命令创建本地dev分支：

> git switch -c dev origin/dev

现在，他就可以在dev上继续修改，然后，时不时地把dev分支push到远程：

... 

现在，你的小伙伴已经向origin/dev分支推送了他的提交，而碰巧你也对同样的文件作了修改，并试图推送：
``` bash
git push origin dev
To github.com:michaelliao/learngit.git
 ! [rejected]        dev -> dev (non-fast-forward)
error: failed to push some refs to 'git@github.com:michaelliao/learngit.git'
hint: Updates were rejected because the tip of your current branch is behind
hint: its remote counterpart. Integrate the remote changes (e.g.
hint: 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
```

推送失败，因为你的小伙伴的最新提交和你试图推送的提交有冲突，解决办法也很简单，Git已经提示我们，先用git pull把最新的提交从origin/dev抓下来，然后，在本地合并，解决冲突，再推送：

``` bash
$ git pull
There is no tracking information for the current branch.
Please specify which branch you want to merge with.
See git-pull(1) for details.

    git pull <remote> <branch>

If you wish to set tracking information for this branch you can do so with:

    git branch --set-upstream-to=origin/<branch> dev
```
`git pull`也失败了，原因是没有指定本地dev分支与远程origin/dev分支的链接，根据提示，设置dev和origin/dev的链接：

``` bash
git branch --set-upstream-to=origin/dev dev
Branch 'dev' set up to track remote branch 'dev' from 'origin'.
```

再pull：

``` bash
$ git pull
Auto-merging env.txt
CONFLICT (add/add): Merge conflict in env.txt
Automatic merge failed; fix conflicts and then commit the result.
```

这回git pull成功，但是合并有冲突，需要手动解决，解决的方法和分支管理中的解决冲突完全一样。解决后，提交，再push：

因此，多人协作的工作模式通常是这样：
1. 首先，可以试图用`git push origin <branch-name>`推送自己的修改；
2. 如果推送失败，则因为远程分支比你的本地更新，需要先用git pull试图合并；
3. 如果合并有冲突，则解决冲突，并在本地提交；
4. 没有冲突或者解决掉冲突后，再用`git push origin <branch-name>`推送就能成功！

如果`git pull`提示`no tracking information`，则说明本地分支和远程分支的链接关系没有创建，用命令`git branch --set-upstream-to <branch-name> origin/<branch-name>`。

### rebase

- rebase操作可以把本地未push的分叉提交历史整理成直线；
- rebase的目的是使得我们在查看历史提交的变化时更容易，因为分叉的提交需要三方对比。

# 标签管理
发布一个版本时，我们通常先在版本库中打一个标签（tag），这样，就唯一确定了打标签时刻的版本。将来无论什么时候，取某个标签的版本，就是把那个打标签的时刻的历史版本取出来。所以，标签也是版本库的一个**快照**。

Git的标签虽然是版本库的快照，但其实它就是指向某个commit的指针（跟分支很像对不对？但是分支可以移动，标签不能移动），所以，创建和删除标签都是瞬间完成的。

在Git中打标签非常简单，首先，切换到需要打标签的分支上：

然后，敲命令`git tag <name>`就可以打一个新标签：

可以用命令`git tag`查看所有标签：

默认标签是打在最新提交的commit上的。有时候，如果忘了打标签，比如，现在已经是周五了，但应该在周一打的标签没有打，怎么办？方法是找到历史提交的commit id，然后打上就可以了：
``` bash
git log --pretty=oneline --abbrev-commit
2dff384 (HEAD -> dev, tag: v1.0, origin/dev) fix confilct 6.21 15.27
c6d0545 cmt 6.21 15.25
51e6b23 cmt 6.21 15.19
3ecce4b feature 6.21 14.45
0b9e615 cmt dev in dorm 6.20 22.43
b253c9a cmt in dorm 6.20 22.38
...
```

比方说要对`cmt 6.21 15.19`这次提交打标签，它对应的commit id是`51e6b23`，敲入命令：
> `git tag v0.9 51e6b23`

再用命令git tag查看标签：
``` bash
$ git tag
v0.9
v1.0
```

注意，标签不是按时间顺序列出，而是按字母排序的。可以用`git show <tagname>`查看标签信息。
还可以创建带有说明的标签，用-a指定标签名，-m指定说明文字：
例： ` git tag -a v0.1 -m "version 0.1 released" 1094adb`

 - 注意：标签总是和某个commit挂钩。如果这个commit既出现在master分支，又出现在dev分支，那么在这两个分支上都可以看到这个标签。

删除标签：
`git tag -d v0.1`

因为创建的标签都只存储在本地，不会自动推送到远程。所以，打错的标签可以在本地安全删除。

如果要推送某个标签到远程，使用命令

`git push origin <tagname>`：

或者，一次性推送全部尚未推送到远程的本地标签：

`git push origin --tags`

如果标签已经推送到远程，要删除远程标签就麻烦一点，先从本地删除：
然后，从远程删除。删除命令也是push，但是格式如下：

`git push origin :refs/tags/v0.9`

要看看是否真的从远程库删除了标签，可以登陆GitHub查看。

