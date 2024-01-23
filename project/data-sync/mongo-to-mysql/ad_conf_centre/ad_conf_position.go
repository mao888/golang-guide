package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfPosition() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgpositions")

	// 2、从mongo查询数据
	mfgPosition := make([]*bean.MfgPosition, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mfgPosition)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mfgPosition)

	// 3、将mongo数据装入切片
	//adConfAudience := make([]*bean.AdConfAudience, 0)
	for _, position := range mfgPosition {

		// PositionType  1：自定义版位 editPosition，2：自动版位(默认选中自定义版位1)autoPosition
		positionType := 1
		if position.PositionType == "autoPosition" {
			positionType = 2
		}
		// Facebook
		var faceBook string
		if len(position.FacebookPositions) != constants.NumberZero {
			for i, dim := range position.FacebookPositions {
				faceBook += dim
				if i < len(position.FacebookPositions)-1 {
					faceBook += ","
				}
			}
		}
		// Instagram
		var instagram string
		if len(position.InstagramPositions) != constants.NumberZero {
			for i, dim := range position.InstagramPositions {
				instagram += dim
				if i < len(position.InstagramPositions)-1 {
					instagram += ","
				}
			}
		}
		// AudienceNetwork
		var audienceNetwork string
		if len(position.AudienceNetworkPositions) != constants.NumberZero {
			for i, dim := range position.AudienceNetworkPositions {
				audienceNetwork += dim
				if i < len(position.AudienceNetworkPositions)-1 {
					audienceNetwork += ","
				}
			}
		}
		// Messenger
		var messenger string
		if len(position.MessengerPositions) != constants.NumberZero {
			for i, dim := range position.MessengerPositions {
				messenger += dim
				if i < len(position.MessengerPositions)-1 {
					messenger += ","
				}
			}
		}
		adConfPosition := &bean.AdConfPosition{
			ID:              position.Id,
			Name:            position.Name,
			PositionType:    int32(positionType),
			Facebook:        faceBook,
			Instagram:       instagram,
			AudienceNetwork: audienceNetwork,
			Messenger:       messenger,
			CreatedAt:       position.CreateTime.Unix(),
			UpdatedAt:       position.UpdateTime.Unix(),
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_conf_position").Create(adConfPosition).Error
		if err != nil {
			fmt.Println("入mysql/ad_conf_position 错误：", err)
		}
	}
}
