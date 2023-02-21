package bean

import "time"

// MActiveLibrary 动作库 From Mongo/activelibraries
type MActiveLibrary struct {
	ID               int32      `json:"_id" bson:"_id"`
	CompanyId        int32      `json:"company_id" bson:"company_id"`
	Name             string     `json:"name" bson:"name"`
	Desc             string     `json:"desc" bson:"desc"`
	Size             string     `json:"size" bson:"size"`
	CategoryId       int32      `json:"category_id" bson:"category_id"`
	TagArr           []int32    `json:"tag_arr" bson:"tag_arr"`                      //标签，上限10
	ThumbUrl         string     `json:"thumb_url" bson:"thumb_url"`                  //预览图
	Url              string     `json:"url" bson:"url"`                              // 视频URL
	CreatorId        int32      `json:"creator_id" bson:"creator_id"`                //上传用户id
	Author           int32      `json:"author" bson:"author"`                        // 作者
	GameId           int32      `json:"game_id" bson:"game_id"`                      // 游戏id
	UeDownloadUrl    string     `json:"ue_download_url" bson:"ue_download_url"`      //UE下载地址
	MayaDownloadUrl  string     `json:"maya_download_url" bson:"maya_download_url"`  //Maya下载地址
	DoneTime         *time.Time `json:"done_time" bson:"done_time"`                  // 完成时间
	RelationArtneeds []string   `json:"relation_artneeds" bson:"relation_artneeds""` // 关联需求
	DeletedTime      *time.Time `json:"deleted_time" bson:"deleted_time"`            // 删除状态
	CreateTime       *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime       *time.Time `json:"update_time" bson:"update_time"`
}

// MArtSource 资产库 From Mongo/artsources
type MArtSource struct {
	ID               int32      `json:"_id" bson:"_id"`
	CompanyId        int32      `json:"company_id" bson:"company_id"`
	Name             string     `json:"name" bson:"name"`
	Desc             string     `json:"desc" bson:"desc"`
	CategoryId       int32      `json:"category_id" bson:"category_id"`
	TagArr           []int32    `json:"tag_arr" bson:"tag_arr"`                      //标签，上限10
	ThumbArr         []int32    `json:"thumb_arr" bson:"thumb_arr"`                  //预览图，上限20
	DownloadCount    int32      `json:"download_count" bson:"download_count"`        // 下载次数
	CreatorId        int32      `json:"creator_id" bson:"creator_id"`                //上传用户id
	Author           int32      `json:"author" bson:"author"`                        // 作者
	GameId           int32      `json:"game_id" bson:"game_id"`                      // 游戏id
	DownloadUrl      string     `json:"download_url" bson:"download_url"`            //UE下载地址
	MayaDownloadUrl  string     `json:"maya_download_url" bson:"maya_download_url"`  //Maya下载地址
	DoneTime         *time.Time `json:"done_time" bson:"done_time"`                  // 完成时间
	RelationArtneeds []string   `json:"relation_artneeds" bson:"relation_artneeds""` // 关联需求
	CreateTime       *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime       *time.Time `json:"update_time" bson:"update_time"`
}

// MCloudUrls From Mongo/cloudurls
type MCloudUrls struct {
	ID           int32      `bson:"_id" json:"_id"`
	CompanyId    int32      `json:"company_id" bson:"company_id"`
	AssetId      int32      `json:"asset_id" bson:"asset_id"`           //关联到第三方资源类型id，如美术资源库id
	CloudType    string     `json:"cloud_type" bson:"cloud_type"`       // 资源存储的CDN类型，默认Akamai
	UseType      string     `json:"use_type" bson:"use_type"`           // 资源用途：art_store_source美术资源库源文件；art_store_thumb美术资源库预览图
	SourceType   string     `json:"source_type" bson:"source_type"`     // 资源类型video、image、zip、file
	Name         string     `json:"name" bson:"name"`                   // 资源名称
	CloudDir     string     `json:"cloud_dir" bson:"cloud_dir"`         // 云地址路径
	Url          string     `json:"url" bson:"url"`                     // 资源地址
	ThumbnailUrl string     `json:"thumbnail_url" bson:"thumbnail_url"` // 资源缩略图地址
	Size         int32      `json:"size" bson:"size"`                   // 资源大小
	Suffix       string     `json:"suffix" bson:"suffix"`               // 资源后缀名
	CreateTime   *time.Time `json:"create_time" bson:"create_time"`     // 创建时间
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

// MCategory From Mongo/categories
type MCategory struct {
	ID         int32      `bson:"_id" json:"_id"`
	CompanyID  int32      `bson:"company_id" json:"company_id"` // 绑定公司id，若为0则为通用型
	ParentId   int32      `json:"parent_id" bson:"parent_id"`   // 父节点id，若为0则为一级节点
	Ctype      string     `json:"ctype" bson:"ctype"`           // art_store美术库；art_tag美术库标签
	Level      int32      `json:"level" bson:"level"`
	Name       string     `json:"name" bson:"name"`
	Enable     bool       `json:"enable" bson:"enable"` // 是否被禁用
	CreateTime *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime *time.Time `bson:"update_time" json:"update_time"`
}

// MTags From Mongo/tags
type MTags struct {
	Id      string `json:"_id" bson:"id"`
	TagList []struct {
		CategoryId int32      `json:"category_id" bson:"category_id"`
		Id         int32      `json:"_id" bson:"_id"`
		Name       string     `json:"name" bson:"name"`
		CompanyId  int32      `json:"company_id" bson:"company_id"`
		Ttype      string     `json:"ttype" bson:"ttype"`
		CreateTime *time.Time `bson:"create_time" json:"create_time"`
		UpdateTime *time.Time `bson:"update_time" json:"update_time"`
	} `json:"tag_list" bson:"tag_list"`
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
