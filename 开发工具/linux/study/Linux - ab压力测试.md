# Linux - ab压力测试

> ab是apachebench命令的缩写, apache自带的压力测试工具。ab非常实用，它不仅可以对apache服务器进行网站访问压力测试，也可以对或其它类型的服务器进行压力测试。

- Linux - ab压力测试
  - [ab的简介](#ab的简介)
  - [ab的原理](#ab的原理)
  - ab的安装
    - [ab的参数说明](#ab的参数说明)
  - 性能指标
    - [吞吐量（Requests per second）](#吞吐量requests-per-second)
    - [并发连接数（The number of concurrent connections）](#并发连接数the-number-of-concurrent-connections)
    - [并发用户数（Concurrency Level）](#并发用户数concurrency-level)
    - [用户平均请求等待时间（Time per request）](#用户平均请求等待时间time-per-request)
    - [服务器平均请求等待时间（Time per request:across all concurrent requests）](#服务器平均请求等待时间time-per-requestacross-all-concurrent-requests)
  - [ab的应用](#ab的应用)
  - [文章来源](#文章来源)

## [#](#ab的简介) ab的简介

ab是apachebench命令的缩写。

ab是apache自带的压力测试工具。ab非常实用，它不仅可以对apache服务器进行网站访问压力测试，也可以对或其它类型的服务器进行压力测试。比如nginx、tomcat、IIS等

## [#](#ab的原理) ab的原理

ab的原理：ab命令会创建多个并发访问线程，模拟多个访问者同时对某一URL地址进行访问。它的测试目标是基于URL的，因此，它既可以用来测试apache的负载压力，也可以测试nginx、lighthttp、tomcat、IIS等其它Web服务器的压力。

ab命令对发出负载的计算机要求很低，它既不会占用很高CPU，也不会占用很多内存。但却会给目标服务器造成巨大的负载，其原理类似CC攻击。自己测试使用也需要注意，否则一次上太多的负载。可能造成目标服务器资源耗完，严重时甚至导致死机。

## [#](#ab的安装) ab的安装

```java
yum -y install httpd-tools
```

测试安装是否成功：

```bash
[root@vic html]# ab -V
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/
```

### [#](#ab的参数说明) ab的参数说明

```bash
[root@vic html]# ab --help
ab: wrong number of arguments
Usage: ab [options] [http[s]://]hostname[:port]/path
Options are:
    -n requests     Number of requests to perform
    -c concurrency  Number of multiple requests to make
    -t timelimit    Seconds to max. wait for responses
    -b windowsize   Size of TCP send/receive buffer, in bytes
    -p postfile     File containing data to POST. Remember also to set -T
    -u putfile      File containing data to PUT. Remember also to set -T
    -T content-type Content-type header for POSTing, eg.
                    'application/x-www-form-urlencoded'
                    Default is 'text/plain'
    -v verbosity    How much troubleshooting info to print
    -w              Print out results in HTML tables
    -i              Use HEAD instead of GET
    -x attributes   String to insert as table attributes
    -y attributes   String to insert as tr attributes
    -z attributes   String to insert as td or th attributes
    -C attribute    Add cookie, eg. 'Apache=1234. (repeatable)
    -H attribute    Add Arbitrary header line, eg. 'Accept-Encoding: gzip'
                    Inserted after all normal header lines. (repeatable)
    -A attribute    Add Basic WWW Authentication, the attributes
                    are a colon separated username and password.
    -P attribute    Add Basic Proxy Authentication, the attributes
                    are a colon separated username and password.
    -X proxy:port   Proxyserver and port number to use
    -V              Print version number and exit
    -k              Use HTTP KeepAlive feature
    -d              Do not show percentiles served table.
    -S              Do not show confidence estimators and warnings.
    -g filename     Output collected data to gnuplot format file.
    -e filename     Output CSV file with percentages served
    -r              Don't exit on socket receive errors.
    -h              Display usage information (this message)
    -Z ciphersuite  Specify SSL/TLS cipher suite (See openssl ciphers)
    -f protocol     Specify SSL/TLS protocol (SSL2, SSL3, TLS1, or ALL)
```

详情说明：

-n在测试会话中所执行的请求个数。默认时，仅执行一个请求。请求的总数量

-c一次产生的请求个数。默认是一次一个。请求的用户量

-t测试所进行的最大秒数。其内部隐含值是-n 50000，它可以使对服务器的测试限制在一个固定的总时间以内。默认时，没有时间限制。

-V显示版本号并退出。

## [#](#性能指标) 性能指标

### [#](#吞吐量-requests-per-second) 吞吐量（Requests per second）

服务器并发处理能力的量化描述，单位是reqs/s，指的是在某个并发用户数下单位时间内处理的请求数。某个并发用户数下单位时间内能处理的最大请求数，称之为最大吞吐率。 记住：吞吐率是基于并发用户数的。这句话代表了两个含义：

- 吞吐率和并发用户数相关
- 不同的并发用户数下，吞吐率一般是不同的

计算公式：总请求数/处理完成这些请求数所花费的时间，即

```bash
Request per second=Complete requests/Time taken for tests
```

必须要说明的是，这个数值表示当前机器的整体性能，值越大越好。

### [#](#并发连接数-the-number-of-concurrent-connections) 并发连接数（The number of concurrent connections）

并发连接数指的是某个时刻服务器所接受的请求数目，简单的讲，就是一个会话。

### [#](#并发用户数-concurrency-level) 并发用户数（Concurrency Level）

要注意区分这个概念和并发连接数之间的区别，一个用户可能同时会产生多个会话，也即连接数。在HTTP/1.1下，IE7支持两个并发连接，IE8支持6个并发连接，FireFox3支持4个并发连接，所以相应的，我们的并发用户数就得除以这个基数。

### [#](#用户平均请求等待时间-time-per-request) 用户平均请求等待时间（Time per request）

计算公式：处理完成所有请求数所花费的时间/（总请求数/并发用户数），即：

```bash
Time per request=Time taken for tests/（Complete requests/Concurrency Level）
```

### [#](#服务器平均请求等待时间-time-per-request-across-all-concurrent-requests) 服务器平均请求等待时间（Time per request:across all concurrent requests）

计算公式：处理完成所有请求数所花费的时间/总请求数，即：

```bash
Time taken for/testsComplete requests
```

可以看到，它是吞吐率的倒数。

同时，它也等于用户平均请求等待时间/并发用户数，即Time per request/Concurrency Level

## [#](#ab的应用) ab的应用

ab的命令参数比较多，我们经常使用的是-c和-n参数。

`ab -c 10 -n 100 http://www.myvick.cn/index.php`：同时处理100个请求并运行10次index.php

- -c10表示并发用户数为10
- -n100表示请求总数为100

```bash
[root@vic html]# ab -c 10 -n 100 http://www.myvick.cn/index.php
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking www.myvick.cn (be patient).....done


Server Software:        nginx/1.13.6   #测试服务器的名字
Server Hostname:        www.myvick.cn  #请求的URL主机名
Server Port:            80             #web服务器监听的端口

Document Path:          /index.php　　  #请求的URL中的根绝对路径，通过该文件的后缀名，我们一般可以了解该请求的类型
Document Length:        799 bytes       #HTTP响应数据的正文长度

Concurrency Level:      10　　　　　　　　# 并发用户数，这是我们设置的参数之一
Time taken for tests:   0.668 seconds   #所有这些请求被处理完成所花费的总时间 单位秒
Complete requests:      100 　　　　　　  # 总请求数量，这是我们设置的参数之一
Failed requests:        0　　　　　　　　  # 表示失败的请求数量，这里的失败是指请求在连接服务器、发送数据等环节发生异常，以及无响应后超时的情况
Write errors:           0
Total transferred:      96200 bytes　　　 #所有请求的响应数据长度总和。包括每个HTTP响应数据的头信息和正文数据的长度
HTML transferred:       79900 bytes　　　　# 所有请求的响应数据中正文数据的总和，也就是减去了Total transferred中HTTP响应数据中的头信息的长度
Requests per second:    149.71 [#/sec] (mean) #吞吐率，计算公式：Complete requests/Time taken for tests  总请求数/处理完成这些请求数所花费的时间
Time per request:       66.797 [ms] (mean)   # 用户平均请求等待时间，计算公式：Time token for tests/（Complete requests/Concurrency Level）。处理完成所有请求数所花费的时间/（总请求数/并发用户数）
Time per request:       6.680 [ms] (mean, across all concurrent requests) #服务器平均请求等待时间，计算公式：Time taken for tests/Complete requests，正好是吞吐率的倒数。也可以这么统计：Time per request/Concurrency Level
Transfer rate:          140.64 [Kbytes/sec] received  #表示这些请求在单位时间内从服务器获取的数据长度，计算公式：Total trnasferred/ Time taken for tests，这个统计很好的说明服务器的处理能力达到极限时，其出口宽带的需求量。

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    2   0.7      2       5
Processing:     2   26  81.3      3     615
Waiting:        1   26  81.3      3     615
Total:          3   28  81.3      6     618

Percentage of the requests served within a certain time (ms)
  50%      6
  66%      6
  75%      7
  80%      7
  90%     10
  95%    209
  98%    209
  99%    618
 100%    618 (longest request)

#Percentage of requests served within a certain time（ms）这部分数据用于描述每个请求处理时间的分布情况，比如以上测试，80%的请求处理时间都不超过7ms，这个处理时间是指前面的Time per request，即对于单个用户而言，平均每个请求的处理时间
```

## [#](#文章来源) 文章来源

参考资料：

https://www.cnblogs.com/myvic/p/7703973.html

http://www.jb51.net/article/59469.htm

http://blog.csdn.net/caotianyin/article/details/49253055

