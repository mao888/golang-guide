# Linux - Linux 常用

> 本文记录常用的Linux命令, 主要使用CentOS7 

- Linux - Linux 常用
  - 常用
    - [Network](#network)
    - [设置host-name](#设置host-name)
  - 工具
    - [查看端口](#查看端口)
    - [top/htop](#tophtop)

## [#](#常用) 常用

### [#](#network) Network

```bash
[root@docker ~]#  vi /etc/sysconfig/network-scripts/ifcfg-ens160
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=ens160
UUID=ae63abaa-be93-4f70-a7b8-8da53e1c3aa8
DEVICE=ens160
ONBOOT=yes
IPADDR=10.11.39.21
PREFIX=24
GATEWAY=10.11.39.254
DNS1=10.11.105.201
IPV6_PRIVACY=no
```

### [#](#设置host-name) 设置host-name

**centos6**

需要修改两处: 一处是/etc/sysconfig/network，另一处是/etc/hosts，只修改任一处会导致系统启动异常。首先切换到root用户。

```bash
/etc/sysconfig/network
打开该文件，里面有一行 `HOSTNAME=localhost.localdomain，修改 localhost.localdomain 为你的主机名。
/etc/hosts
 打开该文件，会有一行 127.0.0.1 localhost.localdomain localhost其中 127.0.0.1 是本地环路地址， localhost.localdomain 是主机名(hostname)，也就是你待修改的。
```

将上面两个文件修改完后，并不能立刻生效。如果要立刻生效的话，可以用 hostname your-hostname 作临时修改(或者修改完之后执行exec bash立即改变)，它只是临时地修改主机名，查看主机名uname -n系统重启后会恢复原样的。

但修改上面两个文件是永久的，重启系统会得到新的主机名。 最后，重启后查看主机名 uname -n 。

**centos7**

临时生效

```bash
[root@centos7 ~]# hostname 132
[root@centos7 ~]# hostname
132
```

永久生效(不会立刻修改，需重启)

```bash
[root@centos7 ~]# hostnamectl set-hostname centos7
(或者修改完之后执行`exec bash`立即改变)，
```

可以参考如下工具: hostnamectl https://blog.csdn.net/younger_china/article/details/51757979

```bash
ceph@client-node ~]$ hostnamectlstatus
  Static hostname: client-node
         Icon name: computer-vm
           Chassis: vm
        Machine ID:cfc5689e4c90435dbf037c4b600bdba2
           Boot ID:0723cc481fd34048ab20036d0367ffc2
   Virtualization: vmware
 Operating System: CentOS Linux 7 (Core)
       CPE OS Name: cpe:/o:centos:centos:7
            Kernel: Linux 3.10.0-327.el7.x86_64
     Architecture: x86-64
[ceph@client-node ~]$
```

## [#](#工具) 工具

### [#](#查看端口) 查看端口

netstat -nltp

需安装net-tools yum install net-tools

参考

### [#](#top-htop) top/htop

htop

需安装htop yum install htop

**参考**

- 官网 http://hisham.hm/htop/
- htop详解 https://www.cnblogs.com/lazyfang/p/7650010.html
- htop的安装和使用！ https://www.cnblogs.com/enet01/p/8316006.html