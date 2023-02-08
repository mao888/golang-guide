package art_need

// ArtNeed 美术需求主表 mapped from table cruiser_console <art_need>
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

// ArtNeedTagRelation 美术需求标签多对多关联表 mapped from table cruiser_console <art_need_tag_relations>
type ArtNeedTagRelation struct {
	ID     int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID int32 `gorm:"column:need_id;not null" json:"need_id"` // 需求id
	TagID  int32 `gorm:"column:tag_id;not null" json:"tag_id"`   // 标签id
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

// ArtNeedRelation 美术需求关联需求多对多关联表 mapped from table cruiser_console <art_need_relations>
type ArtNeedRelation struct {
	ID             int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MainNeedID     int32 `gorm:"column:main_need_id;not null" json:"main_need_id"`         // 主需求id
	RelationNeedID int32 `gorm:"column:relation_need_id;not null" json:"relation_need_id"` // 关联需求id
}

// ArtNeedPersonRelation 美术需求创意负责人多对多关联表 mapped from table cruiser_console <art_need_person_relations>
type ArtNeedPersonRelation struct {
	ID       int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID   int32 `gorm:"column:need_id;not null" json:"need_id"`     // 需求id
	PersonID int32 `gorm:"column:person_id;not null" json:"person_id"` // 创意人id
	Weight   int32 `gorm:"column:weight;not null" json:"weight"`       // 创意人权重
}

// ArtNeedMaterialSizeRelation 美术需求素材尺寸多对多关联表 mapped from table cruiser_console <art_need_material_size_relations>
type ArtNeedMaterialSizeRelation struct {
	ID             int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID         int32 `gorm:"column:need_id;not null" json:"need_id"`                   // 需求id
	MaterialSizeID int32 `gorm:"column:material_size_id;not null" json:"material_size_id"` // 尺寸id
}

// ArtNeedLog 美术需求操作日志表 mapped from table cruiser_console <art_need_logs>
type ArtNeedLog struct {
	ID         int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID     int32  `gorm:"column:need_id;not null" json:"need_id"`             // 需求id
	OpUser     int32  `gorm:"column:op_user;not null" json:"op_user"`             // 操作人
	OpUsername string `gorm:"column:op_username;not null" json:"op_username"`     // 操作人
	OpAction   string `gorm:"column:op_action;not null" json:"op_action"`         // 操作动作
	OpSubject  string `gorm:"column:op_subject;not null" json:"op_subject"`       // 操作主题
	MainDesc   string `gorm:"column:main_desc;not null" json:"main_desc"`         // 需求描述
	ExtraDesc  string `gorm:"column:extra_desc;not null" json:"extra_desc"`       // 补充描述
	CreatedAt  int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建日期
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新日期
}

// ArtNeedLanguageRelation 美术需求语种多对多关联表 mapped from table cruiser_console <art_need_language_relations>
type ArtNeedLanguageRelation struct {
	ID         int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID     int32 `gorm:"column:need_id;not null" json:"need_id"`         // 美术需求id
	LanguageID int32 `gorm:"column:language_id;not null" json:"language_id"` // 语言id
}

// ArtNeedKeywordRelation 美术需求关键词多对多关联表 mapped from table cruiser_console <art_need_keyword_relations>
type ArtNeedKeywordRelation struct {
	ID        int32 `gorm:"column:id;primaryKey" json:"id"`
	NeedID    int32 `gorm:"column:need_id;not null" json:"need_id"`       // 需求id
	KeywordID int32 `gorm:"column:keyword_id;not null" json:"keyword_id"` // 关键词id
}

// ArtNeedAssetRelation 美术需求-资产表多对多关联表 mapped from table cruiser_console <art_need_asset_relations>
type ArtNeedAssetRelation struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID  int32 `gorm:"column:need_id;not null" json:"need_id"`   // 美术需求id
	AssetID int32 `gorm:"column:asset_id;not null" json:"asset_id"` // 资产id
}

// ArtTask 美术需求子任务表 mapped from table cruiser_console <art_tasks>
type ArtTask struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskType  int32  `gorm:"column:task_type;not null;default:1" json:"task_type"` // 子任务类型编号（1：分镜、2：场景、3：模型、4：绑定:5：动画、6：特效、7：渲染、8：平面、9：合成）
	NeedID    int32  `gorm:"column:need_id;not null" json:"need_id"`               // 需求id
	PersonID  int32  `gorm:"column:person_id;not null" json:"person_id"`           // 人员ID，设计负责人
	Desc      string `gorm:"column:desc;not null" json:"desc"`                     // 任务描述
	Status    int32  `gorm:"column:status;not null;default:1" json:"status"`       // 子任务状态（1：待分配、2：制作中、3：制作完成）
	UeURL     string `gorm:"column:ue_url;not null" json:"ue_url"`                 // UE下载地址
	MayaURL   string `gorm:"column:maya_url;not null" json:"maya_url"`             // Maya下载地址
	Remark    string `gorm:"column:remark;not null" json:"remark"`                 // 备注
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 创建日期
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 更新日期
	IsDeleted bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 1: deleted, 0: normal
}

// ArtTaskSchedule 美术需求子任务实际排期表 mapped from table cruiser_console <art_task_schedule>
type ArtTaskSchedule struct {
	ID          int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	NeedID      int32 `gorm:"column:need_id;not null" json:"need_id"`             // 需求ID
	TaskID      int32 `gorm:"column:task_id;not null" json:"task_id"`             // 子任务id
	StartDateAt int64 `gorm:"column:start_date_at;not null" json:"start_date_at"` // 开始日期
	EndDateAt   int64 `gorm:"column:end_date_at;not null" json:"end_date_at"`     //  结束日期
	WorkHour    int32 `gorm:"column:work_hour;not null" json:"work_hour"`         // 实际工时
}
