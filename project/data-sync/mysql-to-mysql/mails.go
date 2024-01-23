package mysql_to_mysql

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	gutil "github.com/mao888/mao-gutils/json"
	"go.mongodb.org/mongo-driver/bson"
)

// MMails From gm-system
type MMails struct {
	ID           int32  `gorm:"column:id;NOT NULL;" json:"id"`
	CompanyID    int32  `gorm:"column:company_id;" json:"company_id"`
	GameID       int32  `gorm:"column:game_id;" json:"game_id"`
	TargetEnv    int32  `gorm:"column:target_env; DEFAULT:1;" json:"target_env"`   // 目标服务器: master debug
	TargetType   int32  `gorm:"column:target_type; DEFAULT:1;" json:"target_type"` // 发送对象类型 all many single
	RecipientID  string `gorm:"column:recipient_id; " json:"recipient_id"`         // 收件人id
	MailStatus   string `gorm:"column:mail_status; " json:"mail_status"`           // 状态:待审核pending 已驳回reject 通过审核pass 已发送finish
	CarryItems   string `gorm:"column:carry_items; " json:"carry_items"`           // 携带物品
	IsHandEntry  bool   `gorm:"column:is_hand_entry; " json:"is_hand_entry"`       // 是否手动录入物品
	EntryItems   string `gorm:"column:entry_items; " json:"entry_items"`           // 录入的物品
	ActionButton string `gorm:"column:action_button; " json:"action_button"`       // 操作按钮
	RedirectUri  string `gorm:"column:redirect_uri; " json:"redirect_uri"`         // 跳转链接
	SendType     int32  `gorm:"column:send_type; DEFAULT:1;" json:"send_type"`     // 发送类型 立即now 定时timing
	SendAt       int64  `gorm:"column:send_at; " json:"send_at"`                   // 发送时间
	ServerSendAt int64  `gorm:"column:server_send_at; " json:"server_send_at"`     // 服务器发送时间
	ExpiredAt    int64  `gorm:"column:expired_at; " json:"expired_at"`             // 过期时间
	ValidDate    int32  `gorm:"column:valid_date; " json:"valid_date"`             // 有效天数
	OperatorID   int32  `gorm:"column:operator_id; " json:"operator_id"`           // 创建人
	CreatedAt    int64  `gorm:"column:created_at; " json:"created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at; " json:"updated_at"`
	IsDeleted    bool   `gorm:"column:is_deleted;softDelete:flag" json:"is_deleted"`
	// MailLanguages []MMailLanguages `gorm:"foreignKey:MailID"  json:"mail_languages"`
}

// Mail mapped from table application_console <mails>
type Mail struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID       string `gorm:"column:game_id" json:"game_id"`                         // 游戏id
	Env          int32  `gorm:"column:env;default:1" json:"env"`                       // 目标服务器: master debug
	TargetType   int32  `gorm:"column:target_type;default:1" json:"target_type"`       // 发送对象类型 all many single
	RecipientID  string `gorm:"column:recipient_id" json:"recipient_id"`               // 收件人id
	MailStatus   string `gorm:"column:mail_status;default:pending" json:"mail_status"` // 状态:待审核pending 已驳回reject 通过审核pass 已发送finish
	CarryItems   string `gorm:"column:carry_items" json:"carry_items"`                 // 携带物品
	EntryItems   string `gorm:"column:entry_items" json:"entry_items"`                 // 导入文件物品信息
	ActionButton string `gorm:"column:action_button" json:"action_button"`             // 操作按钮
	RedirectURI  string `gorm:"column:redirect_uri" json:"redirect_uri"`               // 跳转链接
	SendType     int32  `gorm:"column:send_type;default:1" json:"send_type"`           // 发送类型 立即now 定时timing
	SendAt       int64  `gorm:"column:send_at" json:"send_at"`                         // 发送时间
	ServerSendAt int64  `gorm:"column:server_send_at" json:"server_send_at"`           // 服务器发送时间
	ValidDate    int32  `gorm:"column:valid_date" json:"valid_date"`                   // 有效天数
	ExpiredAt    int64  `gorm:"column:expired_at" json:"expired_at"`                   // 过期时间
	CreatorID    int32  `gorm:"column:creator_id" json:"creator_id"`                   // 创建人
	CreatedAt    int64  `gorm:"column:created_at" json:"created_at"`                   // 创建时间
	UpdatedAt    int64  `gorm:"column:updated_at" json:"updated_at"`                   // 更新时间
	IsDeleted    bool   `gorm:"column:is_deleted" json:"is_deleted"`                   // 是否删除
}

func RunMails() {
	// 1、建立Mongo连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("games")

	// 2、从mongo查询数据
	var mGame []MGame
	err := coll.Find(context.TODO(), bson.M{}).All(&mGame)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}

	// 3、将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map
	idMap := map[int32]string{}
	for _, game := range mGame {
		idMap[game.ID] = game.GameID
	}

	// 4、从mysql查询 m_mails
	mMails := make([]*MMails, 0)
	err = db2.MySQLClientGM.Table("m_mails").Find(&mMails).Error
	if err != nil {
		fmt.Println("从mysql查询 m_mails 错误：", err)
		return
	}
	fmt.Println("len(mMails):=", len(mMails))
	fmt.Println("mMails:=", mMails)

	// 5、将 gm-system/m_mails 入库 application_console/mails
	mails := make([]*Mail, 0)
	for _, mail := range mMails {
		entryItems, err := gutil.Object2JSONE(mail.EntryItems)
		if err != nil {
			fmt.Println(err)
			return
		}
		m := &Mail{
			ID:           mail.ID,
			GameID:       idMap[mail.GameID],
			Env:          mail.TargetEnv,
			TargetType:   mail.TargetType,
			RecipientID:  mail.RecipientID,
			MailStatus:   mail.MailStatus,
			CarryItems:   mail.CarryItems,
			EntryItems:   entryItems,
			ActionButton: mail.ActionButton,
			RedirectURI:  mail.RedirectUri,
			SendType:     mail.SendType,
			SendAt:       mail.SendAt,
			ServerSendAt: mail.ServerSendAt,
			ValidDate:    mail.ValidDate,
			ExpiredAt:    mail.ExpiredAt,
			CreatorID:    mail.OperatorID,
			CreatedAt:    mail.CreatedAt,
			UpdatedAt:    mail.UpdatedAt,
			IsDeleted:    false,
		}
		mails = append(mails, m)
	}
	fmt.Println(mails)
	err = db2.MySQLClient.Table("mails").CreateInBatches(mails, len(mails)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
		return
	}
}
