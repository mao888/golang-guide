/**
    @author:huchao
    @data:2022/3/12
    @note: string 和 byte[]
**/
package main

import "fmt"

func main()  {
	var s []byte
	s = append(s,'1','2','s','a')
	str := GetStringBySlice(s)
	fmt.Println(len(s))
	fmt.Println(len(str))
	fmt.Println("%p\n,%p\n",s,str)
}

//[]byte转string
func GetStringBySlice(s []byte) string {
	return string(s)
}