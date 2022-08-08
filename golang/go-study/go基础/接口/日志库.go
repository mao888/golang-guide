/**
    @author: HuChao
    @since: 2022/8/4
    @desc: //TODO 使用接口的方式实现一个既可以往终端写日志也可以往文件写日志的简易日志库
**/
package main

import (
	"fmt"
	"os"
	"time"
)

// Logger 接口
type Logger interface {
	consoleLog() // 终端
	fileLog()    // 文件
}

// User 用户结构体
type User struct {
	username string
	password string
}

//	User 实现方法
func (u User) consoleLog() {
	t := time.Now()
	fmt.Printf("用户创建成功！用户名为：%s", u.username)
	fmt.Printf("用户密码是：%s\n", u.password)
	fmt.Printf("创建完成时间：%d-%d-%d %d:%d:%d\n", t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}

func (u User) fileLog() {
	t := time.Now()
	file, err := os.OpenFile("./Go/go基础/接口/"+u.username+".txt", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data := "用户创建成功！ 用户名为：" + fmt.Sprintf("%s\n", u.username) +
		"密码是：" + u.password + "\n" +
		fmt.Sprintf("创建完成时间：%d-%d-%d %d:%d:%d\n",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println(err)
	}
}

// User中的字段初始化
func newUser(username, password string) User {
	return User{
		username: username,
		password: password,
	}
}

// 创建用户对象
func createUser() {
	var (
		username string
		password string
	)
	fmt.Print("请输入用户名: ")
	_, err := fmt.Scan(&username)
	fmt.Print("请输入密码: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("输入错误！！ ERROR", err)
	}
	u := newUser(username, password)
	u.consoleLog()
	u.fileLog()
}

func main() {
	createUser()
}
