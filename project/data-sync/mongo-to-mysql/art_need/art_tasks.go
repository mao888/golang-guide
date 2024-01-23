package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtTask() {
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
	//artTasks := make([]*ArtTask, 0)
	for _, need := range mArtNeeds {

		// 需求id >= 8000000 的都不要
		if need.ID >= 8000000 {
			continue
		}

		if need.Status != constants.NumberThree {
			continue
		}
		if len(need.DesignUser) == constants.NumberZero {
			continue
		}
		for i, s := range need.DesignUser {

			//PersonID
			// 根据简称查询员工id
			user := make([]*User, 0)
			err = db2.MySQLClientUser.Table("user").
				Where("abbreviation = ?", s).Find(&user).Error
			if err != nil {
				fmt.Println("根据简称查询员工id 错误：", err)
				return
			}
			var personId int32
			if len(user) == 0 {
				personId = 1000
			} else {
				personId = user[0].ID
			}

			artTask := &ArtTask{
				//ID:        0,
				TaskType: int32(i + 1),
				NeedID:   need.ID,
				PersonID: personId,
				Desc:     "",
				Status:   constants.NumberThree,
				//UeURL:     "",
				//MayaURL:   "",
				//Remark:    "",
				//CreatedAt: 0,
				//UpdatedAt: 0,
				//IsDeleted: false,
			}
			//artTasks = append(artTasks, artTask)
			// 入库 art_tasks
			err = db2.MySQLClientCruiser.Table("art_tasks").Create(artTask).Error
			if err != nil {
				fmt.Println("入库 art_tasks 错误", err)
				return
			}
		}
	}
}
