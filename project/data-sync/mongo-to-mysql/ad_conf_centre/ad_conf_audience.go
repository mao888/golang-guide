package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfAudience() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	dbu := db2.MongoClient.Database("plat_console")
	collUsers := dbu.Collection("platusers")
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
		// AuthorID
		var authorID int32

		mPlatUser := make([]*bean.MPlatUser, 0)
		if audience.UserId != constants.NumberZero {
			// 根据 source.Author 去mongo查询用户信息
			err := collUsers.Find(context.TODO(), bson.M{"_id": audience.UserId}).All(&mPlatUser)
			if err != nil {
				fmt.Println("Mongo/platusers查询错误：", err)
				return
			}
		} else {
			authorID = 0
		}

		if len(mPlatUser) != constants.NumberZero {
			// 根据用户邮箱和昵称查询mysql/user，拿到user_id
			user := make([]*bean.User, 0)

			err = db2.MySQLClientUser.Table("user").
				Where("name = ?", mPlatUser[0].Name).Or("email = ?", mPlatUser[0].Email).
				Find(&user).Error
			if err != nil {
				fmt.Println("mysql/user 查询错误：", err)
			}

			if len(user) == constants.NumberZero {
				authorID = 0
			} else {
				authorID = user[0].ID
			}
		} else {
			authorID = 0
		}
		cfgAudience := &bean.AdConfAudience{
			ID:             audience.ID,
			Name:           audience.Name,
			ExtendInterest: constants.NumberOne,
			AccountID:      audience.AccountId,
			CreatedAt:      audience.CreateTime.Unix(),
			UpdatedAt:      audience.UpdateTime.Unix(),
			Creator:        authorID,
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
