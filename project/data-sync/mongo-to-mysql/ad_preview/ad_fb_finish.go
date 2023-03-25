package ad_preview

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_preview/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdFbFinish() {
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgaudiences")
	collAdFlowCfg := db.Collection("adflowcfg")

	// 2、从mongo查询数据
	mFBCampaignModel := make([]*bean.MFBCampaignModel, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mFBCampaignModel)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	//fmt.Println(mFBCampaignModel)

	// 3、将mongo数据装入切片
	//adFbFinish := make([]*bean.AdFbFinish, 0)
	for i, model := range mFBCampaignModel {
		fmt.Println("广告系列", i)

		// ID
		//campaignId, err := strconv.Atoi(model.CampaignId)
		//if err != nil {
		//	fmt.Println("model.CampaignId 转 int err：", err)
		//}
		adFbFinishCampaignModel := &bean.AdFbFinish{
			//ID:        int32(campaignId),
			ParentID:  0,
			Type:      1,
			Name:      model.CampaignName,
			AccountID: model.AccountId,
			//HTTPCode:   0,
			//Data:       "",
			FbID:       model.CampaignId,
			AdConfigID: 0,
			CreatedAt:  model.CreateTime.Unix(),
			UpdatedAt:  model.UpdateTime.Unix(),
			IsDeleted:  false,
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_fb_finish").Create(adFbFinishCampaignModel).Error
		if err != nil {
			fmt.Println("入mysql/ad_fb_finish 错误：", err)
		}

		// 根据 campaign_id 查询 adflowcfg
		mAdFlowCfg := make([]*bean.MAdFlowCfg, 0)
		err = collAdFlowCfg.Find(context.TODO(), bson.M{"campaign_id": model.CampaignId}).All(&mAdFlowCfg)
		if err != nil {
			fmt.Println("Mongo/adflowcfg 查询错误：", err)
			return
		}
		if len(mAdFlowCfg) == 0 {
			continue
		}
		// IsWifi
		isWifi := 1
		if mAdFlowCfg[0].IsWifi == false {
			isWifi = 2
		}
		adPreview := bean.AdPreview{
			//ID:            0,
			GameID:        "",
			AppID:         0,
			AccountID:     model.AccountId,
			FbAppID:       0,
			FbHomePageID:  "",
			InstagramID:   "",
			BudgetLevel:   0,
			MarketURL:     "",
			UserOs:        mAdFlowCfg[0].UserOs,
			IncludeDevice: "",
			ExcludeDevice: "",
			IsWifi:        int32(isWifi),
			CreatedAt:     mAdFlowCfg[0].CreateTime.Unix(),
			UpdatedAt:     mAdFlowCfg[0].UpdateTime.Unix(),
			IsDeleted:     0,
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_preview").Create(adPreview).Error
		if err != nil {
			fmt.Println("入mysql/ad_preview 错误：", err)
		}
		adPreviewCampaign := bean.AdPreviewCampaign{
			ID:        0,
			PreviewID: 0,
			SchemeID:  0,
			Name:      "",
			Budget:    0,
			CreatedAt: 0,
			UpdatedAt: 0,
		}
		// 4、将装有mongo数据的切片入库
		err = db2.MySQLClientCruiser.Table("ad_preview_campaign").Create(adPreviewCampaign).Error
		if err != nil {
			fmt.Println("入mysql/ad_preview_campaign 错误：", err)
		}
	}
}
