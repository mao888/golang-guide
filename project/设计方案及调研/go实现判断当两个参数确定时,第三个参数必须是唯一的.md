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
type CfgEventParamsValue struct {
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

### 业务需求

同一 **app_id**,同一 **参数名称** 下的**参数值** 不能相同,有错误就直接return

### 实现逻辑

逻辑：通过组合第一和第二个参数创建一个唯一键，并使用此键在映射中跟踪第三个参数的唯一性

### 代码实现

```go
func checkDuplicates(data []CfgEventParamsValue) bool {
	valuesMap := make(map[string]map[string]bool)

	for _, entry := range data {
		key := fmt.Sprintf("%s:%s", entry.AppID, entry.Params)

		if _, ok := valuesMap[key]; !ok {
			valuesMap[key] = make(map[string]bool)
		}

		if _, exists := valuesMap[key][entry.ParamsValue]; exists {
			// Detect duplicate
			fmt.Printf("Error! Duplicate detected: AppID: %s, Params: %s, ParamsValue: %s\n", entry.AppID, entry.Params, entry.ParamsValue)
			return false // found duplicate
		} else {
			valuesMap[key][entry.ParamsValue] = true
		}
	}

	return true // no duplicates found
}
```

### 代码详解

传入一个结构体数组 **[]CfgEventParamsValue**，该结构体的字段描述了我们要处理的数据的格式。

定义了一个函数 **checkDuplicates(data []CfgEventParamsValue) bool**，该函数的目的是检查提供的数据中是否有重复项。

1. **创建一个映射****valuesMap**：

```go
goCopy code
valuesMap := make(map[string]map[string]bool)
```

这是一个嵌套的映射。最外层的映射键是由 **AppID** 和 **Params** 组成的字符串，而内层的映射键则是 **ParamsValue**。这样的结构使我们能够轻松地跟踪每个 **AppID** 和 **Params** 组合下的所有 **ParamsValue**。

1. **遍历数据**：

```go
for _, entry := range data {
```

我们逐个遍历提供的数据条目。

1. **为每个数据条目创建键**：

```go
key := fmt.Sprintf("%s:%s", entry.AppID, entry.Params)
```

这将 **AppID** 和 **Params** 合并为一个字符串键，例如 "100069:ww"。

1. **检查外部映射中是否已存在此键**：

```go
if _, ok := valuesMap[key]; !ok {
    valuesMap[key] = make(map[string]bool)
}
```

如果此键不存在，则在外部映射中为其创建一个新的内部映射。

1. **检查内部映射中是否已存在** **ParamsValue**：

```go
if _, exists := valuesMap[key][entry.ParamsValue]; exists {
```

如果存在，那么我们就找到了一个重复项。

1. **报告错误并返回**：

```go
fmt.Printf("Error! Duplicate detected: AppID: %s, Params: %s, ParamsValue: %s\n", entry.AppID, entry.Params, entry.ParamsValue)
return false // found duplicate
```

如果检测到重复项，我们打印一条错误消息并返回 **false**。

1. **如果不重复，将值添加到内部映射中**：

```go
} else {
    valuesMap[key][entry.ParamsValue] = true
}
```

这确保我们在之后的迭代中跟踪此值。

在函数的末尾，如果没有检测到任何重复项，我们返回 **true**。