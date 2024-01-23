package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfCopywriting() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	dbg := db2.MongoClient.Database("plat_console")
	coll := db.Collection("adtexts")
	collGames := dbg.Collection("games")

	// 2、从mongo查询数据
	mAdTexts := make([]*bean.MAdTexts, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mAdTexts)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mAdTexts)

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
	for _, text := range mAdTexts {
		var adConfCopywriting bean.AdConfCopywriting

		// GameID
		var gameID string
		if text.GameId == constants.EmptyString {
			gameID = ""
		} else {
			p, ok := text.GameId.(int32)
			if ok {
				gameID = idMap[p]
			}
		}
		adConfCopywriting.GameID = gameID
		adConfCopywriting.DefaultLanguage = bean.LanguagesMapCodeToShortName[text.DefaultLang]
		adConfCopywriting.CreatedAt = text.CreateTime.Unix()
		adConfCopywriting.UpdatedAt = text.UpdateTime.Unix()
		adConfCopywriting.DefaultLanguageContent = text.DefaultText
		var (
			en   string
			de   string
			fr   string
			ja   string
			ko   string
			es   string
			it   string
			zhtw string
			ar   string
			th   string
			ru   string
			pt   string
			nl   string
			tr   string
			vi   string
			ms   string
			id   string
			tl   string
		)
		for _, s := range text.Translation {
			lang := bean.LanguagesMapCodeToShortName[s.Lang]
			//En
			//var en string
			if lang == "en" {
				en = s.Text
			}
			//De
			//var de string
			if lang == "de" {
				de = s.Text
			}
			//Fr
			//var fr string
			if lang == "fr" {
				fr = s.Text
			}
			//Ja
			//var ja string
			if lang == "ja" {
				ja = s.Text
			}
			//Ko
			//var ko string
			if lang == "ko" {
				ko = s.Text
			}
			//Es
			//var es string
			if lang == "es" {
				es = s.Text
			}
			//It
			//var it string
			if lang == "it" {
				it = s.Text
			}
			//ZhTw
			//var zhtw string
			if lang == "zh-tw" {
				zhtw = s.Text
			}
			//Ar
			//var ar string
			if lang == "ar" {
				ar = s.Text
			}
			//Th
			//var th string
			if lang == "th" {
				th = s.Text
			}
			//Ru
			//var ru string
			if lang == "ru" {
				ru = s.Text
			}
			//Pt
			//var pt string
			if lang == "pt" {
				pt = s.Text
			}
			//Nl
			//var nl string
			if lang == "nl" {
				nl = s.Text
			}
			//Tr
			//var tr string
			if lang == "tr" {
				tr = s.Text
			}
			//Vi
			//var vi string
			if lang == "vi" {
				vi = s.Text
			}
			//Ms
			//var ms string
			if lang == "ms" {
				ms = s.Text
			}
			//Iid
			//var id string
			if lang == "id" {
				id = s.Text
			}
			//Tl
			//var tl string
			if lang == "tl" {
				tl = s.Text
			}
			adConfCopywriting.En = en
			adConfCopywriting.De = de
			adConfCopywriting.Fr = fr
			adConfCopywriting.Ja = ja
			adConfCopywriting.Ko = ko
			adConfCopywriting.Es = es
			adConfCopywriting.It = it
			adConfCopywriting.ZhTw = zhtw
			adConfCopywriting.Ar = ar
			adConfCopywriting.Th = th
			adConfCopywriting.Ru = ru
			adConfCopywriting.Pt = pt
			adConfCopywriting.Nl = nl
			adConfCopywriting.Tr = tr
			adConfCopywriting.Vi = vi
			adConfCopywriting.Ms = ms
			adConfCopywriting.Iid = id
			adConfCopywriting.Tl = tl
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_conf_copywriting").Create(&adConfCopywriting).Error
		if err != nil {
			fmt.Println("入mysql/ad_conf_copywriting 错误：", err)
		}
	}
}
