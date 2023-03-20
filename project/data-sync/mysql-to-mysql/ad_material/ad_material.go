package ad_material

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

// AdMaterial mapped from table <ad_material>
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
	CreatedAt    int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`   // 创建日期
	UpdatedAt    int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`   // 更新日期
	IsDeleted    bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`         // 1: deleted, 0: normal
	Src          int32  `gorm:"column:src;not null;default:1" json:"src"`             // 1:美术需求2：素材中心上传
	ExtraName    string `gorm:"column:extra_name" json:"extra_name"`                  // 素材扩展名
	GameID       string `gorm:"column:game_id;not null" json:"game_id"`               // 所属游戏ID
	TagID        int32  `gorm:"column:tag_id;not null" json:"tag_id"`                 //  标签id 美术需求的dictionaries.id
}

func RunAdMaterialMysqlToMysql() {

	var material = make([]*AdMaterial, 0)
	err := db2.MySQLClientCruiser.Table("ad_material").Order("updated_at asc").Find(&material)
	if err != nil {
		fmt.Println("查询错误：", err)
		return
	}
	for i, adMaterial := range material {
		fmt.Println("id: ", i)
		err := db2.MySQLClientCruiserTest.Table("ad_material").Create(&adMaterial).Error
		if err != nil {
			fmt.Println("入mysql错误：", err)
			return
		}
		if i == 199 {
			break
		}
	}
}
