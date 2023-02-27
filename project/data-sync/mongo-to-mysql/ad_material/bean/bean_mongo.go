package bean

import "time"

type MAssetCenter struct {
	Id             int32       `json:"_id" bson:"_id"`
	AssetType      string      `json:"asset_type" bson:"asset_type"`           // 素材类型 ['image', 'video']
	AssetMd5       string      `json:"asset_md5" bson:"asset_md5"`             // md5
	AssetName      string      `json:"asset_name" bson:"asset_name"`           // 素材名
	AssetUrl       string      `json:"asset_url" bson:"asset_url"`             // 素材下载地址
	AssetThumbnail string      `json:"asset_thumbnail" bson:"asset_thumbnail"` // 素材缩略图
	AssetWidth     int32       `json:"asset_width" bson:"asset_width"`         // 素材宽
	AssetHeight    int32       `json:"asset_height" bson:"asset_height"`       // 素材高
	AssetDuration  interface{} `json:"asset_duration" bson:"asset_duration"`   // 素材时长
	AssetLanguage  int32       `json:"asset_language" bson:"asset_language"`
	Tag            string      `json:"tag" bson:"tag"`               // 标签
	AssetSize      string      `json:"asset_size" bson:"asset_size"` // 尺寸
	Fbid           string      `json:"fbid" bson:"fbid"`
	MediaList      []struct {
		AccountId  string `json:"account_id" bson:"account_id"`
		MediaId    string `json:"media_id" bson:"media_id"`
		CreativeId string `json:"creative_id" bson:"creative_id"`
	} `json:"media_list" bson:"media_list"` //该素材在每个账户中的对应ID [{account_id:"123", media_id:"678",creative_id:""}]
	CreateUser        string     `json:"create_user" bson:"create_user"`
	UserId            int32      `json:"user_id" bson:"user_id"`
	CreativeUser      string     `json:"creative_user" bson:"creative_user"`             // 创意人员id
	DesignUser        string     `json:"design_user" bson:"design_user"`                 // 设计人员id
	AssetUrlInfo      string     `json:"asset_url_info" bson:"asset_url_info"`           // 包含完整域名的url
	AssetThumbnailUrl string     `json:"asset_thumbnail_url" bson:"asset_thumbnail_url"` // 包含完整域名的url
	CreateTime        *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime        *time.Time `json:"update_time" bson:"update_time"`
	DeleteTime        *time.Time `json:"delete_time" bson:"delete_time"`         //假删除对应字段
	ThirdPartyUrl     string     `json:"third_party_url" bson:"third_party_url"` // youtube素材地址
	//V                 int    `json:"__v"`
	CompanyId     int32  `json:"company_id" bson:"company_id"`
	GameId        int32  `json:"game_id" bson:"game_id"`
	ArtneedId     string `json:"artneed_id" bson:"artneed_id"`
	AssetLongName string `json:"asset_long_name" bson:"asset_long_name"`
}

// MGame From Mongo/games
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
