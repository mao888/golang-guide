package bean

import "time"

// MDimPermission From Mongo/rambler/dimpermissions
type MDimPermission struct {
	Id          string     `json:"_id" bson:"_id"`
	ChildSystem string     `json:"child_system" bson:"child_system"` // cruiser_v2、pandora、bi、developer、arlington、art_needs
	UserId      int        `json:"user_id" bson:"user_id"`
	CompanyId   int        `json:"company_id" bson:"company_id"` // 公司id
	AppIds      []int      `json:"app_ids" bson:"app_ids"`       // 用户id
	AccountIds  []string   `json:"account_ids" bson:"account_ids"`
	Platforms   []string   `json:"platforms" bson:"platforms""`
	CreateTime  *time.Time `bson:"create_time" json:"create_time"`
	UpdateTime  *time.Time `bson:"update_time" json:"update_time"`
}

// MPlatUser From Mongo/plat_console/platusers
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
