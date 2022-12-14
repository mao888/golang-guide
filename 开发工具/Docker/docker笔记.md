- [x] 
- [ ] 
- [ ] [Docker Buildx](https://docs.docker.com/buildx/working-with-buildx/)
- [ ]  
- [ ] [Docker Cheat Sheet](https://intellipaat.com/blog/tutorial/devops-tutorial/docker-cheat-sheet/)
- [x] [Docker 入门教程](https://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html)——阮一峰
- [x] 
- [x] [Docker 10分钟快速入门](https://www.bilibili.com/video/BV1s54y1n7Ev/?spm_id_from=333.788.recommend_more_video.-1&vd_source=8321160752e4f07c473e11ebc0dd0a28)

---

## 概念：

> Docker，其实就是可以打包程序和运行环境，把环境和程序一起发布的容器，当你需要发布程序时，你可以使用Docker将运行环境一起发布，其他人拿到你的程序后可以直接运行，避免出现一次编译到处调试的尴尬局面。
> Docker的出现主要就是为了解决，“在我的机器上是正常的，为什么到了你的机器上就不正常了的问题”。

### Dockerfile：
 Dockerfile就像是一个自动化脚本，他主要用来创建镜像，这个过程就好比在虚拟机中安装操作系统和软件一样，只不过是通过Dockerfile这个自动化脚本完成了。
### Image/镜像：![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1656295588174-a9165a1e-c568-4b84-9941-b21c37607099.png#clientId=ub8c8146d-5df1-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=685&id=ue645981b&margin=%5Bobject%20Object%5D&name=image.png&originHeight=685&originWidth=1587&originalType=binary&ratio=1&rotation=0&showTitle=false&size=174838&status=done&style=none&taskId=u9185c982-b8ab-4f35-adaa-a99c1cfb83c&title=&width=1587)
 理解为虚拟机的快照，里面包含了所有部署的应用程序以及它所关联的所有库。通过镜像，可以创建许多不同的Container容器。
```shell
# 列出本机的所有 image 文件。
$ docker image ls

# 删除 image 文件
$ docker image rm [imageName]

# 从 Docker 镜像仓库获取镜像的命令是 docker pull
$ docker pull [选项] [Docker Registry 地址[:端口号]/]仓库名[:标签]

===============================================================================
docker search 镜像id或name：在Docker Hub（或其他镜像仓库如阿里镜像）仓库中搜索关键字的镜像
docker pull 镜像id或name：从仓库中下载镜像，若要指定版本，则要在冒号后指定
docker images：列出已经下载的镜像，查看镜像
docker rmi 镜像id或name：删除本地镜像
docker rmi -f 镜像id或name:  删除镜像
docker build：构建镜像
```
### Container/容器
这里的容器就像一天天运行起来的虚拟机，里面运行了容器，每个容器是独立运行的，它们相互之间不影响。
```shell
# 列出本机正在运行的容器
$ docker container ls

# 列出本机所有容器，包括终止运行的容器
$ docker container ls --all

# 删除容器
$ docker container rm [containerID]

# docker container run命令会从 image 文件，生成一个正在运行的容器实例 
$ docker container run hello-world

# docker container kill 命令手动终止
$ docker container kill [containID]

===============================================================================
docker ps：列出运行中的容器
docker ps -a ： 查看所有容器，包括未运行
docker stop 容器id或name：停止容器
docker kill 容器id：强制停止容器
docker start 容器id或name：启动已停止的容器
docker inspect 容器id：查看容器的所有信息
docker container logs 容器id：查看容器日志
docker top 容器id：查看容器里的进程
docker exec -it 容器id /bin/bash：进入容器
exit：退出容器
docker rm 容器id或name：删除已停止的容器
docker rm -f 容器id：删除正在运行的容器
docker exec -it 容器ID sh :进入容器
```
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1656321179310-f79efae4-8d31-4872-b7b1-b31caa9ec156.png#clientId=ueafdd814-a3d3-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=390&id=u9fd50155&margin=%5Bobject%20Object%5D&name=image.png&originHeight=390&originWidth=1579&originalType=binary&ratio=1&rotation=0&showTitle=false&size=142872&status=done&style=none&taskId=u2bf12774-a3ad-4f36-946d-0b4d6f69f5a&title=&width=1579)
```shell
# 如果希望重复使用容器，它用来启动已经生成、已经停止运行的容器文件
$ docker container start [containerID]

# 查看 docker 容器的输出，即容器里面 Shell 的标准输出
$ docker container logs [containerID]

# 进入一个正在运行的 docker 容器。
# 如果docker run命令运行容器的时候，没有使用-it参数，就要用这个命令进入容器。
# 一旦进入了容器，就可以在容器的 Shell 执行命令了
$ docker container exec -it [containerID] /bin/bash
```
### Dockerfile
#### .dockerignore
在项目的根目录下，新建一个文本文件.dockerignore,表示，这三个路径要排除，不要打包进入 image 文件。如果你没有路径要排除，这个文件可以不新建。
```shell
.git
node_modules
npm-debug.log
```
#### Dockerfile
```dockerfile
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/app .

# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
CMD ["/dist/app"]
```
使用golang:alpine作为编译镜像来编译得到二进制可执行文件的过程，并基于scratch生成一个简单的、非常小的新镜像。我们将二进制文件从命名为builder的第一个镜像中复制到新创建的scratch镜像中。
```dockerfile
FROM golang:alpine AS builder：使用基础镜像golang:alpine来创建我们的镜像。这和我们要创建的镜像一样是一个我们能够访问的存储在Docker仓库的基础镜像。这个镜像运行的是alpine Linux发行版，该发行版的大小很小并且内置了Go，非常适合我们的用例。
Env : 用来设置我们编译阶段需要用的环境变量。
COPY . /build：将当前目录下的所有文件（除了.dockerignore排除的路径），都拷贝进入 image 文件的/app目录。
WORKDIR /build：指定接下来的工作路径为/build。
RUN npm install：在/build目录下，运行npm install命令安装依赖。注意，安装后所有的依赖，都将打包进入 image 文件。
EXPOSE 8888：将容器 8888 端口暴露出来， 允许外部连接这个端口。
CMD ["/dist/app"] : 启动容器时运行的命令
```
上面的 Dockerfile 里面，多了最后一行CMD /dist/app，它表示容器启动后自动执行CMD /dist/app。
你可能会问，RUN命令与CMD命令的区别在哪里？简单说，RUN命令在 image 文件的构建阶段执行，执行结果都会打包进入 image 文件；CMD命令则是在容器启动后执行。另外，一个 Dockerfile 可以包含多个RUN命令，但是只能有一个CMD命令。
注意，指定了CMD命令以后，docker container run命令就不能附加命令了（比如前面的/bin/bash），否则它会覆盖CMD命令。现在，启动容器可以使用下面的命令。
#### 构建镜像
在项目目录下，执行下面的命令创建镜像，并指定镜像名称为goweb_app：
```dockerfile
docker build . -t goweb_app
```
**-t **参数用来指定 image 文件的名字，后面还可以用冒号指定标签。如果不指定，默认的标签就是latest。
运行镜像：
```dockerfile
docker run -p 8888:8888 goweb_app
```
标志位-p用来定义端口绑定。由于容器中的应用程序在端口8888上运行，我们将其绑定到主机端口也是8888。如果要绑定到另一个端口，则可以使用-p $HOST_PORT:8888。例如-p 5000:8888。
#### 生成容器
```dockerfile
$ docker container run -p 8000:3000 -it koa-demo /bin/bash
# 或者
$ docker container run -p 8000:3000 -it koa-demo:0.0.1 /bin/bash
```

- -p参数：容器的 3000 端口映射到本机的 8000 端口。
- -it参数：容器的 Shell 映射到当前的 Shell，然后你在本机窗口输入的命令，就会传入容器。
- koa-demo:0.0.1：image 文件的名字（如果有标签，还需要提供标签，默认是 latest 标签）。
- /bin/bash：容器启动以后，内部第一个执行的命令。这里是启动 Bash，保证用户可以使用 Shell。
#### 分阶段构建镜像
```dockerfile
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/app /

# 需要运行的命令
ENTRYPOINT ["/app"]
```
使用这种技术，我们剥离了使用golang:alpine作为编译镜像来编译得到二进制可执行文件的过程，并基于scratch生成一个简单的、非常小的新镜像。我们将二进制文件从命名为builder的第一个镜像中复制到新创建的scratch镜像中。

### DockerCompose
除了像上面一样使用--link的方式来关联两个容器之外，我们还可以使用Docker Compose来定义和运行多个容器。
Compose是用于定义和运行多容器 Docker 应用程序的工具。通过 Compose，你可以使用 YML 文件来配置应用程序需要的所有服务。然后，使用一个命令，就可以从 YML 文件配置中创建并启动所有服务。
使用Compose基本上是一个三步过程：

1. 使用Dockerfile定义你的应用环境以便可以在任何地方复制。
2. 定义组成应用程序的服务，docker-compose.yml 以便它们可以在隔离的环境中一起运行。
3. 执行 docker-compose up命令来启动并运行整个应用程序。

我们的项目需要两个容器分别运行mysql和bubble_app，我们编写的docker-compose.yml文件内容如下：
```shell
# yaml 配置
version: "3.7"
services:
  mysql8019:
    image: "mysql:8.0.19"
    ports:
      - "33061:3306"
    command: "--default-authentication-plugin=mysql_native_password --init-file /data/application/init.sql"
    environment:
      MYSQL_ROOT_PASSWORD: "root1234"
      MYSQL_DATABASE: "bubble"
      MYSQL_PASSWORD: "root1234"
    volumes:
      - ./init.sql:/data/application/init.sql
  bubble_app:
    build: .
    command: sh -c "./wait-for.sh mysql8019:3306 -- ./bubble ./conf/config.ini"
    depends_on:
      - mysql8019
    ports:
      - "8888:8888"
```
这个 Compose 文件定义了两个服务：bubble_app 和 mysql8019。其中：
##### bubble_app
使用当前目录下的Dockerfile文件构建镜像，并通过depends_on指定依赖mysql8019服务，声明服务端口8888并绑定对外8888端口。
##### mysql8019
mysql8019 服务使用 Docker Hub 的公共 mysql:8.0.19 镜像，内部端口3306，外部端口33061。
这里需要注意一个问题就是，我们的bubble_app容器需要等待mysql8019容器正常启动之后再尝试启动，因为我们的web程序在启动的时候会初始化MySQL连接。这里共有两个地方要更改，第一个就是我们Dockerfile中要把最后一句注释掉：
```shell
# Dockerfile
...
# 需要运行的命令（注释掉这一句，因为需要等MySQL启动之后再启动我们的Web程序）
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```
第二个地方是在bubble_app下面添加如下命令，使用提前编写的wait-for.sh脚本检测mysql8019:3306正常后再执行后续启动Web应用程序的命令：
```shell
command: sh -c "./wait-for.sh mysql8019:3306 -- ./bubble ./conf/config.ini" 
```
当然，因为我们现在要在bubble_app镜像中执行sh命令，所以不能在使用scratch镜像构建了，这里改为使用debian:stretch-slim，同时还要安装wait-for.sh脚本用到的netcat，最后不要忘了把wait-for.sh脚本文件COPY到最终的镜像中，并赋予可执行权限哦。更新后的Dockerfile内容如下：
```shell
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o bubble .

###################
# 接下来创建一个小镜像
###################
FROM debian:stretch-slim

COPY ./wait-for.sh /
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf


# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/bubble /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
        chmod 755 wait-for.sh

# 需要运行的命令
# ENTRYPOINT ["/bubble", "conf/config.ini"]
```
所有的条件都准备就绪后，就可以执行下面的命令跑起来了：
```shell
docker-compose up
```
完整版代码示例，请查看我的github仓库：[https://github.com/Q1mi/deploy_bubble_using_docker](https://github.com/Q1mi/deploy_bubble_using_docker)。
