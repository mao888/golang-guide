### 表格数据

| app_id | 参数名称 | 参数值 | 参数值名称 |
| ------ | -------- | ------ | ---------- |
| 100074 | qwe12345 | mao11  | 超超1      |
| 100074 | ssss     | mao3   | 超超1      |
| 100074 | qwe12    | ma01   | 超超2      |
| 100074 | www      | mao3   | 超超3      |
| 100074 | www      | mao4   | 超超4      |
| 100036 | qwe12345 | 1233   | 123        |
| 100036 | ssss     | mao3   | 123        |
| 100036 | qwe12    | 100036 | 100036     |
| 100036 | www      | www1   | dd         |

### 对应结构体

```go
type appEvents struct {
    ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
    AppID       string `gorm:"column:app_id;not null" json:"app_id"`               // 应用id
    Params      string `gorm:"column:params;not null" json:"params"`               // 参数名称
    ParamsValue string `gorm:"column:params_value;not null" json:"params_value"`   // 参数值
    ParamsLabel string `gorm:"column:params_label;not null" json:"params_label"`   // 参数值名称
    Creator     int32  `gorm:"column:creator;not null" json:"creator"`             // 创建人id
    CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
    UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
}
```

### 错误示例

```go
appEventsIDStr := make([]string, 0, len(appEvents))
appEventsParams := make([]string, 0, len(appEvents))
for _, event := range appEvents {
    appEventsIDStr = append(appEventsIDStr, event.AppID)
    appEventsParams = append(appEventsParams, event.Params)
}
_, err = dao.WithContext(ctx).Where(dao.AppID.In(appEventsIDStr...), dao.Params.In(appEventsParams...)).Delete()
if err != nil {
    return err
}
```

这段代码的意图是：直接将表格中的**AppID**和**Params**装入分别**appEventsIDStr**和**appEventsParams**中

然后批量删除那些其 **AppID** 值存在于 **appEventsIDStr** 且其 **Params** 值存在于 **appEventsParams** 的记录。

这会造成的问题如下:

正常删除了如下数据:

| app_id | 参数名称 |
| ------ | -------- |
| 100074 | qwe12345 |

但同时也会误删其他app_id下名称为 qwe12345 的数据:

| app_id | 参数名称 |
| ------ | -------- |
| 100036 | qwe12345 |

只要是 **appEventsIDStr** 包含的应用id 和 **appEventsParams**  中包含的参数名称都会被全部删除!

这显然是不符合业务需求的:删除指定app_id和参数名称下的数据

### 正确逻辑

1. **创建映射**：为 **appEvents** 创建一个映射，其中键由 **AppID** 和 **Params** 组合而成。
2. **从数据库获取数据**：查询数据库，获取那些其 **AppID** 和 **Params** 在指定的列表中的数据。
3. **收集要删除的 IDs**：根据先前创建的映射，筛选出数据库中需要删除的条目的 IDs。
4. **删除数据库条目**：基于筛选出来的 IDs，执行数据库中的删除操作。

### 代码实现

```go
// 将 appEventsIDStr 和 appEventsParams 组合成 map的key ,value为空结构体的map
	appEventsMap := make(map[string]struct{}, len(appEventsIDStr))
	for _, event := range appEvents {
		appEventsMap[event.AppID+"_"+event.Params] = struct{}{}
	}

	// 批量删除旧数据
	dao := tx.CfgEventParamsValue
	find, err := dao.WithContext(ctx).Where(dao.AppID.In(appEventsIDStr...), dao.Params.In(appEventsParams...)).Find()
	if err != nil {
		return err
	}
	var ids []int32
	for _, value := range find {
		// 将数据库中的数据组合成map的key
		// value.AppID + "_" + value.Params 为map的key
		// value.ParamsValue 为map的value
		_, exists := appEventsMap[value.AppID+"_"+value.Params]
		if exists {
			ids = append(ids, value.ID)
		}
	}
	_, err = dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
	if err != nil {
		return err
	}
```

### 代码详解

1. **构建****appEventsMap****映射**:

```go
appEventsMap := make(map[string]struct{}, len(appEventsIDStr))
for _, event := range appEvents {
    appEventsMap[event.AppID+"_"+event.Params] = struct{}{}
}
```

在这段代码中，首先初始化了一个名为 **appEventsMap** 的映射，其键是字符串，值是空结构体（**struct{}** 不占用任何额外的内存空间）。接下来，遍历每个 **appEvents** 条目并构造映射键，该键是 **AppID** 和 **Params** 的组合，并为每个键分配一个空结构体值。**目的**: 将 **appEvents** 列表转换为映射，以便后续可以快速检查某个组合（**AppID** 和 **Params**）是否存在于 **appEvents** 中。

1. **查询出要删除的旧数据**:

```go
appEventsIDStr := make([]string, 0, len(appEvents))
appEventsParams := make([]string, 0, len(appEvents))
for _, event := range appEvents {
    appEventsIDStr = append(appEventsIDStr, event.AppID)
    appEventsParams = append(appEventsParams, event.Params)
}

dao := tx.CfgEventParamsValue
find, err := dao.WithContext(ctx).Where(dao.AppID.In(appEventsIDStr...), dao.Params.In(appEventsParams...)).Find()
if err != nil {
    return err
}
```

使用 **dao** 查询数据库中的数据。此查询的条件是 **AppID** 必须在 **appEventsIDStr** 中，**Params** 必须在 **appEventsParams** 中。

1. **筛选要删除的数据ID**:

```go
var ids []int32
for _, value := range find {
    _, exists := appEventsMap[value.AppID+"_"+value.Params]
    if exists {
        ids = append(ids, value.ID)
    }
}
```

在这个循环中，您检查从数据库查询返回的每个条目是否存在于之前创建的 **appEventsMap** 中。如果存在，那么该条目的 **ID** 被添加到要删除的ID列表中。

1. **执行删除操作**:

```go
_, err = dao.WithContext(ctx).Where(dao.ID.In(ids...)).Delete()
if err != nil {
    return err
}
```

最后，执行实际的删除操作。您通过在**ids**列表中指定的ID来删除所有条目。

**综合分析**: 该代码的主要目的是从数据库中删除特定的旧数据。这些数据由其**AppID**和**Params**确定，并在**appEvents**列表中给出。首先，代码将这个列表转换为一个映射，以便可以快速检查一个给定的**AppID**和**Params**组合是否存在。然后，代码查询与这些**AppID**和**Params**组合匹配的所有数据库条目，并删除这些条目。