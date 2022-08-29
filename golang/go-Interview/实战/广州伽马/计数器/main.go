/**
    @author:Hasee
    @data:2022/3/15
    @note: golang 利用接口实现多态统计计数器
**/
/**
	功能：实现一个计数器模块，不依赖外部三方模块和存储，
	要求进程内协程安全、异步、高性能按指标 key-value 统计。
 */
package main

import (
	"fmt"
	"time"
	"./statistics"
)

func main(){
	go Task()

	for {
		func() {
			fmt.Printf("can: success = %d \n",statistics.CanClient.IncSuccess())
			fmt.Printf("can: error = %d \n",statistics.CanClient.IncError())
		}()

		time.Sleep(time.Second)
		fmt.Printf("can: error = %d \n",statistics.CanClient.GetError())
		fmt.Printf("can: count = %d \n",statistics.CanClient.GetCount())
	}

}

func Task() {
	for {
		func() {
			fmt.Printf("gps:success = %d\n", statistics.GpsClient.IncSuccess())
		}()
		time.Sleep(time.Second)
		fmt.Printf("gps count = %d \n", statistics.GpsClient.GetCount())
	}
}
