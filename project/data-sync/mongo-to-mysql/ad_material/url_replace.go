package ad_material

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_material/bean"
	"strings"
)

func RunUrlReplace() {
	var oldUrl2 string = "https://gia-artneeds.nuclearport.com"

	var adMaterial []*bean.AdMaterial
	db2.MySQLClientCruiser.Table("ad_material").Where("url like ?", oldUrl2+"%").Find(&adMaterial)
	fmt.Println("len(adMaterial):", len(adMaterial))

	var newUrl string = "https://ark-oss.bettagames.com"
	//var oldUrl string = "https://adsres.ftstats.com"

	for i, material := range adMaterial {
		fmt.Println("ad_material:", i)

		materialUrl := strings.Replace(material.URL, oldUrl2, newUrl, -1)

		err := db2.MySQLClientCruiser.Table("ad_material").Where("id = ?", material.ID).
			UpdateColumn("url", materialUrl).Error
		if err != nil {
			fmt.Println("更新数据 错误：", err)
			return
		}
	}
}
