package bean

// AdMaterial 广告素材主表（它的ID会社交关联到广告素材tag，尺寸，语言，负责人等关联表) mapped from table cruiser_console <ad_material>
type AdMaterial struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Type         int32  `gorm:"column:type;not null;default:1" json:"type"`           // 附件文件类型， 1: file,  2: image,3: video
	NeedID       string `gorm:"column:need_id;not null" json:"need_id"`               // 需求id
	Name         string `gorm:"column:name;not null" json:"name"`                     // 素材名称 拼接而成
	Title        string `gorm:"column:title;not null" json:"title"`                   // 素材标题
	URL          string `gorm:"column:url;not null" json:"url"`                       // 素材地址
	YtURL        string `gorm:"column:yt_url;not null" json:"yt_url"`                 // youtube素材地址
	SizeRationID int32  `gorm:"column:size_ration_id;not null" json:"size_ration_id"` // 素材尺寸表 Id
	Size         int32  `gorm:"column:size;not null" json:"size"`                     // 素材大小
	Md5          string `gorm:"column:md5;not null" json:"md5"`                       // 素材md5
	Duration     int32  `gorm:"column:duration;not null" json:"duration"`             // 素材时长
	Remark       string `gorm:"column:remark;not null" json:"remark"`                 // 备注
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`         // 创建日期
	UpdatedAt    int64  `gorm:"column:updated_at;not null" json:"updated_at"`         // 更新日期
	IsDeleted    bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 1: deleted, 0: normal
	Src          int32  `gorm:"column:src;not null;default:1" json:"src"`             // 1:美术需求2：素材中心上传
	ExtraName    string `gorm:"column:extra_name" json:"extra_name"`                  // 素材扩展名
	GameID       string `gorm:"column:game_id;not null" json:"game_id"`               // 所属游戏ID
	TagID        int32  `gorm:"column:tag_id;not null" json:"tag_id"`                 //  标签id 美术需求的dictionaries.id
}

// AdMaterialLanguageRelation 广告素材语言关联表-语言表多对多关联表 mapped from table cruiser_console <ad_material_language_relations>
type AdMaterialLanguageRelation struct {
	ID         int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterialID int32 `gorm:"column:material_id;not null" json:"material_id"` // 素材id
	LanguageID int32 `gorm:"column:language_id;not null" json:"language_id"` // 语言id
}

// AdMaterialPersonRelation 广告素材人员关联表-人员表多对多关联表 mapped from table cruiser_console <ad_material_person_relations>
type AdMaterialPersonRelation struct {
	ID         int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterialID int32 `gorm:"column:material_id;not null" json:"material_id"` // 素材id
	PersonID   int32 `gorm:"column:person_id;not null" json:"person_id"`     // 人员id
	Type       int32 `gorm:"column:type;not null" json:"type"`               // 人员类型，1：创意负责人，2：设计负责人
}

// AdMaterialSyncLog 广告素材 上传同步 记录表 mapped from table cruiser_console <ad_material_sync_log>
type AdMaterialSyncLog struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterilaType int32  `gorm:"column:materila_type;not null;default:1" json:"materila_type"` // '附件文件类型， 1: file,  2: image,3: video',
	MaterialID   int32  `gorm:"column:material_id;not null" json:"material_id"`               // 素材id
	Name         string `gorm:"column:name;not null" json:"name"`                             // 素材名称 拼接而成
	URL          string `gorm:"column:url;not null" json:"url"`                               // 素材源地址
	MaterialMd5  string `gorm:"column:material_md5;not null" json:"material_md5"`             // 素材md5
	AccountID    string `gorm:"column:account_id;not null" json:"account_id"`                 // 所属账户
	ErrorMessage string `gorm:"column:error_message;not null" json:"error_message"`           // 错误信息
	Creator      int32  `gorm:"column:creator;not null" json:"creator"`                       // 创建者
	CreatedAt    int64  `gorm:"column:created_at;not null" json:"created_at"`                 // 同步创建时间
	Type         int32  `gorm:"column:type;not null;default:1" json:"type"`                   // 上传日志类型 1：Facebook 2：YouTube
	Status       int32  `gorm:"column:status;not null;default:1" json:"status"`               // 上传日志状态 1：等待  2:上传中 3:成功 4：失败
	DoneAt       int64  `gorm:"column:done_at;not null" json:"done_at"`                       // 同步完成时间
	BatchID      string `gorm:"column:batch_id;not null" json:"batch_id"`                     // 批处理ID
}

// AdMaterialSyncSuccess 广告素材 上传同步 返回对照表 mapped from table cruiser_console <ad_material_sync_success>
type AdMaterialSyncSuccess struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MaterilaType int32  `gorm:"column:materila_type;not null;default:1" json:"materila_type"` // '附件文件类型， 1: file,  2: image,3: video',
	MaterialID   int32  `gorm:"column:material_id;not null" json:"material_id"`               // 素材id
	Name         string `gorm:"column:name;not null" json:"name"`                             // 素材名称 拼接而成
	URL          string `gorm:"column:url;not null" json:"url"`                               // 素材源地址
	MaterialMd5  string `gorm:"column:material_md5;not null" json:"material_md5"`             // 素材md5
	AccountID    string `gorm:"column:account_id;not null" json:"account_id"`                 // 所属账户
	Creator      int32  `gorm:"column:creator;not null" json:"creator"`                       // 创建者
	Type         int32  `gorm:"column:type;not null;default:1" json:"type"`                   // 上传日志类型 1：Facebook 2：YouTube
	SuccessID    string `gorm:"column:success_id;not null" json:"success_id"`                 // fb 返回 结果id
	BatchID      string `gorm:"column:batch_id;not null" json:"batch_id"`                     // 批处理id
}

// Dictionary 系统码表 mapped from table <dictionaries>
type Dictionary struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Label     string `gorm:"column:label;not null" json:"label"`   // 字段名称
	Code      string `gorm:"column:code;not null" json:"code"`     // 字段编码
	Type      int32  `gorm:"column:type;not null" json:"type"`     // 类型 1: keyword, 2: material_size, 3 task_type, 4: tag
	Remark    string `gorm:"column:remark;not null" json:"remark"` // 备注
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

// ArtNeed 美术需求主表 mapped from table cruiser_console <art_needs>
type ArtNeed struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`    // 需求id
	GameID       string `gorm:"column:game_id;not null" json:"game_id"`               // 游戏ID
	Title        string `gorm:"column:title;not null" json:"title"`                   // 需求标题
	Name         string `gorm:"column:name;not null" json:"name"`                     // 需求名称
	Type         int32  `gorm:"column:type;not null" json:"type"`                     // 素材类型   1: "图片", 2: "2D视频", 3: "3D视频", 4: "2D+3D视频", 5: "试玩广告", 6: "conversion图片",
	Status       int32  `gorm:"column:status;not null" json:"status"`                 // 需求状态 1：待分配 2：制作中，3：制作完成,  4: 已完成， 5： 草稿状态
	Priority     int32  `gorm:"column:priority;not null" json:"priority"`             // 需求优先级 1: 最高。2: 较高. 3: 普通，4：较低
	BaseTag      int32  `gorm:"column:base_tag;not null" json:"base_tag"`             // 基本标签 1: 原始  2： 非原始
	DescTemplate int32  `gorm:"column:desc_template;not null" json:"desc_template"`   // 模板编号 1：默认模板  2：剧情3D模板
	IsUseCruiser bool   `gorm:"column:is_use_cruiser;not null" json:"is_use_cruiser"` // 是否用于广告投放
	IsSchedule   bool   `gorm:"column:is_schedule;not null" json:"is_schedule"`       // 1：用于统计工时，0：未用于统计工时
	DoneAt       int64  `gorm:"column:done_at" json:"done_at"`                        // 需求完成日期
	ExtraDesc    string `gorm:"column:extra_desc;not null" json:"extra_desc"`         // 补充说明
	IsDeleted    bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 1: deleted, 0: normal
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 需求创建日期
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 需求更新日期
	GameName     string `gorm:"column:game_name" json:"game_name"`                    // 需求所属游戏名字（冗余字段）
	AssetRemark  string `gorm:"column:asset_remark" json:"asset_remark"`              // 关联资产备注
}

// ArtLanguage mapped from table cruiser_console <art_languages>
type ArtLanguage struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string `gorm:"column:name;not null" json:"name"`             // 语言名称
	Code      string `gorm:"column:code;not null" json:"code"`             // 语言编号
	ShortName string `gorm:"column:short_name;not null" json:"short_name"` // 语言编号
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
