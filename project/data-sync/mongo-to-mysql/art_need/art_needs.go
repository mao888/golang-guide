package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func RunArtNeeds() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artneeds")
	collGames := db.Collection("games")

	// 2、从mongo查询数据
	mArtNeeds := make([]*MArtNeeds, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtNeeds)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mArtNeeds)

	// Game
	var mGame []MGame
	err = collGames.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
	}
	// 将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	gameIdMap := map[int32]string{}
	for _, game := range mGame {
		gameIdMap[game.ID] = game.GameID
	}
	gameNameMap := map[int32]string{}
	for _, game := range mGame {
		gameNameMap[game.ID] = game.GameName
	}

	// 3、将mongo数据装入切片
	artNeeds := make([]*ArtNeed, 0)
	for _, need := range mArtNeeds {

		// 需求id >= 8000000 的都不要
		if need.ID >= 8000000 {
			continue
		}

		// Status
		var status int32
		if need.Status == 0 { // 0:待分配
			status = 1
		} else if need.Status == 1 { // 1: 已排期
			status = 2
		} else if need.Status == 2 { // 2: 制作中
			status = 2
		} else if need.Status == 3 { // 3: 已完成
			status = 4
		}
		// BaseTag
		var baseTag int32
		tag := strings.Split(need.Tag, "+")
		if tag[0] == "原始" {
			baseTag = 1
		} else if tag[0] == "非原始" {
			baseTag = 2
		}
		// IsDeleted
		isDeleted := false
		if need.DeleteTime != nil {
			isDeleted = true
		}

		// DoneAt
		var doneAt int64
		if need.DoneTime == nil {
			doneAt = 0
		} else {
			doneAt = need.DoneTime.Unix()
		}

		artNeed := &ArtNeed{
			ID:           need.ID,
			GameID:       gameIdMap[need.GameId],
			Title:        need.Title,
			Name:         need.Name,
			Type:         need.Type + 1,
			Status:       status,
			Priority:     need.Priority,
			BaseTag:      baseTag, //
			DescTemplate: 1,
			IsUseCruiser: true,
			IsSchedule:   false,
			DoneAt:       doneAt,
			ExtraDesc:    "", // 暂为空
			IsDeleted:    isDeleted,
			CreatedAt:    need.CreateTime.Unix(),
			UpdatedAt:    need.UpdateTime.Unix(),
			GameName:     gameNameMap[need.GameId],
			AssetRemark:  "",
		}
		artNeeds = append(artNeeds, artNeed)

		// 4、将装有mongo数据的切片入库（单条入库）
		err = db2.MySQLClientCruiser.Table("art_needs").Create(artNeed).Error
		if err != nil {
			fmt.Println("入mysql错误：", err)
			return
		}
	}
	// 4、将装有mongo数据的切片入库
	// 注：数据量较大时无法批量入库，报错：Prepared statement contains too many placeholders（占位符过多~）
	// 18（字段数）* 11645（条数） = 209610 > 65535（mysql对占位符的最高限制）
	//err = db2.MySQLClientCruiser.Table("art_needs").CreateInBatches(artNeeds, len(artNeeds)).Error
	//if err != nil {
	//	fmt.Println("入mysql错误：", err)
	//	return
	//}
}
