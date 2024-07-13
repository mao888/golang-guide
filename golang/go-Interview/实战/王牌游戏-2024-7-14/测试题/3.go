package main

import (
	"fmt"
	"hash/fnv"
)

//设计一个名为”abTestHash"的哈希函数，用于AB测试的分组。该函数接受玩家的UUID（int64）和AB测试的ID（int32）作为参数，并返回分组概率（int32）。要求分组均匀且无耦合。

// abTestHash 接收玩家的UUID和AB测试的ID，并返回一个表示分组的int32值（0为A组，1为B组）
func abTestHash(uuid int64, testID int32) int32 {
	//使用哈希函数对UUID和AB测试ID的组合进行哈希
	// 创建一个FNV哈希器
	hasher := fnv.New32a()

	// 将UUID和测试ID写入哈希器（先写入UUID，再写入测试ID，以确保每个玩家在不同的测试中具有不同的哈希值
	combined := fmt.Sprintf("%d%d", uuid, testID)
	hasher.Write([]byte(combined))

	// 使用哈希值的最低位来决定分组（0或1）
	hashValue := hasher.Sum32()
	return int32(hashValue & 1) // 如果最低位是0，返回0（A组）；如果是1，返回1（B组）
}

func main() {
	// 示例用法
	uuid := int64(123456789012345678)
	testID := int32(4)

	group := abTestHash(uuid, testID)
	fmt.Printf("UUID %d; Test ID %d; group %d\n", uuid, testID, group)
}
