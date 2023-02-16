package bean

import "time"

// MCfgAudienceModel 受众组 数据模型
type MCfgAudienceModel struct {
	ID                      int                       `json:"_id" bson:"_id"`
	CompanyId               int                       `json:"company_id" bson:"company_id"`                               //公司id
	Name                    string                    `json:"name" bson:"name"`                                           // 受众组名称
	AccountId               string                    `json:"account_id" bson:"account_id"`                               // 账户id
	CustomAudiences         []CustomAudiences         `json:"custom_audiences" bson:"custom_audiences"`                   // 包含受众
	ExcludedCustomAudiences []ExcludedCustomAudiences `json:"excluded_custom_audiences" bson:"excluded_custom_audiences"` //排除受众
	FlexibleSpec            []FlexibleSpec            `json:"flexible_spec" bson:"flexible_spec"`                         // 包含细分定位、缩小细分定位
	Exclusions              Exclusions                `json:"exclusions" bson:"exclusions"`                               // 排除细分定位
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
	Interests []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"interests"`
}

type FlexibleSpec struct {
	FamilyStatuses []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"family_statuses"`
}
