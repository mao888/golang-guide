package ad_material

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_material/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
)

// RunAdMaterialSyncSuccess 广告素材 上传同步 返回对照表
func RunAdMaterialSyncSuccess(wg sync.WaitGroup) {
	defer wg.Done()
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	dbu := db2.MongoClient.Database("plat_console")
	coll := db.Collection("assetcenters")
	collUsers := dbu.Collection("platusers")

	// 2、从mongo查询数据
	mAssetCenter := make([]*bean.MAssetCenter, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mAssetCenter)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mAssetCenter)

	// 3、将mongo数据装入切片
	//artAsset := make([]*bean.ArtAsset, 0)
	for _, center := range mAssetCenter {

		// MaterilaType
		var assetType int32
		if center.AssetType == "image" {
			assetType = constants.NumberTwo
		} else if center.AssetType == "video" {
			assetType = constants.NumberThree
		}
		// Creator
		// AuthorID
		var authorID int32

		mPlatUser := make([]*bean.MPlatUser, 0)
		if center.UserId != constants.NumberZero {
			// 根据 source.Author 去mongo查询用户信息
			err := collUsers.Find(context.TODO(), bson.M{"_id": center.UserId}).All(&mPlatUser)
			if err != nil {
				fmt.Println("Mongo/platusers查询错误：", err)
				return
			}
		} else {
			authorID = 1000
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
				authorID = 1000
			} else {
				authorID = user[0].ID
			}
		} else {
			authorID = 1000
		}

		// Type
		t := constants.NumberOne
		if center.ThirdPartyUrl != constants.EmptyString {
			t = constants.NumberTwo
		}
		for _, s := range center.MediaList {

			adMaterialSyncSuccess := &bean.AdMaterialSyncSuccess{
				//ID:           center.Id,
				MaterilaType: assetType,
				MaterialID:   center.Id,
				Name:         center.AssetLongName,
				URL:          center.AssetUrlInfo,
				MaterialMd5:  center.AssetMd5,
				AccountID:    s.AccountId, // AccountId
				Creator:      authorID,
				Type:         int32(t),
				SuccessID:    s.MediaId, // MediaId
				//BatchID:      "",
			}
			// 4、将装有mongo数据的切片入库
			err = db2.MySQLClientCruiser.Table("ad_material_sync_success").Create(adMaterialSyncSuccess).Error
			if err != nil {
				fmt.Println("入mysql/ad_material_sync_success 错误：", err)
			}
		}

	}
}
