package bean

// AdPreview mapped from table cruiser_console <ad_preview>
type AdPreview struct {
	ID            int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID        string `gorm:"column:game_id;not null" json:"game_id"`
	AppID         int32  `gorm:"column:app_id;not null" json:"app_id"`
	AccountID     string `gorm:"column:account_id;not null;default:0" json:"account_id"`
	FbAppID       int64  `gorm:"column:fb_app_id;not null" json:"fb_app_id"`
	FbHomePageID  string `gorm:"column:fb_home_page_id;not null" json:"fb_home_page_id"`
	InstagramID   string `gorm:"column:instagram_id;not null" json:"instagram_id"`
	BudgetLevel   int32  `gorm:"column:budget_level" json:"budget_level"`
	MarketURL     string `gorm:"column:market_url;not null" json:"market_url"`
	UserOs        string `gorm:"column:user_os" json:"user_os"`
	IncludeDevice string `gorm:"column:include_device" json:"include_device"`
	ExcludeDevice string `gorm:"column:exclude_device" json:"exclude_device"`
	IsWifi        int32  `gorm:"column:is_wifi;not null" json:"is_wifi"`
	CreatedAt     int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     int64  `gorm:"column:updated_at;not null" json:"updated_at"`
	IsDeleted     int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// AdPreviewAd mapped from table cruiser_console <ad_preview_ad>
type AdPreviewAd struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterialID int32  `gorm:"column:material_id" json:"material_id"`
	CampaignID int32  `gorm:"column:campaign_id;not null" json:"campaign_id"`
	AdSetID    int32  `gorm:"column:ad_set_id;not null" json:"ad_set_id"`
	Name       string `gorm:"column:name;not null" json:"name"`
	AdTitle    string `gorm:"column:ad_title;not null" json:"ad_title"`
	AdText     string `gorm:"column:ad_text;not null" json:"ad_text"`
	LanguageID string `gorm:"column:language_id;not null" json:"language_id"`
	IsFinish   int32  `gorm:"column:is_finish;not null" json:"is_finish"`
	AdType     int32  `gorm:"column:ad_type;not null" json:"ad_type"`
	CreatedAt  int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
}

// AdPreviewAdSet mapped from table cruiser_console <ad_preview_ad_set>
type AdPreviewAdSet struct {
	ID            int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CampaignID    int32   `gorm:"column:campaign_id" json:"campaign_id"`
	Name          string  `gorm:"column:name" json:"name"`
	Bugdet        int32   `gorm:"column:bugdet" json:"bugdet"`
	Cost          float32 `gorm:"column:cost;not null" json:"cost"`
	TargetingData string  `gorm:"column:targeting_data" json:"targeting_data"`
	IsFinish      int32   `gorm:"column:is_finish" json:"is_finish"` // 是否上传完成
	CreatedAt     int64   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     int64   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// AdPreviewCampaign mapped from table cruiser_console <ad_preview_campaign>
type AdPreviewCampaign struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PreviewID int32  `gorm:"column:preview_id;not null" json:"preview_id"`
	SchemeID  int32  `gorm:"column:scheme_id" json:"scheme_id"`
	Name      string `gorm:"column:name" json:"name"`
	Budget    int32  `gorm:"column:budget" json:"budget"`
	CreatedAt int64  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int64  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// AdPreviewConfig mapped from table cruiser_console <ad_preview_config>
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

// AdPreviewConfigRelationAudience mapped from table cruiser_console <ad_preview_config_relation_audience>
type AdPreviewConfigRelationAudience struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdFbConfigID int32  `gorm:"column:ad_fb_config_id;not null" json:"ad_fb_config_id"` // Fb 广告配置从属 id
	Name         string `gorm:"column:name;not null" json:"name"`                       // 受众组的名称
	AudienceID   int32  `gorm:"column:audience_id" json:"audience_id"`                  // 对应的受众组ID
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`           // 创建日期
	UpdatedAt    int64  `gorm:"column:updated_at;not null" json:"updated_at"`           // 更新日期
}

// AdPreviewConfigRelationCountryGroup mapped from table cruiser_console <ad_preview_config_relation_country_group>
type AdPreviewConfigRelationCountryGroup struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdFbConfigID   int32  `gorm:"column:ad_fb_config_id;not null" json:"ad_fb_config_id"`    // Fb 广告配置从属 id
	Name           string `gorm:"column:name;not null" json:"name"`                          // 该国家组对应的名称或者单个选择的国家的名称
	CountryGroupID string `gorm:"column:country_group_id;default:0" json:"country_group_id"` // 对应的国家组ID，如果是单独选择的国家，该字段为该国家的简写，或者州的简称
	Type           int32  `gorm:"column:type;not null;default:1" json:"type"`                //  1：该条目是 国家组id ，2：单选个别国家 3:国家下面的州
	CreatedAt      int64  `gorm:"column:created_at;not null" json:"created_at"`              // 创建日期
	UpdatedAt      int64  `gorm:"column:updated_at;not null" json:"updated_at"`              // 更新日期
}

// AdPreviewConfigRelationLanguage mapped from table cruiser_console <ad_preview_config_relation_language>
type AdPreviewConfigRelationLanguage struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdFbConfigID   int32  `gorm:"column:ad_fb_config_id;not null" json:"ad_fb_config_id"` // Fb 广告配置从属 id
	Name           string `gorm:"column:name;not null" json:"name"`                       // 该语言对应的名称，国家或者国家组配置
	LanguageID     int32  `gorm:"column:language_id;not null" json:"language_id"`         // FB 语言ID
	LanguageCode   string `gorm:"column:language_code" json:"language_code"`              // 语言对应的简写
	CountryGroupID int32  `gorm:"column:country_group_id" json:"country_group_id"`        // 对应的国家组ID，单选国家为0
	Type           int32  `gorm:"column:type;not null;default:1" json:"type"`             //  1：该条目是 国家组配置的语言 ，2：单选个别国家配置的语言
	CreatedAt      int64  `gorm:"column:created_at;not null" json:"created_at"`           // 创建日期
	UpdatedAt      int64  `gorm:"column:updated_at;not null" json:"updated_at"`           // 更新日期
}

// AdPreviewConfigRelationOpt mapped from table cruiser_console <ad_preview_config_relation_opt>
type AdPreviewConfigRelationOpt struct {
	ID               int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdFbConfigID     int32  `gorm:"column:ad_fb_config_id;not null" json:"ad_fb_config_id"`        // Fb 广告配置从属 id
	BillingEvent     string `gorm:"column:billing_event;default:IMPRESSIONS" json:"billing_event"` // 计费方式 billing_event  IMPRESSIONS 展示次数 暂时写死
	OptimizationGoal string `gorm:"column:optimization_goal;not null" json:"optimization_goal"`    // 优化目标 暂时会有 AEO（OFFSITE_CONVERSIONS），install（APP_INSTALLS），VO(VALUE)
	WindowDays       int32  `gorm:"column:window_days" json:"window_days"`                         // window_days 对应的枚举 1，7 天
	Event            string `gorm:"column:event" json:"event"`                                     //  当优化目标位AEO 的时候 会有14个事件
	EventType        string `gorm:"column:event_type;default:CLICK_THROUGH" json:"event_type"`     // event_type 暂时写死 CLICK_THROUGH
	BidStrategy      string `gorm:"column:bid_strategy;not null" json:"bid_strategy"`              // 竞价策略  枚举 1：auto LOWEST_COST_WITHOUT_CAP，LOWEST_COST_WITH_BID_CAP，LOWEST_COST_WITH_MIN_ROAS，COST_CAP
	CreatedAt        int64  `gorm:"column:created_at;not null" json:"created_at"`                  // 创建日期
	UpdatedAt        int64  `gorm:"column:updated_at;not null" json:"updated_at"`                  // 更新日期
}

// AdPreviewConfigRelationPosition mapped from table cruiser_console <ad_preview_config_relation_position>
type AdPreviewConfigRelationPosition struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdFbConfigID int32  `gorm:"column:ad_fb_config_id;not null" json:"ad_fb_config_id"` // Fb 广告配置从属 id
	Name         string `gorm:"column:name;not null" json:"name"`                       // 版位的名称
	PositionID   int32  `gorm:"column:position_id" json:"position_id"`                  // 对应的版位ID
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`           // 创建日期
	UpdatedAt    int64  `gorm:"column:updated_at;not null" json:"updated_at"`           // 更新日期
}

// AdPreviewParam mapped from table cruiser_console <ad_preview_param>
type AdPreviewParam struct {
	ID      int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AdSetID int32  `gorm:"column:ad_set_id;not null" json:"ad_set_id"`
	Params  string `gorm:"column:params" json:"params"`
}
