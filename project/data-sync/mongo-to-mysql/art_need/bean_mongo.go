package art_need

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MArtAttachments struct {
	ID                int32       `json:"_id" bson:"_id"`
	CompanyId         int32       `json:"company_id" bson:"company_id"`                                       //公司id
	GameId            int32       `json:"game_id" bson:"game_id"`                                             //游戏id
	ArtneedId         int32       `json:"artneed_id" bson:"artneed_id"`                                       //主需求id
	AssetLongName     string      `json:"asset_long_name,omitempty" bson:"asset_long_name,omitempty"`         //素材全名
	AssetName         string      `json:"asset_name,omitempty" bson:"asset_name,omitempty"`                   //素材内容标题
	AssetMd5          string      `json:"asset_md5" bson:"asset_md5"`                                         // 素材md5
	AssetUrlInfo      string      `json:"asset_url_info,omitempty" bson:"asset_url_info,omitempty"`           //素材地址
	AssetThumbnailUrl string      `json:"asset_thumbnail_url,omitempty" bson:"asset_thumbnail_url,omitempty"` // 素材缩略图
	Type              int32       `json:"type" bson:"type"`                                                   //0: 普通附件， 1， 终稿附件
	AssetType         string      `json:"asset_type" bson:"asset_type"`                                       // 素材类型。 video/ image/
	AssetSize         string      `json:"asset_size" bson:"asset_size"`                                       //素材尺寸 如 1:1
	AssetLanguage     interface{} `json:"asset_language,omitempty" bson:"asset_language,omitempty"`           // 语种
	AssetDuration     interface{} `json:"asset_duration,omitempty" bson:"asset_duration,omitempty"`           // 视频时长
	AssetWidth        int32       `json:"asset_width,omitempty" bson:"asset_width,omitempty"`                 // 素材宽度
	AssetHeight       int32       `json:"asset_height,omitempty" bson:"asset_height,omitempty"`               // 素材高度
	FileSize          int32       `json:"file_size,omitempty" bson:"file_size,omitempty"`                     // 1024 * 1024 * 1000 = 1M
	CreateTime        *time.Time  `json:"create_time" bson:"create_time"`
	UpdateTime        *time.Time  `json:"update_time" bson:"update_time"`
	DeleteTime        *time.Time  `json:"delete_time,omitempty" bson:"delete_time,omitempty"`

	//ArtsourcesID int `json:"artsources_id,omitempty" bson:"artsources_id,omitempty"` //针对来源为【美术资源库】的资产字段，对应artsources表的_id
}

type MArtNeeds struct {
	ID           int32         `json:"_id" bson:"_id" mapstructure:"_id"`
	CompanyId    int           `json:"company_id" bson:"company_id" mapstructure:"company_id"`                       //公司id
	GameId       int32         `json:"game_id" bson:"game_id" mapstructure:"game_id"`                                //游戏id
	Name         string        `json:"name" bson:"name" mapstructure:"name"`                                         //需求名称
	Title        string        `json:"title" bson:"title" mapstructure:"title"`                                      //素材内容标题
	Type         int32         `json:"type" bson:"type" mapstructure:"type"`                                         // 需求类型  0: 图片,1: 2D视频,2: 3D视频,3: 2D+3D视频',4: 试玩广告,5: conversion图片,
	Tag          string        `json:"tag,omitempty" bson:"tag" mapstructure:"tag"`                                  // 需求标签
	Size         []string      `json:"size,omitempty" bson:"size" mapstructure:"size"`                               // 需求尺寸
	Language     []string      `json:"language,omitempty" bson:"language" mapstructure:"language"`                   // 需求语种
	StartDate    *time.Time    `json:"start_date,omitempty" bson:"start_date" mapstructure:"start_date"`             // 排期开始时间
	EndDate      *time.Time    `json:"end_date,omitempty" bson:"end_date" mapstructure:"end_date"`                   // 排期结束日期
	MainDesc     string        `json:"main_desc,omitempty" bson:"main_desc,omitempty" mapstructure:"main_desc"`      // 需求描述
	Status       int           `json:"status" bson:"status" mapstructure:"status"`                                   // 需求状态，0:待分配.1: 已排期，2: 制作中, 3: 已完成
	DeleteTime   interface{}   `json:"delete_time,omitempty" bson:"delete_time" mapstructure:"delete_time"`          // 需求删除时间
	CreativeUser []string      `json:"creative_user" bson:"creative_user" mapstructure:"creative_user"`              // 创意负责人
	DesignUser   []string      `json:"design_user" bson:"design_user" mapstructure:"design_user"`                    // 设计负责人,当为乙方需求时，设计负责人为乙方公司id
	PbDesignUser []interface{} `json:"pb_design_user,omitempty" bson:"pb_design_user" mapstructure:"pb_design_user"` //乙方设计负责人
	PbRCompany   interface{}   `json:"pb_r_company,omitempty" bson:"pb_r_company" mapstructure:"pb_r_company"`       //可以查看需求的乙方公司
	PartType     int           `json:"part_type" bson:"part_type" mapstructure:"part_type"`                          // 0或不存在,内部需求， 1,外包需求
	Keywords     []string      `json:"keywords,omitempty" bson:"keywords" mapstructure:"keywords"`                   // 关键词
	Priority     int32         `json:"priority,omitempty" mapstructure:"priority"`                                   // 需求优先级 1: 最高。2: 较高. 3: 普通，4：较低
	Complexity   string        `json:"complexity,omitempty" mapstructure:"complexity"`                               // 复杂度 S、A、B、C
	RelatedList  []int         `json:"relatedList,omitempty" mapstructure:"relatedList"`
	IsArtDone    int           `json:"is_art_done" bson:"is_art_done" mapstructure:"is_art_done"`               // 是否完成美术需求  0  没完成， 1 完成
	DoneTime     *time.Time    `json:"done_time,omitempty" bson:"done_time,omitempty" mapstructure:"done_time"` // 完成时间
	PbStatus     int           `json:"pb_status,omitempty" bson:"-" mapstructure:"pb_status"`                   //
	Source       string        `json:"source,omitempty" bson:"source" mapstructure:"source"`                    // 需求来源  artneed_system/entry_artneed
	SysArtId     int           `json:"sys_art_id" bson:"sys_art_id,omitempty" mapstructure:"sys_art_id"`        // 同步的主需求id
	SkipUrl      string        `json:"skip_url" bson:"skip_url" mapstructure:"skip_url"`
	CreateTime   *time.Time    `json:"create_time,omitempty" bson:"create_time" mapstructure:"create_time"`
	UpdateTime   *time.Time    `json:"update_time,omitempty" bson:"update_time" mapstructure:"update_time"`

	//DoneTimeInt   int64 `json:"done_time_int,omitempty" bson:"done_time_int,omitempty" mapstructure:"done_time_int"`
	//CreateTimeInt int64 `json:"create_time_int,omitempty" bson:"create_time_int" mapstructure:"create_time_int"`
	//GameInfo      []Game `json:"game_info" bson:"game_info" mapstructure:"game_info"` // 游戏名称
}

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

type ArtNeedDrafts struct {
	ID             int        `json:"_id" bson:"_id"`
	CompanyId      int        `json:"company_id" bson:"company_id"`                   //公司id
	GameId         int        `json:"game_id" bson:"game_id"`                         //游戏id
	Title          string     `json:"title" bson:"title"`                             //素材内容标题
	Type           int        `json:"type" bson:"type"`                               // 需求类型  0: 图片,1: 2D视频,2: 3D视频,3: 2D+3D视频',4: 试玩广告,5: conversion图片,
	Tag            string     `json:"tag" bson:"tag"`                                 // 需求标签
	Size           []string   `json:"size" bson:"size"`                               // 需求尺寸
	Language       []string   `json:"language" bson:"language"`                       // 需求语种
	MainDesc       string     `json:"main_desc,omitempty" bson:"main_desc,omitempty"` // 需求描述
	CreativeUser   []string   `json:"creative_user" bson:"creative_user"`             // 创意负责人
	AttachmentInfo string     `json:"attachment_info" bson:"attachment_info"`         // 草稿附件信息
	CreateTime     *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime     *time.Time `json:"update_time" bson:"update_time"`
}

type MArtNeedLogs struct {
	ID              int        `json:"_id" bson:"_id"`
	CompanyId       int        `json:"company_id" bson:"company_id"`             //公司id
	GameId          int        `json:"game_id" bson:"game_id"`                   //游戏id
	ArtneedId       int        `json:"artneed_id" bson:"artneed_id"`             //主需求id
	MainDesc        string     `json:"main_desc" bson:"main_desc"`               //需求描述
	SupplyDesc      string     `json:"supply_desc,omitempty" bson:"supply_desc"` // 补充说明
	Remark          string     `json:"remark,omitempty" bson:"remark"`           //排期备注
	Attachment      []int      `json:"attachment" bson:"attachment"`             // 附件
	AttachmentNames []string   `json:"attachment_names" bson:"attachment_names"` // 附件名称列表
	OpUser          int        `json:"op_user" bson:"op_user"`                   //操作人_id
	OpUsername      string     `json:"op_username" bson:"op_username"`           // 操作人姓名
	OpAction        string     `json:"op_action" bson:"op_action"`               // 操作动作
	OpSubject       string     `json:"op_subject" bson:"op_subject"`             // 操作主题
	CreateTime      *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime      *time.Time `json:"update_time" bson:"update_time"`
}

type ArtneedSearch struct {
	ArtNeedId      int    `json:"artneed_id" bson:"artneed_id"`
	SearchKeywords string `json:"search_keywords" `
	GameId         int    `json:"game_id" bson:"game_id"` //游戏id
}

type ArtNeedTags struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	CompanyId int                `json:"company_id" bson:"company_id"` //公司id
	Name      string             `json:"name" bson:"name"`             //名称
}

type MArtNeedTasks struct {
	ID           int32      `json:"_id" bson:"_id"`
	CompanyId    int        `json:"company_id" bson:"company_id"`           //公司id
	GameId       int        `json:"game_id" bson:"game_id"`                 //游戏id
	ArtneedId    int        `json:"artneed_id" bson:"artneed_id"`           //主需求id
	Title        string     `json:"title" bson:"title"`                     //子任务标题
	StartDate    *time.Time `json:"start_date,omitempty" bson:"start_date"` // 排期开始时间
	EndDate      *time.Time `json:"end_date,omitempty" bson:"end_date"`     // 排期结束日期
	TaskStepNo   int        `json:"task_step_no" bson:"task_step_no"`       //子任务编号
	Status       int        `json:"status" bson:"status"`                   // 需求状态，0:待分配.1: 已排期，2: 制作中, 3: 已完成
	DesignUser   []string   `json:"design_user" bson:"design_user"`         // 设计负责人,当为乙方需求时，设计负责人为乙方公司id
	BeginTime    string     `json:"begin_time" bson:"begin_time"`           // 开始时间
	DoneTime     string     `json:"done_time" bson:"done_time"`             // 完成时间
	WorkingHours float64    `json:"working_hours" bson:"working_hours"`     // 工时
	CreateTime   *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime   *time.Time `json:"update_time" bson:"update_time"`
}
