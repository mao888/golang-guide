package main

import (
	"context"
	"fmt"
)

type keyType string

// Value(key interface{}) interface{} 是 context.Context 接口中的一个方法，用于在 context.Context 中存储和检索键值对信息。
// 这个方法允许你在 context.Context 中关联任意键（key）和值（value），以便将一些请求特定的数据传递给相关的函数和 goroutine。

func main() {
	// 创建一个父 Context
	parent := context.Background()

	// 在父 Context 中使用 context.WithValue 存储键值对信息
	key := keyType("user")
	value := "john_doe"
	ctx := context.WithValue(parent, key, value)

	// 在子函数中检索存储的值
	retrieveValue(ctx, key)
}

func retrieveValue(ctx context.Context, key keyType) {
	// 使用 context.Value 方法来检索存储的值
	if value, ok := ctx.Value(key).(string); ok {
		fmt.Printf("Value for key %s: %s\n", key, value)
	} else {
		fmt.Printf("Key %s not found in the context\n", key)
	}
}
