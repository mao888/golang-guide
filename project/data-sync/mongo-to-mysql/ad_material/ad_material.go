package ad_material

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_material/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// RunAdMaterial 广告素材主表（它的ID会社交关联到广告素材tag，尺寸，语言，负责人等关联表)
func RunAdMaterial(wg sync.WaitGroup) {
	defer wg.Done()

	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	dbg := db2.MongoClient.Database("plat_console")
	coll := db.Collection("assetcenters")
	collGames := dbg.Collection("games")

	// 2、从mongo查询数据
	mAssetCenter := make([]*bean.MAssetCenter, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mAssetCenter)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mAssetCenter)

	var mGame []bean.MGame
	err = collGames.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询MGame错误：", err)
	}

	// 3、将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	idMap := map[int32]string{}
	for _, game := range mGame {
		idMap[game.ID] = game.GameID
	}

	// 3、将mongo数据装入切片
	//artAsset := make([]*bean.ArtAsset, 0)
	for _, center := range mAssetCenter {

		// Type
		var assetType int32
		if center.AssetType == "image" {
			assetType = constants.NumberTwo
		} else if center.AssetType == "video" {
			assetType = constants.NumberThree
		}

		// SizeRationID
		var sizeRationID int32
		dictionary := make([]*bean.Dictionary, 0)
		err = db2.MySQLClientCruiser.Table("dictionaries").
			Where("label = ?", center.AssetSize).Find(&dictionary).Error
		if err != nil {
			fmt.Println("mysql/dictionaries 查询 错误：", err)
			return
		}
		// 若没有则添加
		if len(dictionary) == constants.NumberZero {
			dic := &bean.Dictionary{
				//ID:        0,
				Label:     center.AssetSize,
				Code:      center.AssetSize,
				Type:      constants.NumberTwo,
				Remark:    "",
				CreatedAt: time.Now().Unix(),
			}
			err = db2.MySQLClientCruiser.Table("dictionaries").Create(dic).Error
			if err != nil {
				fmt.Println("向dictionaries中添加标签 错误", err)
				return
			}
			// 将插入后获取到的 dictionary.id 赋值 给 sizeRationID
			sizeRationID = dic.ID
		} else {
			// 若有，则拿到id,赋值 给 sizeRationID
			sizeRationID = dictionary[0].ID
		}
		// Src
		var src int32
		artNeed := make([]*bean.ArtNeed, 0)
		err = db2.MySQLClientCruiser.Table("art_needs").
			Where("id = ?", center.AssetSize).Find(&artNeed).Error
		if err != nil {
			fmt.Println("mysql/art_needs 查询 错误：", err)
			return
		}
		if len(artNeed) == constants.NumberZero {
			src = constants.NumberOne
		} else {
			src = constants.NumberTwo
		}
		// ExtraName
		var extraName string
		extraNames := strings.Split(center.AssetUrl, ".")
		if len(extraNames) == constants.NumberTwo {
			extraName = extraNames[1]
		}

		// TagID
		var tagID int32
		dictionary2 := make([]*bean.Dictionary, 0)
		err = db2.MySQLClientCruiser.Table("dictionaries").
			Where("label = ?", center.Tag).Find(&dictionary2).Error
		if err != nil {
			fmt.Println("mysql/dictionaries 查询 错误：", err)
			return
		}
		// 若没有则添加
		if len(dictionary2) == constants.NumberZero {
			dic := &bean.Dictionary{
				//ID:        0,
				Label:     center.Tag,
				Code:      center.Tag,
				Type:      constants.NumberFour,
				Remark:    "",
				CreatedAt: time.Now().Unix(),
			}
			err = db2.MySQLClientCruiser.Table("dictionaries").Create(dic).Error
			if err != nil {
				fmt.Println("向dictionaries中添加标签 错误", err)
				return
			}
			// 将插入后获取到的 dictionary.id 赋值 给 sizeRationID
			tagID = dic.ID
		} else {
			// 若有，则拿到id,赋值 给 sizeRationID
			tagID = dictionary2[0].ID
		}
		// IsDeleted
		isDeleted := false
		if center.DeleteTime != nil {
			isDeleted = true
		}
		// Duration
		var duration int32
		if center.AssetDuration != constants.EmptyString {
			start := reflect.TypeOf(center.AssetDuration).String()
			if start == "string" {
				durationStr, err := strconv.Atoi(center.AssetDuration.(string))
				if err != nil {
					fmt.Println("字符串转int错误：", err)
					return
				}
				duration = int32(durationStr)
			}
			if start == "int32" {
				duration = center.AssetDuration.(int32)
			}
		}
		adMaterial := &bean.AdMaterial{
			ID:           center.Id,
			Type:         assetType,
			NeedID:       center.ArtneedId,
			Name:         center.AssetLongName,
			Title:        center.AssetName,
			URL:          center.AssetUrlInfo,
			YtURL:        center.ThirdPartyUrl,
			SizeRationID: sizeRationID,
			//Size:         0,
			Md5:      center.AssetMd5,
			Duration: duration,
			//Remark:       "",
			CreatedAt: center.CreateTime.Unix(),
			UpdatedAt: center.UpdateTime.Unix(),
			IsDeleted: isDeleted,
			Src:       src,
			ExtraName: extraName,
			GameID:    idMap[center.GameId],
			TagID:     tagID,
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_material").Create(adMaterial).Error
		if err != nil {
			fmt.Println("入mysql/ad_material错误：", err)
		}
	}
}
