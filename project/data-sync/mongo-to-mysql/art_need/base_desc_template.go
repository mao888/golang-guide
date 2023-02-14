package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// 1000361

func RunBaseDescTemplate() {
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

		baseDescTemplate := &BaseDescTemplate{
			//ID:        0,
			NeedID: need.ID,
			//LogID:     0,
			MainDesc:  need.MainDesc,
			CreatedAt: need.CreateTime.Unix(),
			UpdatedAt: need.UpdateTime.Unix(),
		}
		// 4、将装有mongo数据的切片入库（单条入库）
		err = db2.MySQLClientCruiser.Table("base_desc_template").Create(baseDescTemplate).Error
		if err != nil {
			fmt.Println("入mysql错误：", err)
			return
		}
	}
}
