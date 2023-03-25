package bean

import "time"

// MFBCampaignModel 广告系列列表 From Mongo/cruiser_console_v2/fbcampaigns
type MFBCampaignModel struct {
	Id               string     `json:"_id" bson:"_id"`
	CampaignId       string     `json:"campaign_id" bson:"campaign_id"`
	AdTarget         string     `json:"ad_target" bson:"ad_target"` //  广告目标字段， 1：应用安装量，2: 转化量
	CampaignName     string     `json:"campaign_name" bson:"campaign_name"`
	CfgFrameId       string     `json:"cfg_frame_id" bson:"cfg_frame_id"`   // 广告结构方案id （v3版 新增字段）
	AudienceName     string     `json:"audience_name" bson:"audience_name"` // 受众名称
	CustomSuf        string     `json:"custom_suf" bson:"custom_suf"`       // 自定义后缀
	AccountId        string     `json:"account_id" bson:"account_id"`
	CreateUser       string     `json:"create_user" bson:"create_user"` //创建人
	UserId           int32      `json:"user_id" bson:"user_id"`
	AiLevelCfg       string     `json:"ai_level_cfg" bson:"ai_level_cfg"` //ai广告模板层级配置
	AiRulerId        int32      `json:"ai_ruler_id" bson:"ai_ruler_id"`
	AiRulerOpen      bool       `json:"ai_ruler_open" bson:"ai_ruler_open"`
	CreateTime       *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime       *time.Time `json:"update_time" bson:"update_time"`
	CreatedByCruiser bool       `json:"created_by_cruiser" bson:"created_by_cruiser"` //是否为自动化创建
	Status           string     `json:"status" bson:"status"`                         //状态 ['ACTIVE', 'PAUSED', 'DELETED', 'ARCHIVED']
	BuyingType       string     `json:"buying_type" bson:"buying_type"`               //购买类型	['AUCTION', 'RESERVED']
	Objective        string     `json:"objective" bson:"objective"`                   //营销目标	['APP_INSTALLS', 'CONVERSIONS']
	//CBO特有属性
	DailyBudget    int      `json:"daily_budget" bson:"daily_budget"`
	LifetimeBudget int32    `json:"lifetime_budget" bson:"lifetime_budget"`
	BidStrategy    string   `json:"bid_strategy" bson:"bid_strategy"` //竞价策略	['LOWEST_COST_WITHOUT_CAP', 'LOWEST_COST_WITH_BID_CAP', 'COST_CAP', 'TARGET_COST', 'LOWEST_COST_WITH_MIN_ROAS']
	PacingType     []string `json:"pacing_type" bson:"pacing_type"`   //投放时段或类型说明字段 enum: ['standard', 'no_pacing', 'day_parting']
	CreateType     string   `json:"create_type" bson:"create_type"`
	CompanyId      int32    `json:"company_id" bson:"company_id"`
}

// MFBAdSetModel From Mongo/cruiser_console_v2/fbadsets
type MFBAdSetModel struct {
	Id                string `json:"_id" bson:"_id"`
	Status            string `json:"status" bson:"status"` //状态	['ACTIVE', 'PAUSED', 'DELETED', 'ARCHIVED']
	AdsetId           string `json:"adset_id" bson:"adset_id"`
	AdsetName         string `json:"adset_name" bson:"adset_name"`
	CampaignId        string `json:"campaign_id" bson:"campaign_id"`
	AccountId         string `json:"account_id" bson:"account_id"`
	CreatedByCruiser  bool   `json:"created_by_cruiser" bson:"created_by_cruiser"`   // 是否为自动化创建
	CloneTemplateOpen bool   `json:"clone_template_open" bson:"clone_template_open"` //克隆模板是否开启
	CreateUser        string `json:"create_user" bson:"create_user"`                 //创建人
	UserId            int32  `json:"user_id" bson:"user_id"`
	//受众
	Targeting struct {
		CustomAudiences         []CustomAudiences         `json:"custom_audiences" bson:"custom_audiences"`                   // 自定义受众
		ExcludedCustomAudiences []ExcludedCustomAudiences `json:"excluded_custom_audiences" bson:"excluded_custom_audiences"` // 排除自定义受众
		Connections             []Connections             `json:"connections" bson:"connections"`                             // 关系 与以下内容建立关系网络的用户
		ExcludedConnections     []ExcludedConnections     `json:"excluded_connections" bson:"excluded_connections"`           // 排除与以下内容建立关系网络的用户
		FriendsOfConnections    []FriendsOfConnections    `json:"friends_of_connections" bson:"friends_of_connections"`       // 与以下内容建立关系网络的用户的好友
		AppInstallState         string                    `json:"app_install_state" bson:"app_install_state"`                 // not_installed 设置关系时加入该属性
		GeoLocations            struct {
			LocationTypes []string `json:"location_types" bson:"location_types"`
			Countries     []string `json:"countries" bson:"countries"`
			CountryGroups []string `json:"country_groups" bson:"country_groups"`
		} `json:"geo_locations" bson:"geo_locations"`
		AgeMin                int            `json:"age_min" bson:"age_min"`
		AgeMax                int            `json:"age_max" bson:"age_max"`
		Genders               []int          `json:"genders" bson:"genders"`
		Locales               []int          `json:"locales" bson:"locales"`                               //语言
		FlexibleSpec          []FlexibleSpec `json:"flexible_spec" bson:"flexible_spec"`                   // 细分定位 包含至少符合一项条件的用户
		Exclusions            *Exclusions    `json:"exclusions" bson:"exclusions"`                         // 细分定位 排除符合以下至少一项条件的用户
		TargetingOptimization string         `json:"targeting_optimization" bson:"targeting_optimization"` //expansion_all-当能够以更低的安装费用提高安装量时，系统就会扩展兴趣范围 ['expansion_all', 'none']
		UserOs                []string       `json:"user_os" bson:"user_os"`
		UserDevice            []string       `json:"user_device" bson:"user_device"`                   //包含的设备
		ExcludedUserDevice    []string       `json:"excluded_user_device" bson:"excluded_user_device"` //被排除的设备
		WirelessCarrier       []string       `json:"wireless_carrier" bson:"wireless_carrier"`         //仅在连接wifi时 ["Wifi"]
		//-----------版位---------
		PublisherPlatforms       []string `json:"publisher_platforms" bson:"publisher_platforms"`               // 版位平台  enum: ['facebook', 'instagram', 'messenger', 'audience_network']
		FacebookPositions        []string `json:"facebook_positions" bson:"facebook_positions"`                 // fb版位数组 enum: ['feed', 'right_hand_column', 'instant_article', 'marketplace', 'suggested_video', 'instream_video', 'story']
		InstagramPositions       []string `json:"instagram_positions" bson:"instagram_positions"`               // instagram 版位数组 enum: ['stream', 'story']
		MessengerPositions       []string `json:"messenger_positions" bson:"messenger_positions"`               // messenger 版位数组 enum: ['messenger_home', 'sponsored_messages', 'story']
		AudienceNetworkPositions []string `json:"audience_network_positions" bson:"audience_network_positions"` // audience_network 版位数组 enum: ['classic', 'instream_video', 'rewarded_video']
		//-----------版位---------
	} `json:"targeting" bson:"targeting"`
	//---------未开启cbo时特有字段-----
	BidStrategy string   `json:"bid_strategy" bson:"bid_strategy"` //竞价策略，比campaign层级少两个 ['LOWEST_COST_WITHOUT_CAP', 'LOWEST_COST_WITH_BID_CAP', 'COST_CAP', 'TARGET_COST', 'LOWEST_COST_WITH_MIN_ROAS']
	DailyBudget int      `json:"daily_budget" bson:"daily_budget"` //单日预算
	PacingType  []string `json:"pacing_type" bson:"pacing_type"`   //投放时段或类型说明字段 enum: ['standard', 'no_pacing', 'day_parting']
	//---------未开启cbo时特有字段-----

	//---------开启cbo时字段----------
	DailyMinSpendTarget    int32 `json:"daily_min_spend_target" bson:"daily_min_spend_target"`       //开启cbo后广告组限额下限-每日预算
	DailySpendCap          int32 `json:"daily_spend_cap" bson:"daily_spend_cap"`                     //开启cbo后广告组限额上限-每日预算
	LifetimeMinSpendTarget int32 `json:"lifetime_min_spend_target" bson:"lifetime_min_spend_target"` //开启cbo后广告组限额下限-总预算
	LifetimeSpendCap       int32 `json:"lifetime_spend_cap" bson:"lifetime_spend_cap"`               //开启cbo后广告组限额上限-总预算
	BidConstraints         struct {
		RoasAverageFloor int `json:"roas_average_floor" bson:"roas_average_floor"`
	} `json:"bid_constraints" bson:"bid_constraints"` //roas竞价策略值
	StartTime *time.Time `json:"start_time" bson:"start_time"` //cbo开启时，广告系列选择总预算，广告组排期选择开始和结束时间
	EndTime   *time.Time `json:"end_time" bson:"end_time"`
	//---------开启cbo时字段----------

	AttributionSpec []struct {
		EventType  string `json:"event_type" bson:"event_type"`
		WindowDays int    `json:"window_days" bson:"window_days"`
	} `json:"attribution_spec" bson:"attribution_spec"` //转化时间窗
	BidAmount        int32  `json:"bid_amount" bson:"bid_amount"`
	OptimizationGoal string `json:"optimization_goal" bson:"optimization_goal"` // 优化方式 ['APP_INSTALLS', 'VALUE', 'OFFSITE_CONVERSIONS', 'LINK_CLICKS', 'DERIVED_EVENTS', 'LANDING_PAGE_VIEWS']
	BillingEvent     string `json:"billing_event" bson:"billing_event"`         // 计费方式 ['APP_INSTALLS', 'IMPRESSIONS', 'LINK_CLICKS']
	PositionType     string `json:"positionType" bson:"position_type"`          // 自定义字段 ['autoPosition', 'editPosition']
	PromotedObject   struct {
		ObjectStoreUrl  string `json:"object_store_url" bson:"object_store_url"`
		ApplicationId   string `json:"application_id" bson:"application_id"`
		CustomEventType string `json:"custom_event_type" bson:"custom_event_type"` //当优化目标为OFFSITE_CONVERSIONS optimize_event_value值会填充到custom_event_type属性
	} `json:"promoted_object" bson:"promoted_object"` // 缺少时报错找不到推广对象
	CreateTime *time.Time `json:"create_time" bson:"create_time"`
	CreateType string     `json:"create_type" bson:"create_type"`
	CompanyId  int32      `json:"company_id" bson:"company_id"`
	//V                int    `json:"__v"`
}

type CustomAudiences struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type ExcludedCustomAudiences struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type Connections struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type ExcludedConnections struct {
	ID   string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

type FriendsOfConnections struct {
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
