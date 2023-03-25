package bean

import "time"

// MAdFlowCfg 广告组信息 From Mongo/cruiser_console_v2/adflowcfg
type MAdFlowCfg struct {
	Id             string `json:"_id" bson:"_id"`
	OId            string `json:"o_id" bson:"o_id"`
	Fbid           string `json:"fbid" bson:"fbid"`
	CampaignId     string `json:"campaign_id" bson:"campaign_id"`           // 广告系列id
	AdsetId        string `json:"adset_id" bson:"adset_id"`                 // 广告组id
	CfgFrameId     string `json:"cfg_frame_id" bson:"cfg_frame_id"`         // 结构方案id
	CfgCountryIds  []int  `json:"cfg_country_ids" bson:"cfg_country_ids"`   // 国家组id
	CfgAudienceIds []int  `json:"cfg_audience_ids" bson:"cfg_audience_ids"` // 受众组id
	CfgPositionIds []int  `json:"cfg_position_ids" bson:"cfg_position_ids"` // 版位组id
	Ages           []struct {
		AgeMin int `json:"age_min"`
		AgeMax int `json:"age_max"`
	} `json:"ages" bson:"ages"` // 年龄
	Genders        []int         `json:"genders" bson:"genders"` // 性别
	Contrys        []interface{} `json:"contrys" bson:"contrys"`
	LanguageGroups []struct {
		Name      string `json:"name"`
		NoLimit   bool   `json:"no_limit"`
		Languages int    `json:"languages"`
	} `json:"language_groups" bson:"language_groups"` // 年龄组 [{ name: '国家组名称', languages: ['en', 'zh'] }]
	Strategys []struct {
		Target struct {
			TargetType string `json:"target_type" bson:"target_type"`
		} `json:"target" bson:"target"`
		Option       string `json:"option" bson:"option"`
		BillingEvent string `json:"billing_event" bson:"billingEvent"`
	} `json:"strategys" bson:"strategys"` // 优化方式
	UserOs             string        `json:"user_os" bson:"user_os"`                           // 设备系统
	UserDevice         []string      `json:"user_device" bson:"user_device"`                   // 包含设备
	ExcludedUserDevice []interface{} `json:"excluded_user_device" bson:"excluded_user_device"` // 排除设备
	IsWifi             bool          `json:"is_wifi" bson:"is_wifi"`                           // 链接wifi
	VerMin             string        `json:"ver_min" bson:"ver_min"`                           // 最低版本号
	VerMax             interface{}   `json:"ver_max" bson:"ver_max"`                           // 最高版本号
	CreateTime         *time.Time    `json:"create_time" bson:"create_time"`
	UpdateTime         *time.Time    `json:"update_time" bson:"update_time"`
}

// MAdFlow 素材 From Mongo/cruiser_console_v2/adflow
type MAdFlow struct {
	Id            int `json:"_id" bson:"_id"`
	CampaignDatas []struct {
		Id         string `json:"_id" bson:"_id"`
		AdsetDatas []struct {
			Id             string `json:"_id" bson:"_id"`
			Name           string `json:"name" bson:"name"`
			Status         string `json:"status" bson:"status"`
			AccountId      string `json:"account_id" bson:"account_id"`
			CampaignId     string `json:"campaign_id" bson:"campaign_id"`
			PromotedObject struct {
				ObjectStoreUrl  string `json:"object_store_url" bson:"object_store_url"`
				ApplicationId   string `json:"application_id" bson:"application_id"`
				CustomEventType string `json:"custom_event_type" bson:"custom_event_type"`
			} `json:"promoted_object" bson:"promoted_object"`
			Targeting struct {
				AppInstallState string `json:"app_install_state" bson:"app_install_state"`
				GeoLocations    struct {
					Countries     []string `json:"countries" bson:"countries"`
					LocationTypes []string `json:"location_types" bson:"location_types"`
				} `json:"geo_locations" bson:"geo_locations"`
				CustomAudiences []struct {
					Id   string `json:"id" bson:"id"`
					Name string `json:"name" bson:"name"`
				} `json:"custom_audiences" bson:"custom_audiences"`
				TargetingOptimization string        `json:"targeting_optimization" bson:"targeting_optimization"`
				AgeMin                int           `json:"age_min" bson:"age_min"`
				AgeMax                int           `json:"age_max" bson:"age_max"`
				Genders               []int         `json:"genders" bson:"genders"`
				UserOs                []string      `json:"user_os" bson:"user_os"`
				UserDevice            []string      `json:"user_device" bson:"user_device"`
				ExcludedUserDevice    []interface{} `json:"excluded_user_device" bson:"excluded_user_device"`
			} `json:"targeting" bson:"targeting"`
			BillingEvent     string `json:"billing_event" bson:"billing_event"`
			OptimizationGoal string `json:"optimization_goal" bson:"optimization_goal"`
			AdDatas          []struct {
				AssetId     string `json:"asset_id" bson:"asset_id"`
				Name        string `json:"name" bson:"name"`
				Status      string `json:"status" bson:"status"`
				AdText      string `json:"ad_text" bson:"ad_text"`
				AdTitle     string `json:"ad_title" bson:"ad_title"`
				PerformType string `json:"perform_type" bson:"perform_type"`
			} `json:"ad_datas" bson:"ad_datas"`
			AttributionSpec []struct {
				EventType  string `json:"event_type" bson:"event_type"`
				WindowDays int    `json:"window_days" bson:"window_days"`
			} `json:"attribution_spec" bson:"attribution_spec"`
		} `json:"adset_datas" bson:"adset_datas"`
	} `json:"campaign_datas" bson:"campaign_datas"`
	SuccessCampaign []interface{} `json:"success_campaign" bson:"success_campaign"` //[{name: ,fb_id:}]
	SuccessAdset    []struct {
		Name string `json:"name" bson:"name"`
		FbId string `json:"fb_id" bson:"fb_id"`
	} `json:"success_adset" bson:"success_adset"`
	SuccessAd []struct {
		Name string `json:"name" bson:"name"`
		FbId string `json:"fb_id" bson:"fb_id"`
	} `json:"success_ad" bson:"success_ad"`
	FailCampaign []interface{} `json:"fail_campaign" bson:"fail_campaign"`
	FailAdset    []interface{} `json:"fail_adset" bson:"fail_adset"`
	FailAd       []interface{} `json:"fail_ad" bson:"fail_ad"`
	Fbid         string        `json:"fbid" bson:"fbid"`
	Type         string        `json:"type" bson:"type"`                   //创建类型，0，1， 2， （向目标广告组中新建广告， 复制目标广告组参数并新建广告组， 手动新建广告组）
	CampaignType string        `json:"campaign_type" bson:"campaign_type"` //广告系列类型1，2， （1：新建广告系列， 2：使用已有广告系列）
	MainPageId   string        `json:"main_page_id" bson:"main_page_id"`
	InstagramId  string        `json:"instagram_id" bson:"instagram_id"`
	AccountId    string        `json:"account_id" bson:"account_id"`
	StartTime    string        `json:"start_time" bson:"start_time"` //now为立即执行 或指定一个具体的PST时间执行,投放排期
	EndTime      string        `json:"end_time" bson:"end_time"`     // 2020-09-02 11:11:11
	CreateUser   string        `json:"create_user" bson:"create_user"`
	UserId       int           `json:"user_id" bson:"user_id"`
	StoreUrl     string        `json:"store_url" bson:"store_url"`
	CfgFrameId   string        `json:"cfg_frame_id" bson:"cfg_frame_id"` // 广告结构方案id （v3版 新增字段）
	CompanyId    int           `json:"company_id" bson:"company_id"`
	CreateTime   *time.Time    `json:"create_time" bson:"create_time"`
	UpdateTime   *time.Time    `json:"update_time" bson:"update_time"`
	//V            int           `json:"__v"`
	TaskStatus int `json:"task_status" bson:"task_status"`
}

// MAdflowPreview 广告预览及创建配置信息 From Mongo/cruiser_console_v2/adflowpreviews
type MAdflowPreview struct {
	Id        string `json:"_id"`
	StructCfg struct {
		Target struct {
			TargetType string `json:"target_type"`
			WindowDays int    `json:"window_days"`
			EventName  string `json:"event_name"`
		} `json:"target"`
		IsCBO               bool   `json:"is_CBO"`
		Option              string `json:"option"`
		CampaignLevelLables []int  `json:"campaign_level_lables"`
		AdsetLevelLables    []int  `json:"adset_level_lables"`
		AdsetMaxCount       int    `json:"adset_max_count"`
		AdMaxCount          int    `json:"ad_max_count"`
		BudgetLimit         int    `json:"budget_limit"`
		AttributionSpec     struct {
			EventType  string `json:"event_type"`
			WindowDays int    `json:"window_days"`
		} `json:"attribution_spec"`
		OptimizationGoal string   `json:"optimization_goal"`
		CustomEventType  string   `json:"custom_event_type"`
		BillingEvent     string   `json:"billing_event"`
		PacingType       []string `json:"pacing_type"`
	} `json:"struct_cfg"`
	CampaignDatas []interface{} `json:"campaign_datas"`
	Fbid          string        `json:"fbid"`
	AccountId     string        `json:"account_id"`
	UpdateTime    string        `json:"update_time"`
	CreateTime    string        `json:"create_time"`
	V             int           `json:"__v"`
}
