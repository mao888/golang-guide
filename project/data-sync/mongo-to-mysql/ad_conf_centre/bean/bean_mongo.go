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
	VerifyStr               string                    `json:"verify_str"`                                                 // 受众信息校验参数
	CreateUser              string                    `json:"create_user"`                                                // 创建人
	UserId                  int32                     `json:"user_id"`                                                    // 创建人id
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
	EducationStatuses []int `json:"education_statuses"`
	Interests         []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"interests"`
	CollegeYears         []int `json:"college_years"`
	RelationshipStatuses []int `json:"relationship_statuses"`
	Income               []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"income"`
	FamilyStatuses []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"family_statuses"`
	Behaviors []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"behaviors"`
}

type FlexibleSpec struct {
	EducationStatuses []int `json:"education_statuses"` // 教育程度
	Interests         []struct { // 兴趣
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"interests"`
	CollegeYears         []int `json:"college_years"`         // 大学毕业时间
	RelationshipStatuses []int `json:"relationship_statuses"` // 感情状况
	Income               []struct { // 收入
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"income"`
	FamilyStatuses []struct { // 家庭状态
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"family_statuses"`
	Behaviors []struct { // 行为
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"behaviors"`
}

// MCfgFrame 结构方案数据模型
type MCfgFrame struct {
	Id               int32    `json:"_id" bson:"_id"`                             // 结构方案id
	CompanyId        int32    `json:"company_id" bson:"company_id"`               // 公司id
	Name             string   `json:"name"`                                       // 结构方案名称
	CampaignDims     []string `json:"campaign_dims"`                              // campaign划分维度 ['countries', 'audiences', 'positions', 'age', 'sex', 'language', 'materials', 'tag', 'adtype']
	AdsetDims        []string `json:"adset_dims" bson:"adset_dims"`               // adset划分维度
	CampaignLimit    int      `json:"campaign_limit" bson:"campaign_limit"`       // campaign数量上限
	AdsetLimit       int      `json:"adset_limit" bson:"adset_limit"`             // adset数量上限
	IsCbo            bool     `json:"is_cbo" bson:"is_cbo"`                       // 是否开启cbo
	OptimizationGoal string   `json:"optimization_goal" bson:"optimization_goal"` // 优化目标
	BidStrategy      string   `json:"bid_strategy" bson:"bid_strategy"`           // 竞价策略
	AttributionSpec  struct {
		EventType  string `json:"event_type" bson:"event_type"`
		WindowDays int    `json:"window_days" bson:"window_days"`
	} `json:"attribution_spec" bson:"attribution_spec"`                            // 转化窗口
	CustomEventType string     `json:"custom_event_type" bson:"custom_event_type"` // AEO的14个应用事件
	BillingEvent    string     `json:"billing_event" bson:"billing_event"`         // 计费方式
	PacingType      []string   `json:"pacing_type"`
	BudgetLimit     int        `json:"budget_limit" bson:"budget_limit"` // 单日预算上限
	VerifyStr       string     `json:"verify_str" bson:"verify_str"`     // 结构方案信息校验参数
	CreateUser      string     `json:"create_user" bson:"create_user"`   // 创建人
	UserId          int32      `json:"user_id" bson:"user_id"`           // 创建人
	SourceSys       string     `json:"source_sys" bson:"source_sys""`
	CreateTime      *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime      *time.Time `json:"update_time" bson:"update_time"`
}
