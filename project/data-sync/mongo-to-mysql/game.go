package mongo_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type MGame struct {
	ID         int32      `bson:"_id" json:"_id"`
	GameName   string     `bson:"game_name" json:"game_name"`
	CompanyID  int32      `bson:"company_id" json:"company_id"`
	GameID     string     `bson:"game_id" json:"game_id"`
	CreatorID  int32      `bson:"creator_id" json:"creator_id"`
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time `bson:"update_time" json:"update_time"`

	Enable     bool `bson:"enable" json:"enable"`
	IsArchived bool `bson:"is_archived"`

	SdkToken    string `json:"sdk_token" bson:"sdk_token"`
	ServerToken string `json:"server_token" bson:"server_token"`
}

type Game struct {
	ID              string `gorm:"column:id;not null" json:"id"`                               // 游戏ID
	Name            string `gorm:"column:name;not null" json:"name"`                           // 游戏名称
	Logo            string `gorm:"column:logo" json:"logo"`                                    // Logo
	ProjectTeamID   int32  `gorm:"column:project_team_id" json:"project_team_id"`              // 项目组ID
	ProjectTeamName string `gorm:"column:project_team_name;not null" json:"project_team_name"` // 项目组名称
	Status          int32  `gorm:"column:status;not null" json:"status"`                       // 游戏状态 1 运营中 2 调试中 3 未上线 4 已归档
	CreatorID       int32  `gorm:"column:creator_id;not null" json:"creator_id"`               // 创建人ID
	CreatedAt       int64  `gorm:"column:created_at" json:"created_at"`                        // 创建时间戳（秒）
	UpdatedAt       int64  `gorm:"column:updated_at" json:"updated_at"`                        // 更新时间戳（秒）
	IsDeleted       bool   `gorm:"column:is_deleted" json:"is_deleted"`                        // 是否删除（0：未删除，1：已删除）
}

func RunGame() {
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("games")

	mGame := make([]*MGame, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		return
	}

	fmt.Println(mGame)

	game := make([]*Game, 0)

	for _, g := range mGame {
		ga := &Game{
			ID:   g.GameID,
			Name: g.GameName,
			//Logo:            "",
			//ProjectTeamID:   0,
			//ProjectTeamName: "",
			Status:    1,
			CreatorID: g.CreatorID,
			CreatedAt: g.CreateTime.Unix(),
			UpdatedAt: g.UpdateTime.Unix(),
			IsDeleted: true,
		}
		if g.IsArchived {
			ga.Status = 5 // 已归档
		}
		if g.Enable {
			ga.IsDeleted = false
		}
		game = append(game, ga)
	}

	err = db2.MySQLClient.Table("game").CreateInBatches(game, len(game)).Error
	fmt.Println(err)

}
