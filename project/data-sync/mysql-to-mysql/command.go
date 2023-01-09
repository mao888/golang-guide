package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// MCommand From gm-system
type MCommand struct {
	ID           int32  `gorm:"column:id;NOT NULL;" json:"id"`
	GameID       int32  `gorm:"column:game_id;" json:"game_id"`
	TargetEnv    int32  `gorm:"column:target_env; DEFAULT:1;" json:"target_env"` // 目标服务器: master debug
	PlayerID     string `gorm:"column:player_id;" json:"player_id"`              // 玩家id
	PlayerField  string `gorm:"column:player_field" json:"player_field"`         // 玩家字段
	Command      string `gorm:"column:command;" json:"command"`                  // 通过模板拼接的字符串
	OperatorID   int32  `gorm:"column:operator_id; " json:"operator_id"`         // 创建人
	OperatorName string `gorm:"column:operator_name; " json:"operator_name"`     // 创建人
	SendAt       int64  `gorm:"column:send_at;" json:"send_at"`                  // 发送时间
	CreatedAt    int64  `gorm:"column:created_at;" json:"created_at"`            // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at;" json:"updated_at"`            // 更新日期
	IsDeleted    bool   `gorm:"column:is_deleted;softDelete:flag" json:"is_deleted"`
}

// Command From application_console
type Command struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID       string `gorm:"column:game_id" json:"game_id"`             // 游戏id
	TargetEnv    int32  `gorm:"column:target_env" json:"target_env"`       // 目标服务器
	PlayerID     string `gorm:"column:player_id" json:"player_id"`         // 玩家id
	PlayerField  string `gorm:"column:player_field" json:"player_field"`   // 玩家字段
	Command      string `gorm:"column:command" json:"command"`             // 通过模板拼接的字符串
	OperatorID   int32  `gorm:"column:operator_id" json:"operator_id"`     // 创建人ID
	OperatorName string `gorm:"column:operator_name" json:"operator_name"` // 创建人姓名
	SendAt       int64  `gorm:"column:send_at" json:"send_at"`             // 发送时间
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`       // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at" json:"updated_at"`       // 更新时间
	IsDeleted    bool   `gorm:"column:is_deleted" json:"is_deleted"`       // 是否删除
}

func RunCommand() {
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

	// 4、从mysql查询 m_command
	mCommand := make([]*MCommand, 0)
	err = db2.MySQLClientGM.Table("m_command").Find(&mCommand).Error
	if err != nil {
		fmt.Println("从mysql查询 m_command 错误：", err)
	}

	// 5、将 gm-system/m_command 入库 application_console/command
	command := make([]*Command, 0)
	for _, m := range mCommand {
		co := &Command{
			ID:           m.ID,
			GameID:       idMap[m.GameID],
			TargetEnv:    m.TargetEnv,
			PlayerID:     m.PlayerID,
			PlayerField:  m.PlayerField,
			Command:      m.Command,
			OperatorID:   m.OperatorID,
			OperatorName: m.OperatorName,
			SendAt:       m.SendAt,
			CreatedAt:    m.CreatedAt,
			UpdatedAt:    m.UpdatedAt,
			IsDeleted:    m.IsDeleted,
		}
		command = append(command, co)
	}

	fmt.Println(command)

	err = db2.MySQLClient.Table("command").CreateInBatches(command, len(command)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
