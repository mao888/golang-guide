package ad_conf_centre

import (
	"context"
	"fmt"
	"github.com/mao888/go-utils/constants"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfAudience() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgaudiences")

	// 2、从mongo查询数据
	mCfgAudienceModel := make([]*bean.MCfgAudienceModel, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mCfgAudienceModel)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mCfgAudienceModel)

	// 3、将mongo数据装入切片
	adConfAudience := make([]*bean.AdConfAudience, 0)
	for _, audience := range mCfgAudienceModel {
		cfgAudience := &bean.AdConfAudience{
			ID:             audience.ID,
			Name:           audience.Name,
			ExtendInterest: constants.NumberOne,
			AccountID:      audience.AccountId,
			CreatedAt:      audience.CreateTime.Unix(),
			UpdatedAt:      audience.UpdateTime.Unix(),
			Creator:        audience.UserId,
			Remark:         "",
		}
		adConfAudience = append(adConfAudience, cfgAudience)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientCruiser.Table("ad_conf_audience").CreateInBatches(adConfAudience, len(adConfAudience)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
