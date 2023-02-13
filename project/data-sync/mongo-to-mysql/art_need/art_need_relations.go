package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtNeedRelation() {
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
		
		if len(need.RelatedList) == 0 {
			continue
		}
		for _, i := range need.RelatedList {
			artNeedRelation := &ArtNeedRelation{
				//ID:             0,
				MainNeedID:     need.ID,
				RelationNeedID: int32(i),
			}
			// 入库 art_need_relations
			err = db2.MySQLClientCruiser.Table("art_need_relations").Create(artNeedRelation).Error
			if err != nil {
				fmt.Println("入库 art_need_relations 错误", err)
				return
			}
		}
	}
}
