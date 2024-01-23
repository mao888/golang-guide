package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfCountry() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	dbg := db2.MongoClient.Database("plat_console")
	coll := db.Collection("cfgcountries")
	collGames := dbg.Collection("games")

	// 2、从mongo查询数据
	mCfgCountry := make([]*bean.MCfgCountry, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mCfgCountry)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mCfgCountry)

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
	//adConfAudience := make([]*bean.AdConfAudience, 0)
	for _, position := range mCfgCountry {

		// IncludeCountryCodes
		var includeCountryCodes string
		if len(position.GeoLocations.Countries) != constants.NumberZero {
			for i, country := range position.GeoLocations.Countries {
				includeCountryCodes += country
				if i < len(position.GeoLocations.Countries)-1 {
					includeCountryCodes += ","
				}
			}
		}
		if len(position.GeoLocations.CountryGroups) != constants.NumberZero {
			if len(position.GeoLocations.Countries) != constants.NumberZero {
				includeCountryCodes += ","
			}
			for i, group := range position.GeoLocations.CountryGroups {
				includeCountryCodes += group
				if i < len(position.GeoLocations.CountryGroups)-1 {
					includeCountryCodes += ","
				}
			}
		}

		// ExcludeCountryCodes
		var excludeCountryCodes string
		if len(position.ExcludedGeoLocations.Countries) != constants.NumberZero {
			for i, country := range position.ExcludedGeoLocations.Countries {
				excludeCountryCodes += country
				if i < len(position.ExcludedGeoLocations.Countries)-1 {
					excludeCountryCodes += ","
				}
			}
		}
		if len(position.ExcludedGeoLocations.CountryGroups) != constants.NumberZero {
			if len(position.ExcludedGeoLocations.Countries) != constants.NumberZero {
				excludeCountryCodes += ","
			}
			for i, group := range position.ExcludedGeoLocations.CountryGroups {
				excludeCountryCodes += group
				if i < len(position.ExcludedGeoLocations.CountryGroups)-1 {
					excludeCountryCodes += ","
				}
			}
		}
		// GameID
		var gameID string
		if position.GameId == constants.EmptyString {
			gameID = ""
		} else {
			p, ok := position.GameId.(int32)
			if ok {
				gameID = idMap[p]
			}
		}
		adConfCountry := &bean.AdConfCountry{
			ID:                  position.Id,
			Name:                position.Name,
			IncludeCountryCodes: includeCountryCodes,
			ExcludeCountryCodes: excludeCountryCodes,
			GameID:              gameID,
			CreatedAt:           position.CreateTime.Unix(),
			UpdatedAt:           position.UpdateTime.Unix(),
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_conf_country").Create(adConfCountry).Error
		if err != nil {
			fmt.Println("入mysql/ad_conf_country 错误：", err)
		}
	}
}
