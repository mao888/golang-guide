package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

func RunArtAttachment2() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artattachments")
	collArtneeds := db.Collection("artneeds")

	// 2、从mongo查询数据
	mArtAttachments := make([]*MArtAttachments, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtAttachments)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mArtAttachments)

	// 3、将mongo数据装入切片
	//artAttachments := make([]*ArtAttachment, 0)
	for _, attachment := range mArtAttachments {

		//fmt.Println(attachment.AssetLanguage)

		// 根据 需求id查询 状态是否为 4: 已完成
		mArtNeeds := make([]*MArtNeeds, 0)
		err := collArtneeds.Find(context.TODO(), bson.M{"_id": attachment.ArtneedId}).All(&mArtNeeds)
		if err != nil {
			fmt.Println("Mongo/artneeds查询错误：", err)
			return
		}
		if len(mArtNeeds) == 0 {
			continue
		}
		// 需求已完成则跳过
		if mArtNeeds[0].Status == 3 {
			continue
		}

		fmt.Println("未完成需求id：", attachment.ArtneedId)

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
			TaskID: 0,
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
			fmt.Println(attachment.ID, "---", attachment.ArtneedId, "入库 art_attachments 错误", err)
			continue
		}
	}
}
