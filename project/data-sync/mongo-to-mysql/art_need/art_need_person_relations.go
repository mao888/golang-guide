package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtNeedPersonRelation() {
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

		if len(need.CreativeUser) == constants.NumberZero {
			continue
		}
		for i, s := range need.CreativeUser {
			// 根据简称查询员工id
			user := make([]*User, 0)
			err = db2.MySQLClientUser.Table("user").
				Where("abbreviation = ?", s).Find(&user).Error
			if err != nil {
				fmt.Println("根据简称查询员工id 错误：", err)
				return
			}
			// PersonID
			var personId int32
			if len(user) == 0 {
				personId = 1000
			} else {
				personId = user[0].ID
			}
			artNeedPersonRelation := &ArtNeedPersonRelation{
				//ID:       0,
				NeedID:   need.ID,
				PersonID: personId,
				Weight:   int32(i),
			}
			// 入库 art_need_person_relations
			err = db2.MySQLClientCruiser.Table("art_need_person_relations").Create(artNeedPersonRelation).Error
			if err != nil {
				fmt.Println("入库 art_need_person_relations 错误", err)
				return
			}
		}
	}
}
