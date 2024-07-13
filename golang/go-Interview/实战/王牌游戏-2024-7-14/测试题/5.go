package main

import (
	"errors"
	"fmt"
	"sync"
)

// MemoryBlock 表示内存池中的一个内存块
type MemoryBlock struct {
	Object interface{} // 理论上可以存储任何对象，但这里不直接使用
	Index  int         // 内存块在内存池中的索引
	Size   int         // 内存块的大小（在这个简单示例中，所有块大小相同）
	Used   bool        // 标记该内存块是否已被使用
}

// MemoryPool 表示内存池
type MemoryPool struct {
	Blocks    []*MemoryBlock // 内存块数组
	PoolSize  int            // 内存池的大小（即包含的块数）
	BlockSize int            // 每个内存块的大小（这里仅用于演示，实际不分配具体内存）
	mutex     sync.Mutex     // 锁，用于并发控制
}

// NewMemoryPool 创建一个新的内存池
func NewMemoryPool(poolSize, blockSize int) *MemoryPool {
	blocks := make([]*MemoryBlock, poolSize)
	for i := 0; i < poolSize; i++ {
		blocks[i] = &MemoryBlock{
			Index: i,
			Size:  blockSize,
			Used:  false,
		}
	}
	return &MemoryPool{
		Blocks:    blocks,
		PoolSize:  poolSize,
		BlockSize: blockSize,
	}
}

// Malloc 从内存池中分配一个内存块
// 注意：这个实现是简化的，不实际分配内存，仅标记为已使用
func (mp *MemoryPool) Malloc() (*MemoryBlock, error) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	for _, block := range mp.Blocks {
		if !block.Used {
			block.Used = true
			return block, nil
		}
	}
	return nil, errors.New("no available memory block")
}

// Free 释放一个内存块
func (mp *MemoryPool) Free(block *MemoryBlock) error {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()

	if block.Index < 0 || block.Index >= mp.PoolSize {
		return errors.New("invalid memory block index")
	}
	if mp.Blocks[block.Index] != block {
		return errors.New("memory block does not belong to this pool")
	}
	block.Used = false
	return nil
}

func main() {
	// 创建一个包含10个块，每个块“大小”为10（仅用于演示）的内存池
	mp := NewMemoryPool(10, 10)

	// 分配内存块
	block, err := mp.Malloc()
	if err != nil {
		fmt.Println("Error allocating memory:", err)
		return
	}
	fmt.Printf("Allocated block: Index=%d, Size=%d, Used=%t\n", block.Index, block.Size, block.Used)

	// 释放内存块
	err = mp.Free(block)
	if err != nil {
		fmt.Println("Error freeing memory:", err)
		return
	}
	fmt.Printf("Freed block: Index=%d, Size=%d, Used=%t\n", block.Index, block.Size, block.Used)
}
