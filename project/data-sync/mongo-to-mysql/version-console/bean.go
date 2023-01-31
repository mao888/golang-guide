package version_console

import "time"

// app : env      1 : 多
// env : version  1 : 多

// MEnvironment From Mongo plat_console <environments>
type MEnvironment struct {
	ID         string     `bson:"_id" json:"_id"`
	CompanyID  int32      `bson:"company_id" json:"company_id"`
	GameID     int32      `bson:"game_id" json:"game_id"`
	AppID      int32      `bson:"app_id" json:"app_id"`
	EnvID      int32      `bson:"env_id" json:"env_id"` //0开发(bu)；1测试；2预发布(bu)；3生产；其他为自定义
	Name       string     `bson:"name" json:"name"`     //环境名称
	CreatorID  int32      `bson:"creator_id" json:"creator_id"`
	Enable     bool       `bson:"enable" json:"enable"`
	DeleteTime *time.Time `bson:"delete_time,omitempty" json:"delete_time,omitempty"`
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time `bson:"update_time" json:"update_time"`
}

// MGameLanguageConf From Mongo plat_console <gamelanguageconfs>
type MGameLanguageConf struct {
	ID         string     `bson:"_id" json:"_id"`
	AppID      int32      `bson:"app_id" json:"app_id"`
	NameEn     string     `bson:"name_en" json:"name_en"`
	NameShort  string     `bson:"name_short" json:"name_short"`
	DefaultLng bool       `bson:"default_lng" json:"default_lng"`
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time `bson:"update_time" json:"update_time"`
}

// MVersion From Mongo plat_console <versions>
type MVersion struct {
	Id              string     `bson:"-" json:"id,omitempty"`
	ID              string     `bson:"_id" json:"_id"`
	CompanyID       int32      `bson:"company_id" json:"company_id"`
	GameID          int32      `bson:"game_id" json:"game_id"`
	AppID           int32      `bson:"app_id" json:"app_id"`
	EnvID           int32      `bson:"env_id" json:"env_id"`                     //0开发；1测试；2预发布；3生产；其他为自定义
	ParentID        string     `bson:"parent_id" json:"parent_id"`               //父版本ID
	VersionName     string     `bson:"version_name" json:"version_name"`         //default: '1.0.0'
	UpdateType      int32      `bson:"update_type" json:"update_type"`           //1强更;2非强更;3热更新
	GrayFlag        bool       `bson:"gray_flag" json:"gray_flag"`               //false非灰度发布；true灰度发布  默认false
	GrayScale       int32      `bson:"gray_scale" json:"gray_scale"`             //灰度发布值 默认0
	Status          int32      `bson:"status" json:"status"`                     //0未发布；1已发布；2已发布并且已推送到生产；3已废弃  默认0
	NoticeFlag      bool       `bson:"notice_flag" json:"notice_flag"`           //true显示新版本提示
	NoticeText      string     `bson:"notice_text" json:"notice_text"`           //新版本提示内容
	NoticeLngText   []LngText  `bson:"notice_lng_text" json:"notice_lng_text"`   //选择多种语言时起作用
	CloseFlag       bool       `bson:"close_flag" json:"close_flag"`             //true允许用户关闭提示
	MultiLngFlag    bool       `bson:"multi_lng_flag" json:"multi_lng_flag"`     //true：多语言；false：单一语言
	RestartFlag     bool       `bson:"restart_flag" json:"restart_flag"`         //true更新后需要退出应用重新进入
	GlobalConf      []KV       `bson:"global_conf" json:"global_conf"`           //全局配置
	DefaultLanguage string     `bson:"default_language" json:"default_language"` //默认语言
	LanguageConf    []LngConf  `bson:"language_conf" json:"language_conf"`       //语言配置
	CreatorID       int32      `bson:"creator_id" json:"creator_id"`
	Enable          bool       `bson:"enable" json:"enable"`
	PublishTime     *time.Time `bson:"publish_time,omitempty" json:"publish_time,omitempty"` //发布时间
	DeleteTime      *time.Time `bson:"delete_time,omitempty" json:"delete_time,omitempty"`
	CreateTime      *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime      *time.Time `bson:"update_time" json:"update_time"`

	//ForceStar bool       `bson:"-" json:"force_star"`
	//UserCount int        `bson:"-" json:"user_count"`
	//Children  []*Version `bson:"-" json:"children"`
}

type LngText struct {
	Lng  string `bson:"lng" json:"lng"`
	Text string `bson:"text" json:"text"`
}

type LngConf struct {
	Language   string `bson:"language" json:"language"`
	ConfList   []KV   `bson:"conf_list" json:"conf_list"`
	DefaultLng bool   `bson:"default_lng" json:"default_lng"`
}

type KV struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"value" json:"value"`
}

type MWhiteList struct {
	ID         string     `bson:"_id" json:"_id"`
	CompanyID  int32      `bson:"company_id" json:"company_id"`
	GameID     int32      `bson:"game_id" json:"game_id"`
	AppID      int32      `bson:"app_id" json:"app_id"`
	DevName    string     `bson:"dev_name" json:"dev_name"`
	DevCode    string     `bson:"dev_code" json:"dev_code"`
	CreatorID  int32      `bson:"creator_id" json:"creator_id"`
	Enable     bool       `bson:"enable" json:"enable"`
	DeleteTime *time.Time `bson:"delete_time,omitempty" json:"delete_time,omitempty"`
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
}
