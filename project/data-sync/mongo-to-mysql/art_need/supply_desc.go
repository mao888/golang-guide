package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"sort"
)

func RunSupplyDesc() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artneedlogs")

	// 2、从mysql查询数据
	artNeeds := make([]*ArtNeed, 0)
	err := db2.MySQLClientCruiser.Table("art_needs").Find(&artNeeds).Error
	if err != nil {
		fmt.Println("从mysql查询数据 错误：", err)
		return
	}
	//fmt.Println(artNeeds)

	// 3、将mongo数据装入切片
	for _, need := range artNeeds {

		// 根据需求id去mongo中查询日志
		mArtNeedLogs := make([]*MArtNeedLogs, 0)
		err := coll.Find(context.TODO(), bson.M{"artneed_id": need.ID}).All(&mArtNeedLogs)
		if err != nil {
			fmt.Println("Mongo查询错误：", err)
			return
		}
		if len(mArtNeedLogs) == constants.NumberZero {
			continue
		}
		// 根据时间降序
		sort.Slice(mArtNeedLogs, func(i, j int) bool {
			return mArtNeedLogs[i].UpdateTime.Unix() > mArtNeedLogs[j].UpdateTime.Unix() // 降序
		})

		for _, log := range mArtNeedLogs {
			// SupplyDesc 补充说明 为空跳过
			if log.SupplyDesc == constants.EmptyString {
				continue
			}
			// SupplyDesc 补充说明 不为空，则更新当前需求的补充说明
			err = db2.MySQLClientCruiser.Table("art_needs").Where("id = ?", need.ID).
				UpdateColumn("extra_desc", log.SupplyDesc).Error
			if err != nil {
				fmt.Println("更新数据 错误：", err)
				return
			}
		}
	}
}
