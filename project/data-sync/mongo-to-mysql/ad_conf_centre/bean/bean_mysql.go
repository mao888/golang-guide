package bean

// AdConfAudience 广告配置中心-受众组 mapped from table cruiser_console <ad_conf_audience>
type AdConfAudience struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name           string `gorm:"column:name;not null" json:"name"`                                 // 受众组名称
	ExtendInterest int32  `gorm:"column:extend_interest;not null;default:1" json:"extend_interest"` // 当能够以更低的单次安装费用提高安装量时，系统就会扩展兴趣范围  1：默认不勾选，2：勾选
	AccountID      string `gorm:"column:account_id;not null" json:"account_id"`                     // 所属账户
	CreatedAt      int64  `gorm:"column:created_at;not null" json:"created_at"`                     // 创建日期
	UpdatedAt      int64  `gorm:"column:updated_at;not null" json:"updated_at"`                     // 更新日期
	Creator        int32  `gorm:"column:creator;not null" json:"creator"`                           // 创建者
	Remark         string `gorm:"column:remark" json:"remark"`
}

// AdConfAudienceIncludeRelation 广告配置中心-受众组 排除包含受众关联表 mapped from table cruiser_console <ad_conf_audience_include_relations>
type AdConfAudienceIncludeRelation struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AudienceID int32  `gorm:"column:audience_id;not null" json:"audience_id"` // 受众 ID 归属
	Name       string `gorm:"column:name;not null" json:"name"`               // 受众组名称
	Subtype    string `gorm:"column:subtype;not null" json:"subtype"`         // 受众组条件类型
	FbID       string `gorm:"column:fb_id;not null" json:"fb_id"`             // FB 的受众条件id
	Type       int32  `gorm:"column:type;not null;default:1" json:"type"`     //  1：包含受众，2：排除受众
	CreatedAt  int64  `gorm:"column:created_at;not null" json:"created_at"`   // 创建日期
	UpdatedAt  int64  `gorm:"column:updated_at;not null" json:"updated_at"`   // 更新日期
}

// AdConfAudienceIsegmentationRelation 广告配置中心-受众组细分定位关联表 mapped from table cruiser_console <ad_conf_audience_isegmentation_relations>
type AdConfAudienceIsegmentationRelation struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AudienceID int32  `gorm:"column:audience_id;not null" json:"audience_id"` // 受众 ID 归属
	FbName     string `gorm:"column:fb_name;not null" json:"fb_name"`         // 受众组细分条件名称
	FbType     string `gorm:"column:fb_type;not null" json:"fb_type"`         // FB受众组细分定位类型
	FbID       string `gorm:"column:fb_id;not null" json:"fb_id"`             // FB 的受众条件id
	Type       int32  `gorm:"column:type;not null;default:1" json:"type"`     //  1：包含细分定位，2：排除细分定位
	CreatedAt  int64  `gorm:"column:created_at;not null" json:"created_at"`   // 创建日期
	UpdatedAt  int64  `gorm:"column:updated_at;not null" json:"updated_at"`   // 更新日期
	Label      string `gorm:"column:label;not null" json:"label"`             // 该细分label 名字
}

// AdConfCopywriting (广告投放-配置中心-文案表) mapped from table cruiser_console <ad_conf_copywriting>
type AdConfCopywriting struct {
	ID                     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GameID                 string `gorm:"column:game_id;not null" json:"game_id"`                          // 所属游戏ID
	DefaultLanguage        string `gorm:"column:default_language;not null" json:"default_language"`        // 默认语言简称 例如en
	En                     string `gorm:"column:en" json:"en"`                                             // 英语文案
	De                     string `gorm:"column:de" json:"de"`                                             // 德语文案
	Fr                     string `gorm:"column:fr" json:"fr"`                                             // 法语文案
	Ja                     string `gorm:"column:ja" json:"ja"`                                             // 日语文案
	Ko                     string `gorm:"column:ko" json:"ko"`                                             // 韩语文案
	Es                     string `gorm:"column:es" json:"es"`                                             // 西班牙语文案
	It                     string `gorm:"column:it" json:"it"`                                             // 意大利语言文案
	ZhTw                   string `gorm:"column:zh_tw" json:"zh_tw"`                                       // 繁体中文文案
	Ar                     string `gorm:"column:ar" json:"ar"`                                             // 阿拉伯语文案
	Th                     string `gorm:"column:th" json:"th"`                                             // 泰语文案
	Ru                     string `gorm:"column:ru" json:"ru"`                                             // 俄语文案
	Pt                     string `gorm:"column:pt" json:"pt"`                                             // 葡萄牙语文案
	Nl                     string `gorm:"column:nl" json:"nl"`                                             // 荷兰语文案
	Tr                     string `gorm:"column:tr" json:"tr"`                                             // 土耳其语文案
	Vi                     string `gorm:"column:vi" json:"vi"`                                             // 越南语文案
	Ms                     string `gorm:"column:ms" json:"ms"`                                             // 马来语文案
	Iid                    string `gorm:"column:iid" json:"iid"`                                           // 印尼语文案
	Tl                     string `gorm:"column:tl" json:"tl"`                                             // 菲律宾语文案
	CreatedAt              int64  `gorm:"column:created_at;not null" json:"created_at"`                    // 创建日期
	UpdatedAt              int64  `gorm:"column:updated_at" json:"updated_at"`                             // 更新日期
	DefaultLanguageContent string `gorm:"column:default_language_content" json:"default_language_content"` // 默认语言文案 内容
}

// AdConfCountry (广告投放-配置中心-国家组表) mapped from table cruiser_console <ad_conf_country>
type AdConfCountry struct {
	ID                  int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name                string `gorm:"column:name;not null" json:"name"`                          // 国家组名称
	IncludeCountryCodes string `gorm:"column:include_country_codes" json:"include_country_codes"` // 包含的国家编码 ,分隔  GeoLocations
	ExcludeCountryCodes string `gorm:"column:exclude_country_codes" json:"exclude_country_codes"` // 排除的国家编码 ,分隔
	GameID              string `gorm:"column:game_id;not null" json:"game_id"`                    // 所属游戏ID
	CreatedAt           int64  `gorm:"column:created_at;not null" json:"created_at"`              // 创建日期
	UpdatedAt           int64  `gorm:"column:updated_at;not null" json:"updated_at"`              // 更新日期
}

// AdConfPosition (广告投放-配置中心-版位组表) mapped from table cruiser_console <ad_conf_position>
type AdConfPosition struct {
	ID              int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name            string `gorm:"column:name;not null" json:"name"`                             // 版位组名称
	PositionType    int32  `gorm:"column:position_type;not null;default:1" json:"position_type"` // 版位 包含两种：自定义版位和自动版位 1：自定义版位 ，2：自动版位(默认选中自定义版位1)
	Facebook        string `gorm:"column:facebook" json:"facebook"`                              // facebook 相关 JSON 配置
	Instagram       string `gorm:"column:instagram" json:"instagram"`                            // Instagram 相关 JSON 配置
	AudienceNetwork string `gorm:"column:audience_network" json:"audience_network"`              // Audience Network 相关 JSON 配置
	Messenger       string `gorm:"column:messenger" json:"messenger"`                            // messenger  相关 JSON 配置
	CreatedAt       int64  `gorm:"column:created_at;not null" json:"created_at"`                 // 创建日期
	UpdatedAt       int64  `gorm:"column:updated_at;not null" json:"updated_at"`                 // 更新日期
}

// AdConfScheme (广告投放-配置中心-结构方案表) mapped from table cruiser_console <ad_conf_scheme>
type AdConfScheme struct {
	ID                 int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name               string  `gorm:"column:name;not null" json:"name"`                                           // 方案名称
	CampaignDimension  string  `gorm:"column:campaign_dimension;not null" json:"campaign_dimension"`               // 广告系列维度为固定维度ID拼接如1,2（1：国家/国家组,2：受众组,3:版位组,4:年龄,5:性别,6:语言,7:优化方式,8:素材,9:标签,10:广告类型）
	AdsetDimension     string  `gorm:"column:adset_dimension;not null" json:"adset_dimension"`                     // 广告组维度为固定维度ID拼接如1,2（1：国家/国家组,2：受众组,3:版位组,4:年龄,5:性别,6:语言,7:优化方式,8:素材,9:标签,10:广告类型）
	CampaignLimitAdset int32   `gorm:"column:campaign_limit_adset;not null;default:1" json:"campaign_limit_adset"` // 每个campaign最多有X 个ad set：必填，数值格式，范围为1-999  CampaignLimit
	AdsetLimitAd       int32   `gorm:"column:adset_limit_ad;not null;default:1" json:"adset_limit_ad"`             // 每个ad set 最多有X 个ad：必填，数值格式，范围为1-999		AdsetLimit
	BudgetLevel        int32   `gorm:"column:budget_level;not null;default:1" json:"budget_level"`                 // 必选，单选，默认为CBO广告系列预算 1:CBO广告系列预算 2:ABO广告组预算
	BudgetLimit        float32 `gorm:"column:budget_limit" json:"budget_limit"`                                    // 预算上限 ：非必填可为空，数值格式，正整数
	AuthorID           int32   `gorm:"column:author_id" json:"author_id"`                                          // 人员ID，作者
	CreatedAt          int64   `gorm:"column:created_at;not null" json:"created_at"`                               // 创建日期
	UpdatedAt          int64   `gorm:"column:updated_at;not null" json:"updated_at"`                               // 更新日期
}

// User mapped from table user_console <user>
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
