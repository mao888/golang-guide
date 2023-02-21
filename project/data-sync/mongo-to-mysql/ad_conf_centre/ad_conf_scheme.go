package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfScheme() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgframes")

	// 2、从mongo查询数据
	mCfgFrame := make([]*bean.MCfgFrame, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mCfgFrame)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mCfgFrame)

	// 3、将mongo数据装入切片
	//adConfAudience := make([]*bean.AdConfAudience, 0)
	//for _, frame := range mCfgFrame {
	//
	//	adConfScheme := &bean.AdConfScheme{
	//		ID:                 frame.Id,
	//		Name:               frame.Name,
	//		CampaignDimension:  "",
	//		AdsetDimension:     "",
	//		CampaignLimitAdset: 0,
	//		AdsetLimitAd:       0,
	//		BudgetLevel:        0,
	//		BudgetLimit:        frame.BudgetLimit,
	//		AuthorID:           frame.UserId,
	//		CreatedAt:          frame.CreateTime.Unix(),
	//		UpdatedAt:          frame.UpdateTime.Unix(),
	//	}
	//}

}
