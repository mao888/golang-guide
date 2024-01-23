package art_asset

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/art_asset/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAssetTagRelation() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("activelibraries")
	collArtSource := db.Collection("artsources")

	// 2、从mongo查询数据
	// 动作库
	mActiveLibrary := make([]*bean.MActiveLibrary, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mActiveLibrary)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println("mTags:", mActiveLibrary)

	// 资产库
	mArtSource := make([]*bean.MArtSource, 0)
	err = collArtSource.Find(context.TODO(), bson.M{}).All(&mArtSource)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println("mTags:", mArtSource)

	// 3、将mongo数据装入切片
	// 动作库
	for _, library := range mActiveLibrary {
		if len(library.TagArr) == constants.NumberZero {
			continue
		}
		for _, i := range library.TagArr {
			artAssetTagRelation := &bean.ArtAssetTagRelation{
				//ID:      0,
				AssetID: library.ID,
				TagID:   i,
			}
			// 4、将装有mongo数据的切片入库（单条入库）
			err = db2.MySQLClientCruiser.Table("art_asset_tag_relations").Create(artAssetTagRelation).Error
			if err != nil {
				fmt.Println("入mysql/art_asset_tag_relations 错误：", err)
				return
			}
		}
	}

	// 资产库
	for _, library := range mArtSource {
		if len(library.TagArr) == constants.NumberZero {
			continue
		}
		for _, i := range library.TagArr {
			artAssetTagRelation := &bean.ArtAssetTagRelation{
				//ID:      0,
				AssetID: library.ID,
				TagID:   i,
			}
			// 4、将装有mongo数据的切片入库（单条入库）
			err = db2.MySQLClientCruiser.Table("art_asset_tag_relations").Create(artAssetTagRelation).Error
			if err != nil {
				fmt.Println("入mysql/art_asset_tag_relations 错误：", err)
				return
			}
		}
	}

}
