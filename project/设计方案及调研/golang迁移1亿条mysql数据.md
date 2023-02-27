迁移1亿条MySQL数据需要考虑到数据的大小和迁移的速度。以下是一些使用Golang迁移MySQL数据的建议：

1. 优化数据库：在迁移之前，您可以考虑优化数据库，如使用索引和分区表等技术，以提高读写速度。
2. 分批次处理：将1亿条数据分批处理可以提高迁移的效率。您可以使用SQL语句中的LIMIT和OFFSET关键字来分页读取数据。例如：

```go
limit := 10000
offset := 0
for {
    rows, err := db.Query("SELECT * FROM table LIMIT ? OFFSET ?", limit, offset)
    if err != nil {
        // 处理错误
    }
    // 处理数据
    offset += limit
    if len(rows) == 0 {
        break
    }
}
```

上面的代码会每次读取10000条数据，直到读取完所有数据为止。

3. 并发处理：您可以使用Go语言的goroutine来实现并发处理数据，以加快迁移速度。例如：

```go
limit := 10000
offset := 0
var wg sync.WaitGroup
for {
    rows, err := db.Query("SELECT * FROM table LIMIT ? OFFSET ?", limit, offset)
    if err != nil {
        // 处理错误
    }
    // 处理数据
    offset += limit
    if len(rows) == 0 {
        break
    }
    wg.Add(1)
    go func(rows []Row) {
        defer wg.Done()
        // 并发处理数据
    }(rows)
}
wg.Wait()
```

上面的代码会并发地处理数据，加快迁移速度。

4. 批量插入：在迁移数据时，您可以使用批量插入的方式来提高写入速度。例如：

```go
values := []string{}
for _, row := range rows {
    // 处理数据
    values = append(values, fmt.Sprintf("(%d, %s)", row.ID, row.Name))
}
_, err := db.Exec("INSERT INTO table (id, name) VALUES " + strings.Join(values, ","))
```

上面的代码会将多个值拼接成一个SQL语句，一次性插入多条数据。

5. 优化事务：在迁移数据时，您可以使用事务来保证数据的一致性和可靠性。例如：

```go
tx, err := db.Begin()
if err != nil {
    // 处理错误
}
for _, row := range rows {
    // 处理数据
    _, err := tx.Exec("INSERT INTO table (id, name) VALUES (?, ?)", row.ID, row.Name)
    if err != nil {
        tx.Rollback()
        // 处理错误
    }
}
err = tx.Commit()
if err != nil {
    // 处理错误
}
```

上面的代码会将多条插入操作放在一个事务中，以保证数据的一致性。

以上是一些使用Golang迁移MySQL数据的建议，希望对您有所帮助。