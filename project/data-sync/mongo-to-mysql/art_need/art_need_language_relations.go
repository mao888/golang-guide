package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtNeedLanguageRelation() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artneeds")

	// 2、从mongo查询数据
	mArtNeeds := make([]*MArtNeeds, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtNeeds)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mArtNeeds)

	// 3、将mongo数据装入切片
	for _, need := range mArtNeeds {

		// 需求id >= 8000000 的都不要
		if need.ID >= 8000000 {
			continue
		}

		if len(need.Language) == constants.NumberZero {
			continue
		}
		for _, s := range need.Language {
			// 根据code去mysql/art_languages查询语言
			artLanguage := make([]*ArtLanguage, 0)
			err = db2.MySQLClientCruiser.Table("art_languages").
				Where("code = ?", s).Find(&artLanguage).Error
			if err != nil {
				fmt.Println("根据code去mysql/art_language查询语言 错误：", err)
				return
			}

			artNeedLanguageRelation := &ArtNeedLanguageRelation{
				//ID:         0,
				NeedID:     need.ID,
				LanguageID: artLanguage[0].ID,
			}
			// 入库 art_need_language_relations
			err = db2.MySQLClientCruiser.Table("art_need_language_relations").Create(artNeedLanguageRelation).Error
			if err != nil {
				fmt.Println("入库 art_need_language_relations 错误", err)
				return
			}
		}
	}
}
