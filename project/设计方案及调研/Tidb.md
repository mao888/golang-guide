什么是Tidb：

Tidb是HTAP（Hybrid Transactional and Analytical Processing）类型的数据库，HTAP是OLTP和OLAP的交集。



Tidb再数据中台的优势：

海量存储允许多数据源聚合

支持标准SQL，多表关联快速结果

透明多业务模块，支持分表聚合后按照聚合后的多维度查询

Tidb最大下推机制，以及并行HASH Join等算子，决定的Tidb再表关联上的优势



Tidb主要组件有哪些：

TiKV：

PD：

TiFlash：



使用官网的测试数据如下：

tidb：

[![image.png](https://i.postimg.cc/CKgbGYcQ/image.png)](https://postimg.cc/75KCDv83)

mysql：

[![image.png](https://i.postimg.cc/2jsFxP8h/image.png)](https://postimg.cc/BtCKS7Fn)

同时执行的SQL：

```sql
SELECT
    l_orderkey,
    SUM(
        l_extendedprice * (1 - l_discount)
    ) AS revenue,
    o_orderdate,
    o_shippriority
FROM
    customer,
    orders,
    lineitem
WHERE
    c_mktsegment = 'BUILDING'
AND c_custkey = o_custkey
AND l_orderkey = o_orderkey
AND o_orderdate < DATE '1996-01-01'
AND l_shipdate > DATE '1996-02-01'
GROUP BY
    l_orderkey,
    o_orderdate,
    o_shippriority
ORDER BY
    revenue DESC,
    o_orderdate
limit 10;
```

tidb耗时（OLTP）：

[![image.png](https://i.postimg.cc/02t5vqMJ/image.png)](https://postimg.cc/dLycRp9q)

mysql耗时：

[![image.png](https://i.postimg.cc/CLdTjn5z/image.png)](https://postimg.cc/BXWV36x0)



tidb耗时（OLAP）：

[![image.png](https://i.postimg.cc/1RjY1CR3/image.png)](https://postimg.cc/bd1RQR27)



对比数据量一共800W，3张表关联聚合查询：

| 类型         | 平均查询时间（秒） |
| :----------- | :----------------- |
| Tidb（OLTP） | 0.80               |
| Tidb（OLAP） | 0.08               |
| MySQL        | 4.68               |
