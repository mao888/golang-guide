package art_need

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"strings"
)

func RunUrlReplace() {
	var adMaterial []*ArtAttachment
	db2.MySQLClientCruiser.Table("art_attachments").Find(&adMaterial)
	fmt.Println("len(ArtAttachment):", len(adMaterial))

	var newUrl string = "https://ark-oss.bettagames.com"
	var oldUrl string = "https://gia-artneeds.nuclearport.com"

	for i, material := range adMaterial {
		fmt.Println("ad_material:", i)

		materialUrl := strings.Replace(material.URL, oldUrl, newUrl, -1)

		err := db2.MySQLClientCruiser.Table("art_attachments").Where("id = ?", material.ID).
			UpdateColumn("url", materialUrl).Error
		if err != nil {
			fmt.Println("更新数据 错误：", err)
			return
		}
	}
}
