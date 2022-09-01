/**
    @author: edy
    @since: 2022/8/19
    @desc: //TODO
**/
package main

import (
	"encoding/json"
	"fmt"
)

type Opus struct {
	Date  string
	Title string
}
type Actress struct {
	Name       string
	Birthday   string
	BirthPlace string
	Opus       []Opus
}

func main() {
	// JSON嵌套数组JSON
	jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         {
            "date":"2013",
            "title":"《阿娜尔罕》"
         },
         {
            "date":"2014",
            "title":"《逆光之恋》"
         },
         {
            "date":"2015",
            "title":"《克拉恋人》"
         }w
      ]
   }`)
	var actress Actress
	err := json.Unmarshal(jsonData, &actress)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("姓名：%s\n", actress.Name)
	fmt.Printf("生日：%s\n", actress.Birthday)
	fmt.Printf("出生地：%s\n", actress.BirthPlace)
	fmt.Println("作品：")
	for _, val := range actress.Opus {
		fmt.Printf("\t%s - %s\n", val.Date, val.Title)
	}
}
