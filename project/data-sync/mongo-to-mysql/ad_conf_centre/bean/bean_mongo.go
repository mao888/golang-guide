package bean

import "time"

// MCfgAudienceModel 受众组 数据模型
type MCfgAudienceModel struct {
	ID                      int32                     `json:"_id" bson:"_id"`
	CompanyId               int                       `json:"company_id" bson:"company_id"`                               //公司id
	Name                    string                    `json:"name" bson:"name"`                                           // 受众组名称
	AccountId               string                    `json:"account_id" bson:"account_id"`                               // 账户id
	CustomAudiences         []CustomAudiences         `json:"custom_audiences" bson:"custom_audiences"`                   // 包含受众
	ExcludedCustomAudiences []ExcludedCustomAudiences `json:"excluded_custom_audiences" bson:"excluded_custom_audiences"` // 排除受众
	FlexibleSpec            []FlexibleSpec            `json:"flexible_spec" bson:"flexible_spec"`                         // 包含细分定位、缩小细分定位
	Exclusions              *Exclusions               `json:"exclusions" bson:"exclusions"`                               // 排除细分定位
	TargetingOptimization   string                    `json:"targeting_optimization" bson:"targeting_optimization"`       // 细分定位扩展优化
	VerifyStr               string                    `json:"verify_str" bson:"verify_str"`                               // 受众信息校验参数
	CreateUser              string                    `json:"create_user" bson:"create_user"`                             // 创建人
	UserId                  int32                     `json:"user_id" bson:"user_id"`                                     // 创建人id
	SourceSys               string                    `json:"source_sys" bson:"source_sys"`
	CreateTime              *time.Time                `json:"create_time" bson:"create_time"`
	UpdateTime              *time.Time                `json:"update_time" bson:"update_time"`
}

type CustomAudiences struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type ExcludedCustomAudiences struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Exclusions struct {
	EducationStatuses []int `json:"education_statuses" bson:"education_statuses"`
	Interests         []struct {
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"interests" bson:"interests"`
	CollegeYears         []int `json:"college_years" bson:"college_years"`
	RelationshipStatuses []int `json:"relationship_statuses" bson:"relationship_statuses"`
	Income               []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"income" bson:"income"`
	FamilyStatuses []struct {
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"family_statuses" bson:"family_statuses"`
	Behaviors []struct {
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"behaviors" bson:"behaviors"`
}

type FlexibleSpec struct {
	EducationStatuses []int      `json:"education_statuses" bson:"education_statuses"` // 教育程度
	Interests         []struct { // 兴趣
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"interests" bson:"interests"`
	CollegeYears         []int      `json:"college_years" bson:"college_years"`                 // 大学毕业时间
	RelationshipStatuses []int      `json:"relationship_statuses" bson:"relationship_statuses"` // 感情状况
	Income               []struct { // 收入
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"income" bson:"income"`
	FamilyStatuses []struct { // 家庭状态
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"family_statuses" bson:"family_statuses"`
	Behaviors []struct { // 行为
		Id   string `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"behaviors" bson:"behaviors"`
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
	LoginTime       *time.Time    `json:"login_time" bson:"loginTime"`            // 最后登录时间
	MaintainStatus  bool          `json:"maintain_status" bson:"maintain_status"` // 维护状态
	GuiderStep      int           `json:"guider_step" bson:"guider_step"`         // 新手引导
	AccessSystem    []string      `json:"access_system" bson:"access_system"`     // 可访问的系统
	DefaultCompany  int           `json:"default_company" bson:"default_company"` // 当前选中公司
}

// MCfgFrame 结构方案数据模型
type MCfgFrame struct {
	Id               int32    `json:"_id" bson:"_id"`                             // 结构方案id
	CompanyId        int32    `json:"company_id" bson:"company_id"`               // 公司id
	Name             string   `json:"name" bson:"name"`                           // 结构方案名称
	CampaignDims     []string `json:"campaign_dims" bson:"campaign_dims"`         // campaign划分维度 ['countries', 'audiences', 'positions', 'age', 'sex', 'language', 'strategys','materials', 'tag', 'adtype']
	AdsetDims        []string `json:"adset_dims" bson:"adset_dims"`               // adset划分维度
	CampaignLimit    int32    `json:"campaign_limit" bson:"campaign_limit"`       // campaign数量上限
	AdsetLimit       int32    `json:"adset_limit" bson:"adset_limit"`             // adset数量上限
	IsCbo            bool     `json:"is_cbo" bson:"is_cbo"`                       // 是否开启cbo
	OptimizationGoal string   `json:"optimization_goal" bson:"optimization_goal"` // 优化目标
	BidStrategy      string   `json:"bid_strategy" bson:"bid_strategy"`           // 竞价策略
	AttributionSpec  struct {
		EventType  string `json:"event_type" bson:"event_type"`
		WindowDays int    `json:"window_days" bson:"window_days"`
	} `json:"attribution_spec" bson:"attribution_spec"` // 转化窗口
	CustomEventType string     `json:"custom_event_type" bson:"custom_event_type"` // AEO的14个应用事件
	BillingEvent    string     `json:"billing_event" bson:"billing_event"`         // 计费方式
	PacingType      []string   `json:"pacing_type" bson:"pacing_type"`
	BudgetLimit     float64    `json:"budget_limit" bson:"budget_limit"` // 单日预算上限
	VerifyStr       string     `json:"verify_str" bson:"verify_str"`     // 结构方案信息校验参数
	CreateUser      string     `json:"create_user" bson:"create_user"`   // 创建人
	UserId          int32      `json:"user_id" bson:"user_id"`           // 创建人
	SourceSys       string     `json:"source_sys" bson:"source_sys""`
	CreateTime      *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime      *time.Time `json:"update_time" bson:"update_time"`
}

// MfgPosition 版位组 数据模型
type MfgPosition struct {
	Id                       int32      `json:"_id" bson:"_id""`                                              // 版位组id
	CompanyId                int32      `json:"company_id" bson:"company_id"`                                 //公司id
	Name                     string     `json:"name" bson:"name"`                                             // 版位组名称
	PositionType             string     `json:"positionType" bson:"positionType"`                             // 版位类型  editPosition、autoPosition
	PublisherPlatforms       []string   `json:"publisher_platforms" bson:"publisher_platforms"`               // 版位平台
	FacebookPositions        []string   `json:"facebook_positions" bson:"facebook_positions"`                 // fb版位数组
	InstagramPositions       []string   `json:"instagram_positions" bson:"instagram_positions"`               // nstagram 版位数组
	MessengerPositions       []string   `json:"messenger_positions" bson:"messenger_positions"`               // messenger 版位数组
	AudienceNetworkPositions []string   `json:"audience_network_positions" bson:"audience_network_positions"` // audience_network 版位数组
	ContainPosition          []string   `json:"containPosition" bson:"contain_position"`                      // 所有版位的 对照数组信息
	CreateUser               string     `json:"create_user" bson:"create_user"`
	UserId                   int32      `json:"user_id" bson:"user_id"`
	CreateTime               *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime               *time.Time `json:"update_time" bson:"update_time"`
}

// MAdTexts adtexts 素材 文案库
type MAdTexts struct {
	Id          string `json:"_id" bson:"_id"`
	Fbid        string `json:"fbid"  bson:"fbid"`
	EnText      string `json:"en_text" bson:"en_text"`
	Translation []struct {
		Id   string `json:"_id" bson:"id"`
		Lang int32  `json:"lang" bson:"lang"`
		Text string `json:"text" bson:"text"`
	} `json:"translation" bson:"translation"`
	CreateTime  *time.Time  `json:"create_time" bson:"create_time"`
	UpdateTime  *time.Time  `json:"update_time" bson:"update_time"`
	CompanyId   int32       `json:"company_id" bson:"company_id"`
	GameId      interface{} `json:"game_id" bson:"game_id"`
	DefaultLang int32       `json:"default_lang" bson:"default_lang"`
	DefaultText string      `json:"default_text" bson:"default_text"`
}

// MCfgCountry 国家组 数据模型
type MCfgCountry struct {
	Id           int32       `json:"_id" bson:"_id"`               //国家组id
	CompanyId    int32       `json:"company_id" bson:"company_id"` //公司id
	Name         string      `json:"name" bson:"name"`             //国家组名称
	Fbid         string      `json:"fbid" bson:"fbid"`             //fbId gia迁移后废弃
	GameId       interface{} `json:"game_id" bson:"game_id"`       // 游戏id
	GeoLocations struct {
		Countries     []string `json:"countries" bson:"countries"`
		CountryGroups []string `json:"country_groups" bson:"country_groups"`
		LocationTypes []string `json:"location_types" bson:"location_types"`
	} `json:"geo_locations" bson:"geo_locations"` // 包含国家
	ExcludedGeoLocations ExcludedGeoLocations `json:"excluded_geo_locations" bson:"excluded_geo_locations"` //排除国家
	ContainCountries     []string             `json:"contain_countries" bson:"contain_countries"`           // 包含国家 的中文名称数组
	ExcludedCountries    []string             `json:"excluded_countries" bson:"excluded_countries"`         // 排除国家的中文名称数据
	CreateUser           string               `json:"create_user" bson:"create_user"`
	UserId               int32                `json:"user_id" bson:"user_id"`
	CreateTime           *time.Time           `json:"create_time" bson:"create_time"`
	UpdateTime           *time.Time           `json:"update_time" bson:"update_time"`
}

type ExcludedGeoLocations struct {
	Countries     []string `json:"countries" bson:"countries"`
	CountryGroups []string `json:"country_groups" bson:"country_groups"`
	LocationTypes []string `json:"location_types" bson:"location_types"`
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
