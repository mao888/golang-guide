package bean

import "time"

// MActiveLibrary 动作库
type MActiveLibrary struct {
	ID               int32      `json:"_id" bson:"_id"`
	CompanyId        int32      `json:"company_id" bson:"company_id"`
	Name             string     `json:"name" bson:"name"`
	Desc             string     `json:"desc" bson:"desc"`
	Size             string     `json:"size" bson:"size"`
	CategoryId       int32      `json:"category_id" bson:"category_id"`
	TagArr           []int32    `json:"tag_arr" bson:"tagArr"`                       //标签，上限10
	ThumbUrl         string     `json:"thumb_url" bson:"thumb_url"`                  //预览图
	Url              string     `json:"url" bson:"url"`                              // 视频URL
	CreatorId        int32      `json:"creator_id" bson:"creator_id"`                //上传用户id
	Author           int32      `json:"author" bson:"author"`                        // 作者
	GameId           int32      `json:"game_id" bson:"game_id"`                      // 游戏id
	UeDownloadUrl    string     `json:"ue_download_url" bson:"ue_download_url"`      //UE下载地址
	MayaDownloadUrl  string     `json:"maya_download_url" bson:"maya_download_url"`  //Maya下载地址
	DoneTime         *time.Time `json:"done_time" bson:"done_time"`                  // 完成时间
	RelationArtneeds []string   `json:"relation_artneeds" bson:"relation_artneeds""` // 关联需求
	DeletedTime      *time.Time `json:"deleted_time" bson:"deleted_time"`            // 删除状态
	CreateTime       *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime       *time.Time `json:"update_time" bson:"update_time"`
}

// MArtSource 资产库
type MArtSource struct {
	ID               int32      `json:"_id" bson:"_id"`
	CompanyId        int32      `json:"company_id" bson:"company_id"`
	Name             string     `json:"name" bson:"name"`
	Desc             string     `json:"desc" bson:"desc"`
	CategoryId       int32      `json:"category_id" bson:"category_id"`
	TagArr           []int32    `json:"tag_arr" bson:"tagArr"`                       //标签，上限10
	ThumbArr         []int32    `json:"thumb_arr" bson:"thumb_arr"`                  //预览图，上限20
	DownloadCount    int32      `json:"download_count" bson:"download_count"`        // 下载次数
	CreatorId        int32      `json:"creator_id" bson:"creator_id"`                //上传用户id
	Author           int32      `json:"author" bson:"author"`                        // 作者
	GameId           int32      `json:"game_id" bson:"game_id"`                      // 游戏id
	DownloadUrl      string     `json:"download_url" bson:"download_url"`            //UE下载地址
	MayaDownloadUrl  string     `json:"maya_download_url" bson:"maya_download_url"`  //Maya下载地址
	DoneTime         *time.Time `json:"done_time" bson:"done_time"`                  // 完成时间
	RelationArtneeds []string   `json:"relation_artneeds" bson:"relation_artneeds""` // 关联需求
	CreateTime       *time.Time `json:"create_time" bson:"create_time"`
	UpdateTime       *time.Time `json:"update_time" bson:"update_time"`
}
