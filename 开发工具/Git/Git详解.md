# Git 详解

> Git是分布式代码托管的标杆，这里将提供如何学习Git，以及Git工作流的总结，和常用的操作命令。

- 工具详解 - Git 详解
  - [必读](#必读)
  - 常用小结
    - Git操作流程
      - [代码提交和同步代码](#代码提交和同步代码)
      - [代码撤销和撤销同步](#代码撤销和撤销同步)
    - [Git常用操作命令 - 代码提交和同步代码](#git常用操作命令---代码提交和同步代码)
    - Git常用操作命令 - 代码撤销和撤销同步
      - [已修改，但未暂存](#已修改但未暂存)
      - [已暂存，未提交](#已暂存未提交)
      - [已提交，未推送](#已提交未推送)
      - [已推送到远程](#已推送到远程)
    - Git常用操作命令 - 其它常用命令
      - [关联远程仓库](#关联远程仓库)
      - [切换分支](#切换分支)
      - [撤销操作](#撤销操作)
      - [版本回退与前进](#版本回退与前进)
      - [配置属于你的Git](#配置属于你的git)
  - [gitignore](#gitignore)
  - [使用Gource生成版本记录视频](#使用gource生成版本记录视频)

## [#](#必读) 必读

> 如果要问一本就可以学习Git的书或者文章，毫无疑问我会推荐Git Pro2，绝对是良心之作啊

- **Git Pro 2**
  - [Git Pro2英文Github仓库在新窗口打开](https://github.com/progit/progit2)
  - [Git Pro2中文Gitbook在新窗口打开](https://bingohuang.gitbooks.io/progit2/content/01-introduction/sections/about-version-control.html)
  - [Git Pro2对应的中文Markdown版本的仓库地址在新窗口打开](https://github.com/bingohuang/progit2-gitbook)
  - [Git Pro中文阅读在新窗口打开](http://git.oschina.net/progit/index.html)
- **其它资料**
  - [99%的时间在使用的Git命令在新窗口打开](http://imtuzi.com/post/most-used-git-cmd.html)
  - [GIT分支开发模型规范在新窗口打开](https://www.jianshu.com/p/cbd8cf5e232d)
  - [Git - 简明指南在新窗口打开](http://rogerdudler.github.io/git-guide/index.zh.html)
  - [图解 Git在新窗口打开](http://marklodato.github.io/visual-git-guide/index-zh-cn.html)
  - [廖雪峰 : Git 教程在新窗口打开](https://www.liaoxuefeng.com/wiki/0013739516305929606dd18361248578c67b8067c8c017b000)
  - [Learn Git Branching在新窗口打开](https://learngitbranching.js.org/)
- **在线学习**
  - 有个很好的在线学习网站，推荐给大家： [https://learngitbranching.js.org/在新窗口打开](https://learngitbranching.js.org/)
  - ![img](https://www.pdai.tech/images/tool-git-learn-1.png)

## [#](#常用小结) 常用小结

> 读完Git Pro2, 大部分功能适合在使用的时候作为工具来参考，最需要理解的是常用的命令对应的workflow，以及分支管理和合并操作。

### [#](#git操作流程) Git操作流程

#### [#](#代码提交和同步代码) 代码提交和同步代码

![img](https://www.pdai.tech/images/git-four-areas.png)

#### [#](#代码撤销和撤销同步) 代码撤销和撤销同步

![img](https://www.pdai.tech/images/git-five-states.png)

### [#](#git常用操作命令-代码提交和同步代码) Git常用操作命令 - 代码提交和同步代码

- 第零步: 工作区与仓库保持一致
- 第一步: 文件增删改，变为已修改状态
- 第二步: git add ，变为已暂存状态

```bash
$ git status
$ git add --all # 当前项目下的所有更改
$ git add .  # 当前目录下的所有更改
$ git add xx/xx.py xx/xx2.py  # 添加某几个文件
```

- 第三步: git commit，变为已提交状态

```bash
$ git commit -m"<这里写commit的描述>"
```

- 第四步: git push，变为已推送状态

```bash
$ git push -u origin master # 第一次需要关联上
$ git push # 之后再推送就不用指明应该推送的远程分支了
$ git branch # 可以查看本地仓库的分支
$ git branch -a # 可以查看本地仓库和本地远程仓库(远程仓库的本地镜像)的所有分支
```

> 在某个分支下，我最常用的操作如下

```bash
$ git status
$ git add -a
$ git status
$ git commit -m 'xxx'
$ git pull --rebase
$ git push origin xxbranch
```

### [#](#git常用操作命令-代码撤销和撤销同步) Git常用操作命令 - 代码撤销和撤销同步

#### [#](#已修改-但未暂存) 已修改，但未暂存

```bash
$ git diff # 列出所有的修改
$ git diff xx/xx.py xx/xx2.py # 列出某(几)个文件的修改

$ git checkout # 撤销项目下所有的修改
$ git checkout . # 撤销当前文件夹下所有的修改
$ git checkout xx/xx.py xx/xx2.py # 撤销某几个文件的修改
$ git clean -f # untracked状态，撤销新增的文件
$ git clean -df # untracked状态，撤销新增的文件和文件夹

# Untracked files:
#  (use "git add <file>..." to include in what will be committed)
#
#	xxx.py
```

#### [#](#已暂存-未提交) 已暂存，未提交

> 这个时候已经执行过git add，但未执行git commit，但是用git diff已经看不到任何修改。 因为git diff检查的是工作区与暂存区之间的差异。

```bash
$ git diff --cached # 这个命令显示暂存区和本地仓库的差异

$ git reset # 暂存区的修改恢复到工作区
$ git reset --soft # 与git reset等价，回到已修改状态，修改的内容仍然在工作区中
$ git reset --hard # 回到未修改状态，清空暂存区和工作区
```

> git reset --hard 操作等价于 git reset 和 git checkout 2步操作

#### [#](#已提交-未推送) 已提交，未推送

> 执行完commit之后，会在仓库中生成一个版本号(hash值)，标志这次提交。之后任何时候，都可以借助这个hash值回退到这次提交。

```bash
$ git diff <branch-name1> <branch-name2> # 比较2个分支之间的差异
$ git diff master origin/master # 查看本地仓库与本地远程仓库的差异

$ git reset --hard origin/master # 回退与本地远程仓库一致
$ git reset --hard HEAD^ # 回退到本地仓库上一个版本
$ git reset --hard <hash code> # 回退到任意版本
$ git reset --soft/git reset # 回退且回到已修改状态，修改仍保留在工作区中。
```

#### [#](#已推送到远程) 已推送到远程

```java
$ git push -f orgin master # 强制覆盖远程分支
$ git push -f # 如果之前已经用 -u 关联过，则可省略分支名
```

> 慎用，一般情况下，本地分支比远程要新，所以可以直接推送到远程，但有时推送到远程后发现有问题，进行了版本回退，旧版本或者分叉版本推送到远程，需要添加 -f参数，表示强制覆盖。

### [#](#git常用操作命令-其它常用命令) Git常用操作命令 - 其它常用命令

#### [#](#关联远程仓库) 关联远程仓库

- 如果还没有Git仓库，你需要

```bash
$ git init
```

- 如果你想关联远程仓库

```bash
$ git remote add <name> <git-repo-url>
# 例如 git remote add origin https://github.com/xxxxxx # 是远程仓库的名称，通常为 origin
```

- 如果你想关联多个远程仓库

```bash
$ git remote add <name> <another-git-repo-url>
# 例如 git remote add coding https://coding.net/xxxxxx
```

- 忘了关联了哪些仓库或者地址

```bash
$ git remote -v
# origin https://github.com/gzdaijie/koa-react-server-render-blog.git (fetch)
# origin https://github.com/gzdaijie/koa-react-server-render-blog.git (push)
```

- 如果远程有仓库，你需要clone到本地

```bash
$ git clone <git-repo-url>
# 关联的远程仓库将被命名为origin，这是默认的。
```

- 如果你想把别人仓库的地址改为自己的

```bash
$ git remote set-url origin <your-git-url>
```

#### [#](#切换分支) 切换分支

> 新建仓库后，默认生成了master分支

- 如果你想新建分支并切换

```bash
$ git checkout -b <new-branch-name>
# 例如 git checkout -b dev
# 如果仅新建，不切换，则去掉参数 -b
```

- 看看当前有哪些分支

```bash
$ git branch
# * dev
#   master # 标*号的代表当前所在的分支
```

- 看看当前本地&远程有哪些分支

```bash
$ git branch -a
# * dev
#   master
#   remotes/origin/master
```

- 切换到现有的分支

```bash
$ git checkout master
```

- 你想把dev分支合并到master分支

```bash
$ git merge <branch-name>
# 例如 git merge dev
```

- 你想把本地master分支推送到远程去

```bash
$ git push origin master
# 你可以使用git push -u origin master将本地分支与远程分支关联，之后仅需要使用git push即可。
```

- 远程分支被别人更新了，你需要更新代码

```bash
$ git pull origin <branch-name>
# 之前如果push时使用过-u，那么就可以省略为git pull
```

- 本地有修改，能不能先git pull

```bash
$ git stash # 工作区修改暂存
$ git pull  # 更新分支
$ git stash pop # 暂存修改恢复到工作区
```

#### [#](#撤销操作) 撤销操作

- 恢复暂存区文件到工作区

```bash
$ git checkout <file-name>
```

- 恢复暂存区的所有文件到工作区

```bash
$ git checkout .
```

- 重置暂存区的某文件，与上一次commit保持一致，但工作区不变

```bash
$ git reset <file-name>
```

- 重置暂存区与工作区，与上一次commit保持一致

```bash
$ git reset --hard <file-name>
# 如果是回退版本(commit)，那么file，变成commit的hash码就好了。
```

- 去掉某个commit

```bash
$ git revert <commit-hash>
# 实质是新建了一个与原来完全相反的commit，抵消了原来commit的效果
```

- reset回退错误恢复

```bash
$ git reflog #查看最近操作记录
$ git reset --hard HEAD{5} #恢复到前五笔操作
$ git pull origin backend-log #再次拉取代码
```

#### [#](#版本回退与前进) 版本回退与前进

- 查看历史版本

```bash
$ git log
```

- 你可能觉得这样的log不好看，试试这个

```bash
$ git log --graph --decorate --abbrev-commit --all
```

- 检出到任意版本

```bash
$ git checkout a5d88ea
# hash码很长，通常6-7位就够了
```

- 远程仓库的版本很新，但是你还是想用老版本覆盖

```bash
$ git push origin master --force
# 或者 git push -f origin master
```

- 觉得commit太多了? 多个commit合并为1个

```bash
$ git rebase -i HEAD~4
# 这个命令，将最近4个commit合并为1个，HEAD代表当前版本。将进入VIM界面，你可以修改提交信息。推送到远程分支的commit，不建议这样做，多人合作时，通常不建议修改历史。
```

- 想回退到某一个版本

```bash
$ git reset --hard <hash>
# 例如 git reset --hard a3hd73r
# --hard代表丢弃工作区的修改，让工作区与版本代码一模一样，与之对应，--soft参数代表保留工作区的修改。
```

- 想回退到上一个版本，有没有简便方法?

```bash
$ git reset --hard HEAD^
```

- 回退到上上个版本呢?

```bash
$ git reset --hard HEAD^^
# HEAD^^可以换作具体版本hash值。
```

- 回退错了，能不能前进呀

```bash
$ git reflog
# 这个命令保留了最近执行的操作及所处的版本，每条命令前的hash值，则是对应版本的hash值。使用上述的git checkout 或者 git reset命令 则可以检出或回退到对应版本。
```

- 刚才commit信息写错了，可以修改吗

```bash
$ git commit --amend
```

- 看看当前状态吧

```bash
$ git status
```

#### [#](#配置属于你的git) 配置属于你的Git

- 看看当前的配置

```bash
$ git config --list
```

- 估计你需要配置你的名字

```bash
$ git config --global user.name "<name>"
#  --global为可选参数，该参数表示配置全局信息
```

- 希望别人看到你的commit可以联系到你

```bash
$ git config --global user.email "<email address>"
```

- 有些命令很长，能不能简化一下

```bash
$ git config --global alias.logg "log --graph --decorate --abbrev-commit --all"
# 之后就可以开心地使用 git log了
```

## [#](#gitignore) gitignore

[Git 忽略提交 .gitignore在新窗口打开](https://www.jianshu.com/p/74bd0ceb6182)

## [#](#使用gource生成版本记录视频) 使用Gource生成版本记录视频

- 工具下载https://www.cr173.com/soft/761328.html
- 官网 https://github.com/acaudwell/Gource
- 官网 - 转成视频 https://github.com/acaudwell/Gource/wiki/Videos
- 视频预览 http://www.365yg.com/i6595151386688619022/#mid=1592562064545805

------

著作权归@pdai所有 原文链接：https://pdai.tech/md/devops/tool/tool-git.html