package remote_config

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	gutil "github.com/mao888/mao-gutils/json"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// RemoteConfig mapped from table application_console <remote_config>
type RemoteConfig struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID     string `gorm:"column:game_id" json:"game_id"`         // 游戏ID
	Env        int32  `gorm:"column:env" json:"env"`                 // 环境: 1:master 2:test
	AppVersion string `gorm:"column:app_version" json:"app_version"` // 版本号
	GrayLevel  int32  `gorm:"column:gray_level" json:"gray_level"`   // 灰度
	Modules    string `gorm:"column:modules" json:"modules"`         // 功能模块
	Status     string `gorm:"column:status" json:"status"`           // 状态草稿、发布 stash、publish
	IsModified bool   `gorm:"column:is_modified" json:"is_modified"` // 是否有变更
	IsHidden   bool   `gorm:"column:is_hidden" json:"is_hidden"`     // 是否隐藏
	OriginID   int32  `gorm:"column:origin_id" json:"origin_id"`     // 原始id的拷贝
	Order      string `gorm:"column:order" json:"order"`             // 排序
	CreatorID  int32  `gorm:"column:creator_id" json:"creator_id"`   // 创建人ID
	CreatedAt  int64  `gorm:"column:created_at" json:"created_at"`   // 创建时间
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`   // 更新时间
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`   // 是否删除
}

// MRemoteConfig From Mongo app_console <remote_config>
type MRemoteConfig struct {
	ID           string                   `bson:"_id" json:"id"`
	GameID       int32                    `bson:"game_id" json:"game_id"` // 游戏id
	GameIdCustom string                   `bson:"game_id_custom" json:"game_id_custom"`
	AppVersion   string                   `bson:"app_version" json:"app_version"` // 版本号
	Order        string                   `bson:"order" json:"order"`             // 排序
	GrayLevel    int32                    `bson:"gray_level" json:"gray_level"`   // 灰度
	Modules      []map[string]interface{} `bson:"modules" json:"modules"`         // 功能模块
	Status       string                   `bson:"status" json:"status"`           // 状态草稿、发布 stash、publish
	IsModified   bool                     `bson:"is_modified" json:"is_modified"` // 是否有变更
	IsHidden     bool                     `bson:"is_hidden" json:"is_hidden"`     // 是否隐藏
	//OriginID     string `bson:"origin_id" json:"origin_id"`     // 原始id的拷贝
	OperatorID int32 `bson:"operator_id" json:"operator_id"` // 创建人ID
	CreatedAt  int64 `bson:"created_at" json:"created_at"`   // 创建时间
	UpdatedAt  int64 `bson:"updated_at" json:"updated_at"`   // 更新时间
	IsDeleted  bool  `bson:"is_deleted" json:"is_deleted"`   // 是否删除
	Env        int32 `bson:"env" json:"env"`                 // 环境: 1:master 2:test
}

// MPlatUser From Mongo/platusers
type MPlatUser struct {
	ID              int32         `bson:"_id" json:"_id"`
	Name            string        `json:"name" bson:"name"`         //昵称
	Username        string        `json:"username" bson:"username"` //用户姓名
	Password        string        `json:"password" bson:"password"`
	Email           string        `json:"email" bson:"email"`
	Phone           string        `json:"phone" bson:"phone"`
	Avatar          string        `json:"avatar" bson:"avatar"`
	Role            []interface{} `json:"role" bson:"role"`        // 放账号级别角色
	Enable          bool          `json:"enable" bson:"enable"`    // 该用户是否被激活
	UserTag         int           `json:"user_tag"bson:"user_tag"` // 账户类型 [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10], //0无 1管理员大权限
	Token           string        `json:"token" bson:"token"`
	TokenExpireTime *time.Time    `json:"token_expire_time" bson:"token_expire_time"`
	Comments        string        `json:"comments" bson:"comments"` //备注
	CreateTime      *time.Time    `bson:"create_time" json:"create_time"`
	UpdateTime      *time.Time    `bson:"update_time" json:"update_time"`
	LoginTime       *time.Time    `json:"login_time"`                             // 最后登录时间
	MaintainStatus  bool          `json:"maintain_status" bson:"maintain_status"` // 维护状态
	GuiderStep      int           `json:"guider_step" bson:"guider_step"`         // 新手引导
	AccessSystem    []string      `json:"access_system" bson:"access_system"`     // 可访问的系统
	DefaultCompany  int           `json:"default_company" bson:"default_company"` // 当前选中公司
}

// User 员工表 mapped from table user_console <user>
type User struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`           // 主键id
	DingID       string `gorm:"column:ding_id;not null" json:"ding_id"`                      // 钉钉id
	Name         string `gorm:"column:name;not null" json:"name"`                            // 员工名称
	Email        string `gorm:"column:email;not null;default:''" json:"email"`               // 员工邮箱
	Tel          string `gorm:"column:tel;not null" json:"tel"`                              // 员工手机
	Avatar       string `gorm:"column:avatar;not null;default:''" json:"avatar"`             // 员工头像url
	Password     string `gorm:"column:password;not null" json:"password"`                    // 密码
	Region       int32  `gorm:"column:region;not null;default:0" json:"region"`              // 地域  0：其他，1：北京，2：成都，3：海外
	Abbreviation string `gorm:"column:abbreviation;not null;default:''" json:"abbreviation"` // 名字简称
	Status       int32  `gorm:"column:status;not null;default:0" json:"status"`              // 状态 0在职 1离职
	UpdatedAt    int32  `gorm:"column:updated_at;not null" json:"updated_at"`                // 更新时间
	CreatedAt    int32  `gorm:"column:created_at;not null" json:"created_at"`                // 创建时间
	IsDeleted    int32  `gorm:"column:is_deleted;not null;default:0" json:"is_deleted"`      // 是否删除（0:否，1:是）
}

func RunRemoteConfig() {
	// 1、建立连接
	db := db2.MongoClient.Database("app_console")
	coll := db.Collection("remote_config")
	dbu := db2.MongoClient.Database("plat_console")
	collUsers := dbu.Collection("platusers")

	// 2、从mongo查询数据
	mRemoteConfig := make([]*MRemoteConfig, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mRemoteConfig)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mRemoteConfig)

	// 3、将mongo数据装入切片
	remoteConfig := make([]*RemoteConfig, 0)
	for _, config := range mRemoteConfig {
		if config.Status == "stash" {
			continue
		}
		modules, err := gutil.Object2JSONE(&config.Modules)
		if err != nil {
			return
		}
		// AuthorID
		var authorID int32

		mPlatUser := make([]*MPlatUser, 0)
		if config.OperatorID != constants.NumberZero {
			// 根据 source.Author 去mongo查询用户信息
			err := collUsers.Find(context.TODO(), bson.M{"_id": config.OperatorID}).All(&mPlatUser)
			if err != nil {
				fmt.Println("Mongo/platusers查询错误：", err)
				return
			}
		} else {
			authorID = 10000
		}

		if len(mPlatUser) != constants.NumberZero {
			// 根据用户邮箱和昵称查询mysql/user，拿到user_id
			user := make([]*User, 0)

			err = db2.MySQLClientUser.Table("user").
				Where("name = ?", mPlatUser[0].Name).Or("email = ?", mPlatUser[0].Email).
				Find(&user).Error
			if err != nil {
				fmt.Println("mysql/user 查询错误：", err)
			}

			if len(user) == constants.NumberZero {
				authorID = 1000
			} else {
				authorID = user[0].ID
			}
		} else {
			authorID = 1000
		}
		remote := &RemoteConfig{
			//ID:         0,
			GameID:     config.GameIdCustom,
			Env:        config.Env,
			AppVersion: config.AppVersion,
			GrayLevel:  config.GrayLevel,
			Modules:    modules,
			Status:     config.Status,
			IsModified: config.IsModified,
			IsHidden:   config.IsHidden,
			OriginID:   0,
			Order:      config.Order,
			CreatorID:  authorID,
			CreatedAt:  config.CreatedAt,
			UpdatedAt:  config.UpdatedAt,
			IsDeleted:  config.IsDeleted,
		}
		remoteConfig = append(remoteConfig, remote)
	}

	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClient.Table("remote_config").CreateInBatches(remoteConfig, len(remoteConfig)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
