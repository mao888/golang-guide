package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		// 设置响应头，指定内容类型为text/plain
		c.Header("Content-Type", "text/plain")
		c.Header("Transfer-Encoding", "chunked")

		// 创建一个通道用于模拟流式数据
		dataStream := make(chan string)

		// 启动goroutine，模拟生成流式数据
		go func() {
			defer close(dataStream)
			for i := 0; i < 10; i++ {
				time.Sleep(1 * time.Second)
				dataStream <- fmt.Sprintf("Data %d\n", i)
			}
		}()

		// 将流式数据写入响应主体
		for data := range dataStream {
			_, _ = io.WriteString(c.Writer, data)
			c.Writer.Flush()
		}
	})

	r.Run(":8080")
}
