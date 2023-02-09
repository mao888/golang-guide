package art_need

import (
	"context"
	"fmt"
	"github.com/mao888/go-utils/constants"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAttachment() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artattachments")

	// 2、从mongo查询数据
	mArtAttachments := make([]*MArtAttachments, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtAttachments)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mArtAttachments)

	// 3、将mongo数据装入切片
	//artAttachments := make([]*ArtAttachment, 0)
	for _, attachment := range mArtAttachments {

		// TaskID
		// 根据 需求id查询子任务id
		artTask := make([]*ArtTask, 0)
		err = db2.MySQLClientCruiser.Table("art_tasks").
			Where("need_id = ?", attachment.ArtneedId).Find(&artTask).Error
		if err != nil {
			fmt.Println("根据 需求id查询子任务id 错误：", err)
			return
		}
		if len(artTask) == constants.NumberZero {
			continue
		}
		for _, task := range artTask {
			// FileType
			var fileType int32
			if attachment.AssetType == "file" {
				fileType = 1
			} else if attachment.AssetType == "image" {
				fileType = 2
			} else if attachment.AssetType == "video" {
				fileType = 3
			}
			// IsDeleted
			isDeleted := false
			if attachment.DeleteTime != nil {
				isDeleted = true
			}
			artAttachment := &ArtAttachment{
				ID:     attachment.ID,
				NeedID: attachment.ArtneedId,
				TaskID: task.ID,
				//LogID:     0,
				Type:      attachment.Type + 1,
				Name:      attachment.AssetName,
				URL:       attachment.AssetUrlInfo,
				SizeRatio: attachment.AssetSize,
				Size:      attachment.FileSize,
				Md5:       attachment.AssetMd5,
				Height:    attachment.AssetHeight,
				Width:     attachment.AssetWidth,
				FileType:  fileType,
				IsDeleted: isDeleted,
			}
			// 入库 art_attachments
			err = db2.MySQLClientCruiser.Table("art_attachments").Create(artAttachment).Error
			if err != nil {
				fmt.Println("入库 art_attachments 错误", err)
				return
			}
		}
	}

}
