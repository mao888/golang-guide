对应测试表结构:

user表

```
CREATE` `TABLE` ````user``` (`` ```id` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'用户ID'``,`` `````name``` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'姓名'``,`` ```age` ``int``(4) ``NOT` `NULL` `DEFAULT` `'0'` `COMMENT ``'年龄'``,`` ```gender` tinyint(1) ``NOT` `NULL` `DEFAULT` `'0'` `COMMENT ``'性别：0未知，1男；2女'``,`` ```status` tinyint(1) ``NOT` `NULL` `DEFAULT` `'0'` `COMMENT ``'状态：0未知，1正常；2冻结'``,`` ```test1` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test1'``,`` ```test2` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test2'``,`` ```test3` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test3'``,`` ```ct` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'创建时间'``,`` ```ut` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'更新时间'``,`` ``PRIMARY` `KEY` `(`id`),`` ``KEY` ``idx_ct` (`ct`)``) ENGINE=InnoDB ``DEFAULT` `CHARSET=utf8mb4;
```

order表

```
CREATE` `TABLE` ````order``` (`` ```id` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'订单ID'``,`` ```uid` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'用户ID'``,`` ```type` tinyint(2) ``NOT` `NULL` `COMMENT ``'订单类型'``,`` ```price` ``bigint``(10) ``NOT` `NULL` `COMMENT ``'金额'``,`` ```test1` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test1'``,`` ```test2` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test2'``,`` ```test3` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test3'``,`` ```test4` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test4'``,`` ```test5` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test5'``,`` ```test6` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test6'``,`` ```test7` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test7'``,`` ```test8` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test8'``,`` ```test9` ``varchar``(64) ``NOT` `NULL` `DEFAULT` `''` `COMMENT ``'test9'``,`` ```ct` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'创建时间'``,`` ```ut` ``bigint``(20) ``NOT` `NULL` `COMMENT ``'更新时间'``,`` ``PRIMARY` `KEY` `(`id`),`` ``KEY` ``idx_uid_ct` (`uid`,`ct`)``) ENGINE=InnoDB ``DEFAULT` `CHARSET=utf8mb4;
```



### 测试结果

#### user 插入测试:

mysql:

```
//1024 * 1024 百万级数据,一次1000条批量插入 83.766415328s
//1024 * 1024 百万级数据,一次2000条批量插入 78.305083006s
//1024 * 1024 百万级数据,一次3000条批量插入 出现大量慢日志
//1024 * 1024 * 20 千万级数据,一次2500条批量插入 1641.634111424s
```



postgres:

```
//1024 * 1024 百万级数据,一次1000条批量插入 46.985259923s
//1024 * 1024 百万级数据,一次2000条批量插入 36.57633506s
//1024 * 1024 百万级数据,一次3000条批量插入 34.889204107s
//1024 * 1024 百万级数据,一次5000条批量插入 出现大量慢日志
//1024 * 1024 * 20 千万级数据,一次3000条批量插入 729.143911501s
```



#### order 插入测试:

mysql:

```
//1024 * 1024 * 20 千万级数据,一次2000条批量插入 2598s
```

postgres:

```
//1024 * 1024 * 20 千万级数据,一次3000条批量插入 1701s

```

#### 关联查询测试:

测试语句： explain select * from test.user join test.order on test.order.uid=[test.user.id](http://test.user.id/) where test.user.id=1999;



**mysql: explain出来的结果，两张表都走了索引**

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-28-57.png?version=1&modificationDate=1640946537000&api=v2)

```

mysql: 查询结果
```

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-29-40.png?version=1&modificationDate=1640946580000&api=v2)

不走 idx_uid_ct 索引的结果：

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-49-4.png?version=1&modificationDate=1640947744000&api=v2)





**postgres: explain 出来的结果，也都走了索引**

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-40-10.png?version=1&modificationDate=1640947210000&api=v2)



postgres：查询结果

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-39-19.png?version=1&modificationDate=1640947159000&api=v2)



不走 idx_uid_ct 索引的结果：

![img](https://confluence.ftsview.com/download/attachments/45210004/image2021-12-31_18-52-15.png?version=1&modificationDate=1640947935000&api=v2)

#### 存储空间对比:

MySQL：5.1G

PostGreSQL：5.5G
