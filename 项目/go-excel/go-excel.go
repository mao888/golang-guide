package main

import (
	"fmt"
	"os"

	gexcel "github.com/mao888/go-excel"
)

const HeaderAlignCenter = "center"

func main() {

	sheetName := "测试sheet"
	header := []string{"姓名", "性别", "年龄"}
	content := [][]interface{}{
		{"胡超", "男", 18},
	}
	file := gexcel.NewExcel(sheetName, header, content)
	file.SetAlign(HeaderAlignCenter)
	file.SetColWidth(50)
	data, err := file.Export()
	if err != nil {
		fmt.Println("TestNewExcel err: %v ", err)
		return
	}
	err = os.WriteFile("testExcelExport.xlsx", data.Bytes(), 777)
	if err != nil {
		fmt.Println("TestNewExcel WriteFile err: %v ", err)
		return
	}
	fmt.Println("TestNewExcel success")
}
