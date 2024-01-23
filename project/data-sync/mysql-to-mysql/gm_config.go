package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
)

// MCommandProjectTeam From gm-system
type MCommandProjectTeam struct {
	ID        int32  `gorm:"column:id;NOT NULL;" json:"id"` // ID 为 GameID TargetEnv 拼接
	GameID    int32  `gorm:"column:game_id;" json:"game_id"`
	TargetEnv int32  `gorm:"column:target_env; DEFAULT:1;" json:"target_env"` // 目标服务器: master debug
	SendUrl   string `gorm:"column:send_url;" json:"send_url"`                // 发送Http地址
	//CommandOption []MCommandOption `gorm:"foreignKey:ProjectTeamID;"  json:"command_options"`
	OperatorID   int32  `gorm:"column:operator_id; " json:"operator_id"`     // 创建人
	OperatorName string `gorm:"column:operator_name; " json:"operator_name"` // 创建名称
	CreatedAt    int64  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at;" json:"updated_at"`
	IsDeleted    bool   `gorm:"column:is_deleted;softDelete:flag" json:"is_deleted"`
}

// MProjectTeamPlayer From gm-system
type MProjectTeamPlayer struct {
	ID           int32  `gorm:"column:id;NOT NULL;" json:"id"` // ID 为 GameID TargetEnv 拼接
	GameID       int32  `gorm:"column:game_id;" json:"game_id"`
	TargetEnv    int32  `gorm:"column:target_env; DEFAULT:1;" json:"target_env"` // 目标服务器: master debug
	SendUrl      string `gorm:"column:send_url;" json:"send_url"`                // 请求url
	OperatorID   int32  `gorm:"column:operator_id; " json:"operator_id"`         // 创建人
	OperatorName string `gorm:"column:operator_name; " json:"operator_name"`     // 创建名称
	CreatedAt    int64  `gorm:"column:created_at;" json:"created_at"`            // 发送时间
	UpdatedAt    int64  `gorm:"column:updated_at;" json:"updated_at"`            // 创建日期
	IsDeleted    bool   `gorm:"column:is_deleted;softDelete:flag" json:"is_deleted"`
}

// GmConfig From application_console
type GmConfig struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID    string `gorm:"column:game_id" json:"game_id"`              // 游戏ID
	Env       int32  `gorm:"column:env;not null;default:1" json:"env"`   // 1 正式环境 2 测试环境
	Type      int32  `gorm:"column:type;not null;default:1" json:"type"` // 1 邮件管理 2 活动管理 3 玩家查询 4 GM命令
	URL       string `gorm:"column:url" json:"url"`                      // URL
	CreatorID int32  `gorm:"column:creator_id" json:"creator_id"`        // 创建人
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`        // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"`        // 更新时间
	IsDeleted bool   `gorm:"column:is_deleted" json:"is_deleted"`        // 是否删除
}

func RunGmConfig() {
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

	// 4、从mysql查询 m_command_project_team (4、GM命令)
	mCommandProjectTeam := make([]*MCommandProjectTeam, 0)
	err = db2.MySQLClientGM.Table("m_command_project_team").Find(&mCommandProjectTeam).Error
	if err != nil {
		fmt.Println("从mysql查询 m_command_project_team 错误：", err)
	}

	// 5、从mysql查询 m_project_team_player (3、玩家查询)
	mProjectTeamPlayer := make([]MProjectTeamPlayer, 0)
	err = db2.MySQLClientGM.Table("m_project_team_player").Find(&mProjectTeamPlayer).Error
	if err != nil {
		fmt.Println("从mysql查询 m_project_team_player 错误：", err)
	}

	// 6、将 gm-system/m_command_project_team 入库 application_console/gm_config
	gmConfig4 := make([]*GmConfig, 0)
	for _, team := range mCommandProjectTeam {
		gm4 := &GmConfig{
			//ID:        team.ID,
			GameID:    idMap[team.GameID],
			Env:       team.TargetEnv,
			Type:      constants.NumberFour,
			URL:       team.SendUrl,
			CreatorID: team.OperatorID,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
			IsDeleted: team.IsDeleted,
		}
		gmConfig4 = append(gmConfig4, gm4)
	}
	fmt.Println(gmConfig4)
	err = db2.MySQLClient.Table("gm_config").CreateInBatches(gmConfig4, len(gmConfig4)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}

	// 7、gm-system/m_project_team_player 入库 application_console/gm_config
	gmConfig3 := make([]*GmConfig, 0)
	for _, player := range mProjectTeamPlayer {
		gm5 := &GmConfig{
			//ID:        player.ID,
			GameID:    idMap[player.GameID],
			Env:       player.TargetEnv,
			Type:      constants.NumberThree,
			URL:       player.SendUrl,
			CreatorID: player.OperatorID,
			CreatedAt: player.CreatedAt,
			UpdatedAt: player.UpdatedAt,
			IsDeleted: player.IsDeleted,
		}
		gmConfig3 = append(gmConfig3, gm5)
	}
	fmt.Println(gmConfig3)
	err = db2.MySQLClient.Table("gm_config").CreateInBatches(gmConfig3, len(gmConfig3)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
