package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// MGame From Mongo
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

// MActivities From gm-system
type MActivities struct {
	ID         int32  `gorm:"column:id;NOT NULL;" json:"id"`
	CompanyID  int64  `gorm:"column:company_id;" json:"company_id"`
	GameID     int32  `gorm:"column:game_id;" json:"game_id"`
	Title      string `gorm:"column:title;unique" json:"title"`
	TargetEnv  int32  `gorm:"column:target_env; DEFAULT:1;" json:"target_env"` // 目标服务器: master debug
	ActStatus  bool   `gorm:"column:act_status;" json:"act_status"`            // 活动状态
	Params     string `gorm:"column:params;" json:"params"`                    // 活动参数
	Remark     string `gorm:"column:remark;" json:"remark"`                    // 备注
	StartAt    int64  `gorm:"column:start_at;" json:"start_at"`                // 开始时间
	EndAt      int64  `gorm:"column:end_at;" json:"end_at"`                    // 结束时间
	OperatorID int32  `gorm:"column:operator_id; " json:"operator_id"`         // 创建人
	CreatedAt  int64  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at;" json:"updated_at"`
	IsDeleted  bool   `gorm:"column:is_deleted;softDelete:flag" json:"is_deleted"`
}

// Activity From application_console
type Activity struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID    string `gorm:"column:game_id" json:"game_id"`       // 游戏id
	Env       int32  `gorm:"column:env;default:1" json:"env"`     // 目标服务器
	Title     string `gorm:"column:title" json:"title"`           // 活动名称
	ActStatus bool   `gorm:"column:act_status" json:"act_status"` // 活动状态
	StartAt   int64  `gorm:"column:start_at" json:"start_at"`     // 开始时间
	EndAt     int64  `gorm:"column:end_at" json:"end_at"`         // 结束时间
	CreatorID int32  `gorm:"column:creator_id" json:"creator_id"` // 创建人
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"` // 创建时间
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	IsDeleted bool   `gorm:"column:is_deleted" json:"is_deleted"` // 是否删除
	Remark    string `gorm:"column:remark" json:"remark"`         // 备注
}

func RunActivity() {
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

	// 4、从mysql查询 m_activities
	mActivities := make([]*MActivities, 0)
	err = db2.MySQLClientGM.Table("m_activities").Find(&mActivities).Error
	if err != nil {
		fmt.Println("从mysql查询 m_activities 错误：", err)
	}

	// 5、将 gm-system/m_activities 入库 application_console/activities
	activity := make([]*Activity, 0)
	for _, mActivity := range mActivities {
		ac := &Activity{
			ID:        mActivity.ID,
			GameID:    idMap[mActivity.GameID],
			Env:       mActivity.TargetEnv,
			Title:     mActivity.Title,
			ActStatus: mActivity.ActStatus,
			StartAt:   mActivity.StartAt,
			EndAt:     mActivity.EndAt,
			CreatorID: mActivity.OperatorID,
			CreatedAt: mActivity.CreatedAt,
			UpdatedAt: mActivity.UpdatedAt,
			IsDeleted: mActivity.IsDeleted,
			Remark:    mActivity.Remark,
		}
		activity = append(activity, ac)
	}

	fmt.Println(activity)

	err = db2.MySQLClient.Table("activities").CreateInBatches(activity, len(activity)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
