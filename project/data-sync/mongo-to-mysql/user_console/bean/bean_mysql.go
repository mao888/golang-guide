package bean

// Permission 权限表 mapped from table user_console <permission>
type Permission struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`      // 主键id
	Code      string `gorm:"column:code;not null;default:''" json:"code"`            // 权限码
	Name      string `gorm:"column:name;not null;default:''" json:"name"`            // 权限名称
	Type      int32  `gorm:"column:type;not null;default:0" json:"type"`             // 权限类型: 0子系统默认全部权限 1页面 2操作 3URL
	SystemID  int32  `gorm:"column:system_id;not null;default:0" json:"system_id"`   // 系统id,0表示不属于某个系统
	Creator   string `gorm:"column:creator;not null" json:"creator"`                 // 创建人
	UpdatedAt int32  `gorm:"column:updated_at;not null" json:"updated_at"`           // 更新时间
	CreatedAt int32  `gorm:"column:created_at;not null" json:"created_at"`           // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null;default:0" json:"is_deleted"` // 是否删除（0:否，1:是）
}

// UserPerm 员工权限关联表 mapped from table user_console <user_perm>
type UserPerm struct {
	ID        int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // 主键id
	UserID    int32 `gorm:"column:user_id;not null" json:"user_id"`               // 员工id
	PermID    int32 `gorm:"column:perm_id;not null" json:"perm_id"`               // 权限id
	PolicyID  int32 `gorm:"column:policy_id;not null;default:0" json:"policy_id"` // 策略id,0表示没有策略，即任意
	ScopeID   int32 `gorm:"column:scope_id;not null;default:0" json:"scope_id"`   // 资源范围policy id,0表示没有范围，即任意
	UpdatedAt int32 `gorm:"column:updated_at;not null" json:"updated_at"`         // 更新时间
	CreatedAt int32 `gorm:"column:created_at;not null" json:"created_at"`         // 创建时间
}

// Policy 策略表 mapped from table user_console <policy>
type Policy struct {
	ID        int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	UpdatedAt int32 `gorm:"column:updated_at;not null" json:"updated_at"`      // 更新时间
	CreatedAt int32 `gorm:"column:created_at;not null" json:"created_at"`      // 创建时间
}

// PolicyResource 策略资源关联表 mapped from table user_console <policy_resource>
type PolicyResource struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	PolicyID   int32  `gorm:"column:policy_id;not null" json:"policy_id"`        // 策略id
	ResourceID int32  `gorm:"column:resource_id;not null" json:"resource_id"`    // 资源id
	EntityID   string `gorm:"column:entity_id;not null" json:"entity_id"`        // 资源实体id
	UpdatedAt  int32  `gorm:"column:updated_at;not null" json:"updated_at"`      // 更新时间
	CreatedAt  int32  `gorm:"column:created_at;not null" json:"created_at"`      // 创建时间
}

// Resource 资源表 mapped from table user_console <resource>
type Resource struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`      // 主键id
	SystemID  int32  `gorm:"column:system_id;not null" json:"system_id"`             // 系统id
	Code      string `gorm:"column:code;not null;default:''" json:"code"`            // 资源码
	Name      string `gorm:"column:name;not null;default:''" json:"name"`            // 资源名称
	Creator   string `gorm:"column:creator;not null;default:''" json:"creator"`      // 创建人
	UpdatedAt int32  `gorm:"column:updated_at;not null" json:"updated_at"`           // 更新时间
	CreatedAt int32  `gorm:"column:created_at;not null" json:"created_at"`           // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null;default:0" json:"is_deleted"` // 是否删除（0:否，1:是）
}

// User 员工表 mapped from table user_console <user>
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

// App mapped from table application_console <app>
type App struct {
	ID         int32  `gorm:"column:id;not null" json:"id"`                               // 应用ID
	GameID     string `gorm:"column:game_id;not null" json:"game_id"`                     // 游戏ID
	Platform   int32  `gorm:"column:platform;not null;default:1" json:"platform"`         // 平台 1 IOS 2 Android 3 Web
	AppAlias   string `gorm:"column:app_alias;not null" json:"app_alias"`                 // 应用别名
	Package    string `gorm:"column:package;not null" json:"package"`                     // 包
	AppStoreID string `gorm:"column:app_store_id;not null;default:0" json:"app_store_id"` // 苹果商店ID
	Icon       string `gorm:"column:icon" json:"icon"`
	Market     string `gorm:"column:market" json:"market"`                        // 应用市场渠道
	MarketURL  string `gorm:"column:market_url" json:"market_url"`                // 应用市场链接
	CreatorID  int32  `gorm:"column:creator_id" json:"creator_id"`                // 创建人ID
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间戳（秒）
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间戳（秒）
	IsDeleted  bool   `gorm:"column:is_deleted" json:"is_deleted"`                // 是否删除（0：未删除，1：已删除）
}
