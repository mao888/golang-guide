package bean

// ArtAsset 美术资产库表，该表保护（美术资产，动作资产，音乐资产等) mapped from table cruiser_console <art_asset>
type ArtAsset struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Type        int32  `gorm:"column:type;not null;default:1" json:"type"`         // 美术资产类型（1：美术资产、2：动作资产、3：音乐资产）
	AuthorID    int32  `gorm:"column:author_id" json:"author_id"`                  // 人员ID，作者
	Desc        string `gorm:"column:desc" json:"desc"`                            // 资源描述 空
	Name        string `gorm:"column:name;not null" json:"name"`                   // 资源名称 name + " " + desc
	MainURL     string `gorm:"column:main_url" json:"main_url"`                    // 主图URL地址
	UeURL       string `gorm:"column:ue_url;not null" json:"ue_url"`               // UE下载地址
	MayaURL     string `gorm:"column:maya_url;not null" json:"maya_url"`           // Maya下载地址
	Remark      string `gorm:"column:remark;not null" json:"remark"`               // 备注
	GameID      string `gorm:"column:game_id" json:"game_id"`                      // 所属游戏ID
	IsInternal  bool   `gorm:"column:is_internal;not null" json:"is_internal"`     // 0:共享 1:内部
	CategorieID int32  `gorm:"column:categorie_id" json:"categorie_id"`            // 所属分类ID
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建日期
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新日期
	DoneAt      int64  `gorm:"column:done_at" json:"done_at"`                      // 完成时间
	IsDeleted   bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 1: deleted, 0: normal
}

// AssetImg 资源库-主图表 mapped from table cruiser_console <asset_imgs>
type AssetImg struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID   int32  `gorm:"column:asset_id;not null" json:"asset_id"`     // 附属资产id
	Name      string `gorm:"column:name;not null" json:"name"`             // 资产图片名称
	URL       string `gorm:"column:url;not null" json:"url"`               // 资产图片地址
	SizeRatio string `gorm:"column:size_ratio;not null" json:"size_ratio"` // 资产图片尺寸比例
	Size      int32  `gorm:"column:size;not null" json:"size"`             // 图片大小
	Md5       string `gorm:"column:md5;not null" json:"md5"`               // 图片md5
	Height    int32  `gorm:"column:height;not null" json:"height"`         // 图片高度
	Width     int32  `gorm:"column:width;not null" json:"width"`           // 图片宽度
	IsDeleted bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`
}

// ArtAssetAction mapped from table cruiser_console <art_asset_action>
type ArtAssetAction struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// ArtAssetActionCategory 资产动作分类表 mapped from table cruiser_console <art_asset_action_category>
type ArtAssetActionCategory struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int32  `gorm:"column:parent_id;not null" json:"parent_id"`         // 上级分类
	Name      string `gorm:"column:name;not null" json:"name"`                   // 分类名称
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建日期
	Remark    string `gorm:"column:remark;not null" json:"remark"`               // 备注
	Order     int32  `gorm:"column:order;not null" json:"order"`                 // 位置
}

// ArtAssetActionCategoryRelation mapped from table cruiser_console <art_asset_action_category_relations>
type ArtAssetActionCategoryRelation struct {
	ID         int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID    int32 `gorm:"column:asset_id" json:"asset_id"`
	CategoryID int32 `gorm:"column:category_id" json:"category_id"`
}

// ArtAssetActionRelation mapped from table cruiser_console <art_asset_action_relations>
type ArtAssetActionRelation struct {
	ID       int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID  int32 `gorm:"column:asset_id;not null" json:"asset_id"`
	ActionID int32 `gorm:"column:action_id;not null" json:"action_id"`
}

// ArtAssetCategory 资产分类表 mapped from table cruiser_console <art_asset_category>
type ArtAssetCategory struct {
	ID        int32  `gorm:"column:id;primaryKey" json:"id"`
	ParentID  int32  `gorm:"column:parent_id;not null" json:"parent_id"`         // 上级分类
	Name      string `gorm:"column:name;not null" json:"name"`                   // 分类名称
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建日期
	Remark    string `gorm:"column:remark;not null" json:"remark"`               // 备注
	Order     int32  `gorm:"column:order;not null" json:"order"`                 // 位置
}

// ArtAssetCategoryRelation mapped from table cruiser_console <art_asset_category_relations>
type ArtAssetCategoryRelation struct {
	ID         int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID    int32 `gorm:"column:asset_id" json:"asset_id"`
	CategoryID int32 `gorm:"column:category_id;not null" json:"category_id"`
}

// ArtAssetStyle 美术素材风格 mapped from table cruiser_console <art_asset_style>
type ArtAssetStyle struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// ArtAssetStyleRelation mapped from table cruiser_console <art_asset_style_relations>
type ArtAssetStyleRelation struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID int32 `gorm:"column:asset_id;not null" json:"asset_id"`
	StyleID int32 `gorm:"column:style_id;not null" json:"style_id"`
}

// ArtAssetTagRelation 美术资产-标签表多对多关联表 mapped from table cruiser_console <art_asset_tag_relations>
type ArtAssetTagRelation struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AssetID int32 `gorm:"column:asset_id;not null" json:"asset_id"` // 资产id
	TagID   int32 `gorm:"column:tag_id;not null" json:"tag_id"`     //  标签id
}

// ArtAssetTag 资产标签表 mapped from table cruiser_console <art_asset_tags>
type ArtAssetTag struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Label     string `gorm:"column:label;not null" json:"label"`           // 标签名称
	Code      int32  `gorm:"column:code;not null" json:"code"`             // 字段编码 做后期排序或者其他备用
	CreatedAt int64  `gorm:"column:created_at;not null" json:"created_at"` // 创建日期
	Remark    string `gorm:"column:remark;not null" json:"remark"`         // 备注
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

// ArtNeedAssetRelation 美术需求-资产多对多关联表 mapped from table cruiser_console <art_need_asset_relations>
type ArtNeedAssetRelation struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID  int32 `gorm:"column:need_id;not null" json:"need_id"`   // 美术需求id
	AssetID int32 `gorm:"column:asset_id;not null" json:"asset_id"` // 资产id
}
