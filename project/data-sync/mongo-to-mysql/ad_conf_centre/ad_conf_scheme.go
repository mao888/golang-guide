package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfScheme() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgframes")
	dbp := db2.MongoClient.Database("plat_console")
	collUsers := dbp.Collection("platusers")

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
	for _, frame := range mCfgFrame {

		// CampaignDimension
		var campaignDimensions string
		if len(frame.CampaignDims) != constants.NumberZero {
			for i, dim := range frame.CampaignDims {
				campaignDimensions += bean.DimensionMap[dim]
				if i < len(frame.CampaignDims)-1 {
					campaignDimensions += ","
				}
			}
		}
		// AdsetDimension
		var adsetDimension string
		if len(frame.AdsetDims) != constants.NumberZero {
			for i, dim := range frame.AdsetDims {
				adsetDimension += bean.DimensionMap[dim]
				if i < len(frame.AdsetDims)-1 {
					adsetDimension += ","
				}
			}
		}
		// BudgetLevel
		budgetLevel := 1
		if frame.IsCbo == false {
			budgetLevel = 2
		}
		// AuthorID
		var authorID int32

		mPlatUser := make([]*bean.MPlatUser, 0)
		if frame.UserId != constants.NumberZero {
			// 根据 source.Author 去mongo查询用户信息
			err := collUsers.Find(context.TODO(), bson.M{"_id": frame.UserId}).All(&mPlatUser)
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
		adConfScheme := &bean.AdConfScheme{
			ID:                 frame.Id,
			Name:               frame.Name,
			CampaignDimension:  campaignDimensions,
			AdsetDimension:     adsetDimension,
			CampaignLimitAdset: frame.CampaignLimit,
			AdsetLimitAd:       frame.AdsetLimit,
			BudgetLevel:        int32(budgetLevel),
			BudgetLimit:        float32(frame.BudgetLimit),
			AuthorID:           authorID,
			CreatedAt:          frame.CreateTime.Unix(),
			UpdatedAt:          frame.UpdateTime.Unix(),
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_conf_scheme").Create(adConfScheme).Error
		if err != nil {
			fmt.Println("入mysql/ad_conf_scheme错误：", err)
		}
	}

}
