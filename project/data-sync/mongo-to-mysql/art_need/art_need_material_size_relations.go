package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func RunArtNeedMaterialSizeRelation() {
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

		if len(need.Size) == constants.NumberZero {
			continue
		}
		for _, s := range need.Size {
			// 根据 size 去mysql/dictionaries中查type为2，若没有则添加，若有，则拿到id
			dictionary := make([]*Dictionary, 0)
			err = db2.MySQLClientCruiser.Table("dictionaries").
				Where("label = ?", s).Find(&dictionary).Error
			if err != nil {
				fmt.Println("根据 size 去mysql/dictionaries中查type为2 错误：", err)
				return
			}
			// 若没有则添加
			if len(dictionary) == constants.NumberZero {
				dic := &Dictionary{
					//ID:        0,
					Label:     s,
					Code:      s,
					Type:      constants.NumberTwo,
					Remark:    "",
					CreatedAt: time.Now().Unix(),
				}
				err = db2.MySQLClientCruiser.Table("dictionaries").Create(dic).Error
				if err != nil {
					fmt.Println("向dictionaries中添加标签 错误", err)
					return
				}
				// 并插入 art_need_material_size_relations
				artNeedMaterialSizeRelation := &ArtNeedMaterialSizeRelation{
					//ID:             0,
					NeedID:         need.ID,
					MaterialSizeID: dic.ID,
				}
				err = db2.MySQLClientCruiser.Table("art_need_material_size_relations").Create(artNeedMaterialSizeRelation).Error
				if err != nil {
					fmt.Println("插入 art_need_material_size_relations 错误", err)
					return
				}
				continue
			}
			// 若有，则拿到id,并插入 art_need_material_size_relations
			artNeedMaterialSizeRelation := &ArtNeedMaterialSizeRelation{
				//ID:             0,
				NeedID:         need.ID,
				MaterialSizeID: dictionary[0].ID,
			}
			err = db2.MySQLClientCruiser.Table("art_need_material_size_relations").Create(artNeedMaterialSizeRelation).Error
			if err != nil {
				fmt.Println("插入 art_need_material_size_relations 错误", err)
				return
			}
		}
	}
}
