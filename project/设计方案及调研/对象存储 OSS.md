### OSS工作原理

数据以对象（Object）的形式存储在OSS的存储空间（Bucket ）中。如果要使用OSS存储数据，您需要先创建Bucket，并指定Bucket的地域、访问权限、存储类型等属性。创建Bucket后，您可以将数据以Object的形式上传到Bucket，并指定Object的文件名（Key）作为其唯一标识。

当前使用Go接入: [aliyun/aliyun-oss-go-sdk: Aliyun OSS SDK for Go (github.com)](https://github.com/aliyun/aliyun-oss-go-sdk)

服务端使用流程：[服务端签名直传并设置上传回调概述 (aliyun.com)](https://help.aliyun.com/document_detail/31927.html)

缩略图使用教程：[图片缩放 (aliyun.com)](https://help.aliyun.com/document_detail/44688.html)

视频截图使用教程：[视频截帧 (aliyun.com)](https://help.aliyun.com/document_detail/64555.html)



![process](https://static-aliyun-doc.oss-accelerate.aliyuncs.com/assets/img/zh-CN/0844458061/p203792.jpg)

### Web端上传介绍

市面常见方案



![时序图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/7354449951/p140018.png)

和数据直传到OSS相比，以上方法存在以下缺点：

- 上传慢：用户数据需先上传到应用服务器，之后再上传到OSS，网络传输时间比直传到OSS多一倍。如果用户数据不通过应用服务器中转，而是直传到OSS，速度将大大提升。而且OSS采用BGP带宽，能保证各地各运营商之间的传输速度。
- 扩展性差：如果后续用户数量逐渐增加，则应用服务器会成为瓶颈。
- 费用高：需要准备多台应用服务器。由于OSS上行流量是免费的，如果数据直传到OSS，将节省多台应用服务器的费用。

### 服务端签名直传并设置上传回调

客户端直传方案



![流程图CN](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/4172710461/p374419.png)

返回结果:

```json
{

 "accessid":"LTAI5tAzivUnv4ZF1azP****",

 "host":"[http://post-test.oss-cn-hangzhou.aliyuncs.com](http://post-test.oss-cn-hangzhou.aliyuncs.com/)",  

 "policy":"eyJleHBpcmF0aW9uIjoiMjAxNS0xMS0wNVQyMDoyMzoyM1oiLCJjxb25kaXRpb25zIjpbWyJjcb25XC8i****",

 "signature":"I2u57FWjTKqX/AE6doIdyff1****",     

 "expire":1446727949,     

 "callback":"eyJjYWxsYmFja1VybCI6Imh0dHA6Ly9vc3MtZGVtby5hbGl5dW5jcy5jb206MjM0NTAiLAoiY2FsbGJhY2tCb2R5" // base 64

 "dir":"user-dirs/"

}
```



base64 解析如下

| callbackUrl      | OSS向服务器发送的URL请求。                                   |
| ---------------- | ------------------------------------------------------------ |
| callbackHost     | OSS发送该请求时，请求头部所带的Host头。                      |
| callbackBody     | OSS发送给应用服务器的内容。如果是文件，可以是文件的名称、大小、类型等。如果是图片，可以是图片的高度、宽度等。 |
| callbackBodyType | 请求发送的Content-Type。                                     |

```json
{

 "callbackUrl":"[http://oss-demo.aliyuncs.com:23450](http://oss-demo.aliyuncs.com:23450/)",

 "callbackBody":"filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}",

 "callbackBodyType":"application/x-www-form-urlencoded"

}
```



### 创建Bucket



```go
    client, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
    if err != nil {
        // HandleError(err)
    }
    
    err = client.CreateBucket("my-bucket")
    if err != nil {
        // HandleError(err)
    }
```



   注意事项：

- 一旦创建，则无法更改其名称。
- Bucket名称必须全局唯一。
- 只能包括小写字母、数字和短划线（-）。
- 必须以小写字母或者数字开头和结尾。
- 长度必须在3~63字符之间。

### 上传文件

```go
    client, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
    if err != nil {
        // HandleError(err)
    }
    
    bucket, err := client.Bucket("my-bucket")
    if err != nil {
        // HandleError(err)
    }
    
    err = bucket.PutObjectFromFile("my-object", "LocalFile")
    if err != nil {
        // HandleError(err)
    }
```

### 下载文件:

```go
    client, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
    if err != nil {
        // HandleError(err)
    }
    
    bucket, err := client.Bucket("my-bucket")
    if err != nil {
        // HandleError(err)
    }
    
    err = bucket.GetObjectToFile("my-object", "LocalFile")
    if err != nil {
        // HandleError(err)
    }
```



###   数据存储表

[![image.png](https://i.postimg.cc/Y03RwPDb/image.png)](https://postimg.cc/WqdgG5vk)

```sql
CREATE TABLE IF NOT EXISTS `mydb`.`storage` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(256) NOT NULL COMMENT '文件名称',
`path` VARCHAR(256) NOT NULL COMMENT 'OOS 路径',
`ext` VARCHAR(32) NOT NULL DEFAULT 0 COMMENT '拓展名称',
`md5` VARCHAR(32) NOT NULL COMMENT '资源的MD5 ',
`creator_id` INT(16) NOT NULL DEFAULT 0 COMMENT '创建人id',
`created_at` INT(10) NOT NULL DEFAULT 0,
PRIMARY KEY (`id`),
INDEX `md5` (`md5` ASC) VISIBLE)
ENGINE = InnoDB
```



### 问题记录

- 使用服务端签名直传上传回调的方式，同一资源上传，是否可以去重？
  OSS 没有去重的功能，只有重名上传禁止覆盖的功能
- 应用服务器回调失败后，是否重新发送，间隔时间是什么？
  出于安全考虑，OSS的回调请求只会等待5秒。如果5秒后还没有返回，那么OSS就会主动断开与应用服务器的连接，并返回给客户端超时错误
