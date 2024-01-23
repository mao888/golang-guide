package art_asset

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/art_asset/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
	"strings"
)

func RunArtNeedAssetRelation() {
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
	//fmt.Println("mTags:", mActiveLibrary)

	// 资产库
	mArtSource := make([]*bean.MArtSource, 0)
	err = collArtSource.Find(context.TODO(), bson.M{}).All(&mArtSource)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println("mTags:", mArtSource)

	// 3、将mongo数据装入切片
	for _, library := range mActiveLibrary {
		if len(library.RelationArtneeds) != constants.NumberZero {
			for _, artneed := range library.RelationArtneeds {
				str := strings.Split(artneed, "-")
				fmt.Println("str", str)
				needID, err := strconv.ParseInt(str[0], 10, 32)
				if err != nil {
					fmt.Println("字符串转int32错误:", err)
					return
				}
				artNeedAssetRelation := &bean.ArtNeedAssetRelation{
					//ID:      0,
					NeedID:  int32(needID),
					AssetID: library.ID,
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("art_need_asset_relations").Create(artNeedAssetRelation).Error
				if err != nil {
					fmt.Println("入mysql/art_need_asset_relations 错误：", err)
				}
			}
		}
	}

	for _, library := range mArtSource {
		if len(library.RelationArtneeds) != constants.NumberZero {
			for _, artneed := range library.RelationArtneeds {
				str := strings.Split(artneed, "-")
				if len(str) == constants.NumberZero {
					continue
				}
				fmt.Println("str", str)
				needID, err := strconv.ParseInt(str[0], 10, 32)
				if err != nil {
					fmt.Println("字符串转int32错误:", err)
					return
				}
				artNeedAssetRelation := &bean.ArtNeedAssetRelation{
					//ID:      0,
					NeedID:  int32(needID),
					AssetID: library.ID,
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("art_need_asset_relations").Create(artNeedAssetRelation).Error
				if err != nil {
					fmt.Println("入mysql/art_need_asset_relations 错误：", err)
				}
			}
		}
	}
}
