## 基础
### MySQL字符集
字符集规定了字符在数据库中的存储格式，比如占多少空间，支持哪些字符等等。不同的字符集有不同的编码规则，在有些情况下，甚至还有校对规则的存在。在运维和使用MySQL数据库中，选取合适的字符集非常重要，如果选择不恰当，轻则影响数据库性能，严重的可能导致数据存储乱码。
常见的MySQl字符集主要有以下四种：

| **字符集** | **长度** | **说明** |
| --- | --- | --- |
| GBK | 2 | 支持中文，但是不是国际通用字符集 |
| UTF-8 | 3 | 支持中英文混合场景，是国际通用字符集 |
| latin1 | 1 | MySQL默认字符集 |
| utf8mb4 | 4 | 完全兼容UTF-8，用四个字节存储更多的字符 |

MySQL数据库在开发运维中，字符集选用规则如下：
1、如果系统开发面向国外业务，需要处理不同国家、不同语言，则应该选择utf-8或者utf8mb4。
2、如果只需要支持中文，没有国外业务，则为了性能考虑，可以采用GBK。

### 设置数据库的字符集和设置表字段字符集的区别是什么？
mysql提供4种不同的颗粒度

1. server,全局级
2. 数据库级
3. 表级
4. 列级

4个级别作用域依次递减，优先级依次递增。就是设置的列字符集>表>库>server
如果修改库的字符集不会对原有的数据字符集进行改变，只会影响新增字符集。
如果修改表和列字符集，会对历史数据进行修改
### 数据库三范式

- **第一范式**
强调列的原子性, 数据库表的每一列都是不可分割的原子数据项
- **第二范式**
属性完全依赖于主键. 不能存在仅依赖主关键字一部分的属性.
- **第三范式**
确保每列都和主键列直接相关, 属性不依赖于其他非主属性.

### InnoDB与MyISAM
[存储引擎](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951780238-05dae39b-14b3-498f-aaea-847402b03025.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)

|  | **InnoDB** | **MyISAM** |
| --- | --- | --- |
| 事务 | 支持 | 不支持 |
| 外键 | 支持 | 不支持 |
| 行锁 | 支持 | 不支持 |
| 行表锁 | 行锁，操作时只锁某一行，不对其它行有影响，
适合高并发的操作 | 表锁，即使操作一条记录也会锁住
整个表，不适合高并发的操作 |
| 缓存 | 不仅缓存索引还要缓存真实数据，对内存要求较高，而且内存大小对性能有决定性的影响 | 只缓存索引，不缓存真实数据 |
| crash-safe能力 | 支持 | 不支持 |
| MVCC | 支持 | 不支持 |
| 索引存储类型 | 聚簇索引 | 非聚簇索引 |
| 是否保存表行数 | 不保存 | 保存 |
| 关注点 | 事务：并发写、事务、更大资源 | 性能：节省资源、消耗少、简单业
务 |

- InnoDB支持事物，而MyISAM不支持事物
- InnoDB支持行级锁，而MyISAM支持表级锁
- InnoDB支持MVCC, 而MyISAM不支持
- InnoDB支持外键，而MyISAM不支持
- InnoDB不支持全文索引，而MyISAM支持。

MVCC，全称Multi-Version Concurrency Control，即多版本并发控制。MVCC是一种并发控制的方法，一般在数据库管理系统中，实现对数据库的并发访问，在编程语言中实现事务内存。

### MySQL执行查询过程
![image-20220305154020996.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468335457-8f286a22-e1a3-4dbf-9a2a-e9f4d62371d3.png#averageHue=%23ebeee1&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=srts0&margin=%5Bobject%20Object%5D&name=image-20220305154020996.png&originHeight=507&originWidth=648&originalType=binary&ratio=1&rotation=0&showTitle=false&size=199722&status=done&style=none&taskId=u3abf2dfb-ddf9-48a2-9c0f-577a1ff1957&title=)

1. 客户端通过TCP连接发生连接请求到 MySQL 连接器, 连接器会对该请求进行权限验证以及连接资源分配
2. **查询缓存**(8.0之后没了, 原因是一般失效会非常频繁)。
   1. 当判断缓存是否命中时，MySQL不会进行解析查询语句，而是直接使用SQL语句和客户端发送过来的其他原始信息。所以，任何字符上的不同，例如空格、注解等都会导致缓存的不命中。）
3. 分析器(词法分析 ->** 语法分析**)
   1. （SQL语法是否写错了）。 如何把语句给到预处理器，检查数据表和数据列是否存在，解析别名看是否存在歧义。
4. **优化**器(决定索引的最佳使用方案)。是否使用索引，生成执行计划。
5. 执行器(检查权限 -> 执行语句 -> 返回结果集)
   1. 交给执行器，将数据保存到结果集中，同时会逐步将数据缓存到查询缓存中，最终将结果集返回给客户端。

### MySQL建表的约束条件有哪些？

- 主键约束（Primay Key Coustraint） 唯一性，非空性
- 唯一约束 （Unique Counstraint）唯一性，可以空，但只能有一个
- 检查约束 (Check Counstraint) 对该列数据的范围、格式的限制
- 默认约束 (Default Counstraint) 该数据的默认值
- 外键约束 (Foreign Key Counstraint) 需要建立两表间的关系并引用主表的列
- [x] 

---

## 索引
[索引的数据结构](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951782284-4a1be251-d755-4b55-863c-68f2b24946f0.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
### 索引模型

#### 哈希模型
哈希表以 **键-值对(key - value)** 存储数据, key经过哈希函数的换算, 确定其在数组中存储的位置, 但是哈希存在冲突, 可以用采用拉链法来解决哈希冲突.
![image-20220304135352101.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468348523-63b383f3-7aa6-4aa4-a540-63aa4ee0af76.png#averageHue=%23dbe3cf&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&height=446&id=tqEvJ&margin=%5Bobject%20Object%5D&name=image-20220304135352101.png&originHeight=446&originWidth=688&originalType=binary&ratio=1&rotation=0&showTitle=false&size=162646&status=done&style=none&taskId=u47c89065-8276-41af-8167-3472ba036c1&title=&width=688)
但是哈希后的数据不是有序的, 如果用于区间查询, 那么就必须一个个哈希查找了, 性能非常低, 所以哈希表这种存储结构只适用于**等值查询**的场景.

#### 有序数组模型
![image-20220304135919985.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468368777-adb059d2-9d60-4de6-b36d-3282405630ef.png#averageHue=%23d6e1ca&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=TeNHh&margin=%5Bobject%20Object%5D&name=image-20220304135919985.png&originHeight=225&originWidth=679&originalType=binary&ratio=1&rotation=0&showTitle=false&size=108377&status=done&style=none&taskId=ub4e6dfcf-89c1-4bc8-a3b0-4d2481a9ef4&title=)
有序数组无论是在**等值查询和范围查询**的场景都非常优秀, 因为有序可以使用二分法来查找, 时间复杂度是 ![](https://g.yuque.com/gr/latex?O(logN)#card=math&code=O%28logN%29&id=t3Usk). 范围查找 ![](https://g.yuque.com/gr/latex?k#card=math&code=k&id=d0DyN) 条数据, 只需要先二分查找首条数据, 之后向右遍历, 时间复杂度也就是 ![](https://g.yuque.com/gr/latex?O(klogN)#card=math&code=O%28klogN%29&id=En6iS)  .
但是有序数组为了保持有序, 若在中间插入数据时, 必须移动后面所有数据, 成本开销大.
所以, 有序数组索引只适用于**静态存储**引擎, 保存一些存储后就不会再去修改的数据.

#### 搜索树模型  
[6.MySQL数据结构选择的合理性](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951782284-4a1be251-d755-4b55-863c-68f2b24946f0.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
##### BST和AVL等二叉树模型
![image-20220304140528743.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468377486-65591b12-83e9-4431-82f9-246a81dbcb5b.png#averageHue=%23e5eadc&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=B1hjx&margin=%5Bobject%20Object%5D&name=image-20220304140528743.png&originHeight=506&originWidth=684&originalType=binary&ratio=1&rotation=0&showTitle=false&size=195417&status=done&style=none&taskId=u9cde14c9-87db-45fe-8243-b26b6ea1b49&title=)
BST不管是查询还是更新, 都只需要 ![](https://g.yuque.com/gr/latex?O(logN)#card=math&code=O%28logN%29&id=mZrL0) 的时间复杂度. 但是BST在某种情况下, 会使得其退化成链表. 如果想让他保持平衡, 那么就可以使用AVL. 对于二叉树来说, 如果数据量十分大, 那么这个层数就会越堆越高, 而数据是存放在磁盘中, 那么意味着要访问非常多的数据块, 就非常影响性能.

##### B树模型
![image-20220304142450189.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468389212-b9ca5523-2c1a-4e06-a11c-364b084240a8.png#averageHue=%23fafafa&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=xGYdZ&margin=%5Bobject%20Object%5D&name=image-20220304142450189.png&originHeight=222&originWidth=357&originalType=binary&ratio=1&rotation=0&showTitle=false&size=49245&status=done&style=none&taskId=u9dbce82b-d272-4c5f-a504-daa2e8700ac&title=)

B树是多路搜索树, B树每个结点都存储着数据, 解决了二叉树随数据变大而层数变高导致对磁盘IO时的性能低下问题. 
但是很明显, B树还不是最理想的存储结构, 试想一下如果进行范围查询, 对于范围中的数据来说, 那么不是每次都要从根节点开始往下找么, 必然有性能的问题.

##### B+树

![image-20220304143146009.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468399152-97e94467-955e-4b83-8a87-a0ee42a69d18.png#averageHue=%23fbfafa&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=EWNiZ&margin=%5Bobject%20Object%5D&name=image-20220304143146009.png&originHeight=234&originWidth=410&originalType=binary&ratio=1&rotation=0&showTitle=false&size=51302&status=done&style=none&taskId=u90b8f42e-6fd0-4000-abda-66c1d32e65a&title=)
于是在基于B树的模型上, 出现了B+树, 
B+树只有叶子节点是存储数据的, 而其他非叶子节点均为索引, 叶子节点用链表串起来, 且保证了有序. 在范围查询就只需找到其中一个数据, 之后向后遍历即可.

### [主键索引和非主键索引](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951782284-4a1be251-d755-4b55-863c-68f2b24946f0.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
|  | **TYPE** | **INDEX** |
| --- | --- | --- |
| id | int | id(primary key) |
| k | int | k |
| name | varchar |  |

假设有如上表结构, 那么建立起的索引结构如下图
![image-20220304144720568.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468413069-f47294ff-f9e9-494d-ae20-f7e6fb44de8c.png#averageHue=%23e5ebdc&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=rhgRE&margin=%5Bobject%20Object%5D&name=image-20220304144720568.png&originHeight=321&originWidth=669&originalType=binary&ratio=1&rotation=0&showTitle=false&size=121499&status=done&style=none&taskId=u87ddf3d4-a455-48c3-9d9f-42b6d34bb8d&title=)
从图中看出, 根据叶子节点内容的不同, 索引类型分为主键索引和非主键索引.

- **主键索引**的叶子节点存储的是整行数据. 在InnoDB里, 主键索引也被称为**聚簇索引**
- **非主键索引**的叶子节点存储的是主键的值. 在InnoDB里, 非主键索引也被称为**二级索引 / 非聚簇索引**

### [回表](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951782284-4a1be251-d755-4b55-863c-68f2b24946f0.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
当执行 `SELECT * FROM t WHERE id = 500` 时, 即主键查询方式, 则需要搜索ID这颗B+树; 当执行 `SELECT * FROM t WHERE k = 5` 时, 即普通索引查询方式, 则先在k这棵树查找到主键的值, 再从ID这棵树中查找到对应的行.
当我们执行SQL搜索数据时, 如果需要先从非主键索引中查询到主键的值, 再从主键索引中查询到对应的数据, 这个过程就被称为**回表**. 所以应该尽量使用主键查询.

### 索引维护 (页分裂与页合并)
B+树为了有序性, 需要对插入和删除数据时做出对应的维护. 当插入数据时, 如在上图中插入ID=400的数据, 那么从逻辑上来说, 需要移动后面的数据, 空出位置.

若此时R5所在数据页满了, 则需要申请一个新的数据页, 然后移动部分数据到新数据页中, 这个过程被称为**页分裂**. 页分裂影响了数据页的空间利用率, 而且在分裂过程中, 性能也会有所影响.

若相邻两个数据页因为删除导致利用率很低后, 那么会将这两个数据页的数据合并到一个数据页中, 这个过程被称为**页合并**. 即页分裂的逆过程.

### 覆盖索引
如果执行了语句 `SELECT id FROM t WHERE k between 3 and 5` 时, 只需要查询 id 的值, 而 id 已经在 k 的索引树上, 所以不需要再回表去查询整行, 直接返回查询结果, 索引 k 已经覆盖了这条SQL查询的需求, 被称为 **覆盖索引**. 覆盖索引能够减少树的搜索次数, 不需要再次回表查询整行, 所以是一个常用的性能优化手段.

### 最左前缀原则
**最左前缀原则** 就是利用索引列中最左的字段优先进行匹配

|  | **TYPE** | **INDEX** |
| --- | --- | --- |
| id | int | id(primary key) |
| id_card | varchar | id_card |
| name | varchar | (name, age) |
| age | int |  |
| ismale | tinyint |  |

若有如上表结构, 对于INDEX(name, age)来说, 索引树结构如下, 可以看到, 索引项是按照索引定义里面出现的顺序排序的.
![image-20220304154001675.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468426242-23685a2f-6a69-4c6f-9fe0-b78137a75cc4.png#averageHue=%23e5eadc&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=QrlWf&margin=%5Bobject%20Object%5D&name=image-20220304154001675.png&originHeight=381&originWidth=683&originalType=binary&ratio=1&rotation=0&showTitle=false&size=124428&status=done&style=none&taskId=uf60e76ab-e8e3-4aef-af80-3dec18840e8&title=)
对于SQL语句 `SELECT * FROM t WHERE name LIKE '张%'` 来说, 也是能够用到INDEX(name, age)这个索引的, 只需检索到第一个姓为张的人, 之后向后遍历即可, 所以可以利用最左前缀来加速检索. **最左前缀**可以是**联合索引的最左N个字段**, 也可以是**字符串索引的最左M个字符**.

其效果和单独创建一个INDEX(name)的效果是一样的, 如果你想使用INDEX(name, age)也想让name也拥有索引INDEX(name), 那么只需保留前者即可, **若通过调整索引字段的顺序, 可以少维护一个索引树, 那么这个顺序就是需要优先考虑采用的**. 但如果也有SQL语句条件类似 `WHERE age = 1` , 那么最好再维护一个INDEX(age)的索引.

### 前缀索引
在对字符串创建索引, 如INDEX(name)中, 若字符串非常大, 那么响应的空间使用和维护开销也非常大, 就可以**使用字符串从左开始的部分字符创建索引**, 减少空间和维护的成本, 但是也会降低索引的选择性. **索引的选择性**指的是 : 不重复的索引值和数据表的记录总数(#T)的比值, 范围为 1/#T 到 1 之间, 索引选择性越高则查询效率越高. 对于BLOB, TEXT, VARCHAR等类型的列, 必须使用前缀索引, MySQL不允许索引这些列的完整长度.

1. 先计算完整列的选择性 `SELECT COUNT(DISTINCT name)/COUNT(1) FROM t`
2. 在计算不同前缀长度N的选择性 `SELECT COUNT(DISCTINCT LEFT(name, N)) / COUNT(1) FROM t`
3. 看哪个N更靠近1, 进行索引的创建

### 索引下推
对于SQL语句 `SELECT * FROM t WHERE name LIKE '陈%' AND age = 10` , INDEX(name, age) 情况来说
在 MySQL5.6 之前没有引入索引下推优化时, 执行流程如下图, 在定位完name字段的索引后, 需要一条条进行回表查询, 然后再判断其他字段是否满足条件.
![image-20220304160409381.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468447625-ea605667-5913-4f4f-9a12-4dfc3cbb5030.png#averageHue=%23f0f0e9&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=dEDJf&margin=%5Bobject%20Object%5D&name=image-20220304160409381.png&originHeight=258&originWidth=666&originalType=binary&ratio=1&rotation=0&showTitle=false&size=119734&status=done&style=none&taskId=u001675a9-6ed7-480d-be40-64f2149c5f8&title=)
而 MySQL5.6 引入了索引下推优化后, 可以在所有遍历过程中, **对索引中包含的字段先进行判断过滤**, 然后再进行后续操作, 减少了回表次数.
![image-20220304160654992.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468451640-b30d74f6-06bb-4385-bb02-23db2d0368ce.png#averageHue=%23efefe8&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=RcWRF&margin=%5Bobject%20Object%5D&name=image-20220304160654992.png&originHeight=246&originWidth=656&originalType=binary&ratio=1&rotation=0&showTitle=false&size=123478&status=done&style=none&taskId=u10ab9e08-fc43-4d31-9eb1-283cc212fcf&title=)

### 自适应哈希索引
InnoDB中不存在哈希索引, 但是哈希索引确实有利于快速查找, 于是InnoDB引入了"**自适应哈希索引**", 在某些索引值被使用的非常频繁时, InnoDB会在内存中的B+树结构之上创建一个哈希索引, 用于这些频繁使用的索引值的快速查找, 使得其存有哈希快速查找的特点.

### 索引相关高频面试题
> 1. 索引是什么? 索引优缺点? 
>    - 索引类似于目录, 进行数据的快速定位
>    - 优点: 加快数据检索速度
>    - 缺点: 创建索引和维护索引需要消耗空间和时间
> 
 
> 2. MySQL索引类型 
>    - 按存储结构划分 : B+Tree索引, Hash索引, FULLINDEX全文索引, R-TREE索引
>    - 按应用层次划分: 普通索引, 唯一索引, 联合索引, 聚簇索引, 非聚簇索引
> 
 
> 3. 索引底层实现? 为什么使用B+树, 而不是B树, BST, AVL, 红黑树等等?
> 4. 什么是聚簇索引和非聚簇索引?
> 5. 非聚簇索引一定会回表吗?
(不一定, 覆盖索引不会回表)
> 6. 什么是联合索引?为什么需要注意联合索引中的字段顺序?
> 7. 什么是最左前缀原则?
> 8. 什么是前缀索引?
> 9. 什么是索引下推?
> 10. 如何查看MySQL语句是否使用到索引?
EXPLAIN SQL语句
possible_key: 可能用到的索引(可以查看是否有冗余索引)
key: 真正使用到的索引
> 11. 为什么建议使用自增主键作为索引?
(索引维护可能造成页分裂, 自增主键减少数据的移动和分裂)
> 12. 建立索引的原则 
>    - 建立索引的字段最好为NOT NULL
>    - 索引字段占用空间越小越好
>    - 最左匹配原则
>    - =和in建立索引时顺序可以任意, 比如a = 1 and b = 2 and c = 3 建立(a, b, c)和(b, a, c)索引效果是一样的, MySQL查询优化器会进行优化
>    - 建立的索引让索引的选择性尽可能接近1, 唯一索引的索引选择性为1
>    - 尽量扩展索引, 不要让索引冗余, 如有SQL需要对单个a进行索引, 那么上述条件建立的索引应该为(a, b, c)或(a, c, b)
>    - 索引列不能参与计算
> 
 
> 13. 什么情况下索引失效? 
>    - 使用 != 或  <>
>    - 类型不一致导致索引失效
>    - 函数导致的索引失效, 函数用在索引列时, 不走索引
如 `SELECT * FROM t WHERE DATE(create_time) = 'yyyy-MM-dd'`
>    - 运算符导致的索引失效
如 `SELECT * FROM t WHERE k - 1 = 2`, 若有INDEX(k), 则不走索引
>    - OR引起的索引失效
如 `SELECT * FROM t WHERE k = 1 OR j = 2`, 若有INDEX(k), 则不走索引, 如果OR连接的时同一个字段, 则不会失效
>    - 模糊查询导致的索引失效
如 `SELECT * FROM t WHERE name = '%三'`, %放字符串字段前匹配不走索引
>    - NOT IN, NOT EXISTS导致索引失效


---

## [事务](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951782526-a2a8a3d3-fc3e-4526-b57c-a0fa9f78b87f.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
对于一个事务, 要么事务内的SQL全部执行, 要么都不执行
```sql
START TRANSACTION;
SELECT balance FROM checking WHERE customer_id = 10233276;
UPDATE checking SET balance = balance - 200.00 WHERE customer_id = 10233276;
UPDATE savings SET balance = balance + 200.00 WHERE customer_id = 10233276;
COMMIT;
```
### 什么是数据库事务？
**事务**：一组逻辑操作单元，使数据从一种状态变换到另一种状态。
**事务处理的原则**：保证所有事务都作为 一个工作单元 来执行，即使出现了故障，都不能改变这种执行方
式。当在一个事务中执行多个操作时，要么所有的事务都被提交( commit )，那么这些修改就 永久 地保
存下来；要么数据库管理系统将 放弃 所作的所有 修改 ，整个事务回滚( rollback )到最初状态。

事务最经典也经常被拿出来说例子就是转账了。
假如小明要给小红转账1000元，这个转账会涉及到两个关键操作就是：将小明的余额减少1000元，将小红的余额增加1000元。万一在这两个操作之间突然出现错误比如银行系统崩溃，导致小明余额减少而小红的余额没有增加，这样就不对了。事务就是保证这两个关键操作要么都成功，要么都要失败。
### 事务的特性 ACID

- **原子性**：原子性是指事务是一个不可分割的工作单位，要么全部提交，要么全部失败回滚。
- **一致性**：根据定义，一致性是指事务执行前后，数据从一个 合法性状态 变换到另外一个 合法性状态 。这种状态是 语义上 的而不是语法上的，跟具体的业务有关。

那什么是合法的数据状态呢？满足 预定的约束 的状态就叫做合法的状态。通俗一点，这状态是由你自己来定义的（比如满足现实世界中的约束）。满足这个状态，数据就是一致的，不满足这个状态，数据就是不一致的！如果事务中的某个操作失败了，系统就会自动撤销当前正在执行的事务，返回到事务操作之前的状态。

- **隔离性**：事务的隔离性是指一个事务的执行 不能被其他事务干扰 ，即一个事务内部的操作及使用的数据对 并发 的其他事务是隔离的，并发执行的各个事务之间不能互相干扰。
- **持久性：**持久性是指一个事务一旦被提交，它对数据库中数据的改变就是 永久性的 ，接下来的其他操作和数据库故障不应该对其有任何影响。

持久性是通过 事务日志 来保证的。日志包括了 重做日志 和 回滚日志 。当我们通过事务对数据进行修改的时候，首先会将数据库的变化信息记录到重做日志中，然后再对数据库中对应的行进行修改。这样做的好处是，即使数据库系统崩溃，数据库重启后也能找到没有更新到数据库系统中的重做日志，重新执行，从而使事务具有持久性。
### [事务ACID怎么实现的](https://www.cnblogs.com/rjzheng/p/10841031.html)
mysql具有**redo log** 和**undo log**，这二种文件都是事务相关的问题。
每次开启一个事务，则mysql的innodb引擎就会生成一张**undo log**文件，该文件主要记录这个事务ID所产生的一些更新、删除、插入操作。

当事务1执行update的时候，就会将udpate记录到undo log文件，当事务进行commit的时候，就会将undo log文件删除，如果回滚时，则会根据undo log文件的内容进行执行插入回滚SQL脚本。

**redo log** 文件是数据库的一个共享的文件，也是一份写缓存的文件，试想一下，每次操作读写都需要去访问磁盘的随机IO，其实会很耗时，因此可以将一些频繁的页内容加载到内存的一个缓存buffer中，当进行读操作时去查看缓存buffer是否有对应的数据，如果没有，则去磁盘查询，查询后再将查询的结果写入到缓存buffer。当执行写操作的时候，就先去更新缓存buffer,等到一定时间，将缓存buffer的数据再一次写入到磁盘中。但是这样就会有一个数据一致性的问题了，假设buffer的数据没有flush到磁盘，mysql服务器就宕机了，那内存的buffer的数据也会清空，redo log就解决了数据一致性的问题。

写操作的时候，不会写入到缓存buffer中，而是写入到redo log中，当事务提交后，redo log的内容就flush到磁盘中，redo log是一个文件，当服务器宕机了，也不影响redo log已保存后的数据，当mysql进行宕机后，redo log还是有内容的，如果想事务提交，则就执行redo log的数据到磁盘，一般而言是会将redo log的数据进行回滚，也就是删除数据。

#### 原子性
利用Innodb的undo log。
undo log名为回滚日志，是实现原子性的关键，当事务回滚时能够撤销所有已经成功执行的sql语句，他需要记录你要回滚的相应日志信息。
例如

- (1)当你delete一条数据的时候，就需要记录这条数据的信息，回滚的时候，insert这条旧数据
- (2)当你update一条数据的时候，就需要记录之前的旧值，回滚的时候，根据旧值执行update操作
- (3)当年insert一条数据的时候，就需要这条记录的主键，回滚的时候，根据主键执行delete操作

undo log记录了这些回滚需要的信息，当事务执行失败或调用了rollback，导致事务需要回滚，便可以利用undo log中的信息将数据回滚到修改之前的样子。
#### 持久性
redo log 保证了持久性，事务提交了，redo log的内容就会flush到磁盘中。
是利用Innodb的redo log。
正如之前说的，Mysql是先把磁盘上的数据加载到内存中，在内存中对数据进行修改，再刷回磁盘上。如果此时突然宕机，内存中的数据就会丢失。
_怎么解决这个问题？_
简单啊，事务提交前直接把数据写入磁盘就行啊。
_这么做有什么问题？_

- 只修改一个页面里的一个字节，就要将整个页面刷入磁盘，太浪费资源了。毕竟一个页面16kb大小，你只改其中一点点东西，就要将16kb的内容刷入磁盘，听着也不合理。
- 毕竟一个事务里的SQL可能牵涉到多个数据页的修改，而这些数据页可能不是相邻的，也就是属于随机IO。显然操作随机IO，速度会比较慢。

于是，决定采用redo log解决上面的问题。当做数据修改的时候，不仅在内存中操作，还会在redo log中记录这次操作。当事务提交的时候，会将redo log日志进行刷盘(redo log一部分在内存中，一部分在磁盘上)。当数据库宕机重启的时候，会将redo log中的内容恢复到数据库中，再根据undo log和binlog内容决定回滚数据还是提交数据。
_采用redo log的好处？_
其实好处就是将redo log进行刷盘比对数据页刷盘效率高，具体表现如下

- redo log体积小，毕竟只记录了哪一页修改了啥，因此体积小，刷盘快。
- redo log是一直往末尾进行追加，属于顺序IO。效率显然比随机IO来的快。
#### 隔离性
> mysql采用mvcc进行，通过版本链、read view以及隐藏的三个字段来实现，具体可以百度

利用的是锁和MVCC机制。还是拿转账例子来说明，有一个账户表如下
表名t_balance

| id | user_id | balance |
| --- | --- | --- |
| 1 | A | 200 |
| 2 | B | 0 |

其中id是主键，user_id为账户名，balance为余额。还是以转账两次为例，如下图所示
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663295801787-fb42a2ce-82bb-4653-b9a2-6ee5b8ee546f.png#averageHue=%23f3f2f2&clientId=ub06f29a5-b09f-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=323&id=uf3a4683b&margin=%5Bobject%20Object%5D&name=image.png&originHeight=645&originWidth=781&originalType=binary&ratio=1&rotation=0&showTitle=false&size=142118&status=done&style=none&taskId=u1f7cd164-d1f3-440c-bd6e-35452c2553c&title=&width=390.5)
至于MVCC,即多版本并发控制(Multi Version Concurrency Control),一个行记录数据有多个版本对快照数据，这些快照数据在undo log中。
如果一个事务读取的行正在做DELELE或者UPDATE操作，读取操作不会等行上的锁释放，而是读取该行的快照版本。
由于MVCC机制在可重复读(Repeateable Read)和读已提交(Read Commited)的MVCC表现形式不同，就不赘述了。
但是有一点说明一下，在事务隔离级别为读已提交(Read Commited)时，一个事务能够读到另一个事务已经提交的数据，是不满足隔离性的。但是当事务隔离级别为可重复读(Repeateable Read)中，是满足隔离性的。

#### 一致性
这个问题分为两个层面来说。
从**数据库层面**，数据库通过原子性、隔离性、持久性来保证一致性。也就是说ACID四大特性之中，C(一致性)是目的，A(原子性)、I(隔离性)、D(持久性)是手段，是为了保证一致性，数据库提供的手段。数据库必须要实现AID三大特性，才有可能实现一致性。例如，原子性无法保证，显然一致性也无法保证。
但是，如果你在事务里故意写出违反约束的代码，一致性还是无法保证的。例如，你在转账的例子中，你的代码里故意不给B账户加钱，那一致性还是无法保证。因此，还必须从应用层角度考虑。
从**应用层面**，通过代码判断数据库数据是否有效，然后决定回滚还是提交数据！
[
](https://blog.csdn.net/qq_38240227/article/details/123798481)
### 事务的状态
我们现在知道 事务 是一个抽象的概念，它其实对应着一个或多个数据库操作，MySQL根据这些操作所执
行的不同阶段把 事务 大致划分成几个状态：

- **活动的（active）**

事务对应的数据库操作正在执行过程中时，我们就说该事务处在 活动的 状态。

- **部分提交的（partially committed）**

当事务中的最后一个操作执行完成，但由于操作都在内存中执行，所造成的影响并 没有刷新到磁盘
时，我们就说该事务处在 部分提交的 状态。

- **失败的（failed）**

当事务处在 活动的 或者 部分提交的 状态时，可能遇到了某些错误（数据库自身的错误、操作系统
错误或者直接断电等）而无法继续执行，或者人为的停止当前事务的执行，我们就说该事务处在 失
败的 状态。

- **

如果事务执行了一部分而变为 失败的 状态，那么就需要把已经修改的事务中的操作还原到事务执
行前的状态。换句话说，就是要撤销失败事务对当前数据库造成的影响。我们把这个撤销的过程称
之为 回滚 。当 回滚 操作执行完毕时，也就是数据库恢复到了执行事务之前的状态，我们就说该事
务处在了 中止的 状态。

- **提交的（committed）**

当一个处在 部分提交的 状态的事务将修改过的数据都 同步到磁盘 上之后，我们就可以说该事务处
在了 提交的 状态。
一个基本的状态转换图如下所示：
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1655343228326-9c236f52-d4e0-4fe7-84c9-9ce989e13e2e.png#averageHue=%23fbfbfb&clientId=u93f3805a-794b-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=404&id=ue26b1dc2&margin=%5Bobject%20Object%5D&name=image.png&originHeight=404&originWidth=555&originalType=binary&ratio=1&rotation=0&showTitle=false&size=57775&status=done&style=none&taskId=u7b06426a-2d47-4624-8efa-c347eff5066&title=&width=555)
### 并发事务带来的问题

- **脏读 Dirty Read**
当一个事务正在访问数据并且对数据进行了修改，而这种修改还没有提交到数据库中，这时另外一个事务也访问了这个数据，然后使用了这个数据。因为这个数据是还没有提交的数据，那么另外一个事务读到的这个数据是“脏数据”，依据“脏数据”所做的操作可能是不正确的
- **丢失修改 Lost to modify**
指在一个事务读取一个数据时，另外一个事务也访问了该数据，那么在第一个事务中修改了这个数据后，第二个事务也修改了这个数据。这样第一个事务内的修改结果就被丢失，因此称为丢失修改
- **不可重复读 Unrepeatable Read**
指在一个事务内多次读同一数据。在这个事务还没有结束时，另一个事务也访问该数据。那么，在第一个事务中的两次读数据之间，由于第二个事务的修改导致第一个事务两次读取的数据可能不太一样。这就发生了在一个事务内两次读到的数据是不一样的情况，因此称为不可重复读. 侧重点为修改.
- **幻读 Phantom Read**
幻读与不可重复读类似。它发生在一个事务（T1）读取了几行数据，接着另一个并发事务（T2）插入了一些数据时。在随后的查询中，第一个事务（T1）就会发现多了一些原本不存在的记录，就好像发生了幻觉一样，所以称为幻读. 侧重点为新增或删除.

### 隔离性与隔离级别

- **串行化 Serializable**
   - 最高隔离级别. 强制事务串行执行. 避免幻读问题. SERIALIZABLE会在读取的每一行数据上都加锁, 所以可能导致大量的超时和锁竞争问题。
   - 可避免脏读、不可重复读、幻读。
- **可重复读 Repeatable Read** (MySQL默认隔离级别)
   - 保证了在同一个事务中多次读取同样的记录的结果是一致的. 可能会发生幻读. InnoDB通过MVCC多并发版本控制来解决幻读问题.
   - 可避免脏读、不可重复读的发生。
- **读提交 Read Committed**
   - 一个事务从开始直到提交之前, 所做的任何修改都是对其他事务不可见的. 可能会发生幻读, 不可重复读.
   - 可避免脏读的发生。
- **读未提交 Read Uncommitted**
   - 在这个级别下, 事务中的修改, 即使没有提交, 对其他事务也是可见的. 可能会导致脏读, 不可重复读或幻读.
   - 最低级别，任何情况都无法保证。

在 MySQL 数据库中，支持上面四种隔离级别，默认的为 Repeatable read (可重复读)；而在 Oracle 数据库中，只支持 Serializable (串行化)级别和 Read committed (读已提交)这两种级别，其中默认的为 Read committed 级别。
#### 隔离性
| 隔离级别 | 说明 |
| --- | --- |
| 读未提交 | 一个事务还没提交时，它做的变更就能被别的事务看到 |
| 读提交 | 一个事务提交之后，它做的变更才会被其他事务看到 |
| 可重复读 | 一个事务中，对同一份数据的读取结果总是相同的，无论是否有其他事务对这份数据进行操作，以及这个事务是否提交。**InnoDB默认级别**。 |
| 串行化 | 事务串行化执行，每次读都需要获得表级共享锁，读写相互都会阻塞，隔离级别最高，牺牲系统并发性。 |

不同的隔离级别是为了解决不同的问题。也就是脏读、幻读、不可重复读。

| 隔离级别 | **脏读** | **不可重复读** | **幻读** |
| --- | --- | --- | --- |
| 读未提交 | 可以出现 | 可以出现 | 可以出现 |
| 读提交 | 不允许出现 | 可以出现 | 可以出现 |
| 可重复读 | 不允许出现 | 不允许出现 | 可以出现 |
| 串行化 | 不允许出现 | 不允许出现 | 不允许出现 |

那么不同的隔离级别，隔离性是如何实现的，为什么不同事物间能够互不干扰？ 答案是 **锁 和 MVCC。**
### 什么是脏读？幻读？不可重复读？
1、脏读：事务 A 读取了事务 B 更新的数据，然后 B 回滚操作，那么 A 读取到的数据是脏数据
2、不可重复读：事务 A 多次读取同一数据，事务 B 在事务 A 多次读取的过程中，对数据作了更新并提交，导致事务 A 多次读取同一数据时，结果 不一致。
3、幻读：系统管理员 A 将数据库中所有学生的成绩从具体分数改为 ABCDE 等级，但是系统管理员 B 就在这个时候插入了一条具体分数的记录，当系统管理员 A 改结束后发现还有一条记录没有改过来，就好像发生了幻觉一样，这就叫幻读。
不可重复读侧重于修改，幻读侧重于新增或删除（多了或少量行），脏读是一个事务回滚影响另外一个事务。
### 事务的实现原理
事务是基于重做日志文件(redo log)和回滚日志(undo log)实现的。
每提交一个事务必须先将该事务的所有日志写入到重做日志文件进行持久化，数据库就可以通过重做日志来保证事务的原子性和持久性。
每当有修改事务时，还会产生 undo log，如果需要回滚，则根据 undo log 的反向语句进行逻辑操作，比如 insert 一条记录就 delete 一条记录。undo log 主要实现数据库的一致性。

### 事务相关高频面试题
> 1. 什么是事务
> 2. 事务的四个特征
> 3. MySQL四种隔离级别
> 4. 什么是脏读? 幻读? 不可重复读?
> 5. 事务是如何实现的(原理) ?
redo log 实现原子和持久性
undo log 实现一致性


---

## [事务日志](https://www.yuque.com/office/yuque/0/2022/pdf/22219483/1652951783650-2beb33b4-dd26-4778-92d7-f0380a52f552.pdf?from=file%3A%2F%2F%2FE%3A%2FProgram%2520Files%2F%25E8%25AF%25AD%25E9%259B%2580%2Fyuque-desktop%2Fresources%2Fapp.asar%2Fbuild%2Frenderer%2Findex.html%3Flocale%3Dzh-CN%26isYuque%3Dtrue%26theme%3D%26isWebview%3Dtrue%26editorType%3Deditor%26useLocalPath%3Dundefined%23%2Feditor)
事务有4种特性：原子性、一致性、隔离性和持久性。那么事务的四种特性到底是基于什么机制实现呢？

- 事务的隔离性由 锁机制 实现。
- 而事务的原子性、一致性和持久性由事务的 redo 日志和undo 日志来保证。
   - REDO LOG 称为 重做日志 ，提供再写入操作，恢复提交事务修改的页操作，用来保证事务的持久性。
   - UNDO LOG 称为 回滚日志 ，回滚行记录到某个特定版本，用来保证事务的原子性、一致性。

有的DBA或许会认为 UNDO 是 REDO 的逆过程，其实不然。
### redo log（重写日志）
redo log 是物理日志, 记录的是"在某个数据页做出了什么修改", 属于 `InnoDB存储引擎` 层面.

当有一条记录需要更新的时候, InnoDB引擎会先把记录写到 redo log 中, 并更新内存, 这时候更新就算完成. 同时, InnoDB引擎会在适当的时候, 将这个操作记录更新到磁盘里面, 往往是系统较为空闲时.

InnoDB的redo log是固定大小的, 如可以配置为一组4个文件, 每个文件1GB, 那么redo log总共就可以记录4GB的操作. 从头开始写, 写到末尾又回到开头循环写.
![image-20220304195247657.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468474598-8d818589-d1ec-4e9a-ae1b-bbb9540d96ea.png#averageHue=%23dfe6d4&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=GuR8t&margin=%5Bobject%20Object%5D&name=image-20220304195247657.png&originHeight=214&originWidth=547&originalType=binary&ratio=1&rotation=0&showTitle=false&size=62322&status=done&style=none&taskId=ue7ee8264-f69f-4f72-9b9d-b2c439271c0&title=)
**write pos** 是当前记录的位置, 边写边后移, 写到第三号文件末尾后就回到0号文件开头.
**checkpoint** 是当前要擦除的位置, 也是往后推移并循环的, 擦除记录前要把记录更新到数据文件.

**write pos** 和 **checkpoint** 之间是空闲的部分, 可以用来记录新的操作, 如果 **write pos** 追上 **checkpoint** ,表示 redo log 满了, 这时不能再执行新的更新, 得停下先擦掉一些记录, 把 **checkpoint** 推进.

有了 redo log, InnoDB 可以保证即使数据库发生异常重启, 之前提交的记录都不会丢失, 这个能力被称为 **crash-safe**

#### 1 为什么需要REDO日志
一方面，缓冲池可以帮助我们消除CPU和磁盘之间的鸿沟，checkpoint机制可以保证数据的最终落盘，然
而由于checkpoint 并不是每次变更的时候就触发 的，而是master线程隔一段时间去处理的。所以最坏的情
况就是事务提交后，刚写完缓冲池，数据库宕机了，那么这段数据就是丢失的，无法恢复。

另一方面，事务包含 持久性 的特性，就是说对于一个已经提交的事务，在事务提交后即使系统发生了崩
溃，这个事务对数据库中所做的更改也不能丢失。

那么如何保证这个持久性呢？ 一个简单的做法 ：在事务提交完成之前把该事务所修改的所有页面都刷新
到磁盘，但是这个简单粗暴的做法有些问题

另一个解决的思路 ：我们只是想让已经提交了的事务对数据库中数据所做的修改永久生效，即使后来系
统崩溃，在重启后也能把这种修改恢复出来。所以我们其实没有必要在每次事务提交时就把该事务在内
存中修改过的全部页面刷新到磁盘，只需要把 修改 了哪些东西 记录一下 就好。比如，某个事务将系统
表空间中 第10号 页面中偏移量为 100 处的那个字节的值 1 改成 2 。我们只需要记录一下：将第0号表
空间的10号页面的偏移量为100处的值更新为 2 。
#### 2 REDO日志的好处、特点
**好处**

- redo日志降低了刷盘频率
- redo日志占用的空间非常小

**特点**

- redo日志是顺序写入磁盘的
- 事务执行过程中，redo log不断记录
#### 3 redo的组成
Redo log可以简单分为以下两个部分：

- 重做日志的缓冲 (redo log buffer) ，保存在内存中，是易失的。

**参数设置：innodb_log_buffer_size：**
redo log buffer 大小，默认 16M ，最大值是4096M，最小值为1M。
```sql
mysql> show variables like '%innodb_log_buffer_size%';
+------------------------+----------+
| Variable_name          | Value
+------------------------+----------+
| innodb_log_buffer_size | 16777216 |
+------------------------+----------+
```

- 重做日志文件 (redo log file) ，保存在硬盘中，是持久的。
#### 4 redo的整体流程
以一个更新事务为例，redo log 流转过程，如下图所示：
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1655385880949-8b99c805-e7fe-48f5-90b6-45b5575da558.png#averageHue=%23f8f9f6&clientId=u3fe4f8a9-affe-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=237&id=ub6795cbc&margin=%5Bobject%20Object%5D&name=image.png&originHeight=237&originWidth=776&originalType=binary&ratio=1&rotation=0&showTitle=false&size=85093&status=done&style=none&taskId=u4461b7e8-dc60-498d-964a-ed6fa9a0e09&title=&width=776)
> 第1步：先将原始数据从磁盘中读入内存中来，修改数据的内存拷贝
> 第2步：生成一条重做日志并写入redo log buffer，记录的是数据被修改后的值
> 第3步：当事务commit时，将redo log buffer中的内容刷新到 redo log file，对 redo log file采用追加
> 写的方式
> 第4步：定期将内存中修改的数据刷新到磁盘中

> 体会：
> Write-Ahead Log(预先日志持久化)：在持久化一个数据页之前，先将内存中相应的日志页持久化。

#### 5 redo log的刷盘策略
redo log的写入并不是直接写入磁盘的，InnoDB引擎会在写redo log的时候先写redo log buffer，之后以 一
定的频率 刷入到真正的redo log file 中。这里的一定频率怎么看待呢？这就是我们要说的刷盘策略。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1655385973460-052d91a1-c4df-410c-8e4d-44ccbb16b761.png#averageHue=%23eff0d4&clientId=u3fe4f8a9-affe-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=427&id=u8fe0e39d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=427&originWidth=741&originalType=binary&ratio=1&rotation=0&showTitle=false&size=79870&status=done&style=none&taskId=ua711ee70-40ee-4bf8-badc-5774a6a2f6f&title=&width=741)
注意，redo log buffer刷盘到redo log file的过程并不是真正的刷到磁盘中去，只是刷入到 文件系统缓存
（page cache）中去（这是现代操作系统为了提高文件写入效率做的一个优化），真正的写入会交给系
统自己来决定（比如page cache足够大了）。那么对于InnoDB来说就存在一个问题，如果交给系统来同
步，同样如果系统宕机，那么数据也丢失了（虽然整个系统宕机的概率还是比较小的）。
针对这种情况，InnoDB给出 innodb_flush_log_at_trx_commit 参数，该参数控制 commit提交事务
时，如何将 redo log buffer 中的日志刷新到 redo log file 中。它支持三种策略：

- 设置为0 ：表示每次事务提交时不进行刷盘操作。（系统默认master thread每隔1s进行一次重做日志的同步）
- 设置为1 ：表示每次事务提交时都将进行同步，刷盘操作（ 默认值 ）
- 设置为2 ：表示每次事务提交时都只把 redo log buffer 内容写入 page cache，不进行同步。由os自己决定什么时候同步到磁盘文件。
- ![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1655386113630-5f0c3558-c6db-40c4-a5bc-d4da7ca9c946.png#averageHue=%23fcfbfa&clientId=u3fe4f8a9-affe-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=656&id=u316e4578&margin=%5Bobject%20Object%5D&name=image.png&originHeight=656&originWidth=871&originalType=binary&ratio=1&rotation=0&showTitle=false&size=83974&status=done&style=none&taskId=ucbb01204-7507-4164-8b8d-3793c1b7c44&title=&width=871)


### binlog（二进制日志）
binlog是逻辑日志, 记录的是SQL语句的原始逻辑, 属于 `MySQL Server` 层面.
binlog 主要用来保证数据的一致性, 在主从等环境下, 需要通过 binlog 来进行数据的同步.
![image-20220304200528837.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468485491-e2e4e686-f1b2-4bf0-ae12-05538a835fe8.png#averageHue=%23cafbc9&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=woc8w&margin=%5Bobject%20Object%5D&name=image-20220304200528837.png&originHeight=434&originWidth=983&originalType=binary&ratio=1&rotation=0&showTitle=false&size=85666&status=done&style=none&taskId=ue9b2d913-ea22-4e58-9596-baa453fdd55&title=)
binlog 日志有三种记录格式

- **statement**
这个记录的内容是SQL语句的原文, 同步数据时, 会执行记录的SQL语句, 但如果存在 `update_time = now()` 这种实时性强的SQL语句, 那么两次操作的时间不一样就会导致数据不一致问题.
- **row**
指定为 row 时, 记录的内容包含了操作的具体数据, 解决了 statement 格式的问题, 但是有数据的存在说明需要空间占用, 恢复与同步时会更消耗 IO 资源, 影响执行速度.
- **mixed**
作为以上两种的折中方案, 通过判断SQL语句是否会带来数据不一致问题而采用 statement 或 row
### MySQL的binlog有几种录入格式?分别有什么区别?
有三种格式,statement,row和mixed.

- statement模式下,记录单元为语句.即每一个sql造成的影响会记录.由于sql的执行是有上下文的,因此在保存的时候需要保存相关的信息,同时还有一些使用了函数之类的语句无法被记录复制.
- row级别下,记录单元为每一行的改动,基本是可以全部记下来但是由于很多操作,会导致大量行的改动(比如alter table),因此这种模式的文件保存的信息太多,日志量太大。
- mixed. 一种折中的方案,普通操作使用statement记录,当无法使用statement的时候使用row. 此外,新版的MySQL中对row级别也做了一些优化,当表结构发生变化的时候,会记录语句而不是逐行记录.
### 两阶段提交
redo log 让 InnoDB 存储引擎拥有 crash-safe 能力; binlog 保证了 MySQL 集群下的数据一致性.
redo log 在事务执行过程中可以不断写入, 而 binlog 只有在提交事务时才写入, 两者写入时机不同.

假设有一个事务正在执行, 执行过程中已经写入了 redo log, 而提交完后 binlog写入时发生异常, 那么在 binlog 中可能就没有对应的更新记录, 之后从库使用 binlog 恢复时, 导致少一次更新操作. 而主库用 redo log 进行恢复, 操作则正常. 最终导致这两个库的数据不一致.

于是 InnoDB存储引擎 使用**两阶段提交**方案 : 将 redo log 的写入拆成了两个步骤 **prepare** 和 **commit**

1. 执行事务时写入redo log (这时处于prepare)
2. 提交事务之前, 先写入 binlog
3. 最后提交事务, 并将 redo log 进行 commit

若使用 redo log 恢复数据时, 发现处于 prepare 阶段, 且没有 binlog, 则会回滚该事务. 若 redo log commit 时异常, 但是存在对应 binlog, MySQL还是认为这一组操作是有效的, 并不会进行回滚.

![image-20220304202100774.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468496177-312faa13-aa9e-4568-bb48-ad3e960f7dd9.png#averageHue=%23e1e7d6&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=Ux6md&margin=%5Bobject%20Object%5D&name=image-20220304202100774.png&originHeight=640&originWidth=433&originalType=binary&ratio=1&rotation=0&showTitle=false&size=141141&status=done&style=none&taskId=u4adb75c5-7cc8-47c7-a651-64c8aa2de9b&title=)

### undo log（回滚日志）
如果需要保证事务的原子性, 就需要在异常发生时, 对已执行操作进行回滚. undo log 会保存事务未提交之前的版本数据, 在执行过程中异常时, 就可以直接利用 undo log 中的信息将数据回滚到未修改之前. 并且 undo log 中的数据可以作为数据的旧版本快照供其他并发事务进行快照读. 在 InnoDB 中也用于实现 MVCC.

#### 1. 如何理解Undo日志
事务需要保证 原子性 ，也就是事务中的操作要么全部完成，要么什么也不做。但有时候事务执行到一半
会出现一些情况，比如：

- 情况一：事务执行过程中可能遇到各种错误，比如 服务器本身的错误 ， 操作系统错误 ，甚至是突然 断电 导致的错误。
- 情况二：程序员可以在事务执行过程中手动输入 ROLLBACK 语句结束当前事务的执行。

以上情况出现，我们需要把数据改回原先的样子，这个过程称之为 回滚 ，这样就可以造成一个假象：这
个事务看起来什么都没做，所以符合 原子性 要求。
#### 2. Undo日志的作用

- 作用1：回滚数据
- 作用2：MVCC
#### 3. undo的存储结构

1. 回滚段与undo页
InnoDB对undo log的管理采用段的方式，也就是 回滚段（rollback segment） 。每个回滚段记录了
1024 个 undo log segment ，而在每个undo log segment段中进行 undo页 的申请。
- 在 InnoDB1.1版本之前 （不包括1.1版本），只有一个rollback segment，因此支持同时在线的事务
限制为 1024 。虽然对绝大多数的应用来说都已经够用。
- 从1.1版本开始InnoDB支持最大 128个rollback segment ，故其支持同时在线的事务限制提高到
了 128*1024 。
```sql
mysql> show variables like 'innodb_undo_logs';
+------------------+-------+
| Variable_name | Value |
+------------------+-------+
| innodb_undo_logs | 128 |
+------------------+-------+
```

2. 回滚段与事务
   1. 每个事务只会使用一个回滚段，一个回滚段在同一时刻可能会服务于多个事务。
   2. 当一个事务开始的时候，会制定一个回滚段，在事务进行的过程中，当数据被修改时，原始的数
据会被复制到回滚段。
   3. 在回滚段中，事务会不断填充盘区，直到事务结束或所有的空间被用完。如果当前的盘区不够
用，事务会在段中请求扩展下一个盘区，如果所有已分配的盘区都被用完，事务会覆盖最初的盘
区或者在回滚段允许的情况下扩展新的盘区来使用。
   4. 回滚段存在于undo表空间中，在数据库中可以存在多个undo表空间，但同一时刻只能使用一个
undo表空间。
   5. 当事务提交时，InnoDB存储引擎会做以下两件事情：
将undo log放入列表中，以供之后的purge操作
判断undo log所在的页是否可以重用，若可以分配给下个事务使用
3. 回滚段中的数据分类	
>    1. 未提交的回滚数据(uncommitted undo information)
>    2. 已经提交但未过期的回滚数据(committed undo information)
>    3. 事务已经提交并过期的数据(expired undo information)


#### 4. undo的类型
在InnoDB存储引擎中，undo log分为：

- insert undo log
- update undo log

#### 5. undo log总结
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1655387070230-7902c807-f5c4-4cc3-b2b9-7de80264d326.png#averageHue=%2389ada8&clientId=u3fe4f8a9-affe-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=410&id=ue5129c3a&margin=%5Bobject%20Object%5D&name=image.png&originHeight=410&originWidth=829&originalType=binary&ratio=1&rotation=0&showTitle=false&size=163277&status=done&style=none&taskId=u7fe50653-4d50-45cb-8072-6333e03242a&title=&width=829)

undo log是逻辑日志，对事务回滚时，只是将数据库逻辑地恢复到原来的样子。
redo log是物理日志，记录的是数据页的物理变化，undo log不是redo log的逆过程。

### MySQL中的六种日志
#### （一）概述
MySQL中存在着以下几种日志：重写日志（redo log）、回滚日志（undo log）、二进制日志（bin log）、错误日志（error log）、慢查询日志（slow query log）、一般查询日志（general log）。
MySQL中的数据变化会体现在上面这些日志中，比如事务操作会体现在redo log、undo log以及bin log中，数据的增删改查会体现在 binlog 中。本章是对MySQL日志文件的概念及基本使用介绍，不涉及底层内容。针对开发人员而言，这几种日志中最有可能使用到的是慢查询日志。
#### （二）redo log
redo log是一种基于磁盘的数据结构，用来在MySQL宕机情况下将不完整的事务执行数据纠正，redo日志记录事务执行后的状态。
当事务开始后，redo log就开始产生，并且随着事务的执行不断写入redo log file中。redo log file中记录了xxx页做了xx修改的信息，我们都知道数据库的更新操作会在内存中先执行，最后刷入磁盘。
redo log就是为了恢复更新了内存但是由于宕机等原因没有刷入磁盘中的那部分数据。
#### （三）undo log

1. undo log主要用来回滚到某一个版本，是一种逻辑日志。undo log记录的是修改之前的数据，比如：当delete一条记录时，undolog中会记录一条对应的insert记录，从而保证能恢复到数据修改之前。在执行事务回滚的时候，就可以通过undo log中的记录内容并以此进行回滚。
2. undo log还可以提供多版本并发控制下的读取（MVCC）。
#### （四）bin log
MySQL的bin log日志是用来记录MySQL中增删改时的记录日志。简单来讲，就是当你的一条sql操作对数据库中的内容进行了更新，就会增加一条bin log日志。查询操作不会记录到bin log中。bin log最大的用处就是进行**主从复制**，以及数据库的恢复。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663236050094-37f9663c-c825-4486-a07d-46a1ab73b3d4.png#averageHue=%23828098&clientId=uc9b92f8f-761c-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u606e8b9d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=326&originWidth=640&originalType=url&ratio=1&rotation=0&showTitle=false&size=84791&status=done&style=none&taskId=u84bf4b2d-5dde-4c9c-8bfc-9b001527a2d&title=)
通过下面的命令可以查看是否开启binlog日志
```shell
show VARIABLES like '%log_bin%'
```
开启binlog的方式如下：
```shell
log-bin=mysql-bin
server-id=1
binlog_format=ROW
```
其中log-bin指定日志文件的名称，默认会放到数据库目录下，可通过以下命令查看
```shell
show VARIABLES like '%datadir%'
```
#### （五）error log
error log主要记录MySQL在启动、关闭或者运行过程中的错误信息，在MySQL的配置文件my.cnf中，可以通过log-error=/var/log/mysqld.log 执行mysql错误日志的位置。
通过MySQL的命令
```shell
show variables like "%log_error%";
```
也可以获取到错误日志的位置。
#### （六）slow query log
慢查询日志用来记录执行时间超过指定阈值的SQL语句，慢查询日志往往用于优化生产环境的SQL语句。可以通过以下语句查看慢查询日志是否开启以及日志的位置：
```shell
show variables like "%slow_query%";
```
慢查询日志的常用配置参数如下：
```shell
slow_query_log=1 #是否开启慢查询日志，0关闭，1开启
slow_query_log_file=/usr/local/mysql/mysql-8.0.20/data/slow-log.log #慢查询日志地址（5.6及以上版本）
long_query_time=1 #慢查询日志阈值，指超过阈值时间的SQL会被记录
log_queries_not_using_indexes #表示未走索引的SQL也会被记录
```
分析慢查询日志一般会用专门的日志分析工具。找出慢SQL后可以通过explain关键字进行SQL分析，找出慢的原因。
#### （七）general log
general log 记录了客户端连接信息以及执行的SQL语句信息，通过MySQL的命令
```shell
show variables like '%general_log%';
```
可以查看general log是否开启以及日志的位置。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663236049994-8295d29e-98cb-4788-9751-0e7fba7e5ae5.png#averageHue=%230b0908&clientId=uc9b92f8f-761c-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u62c45274&margin=%5Bobject%20Object%5D&name=image.png&originHeight=152&originWidth=640&originalType=url&ratio=1&rotation=0&showTitle=false&size=62203&status=done&style=none&taskId=ufcd8b81a-05eb-41bf-bec0-49b1b989ed6&title=)
general log 可通过配置文件启动，配置参数如下：
general_log = on
general_log_file = /usr/local/mysql/mysql-8.0.20/data/hecs-78422.log
普通查询日志会记录增删改查的信息，因此一般是关闭的。

### 事务日志相关高频面试题
> 1. 介绍下MySQL事务日志? redo log和undo log?
> 2. 什么是binlog?


---

## MVCC
有了锁，当前事务没有写锁就不能修改数据，但还是能读的，而且读的时候，即使该行数据其他事务已修改且提交，还是可以重复读到同样的值。这就是**MVCC，多版本的并发控制，Multi-Version Concurrency Control。**

### 一致性非锁定读和锁定读
#### 一致性非锁定读
对于一致性非锁定读(MVCC)的实现, 通常时加一个版本号或时间戳. 查询时, 将当前可见的版本号和对应的版本号进行比对, 若记录的版本号小于可见版本号, 则表示该记录可见.

在 InnoDB 中, 多版本控制(Multi Versioning)就是对非锁定读的实现. 若读取的行正在执行 DELETE 或 UPDATE, 这时读操作不会去等待行锁的释放, 而是读取行的一个快照, 被称为**快照读**

#### 锁定读
也被称为 **当前读**. 锁定读会对读取到的记录加锁.

- `select ... lock in share mode` : 对记录加 S 锁, 其它事务也可以加 S 锁, 但是加 X 锁会被阻塞
- `select ... for update`、`insert`、`update`、`delete` : 对记录加 X 锁

当前读每次读取的都是最新数据, 两次查询中间如果有其他事务插入数据, 就会产生幻读.

### MVCC 实现原理

MVCC是通过保存数据在某个时间点的快照来实现的. 根据事务开始的时间不同, 每个事务对同一张表, 同一时刻看到数据可能是不一样的.
![image-20220304210135517.png](https://cdn.nlark.com/yuque/0/2022/png/21380271/1646468515309-09d9dfe3-929e-46fb-91b7-3a169f96bd6e.png#averageHue=%23eaede1&clientId=ub63f7278-093b-4&crop=0&crop=0&crop=1&crop=1&from=ui&id=UWfiP&margin=%5Bobject%20Object%5D&name=image-20220304210135517.png&originHeight=299&originWidth=681&originalType=binary&ratio=1&rotation=0&showTitle=false&size=121327&status=done&style=none&taskId=u17740057-d3c4-4067-9dff-5b35436a6a3&title=)
MVCC实现依赖于: **隐藏字段**, **Read View**, **undo log**

#### 隐藏字段主要包含:

- ROW ID : 隐藏的自增ID, 如果表没有主键, InnoDB 会自动按 ROW ID 产生一个聚簇索引树
- 事务 ID : 记录最后一次修改该记录的事务ID
- 回滚指针 : 指向这条记录的上一个版本

InnoDB 每行数据都有一个隐藏的回滚指针, 用于指向该行数据修改前的最后一个历史版本, 这个历史版本会存放在 undo log 中. 如果要执行更新操作, 会将原记录放入 undo log 中, 并通过隐藏指针指向 undo log 中的原记录. 其他事务此时需要查询时, 就是查询 undo log 中这行数据的最后一个历史版本.

但是 undo log 总不可能一直保留. 在不需要的时候它应该被删除, 这时就交由系统自动判断, 即当系统没有比这个 undo log 更早的 read-view 的时候. 所以尽量不要使用长事务, 长事务意味着系统里会存在非常古老的事务视图. 由于这些事务随时可能访问数据库中任何数据, 所以这个事务提交前, 数据库里它可能使用到的 undo log 都必须保存, 导致占用大量存储空间.
#### 版本链
Innodb 中行记录的存储格式，有一些额外的字段：**DATA_TRX_ID **和 **DATA_ROLL_PTR**。

- **DATA_TRX_ID**：数据行版本号。用来标识最近对本行记录做修改的事务 id。
- **DATA_ROLL_PTR**：指向该行回滚段的指针。该行记录上所有旧版本，在 undo log 中都通过链表的形式组织。
> undo log : 记录数据被修改之前的日志，后面会详细说。

![image.png](https://cdn.nlark.com/yuque/0/2022/png/22219483/1663236594178-8dd5c313-e422-41c0-9602-f5f68a2b6ce5.png#averageHue=%236eac44&clientId=ub06f29a5-b09f-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=37&id=u0a2c85e4&margin=%5Bobject%20Object%5D&name=image.png&originHeight=73&originWidth=600&originalType=binary&ratio=1&rotation=0&showTitle=false&size=7568&status=done&style=none&taskId=uaec341e2-f127-493f-b51f-e3c91ad4772&title=&width=300)
#### ReadView
在每一条 SQL 开始的时候被创建，有几个重要属性：

**trx_ids: **当前系统活跃(未提交)事务版本号集合。
**low_limit_id:** 创建当前 read view 时“当前系统最大事务版本号+1”。
**up_limit_id:** 创建当前read view 时“系统正处于活跃事务最小版本号”
**creator_trx_id: **创建当前read view的事务版本号；

**开始查询**
现在开始查询，一个 select 过来了，找到了一行数据。

- **DATA_TRX_ID <up_limit_id ：说明数据在当前事务之前就存在了，显示。**
- **DATA_TRX_ID >= low_limit_id：**说明该数据是在当前read view 创建后才产生的，数据不显示。
   - 不显示怎么办，根据 DATA_ROLL_PTR 从 undo log 中找到历史版本，找不到就空。
- ** up_limit_id <DATA_TRX_ID <low_limit_id ：**就要看隔离级别了。

[
](https://blog.csdn.net/liyifan687/article/details/115051198)
### RC(**读提交**)和RR(可重复读)级别下MVCC的差异

- RC级别 : 每次SELECT查询前都生成一个Read View
- RR级别  : 事务开启后第一次SELECT数据前生成一个Read View

### MVCC + Next-key Lock 防止幻读
InnoDB在RR级别下通过 `MVCC` 和 `Next-key Lock` 解决幻读问题

- 在快照读（snapshot read）的情况下，MySQL通过MVCC（多版本并发控制）来避免幻读。
>    - 快照读，读取的是记录的可见版本 (有可能是历史版本)，不用加锁。主要应用于无需加锁的普通查询（select）操作。

- 在当前读（current read）的情况下，MySQL通过next-key lock来避免幻读。
>    - 当前读，读取的是记录的最新版本，并且会对当前记录加锁，防止其他事务发修改这条记录。
>    - 加行共享锁（SELECT ... LOCK IN SHARE MODE ）、加行排他锁（SELECT ... FOR UPDATE / INSERT / UPDATE / DELETE）的操作都会用到当前度。行锁可参看 [MySQL行锁](https://links.jianshu.com/go?to=https%3A%2F%2Fblog.csdn.net%2Fzmflying8177%2Farticle%2Fdetails%2F104826872)。


1. **执行普通 **`**SELECT**`**, 此时会以 **`**MVCC**`** 快照读方式读取数据.**
在快照读的情况下，RR 隔离级别只会在事务开启后的第一次查询生成 `Read View` ，并使用至事务提交。所以在生成 `Read View` 之后其它事务所做的更新、插入记录版本对当前事务并不可见，实现了可重复读和防止快照读下的 “幻读”
2. **执行 select...for update/lock in share mode、insert、update、delete 等当前读**
在当前读下，读取的都是最新的数据，如果其它事务有插入新的记录，并且刚好在当前事务查询范围内，就会产生幻读！

      `InnoDB` 使用 `Next-key Lock` 来防止这种情况。当执行当前读时，会锁定读取到的记录的同时，锁定它们的间隙，防止其它事务在查询范围内插入数据。只要我不让你插入，就不会发生幻读

### [MySQL间隙锁：Next-Key Lock主要知识点](https://www.jianshu.com/p/d5c2613cbb81)
### innoDB的间隙锁/Next-Key Lock
**明确前提条件**

- innoDB的间隙锁只存在于 RR 隔离级别

所以希望禁用间隙锁，提升系统性能的时候，可以考虑将隔离级别降为 RC。

**间隙锁/Next-Key Lock**
间隙锁在innoDB中的唯一作用就是在一定的“间隙”内防止其他事务的插入操作，以此防止幻读的发生：

- 防止间隙内有新数据被插入。
- 防止已存在的数据，更新成间隙内的数据。

### innoDB支持三种行锁定方式：

- 行锁（Record Lock）：锁直接加在索引记录上面（无索引项时演变成表锁）。
- 间隙锁（Gap Lock）：锁定索引记录间隙，确保索引记录的间隙不变。间隙锁是针对事务隔离级别为可重复读或以上级别的。
- Next-Key Lock ：行锁和间隙锁组合起来就是 Next-Key Lock。

1. innoDB默认的隔离级别是可重复读(Repeatable Read)，并且会以Next-Key Lock的方式对数据行进行加锁。
2. Next-Key Lock是行锁和间隙锁的组合，当InnoDB扫描索引记录的时候，会首先对索引记录加上行锁（Record Lock），
3. 再对索引记录两边的间隙加上间隙锁（Gap Lock）。
4. 加上间隙锁之后，其他事务就不能在这个间隙修改或者插入记录。
5. 当查询的索引含有唯一属性（唯一索引，主键索引）时，Innodb存储引擎会对next-key lock进行优化，将其降为record lock,即仅锁住索引本身，而不是范围。
### 何时使用行锁，何时产生间隙锁
对上一节的最后一句做个扩展说明。

1. 只使用唯一索引查询，并且只锁定一条记录时，innoDB会使用行锁。
2. 只使用唯一索引查询，但是检索条件是范围检索，或者是唯一检索然而检索结果不存在（试图锁住不存在的数据）时，会产生 Next-Key Lock。
3. 使用普通索引检索时，不管是何种查询，只要加锁，都会产生间隙锁。
4. 同时使用唯一索引和普通索引时，由于数据行是优先根据普通索引排序，再根据唯一索引排序，所以也会产生间隙锁。

### MVCC相关高频面试题
> 1. 了解MVCC吗?说下什么是MVCC?
> 2. MVCC实现原理? 有什么好处?


---

## 锁

根据加锁范围, MySQL里的锁大致分成 **全局锁** 、**表锁** 和 **行锁**

### 全局锁

全局锁就是对整个数据库实例加锁. MySQL提供了一个加全局读锁的方法, `Flush tables with read lock(FTWRL)` 使整个库都处于只读状态. 一般用于全局备份.

### 表锁

InnoDB中的表锁十分鸡肋, 一般都是通过 `MySQL` 的 server 层下的 **元数据锁 (Metadata Lock)** 来实现当对表执行DDL语句时, 使得其他事务阻塞. InnoDB厉害之处是实现了更细粒度的行锁.

### 行锁
**在 InnoDB 事务中，行锁通过给索引上的索引项加锁来实现。这意味着只有通过索引条件检索数据，InnoDB才使用行级锁，否则将使用表锁。**
行级锁定同样分为两种类型：**共享锁**和**排他锁**，以及加锁前需要先获得的意向共享锁和意向排他锁。

- 共享锁：读锁，允许其他事务再加S锁，不允许其他事务再加X锁，即其他事务只读不可写。select...lock in share mode 加锁。
- 排它锁：写锁，不允许其他事务再加S锁或者X锁。insert、update、delete、for update加锁。

**行锁是在需要的时候才加上的，但并不是不需要了就立刻释放，而是要等到事务结束时才释放。这个就是两阶段锁协议。**
#### 行锁的实现算法

-  **Record Lock 记录锁：**单个行记录上的锁，总是会去锁住索引记录。
仅仅只是把一条记录给锁上, 当一个事务获取了记录的S锁, 其他事务也能够获得S锁, 但无法获得X锁; 而一个事务获取了记录的X锁, 则其他事务不能获得S锁和X锁. 
-  **Gap Lock 间隙锁 ** 解决幻读
在RR级别下产生幻读问题一般有两种解决方案: MVCC或加锁.
假设有记录id = [1, 3, 4], 这时有事务想在[1, 3]之间插入一个id = 2的记录, 可能产生幻读, 所以给id = 3的记录前的间隙加上了 Gap Lock, 意思是在释放这个 Gap Lock 之前都不允许其他事务在这条记录前插入数据. 那如果在最后一条后插入数据呢 ? 在表最后会有一条伪记录Supremum, 对Supremum加 Gap Lock,  这样就会防止在最后插入记录造成的幻读. 
-  **Next-Key Lock**
有时, 既想锁住某条记录, 又想阻止在记录前的间隙插入新记录. 于是就有了 Next-Key Lock. 本质就是一个Record Lock + Gap Lock，左开又闭. 
-  **Insert Intention Lock 插入意向锁**
若插入位置已被别的事务加了 Gap Lock, 则事务在等待时也需要在内存中生成一个锁结构, 被称为 插入意向锁. 当 Gap Lock 释放的时候,插入意向锁就会将等待事务中锁结构内的 is_waiting 的状态改为 false, 然后开始继续往下执行插入操作. 
-  **隐式锁**
隐式锁其实是一种延迟生成锁结构的方案, 通过判断事务id, 确定两个并发事务之间是否真的有必要加锁, 若需要, 则会生成锁结构, 然后进入等待; 不需要, 那么就没必要浪费内存去对事务生成锁结构, 降低维护成本. 类似于乐观锁实现. 
### 锁之于隔离性
大致介绍了下锁，可以看到。有了锁，当某事务正在写数据时，其他事务获取不到写锁，就无法写数据，一定程度上保证了事务间的隔离。但前面说，**加了写锁，为什么其他事务也能读数据呢，不是获取不到读锁吗**？
### 两阶段锁协议

在InnoDB事务中, 行锁是需要的时候才加上的, 但并不是不需要了就立刻释放, 而是等到事务结束时才释放. 这就是 **两阶段锁协议**

如果事务中需要锁住多行, 要把最可能造成锁冲突、最可能影响并发度的锁尽量往后放. 这样就最大程度减少了事务间的锁等待, 提升了并发度.
### 锁相关高频面试题
> 1. 为什么需要加锁?
> 2. MySQL锁粒度?
> 3. MySQL有哪些锁?
> 4. 乐观锁和悲观锁是什么?如何实现?
> 5. InnoDB的行锁是如何实现的?
> 6. 什么是两阶段锁协议?

