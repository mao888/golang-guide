package bean

// AdDeliveryLog 投放日志 mapped from table cruiser_console <ad_delivery_log>
type AdDeliveryLog struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID      string `gorm:"column:game_id;not null" json:"game_id"`
	AppID       int32  `gorm:"column:app_id;not null" json:"app_id"`
	AccountID   string `gorm:"column:account_id;not null" json:"account_id"`
	Status      int32  `gorm:"column:status;not null" json:"status"` // 0 未开始 1 进行中 2 已完成
	SpendTimeAt int64  `gorm:"column:spend_time_at;not null" json:"spend_time_at"`
	Details     string `gorm:"column:details;not null" json:"details"`
	CreatorID   int32  `gorm:"column:creator_id;not null" json:"creator_id"`
	CreatedAt   int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	IsDeleted   int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// AdFbFinish FB 创建广告系列，广告组，广告成功的表 mapped from table cruiser_console <ad_fb_finish>
type AdFbFinish struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID   int32  `gorm:"column:parent_id;not null" json:"parent_id"`   // 父级ID
	Type       int32  `gorm:"column:type;not null;default:1" json:"type"`   // 广告类型， 1: campaign,  2: adset,3: ad
	Name       string `gorm:"column:name;not null" json:"name"`             // 对应的名字 FB返回
	AccountID  string `gorm:"column:account_id;not null" json:"account_id"` // 所属账户
	HTTPCode   int32  `gorm:"column:http_code;not null" json:"http_code"`
	Data       string `gorm:"column:data;not null" json:"data"`                 // 该广告组，系列，广告 的详细数据 FB返回的
	FbID       string `gorm:"column:fb_id;not null" json:"fb_id"`               // 创建成功FB返回的对应ID
	AdConfigID int32  `gorm:"column:ad_config_id;not null" json:"ad_config_id"` // 该广告组，广告系列，广告 通过那个 配置创建的该广告
	CreatedAt  int64  `gorm:"column:created_at;not null" json:"created_at"`     // 创建日期
	UpdatedAt  int64  `gorm:"column:updated_at;not null" json:"updated_at"`     // 更新日期
	IsDeleted  bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`     // 1: deleted, 0: normal
}

// AdPreview 预览表 mapped from table cruiser_console <ad_preview>
type AdPreview struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID        string `gorm:"column:game_id;not null" json:"game_id"`
	AppID         int32  `gorm:"column:app_id;not null" json:"app_id"`
	AccountID     string `gorm:"column:account_id;not null;default:0" json:"account_id"`
	FbAppID       int64  `gorm:"column:fb_app_id;not null" json:"fb_app_id"`
	FbHomePageID  string `gorm:"column:fb_home_page_id;not null" json:"fb_home_page_id"`
	InstagramID   string `gorm:"column:instagram_id;not null" json:"instagram_id"`
	BudgetLevel   int32  `gorm:"column:budget_level" json:"budget_level"`      // 默认为CBO广告系列预算 1:CBO广告系列预算 2:ABO广告组预算
	MarketURL     string `gorm:"column:market_url;not null" json:"market_url"` // 商店渠道链接
	UserOs        string `gorm:"column:user_os" json:"user_os"`                // 操作系统
	IncludeDevice string `gorm:"column:include_device" json:"include_device"`  // 包含设备
	ExcludeDevice string `gorm:"column:exclude_device" json:"exclude_device"`  // 排除设备
	IsWifi        int32  `gorm:"column:is_wifi;not null" json:"is_wifi"`       // 是否使用wifi 1 不使用 2 使用
	CreatedAt     int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	IsDeleted     int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// AdPreviewAd 预览的广告 mapped from table cruiser_console <ad_preview_ad>
type AdPreviewAd struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterialID int32  `gorm:"column:material_id" json:"material_id"` // 素材ID
	CampaignID int32  `gorm:"column:campaign_id;not null" json:"campaign_id"`
	AdSetID    int32  `gorm:"column:ad_set_id;not null" json:"ad_set_id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	AdTitle    string `gorm:"column:ad_title;not null" json:"ad_title"`
	AdText     string `gorm:"column:ad_text;not null" json:"ad_text"`
	LanguageID string `gorm:"column:language_id;not null" json:"language_id"`
	IsFinish   int32  `gorm:"column:is_finish;not null" json:"is_finish"`
	AdType     int32  `gorm:"column:ad_type;not null" json:"ad_type"` // 附件文件类型， 1: file,  2: image,3: video
	CreatedAt  int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
}

// AdPreviewAdSet 预览的广告组 mapped from table cruiser_console <ad_preview_ad_set>
type AdPreviewAdSet struct {
	ID            int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CampaignID    int32   `gorm:"column:campaign_id" json:"campaign_id"`
	Name          string  `gorm:"column:name" json:"name"`
	Bugdet        int32   `gorm:"column:bugdet" json:"bugdet"`                 // 预算金额
	Cost          float32 `gorm:"column:cost;not null" json:"cost"`            // 竞价策略控制
	TargetingData string  `gorm:"column:targeting_data" json:"targeting_data"` // Fb的定位控制
	IsFinish      int32   `gorm:"column:is_finish" json:"is_finish"`           // 是否上传完成
	CreatedAt     int64   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     int64   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// AdPreviewCampaign 预览的广告系列 mapped from table cruiser_console <ad_preview_campaign>
type AdPreviewCampaign struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PreviewID int32  `gorm:"column:preview_id;not null" json:"preview_id"`
	SchemeID  int32  `gorm:"column:scheme_id" json:"scheme_id"` // 结构方案ID
	Name      string `gorm:"column:name" json:"name"`
	Budget    int32  `gorm:"column:budget" json:"budget"`
	CreatedAt int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// AdPreviewConfig FB 创建广告系列，广告组，广告的配置选项 表 主要用来创建配置广告，回显使用该表，以便重复利用 mapped from table cruiser_console <ad_preview_config>
type AdPreviewConfig struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID         int32  `gorm:"column:app_id;not null" json:"app_id"` // 该配置的appid
	AdSetID       int32  `gorm:"column:ad_set_id;not null" json:"ad_set_id"`
	AccountID     string `gorm:"column:account_id;not null" json:"account_id"`                 // 所属账户
	SchemeID      int32  `gorm:"column:scheme_id;not null" json:"scheme_id"`                   // 该广告配置 结构方案id
	IsWifi        int32  `gorm:"column:is_wifi;not null;default:1" json:"is_wifi"`             // 是否在WiFi状态下 1: false 不选 2:选了，在WiFi状态下
	AdTarget      int32  `gorm:"column:ad_target;not null" json:"ad_target"`                   // 目标广告类型 类型 1：应用安装量 2：转化量
	OsMin         string `gorm:"column:os_min;not null" json:"os_min"`                         // 用户系统最小版本
	OsMax         string `gorm:"column:os_max;not null" json:"os_max"`                         // 用户系统最大版本
	Os            string `gorm:"column:os;not null" json:"os"`                                 // 用户操作系统
	IncludeDevice string `gorm:"column:include_device;not null" json:"include_device"`         // 用户包含设备，用 英文逗号分割 例如：ipad,ipad min
	ExcludeDevice string `gorm:"column:exclude_device" json:"exclude_device"`                  // 用户排除设备，用 英文逗号分割 例如：ipad,ipad min
	CampaignType  int32  `gorm:"column:campaign_type;not null;default:1" json:"campaign_type"` // campagin 类型写死为应用广告，1：应用广告 2：自动应用广告
	Languages     string `gorm:"column:languages;not null" json:"languages"`                   // 语言
	CreatedAt     int64  `gorm:"column:created_at;not null" json:"created_at"`                 // 创建日期
	UpdatedAt     int64  `gorm:"column:updated_at;not null" json:"updated_at"`                 // 更新日期
}
