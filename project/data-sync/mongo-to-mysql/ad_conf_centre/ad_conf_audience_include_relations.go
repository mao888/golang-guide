package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfAudienceIncludeRelation() {
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
	//adConfAudience := make([]*bean.AdConfAudience, 0)
	for _, audience := range mCfgAudienceModel {

		for _, customAudience := range audience.CustomAudiences {
			adConfAudienceIncludeRelation := &bean.AdConfAudienceIncludeRelation{
				//ID:         0,
				AudienceID: audience.ID,
				Name:       customAudience.Name,
				Subtype:    "",
				FbID:       customAudience.ID,
				Type:       constants.NumberOne,
				CreatedAt:  audience.CreateTime.Unix(),
				UpdatedAt:  audience.UpdateTime.Unix(),
			}
			// 4、将装有mongo数据的切片入库
			err = db2.MySQLClientCruiser.Table("ad_conf_audience_include_relations").Create(adConfAudienceIncludeRelation).Error
			if err != nil {
				fmt.Println("入mysql错误：", err)
			}
		}

		for _, customAudience := range audience.ExcludedCustomAudiences {
			adConfAudienceIncludeRelation := &bean.AdConfAudienceIncludeRelation{
				//ID:         0,
				AudienceID: audience.ID,
				Name:       customAudience.Name,
				Subtype:    "",
				FbID:       customAudience.ID,
				Type:       constants.NumberTwo,
				CreatedAt:  audience.CreateTime.Unix(),
				UpdatedAt:  audience.UpdateTime.Unix(),
			}
			// 4、将装有mongo数据的切片入库
			err = db2.MySQLClientCruiser.Table("ad_conf_audience_include_relations").Create(adConfAudienceIncludeRelation).Error
			if err != nil {
				fmt.Println("入mysql错误：", err)
			}
		}
	}
}
