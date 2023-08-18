# golang使用信号量并发安全的迁移千万条数据

由于公司业务需要: 需将 PostgreSQL 数据库中的9百万条数据 迁移到 MySQL.

现将迁移脚本的开发过程记录如下:

### 安装驱动库

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u gorm.io/driver/mysql
```

### 初始化数据库连接

```go
import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var MySQLClientBI *gorm.DB
var PostgreSQLClient *gorm.DB

func init() {
    dsnPg := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    pgDB, err := gorm.Open(postgres.Open(dsnPg), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	PostgreSQLClient = pgDB
    
    dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})  
    if err != nil {
		fmt.Println(err)
	}
	MySQLClientBI = db
}
```



### 批量插入

GORM 的 **CreateInBatches** 方法可以用于批量插入数据，这确实有助于提高大量数据的插入效率。但是，有几点需要注意：

1. **内存使用**：即使使用 **CreateInBatches**，如果你首先从一个数据库中提取900万条记录并尝试将其存储在内存中，那么可能会出现内存使用过高的问题。你应该在查询数据时考虑分页或限制提取的记录数。
2. **批次大小**：为了达到最佳性能和避免潜在的问题，您需要确定合适的批次大小。例如，一次插入1000或5000条记录，而不是所有900万条记录。
3. **MySQL的限制**：MySQL有一个**max_allowed_packet**参数，它定义了单个客户端发送到MySQL服务器的数据包的最大大小。批量插入时可能会触及此限制，导致错误。

基于以上考虑，建议以下方法：

1. 分批从 PostgreSQL 中读取数据，例如每次读取5000条。
2. 使用 **CreateInBatches** 将每批数据插入到 MySQL 中。

这样可以确保不会因为一次性处理大量数据而耗尽内存，并且可以在必要时轻松调整批次大小。

简化代码如下:

```go
const batchSize = 5000
var offset = 0
for {
    // 从pg查,一次查 5000
	var cfgEventParamsValue []EventParamsValue
	result := pgDB.Table("data_cfg.cfg_event_params_value").Limit(batchSize).Offset(offset).Find(&cfgEventParamsValue)
	if result.Error != nil {
		log.Fatalf("Error fetching from PostgreSQL: %v", result.Error)
	}
	if len(cfgEventParamsValue) == 0 {
		break
	}
	// 入mysql,一次入 5000
	err := mysqlDB.Table("cfg_event_params_value").CreateInBatches(cfgEventParamsValue, batchSize).Error
	if err != nil {
		log.Fatalf("Error inserting batch into MySQL: %v", err)
	}
	offset += batchSize
}
```

### 开启协程

使用 Go 的协程可以极大地提高数据迁移的效率，特别是当涉及到网络IO或数据库IO操作时。但请注意，太多的并发可能会对数据库造成压力，导致性能下降或其他问题，所以需要平衡。

代码如下:

```go
package main

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EventParamsValue struct {
	ID    uint   `gorm:"primaryKey"`
	// 根据实际字段和类型调整以下内容
	FieldName string
}

var pgDB *gorm.DB
var mysqlDB *gorm.DB
var wg sync.WaitGroup	// 使用 sync.WaitGroup 来确保主程序等待所有的协程完成

const batchSize = 5000

func migrateBatch(offset int) {
	defer wg.Done()

     // 从pg查,一次查 5000
	var cfgEventParamsValue []EventParamsValue
	result := pgDB.Table("data_cfg.cfg_event_params_value").Limit(batchSize).Offset(offset).Find(&cfgEventParamsValue)
	if result.Error != nil {
		log.Printf("Error fetching from PostgreSQL: %v", result.Error)
		return
	}

	if len(cfgEventParamsValue) == 0 {
		return
	}
    
	// 入mysql,一次入 5000
	err := mysqlDB.Table("cfg_event_params_value").CreateInBatches(cfgEventParamsValue, batchSize).Error
	if err != nil {
		log.Printf("Error inserting batch into MySQL: %v", err)
	}
}

func main() {
	// PostgreSQL 连接 (省略了代码...)

	// MySQL 连接 (省略了代码...)

	// 获取总记录数
	var totalRecords int64
	pgDB.Table("data_cfg.cfg_event_params_value").Count(&totalRecords)

	for offset := 0; offset < int(totalRecords); offset += batchSize {
		wg.Add(1)
		go migrateBatch(offset)  // 使用协程执行数据迁移
	}

	wg.Wait() // 等待所有协程完成
	fmt.Println("Migration complete!")
}
```

协程非常快，所以可能会很快地打开很多协程。如果发现数据库响应缓慢或有其他问题，可能需要引入一个限制并发数量的机制，例如使用通道 (channel) 或第三方库，如 [semaphore](https://pkg.go.dev/golang.org/x/sync/semaphore)。

在执行此程序之前，务必先在非生产环境中测试，确保其行为如预期，并确保它不会对您的数据库产生不良影响。

上述这段代码根据数据总量和每批处理的数据量（**batchSize**）来决定开启多少个协程。

这里是决定开启协程数量的关键部分：

```go
for offset := 0; offset < int(totalRecords); offset += batchSize {
	wg.Add(1)
	go migrateBatch(offset)  // 使用协程执行数据迁移
}
```

每次迭代中，我们都会开启一个新的协程。迭代的次数由总记录数（**totalRecords**）和每批的大小（**batchSize**）决定。

计算开启的协程数量的公式为：

```go
numGoroutines = ceil(totalRecords / batchSize)
```

其中 **ceil** 是向上取整函数。例如，如果您有 9,000,000 条记录，并且每批大小是 5,000，那么您会开启 1,800 个协程。

需要注意的是，尽管协程在 Go 中非常轻量，但同时开启太多协程可能会导致数据库连接的问题，尤其是当数据库的连接池大小有限时。您可能需要考虑增加数据库的连接数限制或使用信号量来限制同时运行的协程数量，以保护数据库不被过度压迫。

### 增加信号量模式

为了手动控制协程数量，可以使用Go中的信号量模式，利用**chan struct{}**来达到这个目的。

```go
package main

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EventParamsValue struct {
	ID    uint   `gorm:"primaryKey"`
	// 根据实际字段和类型调整以下内容
	FieldName string
}

var pgDB *gorm.DB
var mysqlDB *gorm.DB
var wg sync.WaitGroup

const batchSize = 5000
const maxGoroutines = 10 // 手动设置最大并发协程数量

var sem = make(chan struct{}, maxGoroutines)

func migrateBatch(offset int) {
	defer wg.Done()
	defer func() { <-sem }() // 释放一个协程位

	var cfgEventParamsValue []EventParamsValue
	result := pgDB.Table("data_cfg.cfg_event_params_value").Limit(batchSize).Offset(offset).Find(&cfgEventParamsValue)
	if result.Error != nil {
		log.Printf("Error fetching from PostgreSQL: %v", result.Error)
		return
	}

    // 如果该批次没有数据，直接返回
	if len(cfgEventParamsValue) == 0 {
		return
	}

    log.Printf("Migrating records from offset %d to %d", offset, offset+batchSize-1) // 记录每批次迁移数据的起止

	err := mysqlDB.Table("cfg_event_params_value").CreateInBatches(cfgEventParamsValue, batchSize).Error
	if err != nil {
		log.Printf("Error inserting batch into MySQL from offset %d to %d: %v", offset, offset+batchSize-1, err)
	} else {
		log.Printf("Successfully migrated records from offset %d to %d", offset, offset+batchSize-1)
	}
}

func main() {
	// PostgreSQL 连接 (省略了代码...)

	// MySQL 连接 (省略了代码...)

	// 获取总记录数
	var totalRecords int64
	pgDB.Table("data_cfg.cfg_event_params_value").Count(&totalRecords)

	for offset := 0; offset < int(totalRecords); offset += batchSize {
		sem <- struct{}{} // 获取一个协程位
		wg.Add(1)
		go migrateBatch(offset)  // 使用协程执行数据迁移
	}

	wg.Wait() // 等待所有协程完成
	fmt.Println("Migration complete!")
}
```

在上面的代码中：

- **sem** 是一个有限容量的通道，用于控制并发的协程数量。
- **sem <- struct{}{}** 尝试向通道发送一个空结构，如果通道已满，这一行将会阻塞，直到有其他协程完成并释放一个位置。
- **defer func() { <-sem }()** 保证当协程结束时，从**sem**通道中移除一个空结构，从而释放一个协程位置。

这样，一次只有**maxGoroutines**数量的协程能够并发运行。您可以根据需要调整**maxGoroutines**的值来控制并发数量。