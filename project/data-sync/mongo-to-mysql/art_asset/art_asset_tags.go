package art_asset

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/art_asset/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAssetTag() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("tags")

	// 2、从mongo查询数据
	mTags := make([]*bean.MTags, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mTags)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println("mTags:", mTags)

	// 3、将mongo数据装入切片
	//artAssetTag := make([]*bean.ArtAssetTag, 0)
	for _, source := range mTags {
		if len(source.TagList) == constants.NumberZero {
			continue
		}
		for _, s := range source.TagList {
			assetTag := &bean.ArtAssetTag{
				ID:    s.Id,
				Label: s.Name,
				//Code:      0,
				CreatedAt: s.CreateTime.Unix(),
				//Remark:    "",
			}
			//artAssetTag = append(artAssetTag, assetTag)
			// 4、将装有mongo数据的切片入库（单条入库）
			//fmt.Println(assetTag.ID)
			err = db2.MySQLClientCruiser.Table("art_asset_tags").Create(assetTag).Error
			if err != nil {
				fmt.Println("入mysql/art_asset_tags错误：", err)
				return
			}
		}
	}
}
