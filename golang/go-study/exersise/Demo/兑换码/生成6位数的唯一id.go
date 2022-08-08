/**
    @author:Hasee
    @data:2022/4/9
    @note:
**/
package main

import (
	//"database/sql"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {
	RandomTestBase()
}

//生成若干个不重复的随机数
func RandomTestBase() {
	//测试5次
	//for i := 0; i < 1; i++ {
	nums := generateRandomNumber(100000, 999999, 800000)
	fmt.Println(len(nums))
	//fmt.Println(nums)

	var wireteString = ""
	var filename = "./output1.txt"
	var f *os.File
	var err1 error
	/***************************** 第一种方式: 使用 io.WriteString 写入文件 ***********************************************/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	checkErr(err1)

	wireteString = "insert into king_identifier (recordid) values "
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	checkErr(err1)
	fmt.Printf("写入 %d 个字节n", n)

	for k, v := range nums {
		fmt.Println(k)
		fmt.Println(v)
		wireteString = "(" + strconv.Itoa(v) + "),"
		n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
		checkErr(err1)
		fmt.Printf("写入 %d 个字节n", n)
	}
	io.WriteString(f, ";")
	//}
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//db, err := sql.Open("mysql", "game:game@/kingclubdb")
	//checkErr(err)

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
			//插入数据
			//stmt, err := db.Prepare("insert into king_identifier (recordid) values (?)")
			//checkErr(err)
			//res, err := stmt.Exec(num)
			//checkErr(err)
			//id, err := res.LastInsertId()
			//checkErr(err)
			//
			//fmt.Println(id)
			//stmt.Close();
		}
	}

	return nums
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func TestTransRate(t *testing.T) {

	RandomTestBase()

}
