package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// MMailGoods From gm-system
type MMailGoods struct {
	ID         int32  `gorm:"column:id;NOT NULL;" json:"id"`
	GameID     int32  `gorm:"column:game_id; " json:"game_id"`            // 游戏id
	GoodName   string `gorm:"column:good_name; "  json:"good_name"`       // 物品名称
	GoodID     string `gorm:"column:good_id; "  json:"good_id"`           // 物品id
	GoodType   string `gorm:"column:good_type; "  json:"good_type"`       // 物品类型
	GoodTypeID string `gorm:"column:good_type_id; "  json:"good_type_id"` // 物品类型ID
	CreatedAt  int64  `gorm:"column:created_at; "  json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at; "  json:"updated_at"`
}

// MailGood mapped from table application_console <mail_goods>
type MailGood struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID    string `gorm:"column:game_id" json:"game_id"`       // game_id 游戏id
	GoodID    string `gorm:"column:good_id" json:"good_id"`       // 物品id
	Name      string `gorm:"column:name" json:"name"`             // 物品名称
	Type      string `gorm:"column:type" json:"type"`             // 物品类型
	TypeID    string `gorm:"column:type_id" json:"type_id"`       // 物品类型ID
	CreatorID int32  `gorm:"column:creator_id" json:"creator_id"` // 创建人
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

func RunMailGoods() {
	// 1、建立Mongo连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("games")

	// 2、从mongo查询数据
	var mGame []MGame
	err := coll.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
	}

	// 3、将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	idMap := map[int32]string{}
	for _, game := range mGame {
		idMap[game.ID] = game.GameID
	}

	// 4、从mysql查询 m_mail_goods
	mMailGoods := make([]*MMailGoods, 0)
	err = db2.MySQLClientGM.Table("m_mail_goods").Find(&mMailGoods).Error
	if err != nil {
		fmt.Println("从mysql查询 m_mail_goods 错误：", err)
	}

	// 5、将 gm-system/m_mail_goods 入库 application_console/mail_goods
	mailGood := make([]*MailGood, 0)
	for _, good := range mMailGoods {
		mail := &MailGood{
			ID:     good.ID,
			GameID: idMap[good.GameID],
			GoodID: good.GoodID,
			Name:   good.GoodName,
			Type:   good.GoodType,
			TypeID: good.GoodTypeID,
			//CreatorID: 0,
			CreatedAt: good.CreatedAt,
			UpdatedAt: good.UpdatedAt,
		}
		mailGood = append(mailGood, mail)
	}
	fmt.Println(mailGood)

	err = db2.MySQLClient.Table("mail_goods").CreateInBatches(mailGood, len(mailGood)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
