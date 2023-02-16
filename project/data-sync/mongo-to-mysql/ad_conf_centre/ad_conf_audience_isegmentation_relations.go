package ad_conf_centre

import (
	"context"
	"fmt"
	"github.com/mao888/go-utils/constants"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"go.mongodb.org/mongo-driver/bson"
)

func RunAdConfAudienceIsegmentationRelation() {

	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("cfgaudiences")

	// 2、从mongo查询数据
	mCfgAudienceModel := make([]*bean.MCfgAudienceModel, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mCfgAudienceModel)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mCfgAudienceModel)

	// 3、将mongo数据装入切片
	//adConfAudience := make([]*bean.AdConfAudience, 0)
	for _, audience := range mCfgAudienceModel {

		// 包含细分定位、缩小细分定位
		for _, spec := range audience.FlexibleSpec {

			// education_statuses 教育程度
			for _, status := range spec.EducationStatuses {
				adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
					//ID:         0,
					AudienceID: audience.ID,
					FbName:     bean.EducationStatusesMap[status],
					FbType:     "education_statuses",
					FbID:       string(status),
					Type:       constants.NumberOne,
					CreatedAt:  audience.CreateTime.Unix(),
					UpdatedAt:  audience.UpdateTime.Unix(),
					Label:      "",
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
				if err != nil {
					fmt.Println("入mysql错误：", err)
				}
			}

			// interests 兴趣
			for _, interest := range spec.Interests {
				adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
					//ID:         0,
					AudienceID: audience.ID,
					FbName:     interest.Name,
					FbType:     "interests",
					FbID:       interest.Id,
					Type:       constants.NumberOne,
					CreatedAt:  audience.CreateTime.Unix(),
					UpdatedAt:  audience.UpdateTime.Unix(),
					Label:      interest.Name,
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
				if err != nil {
					fmt.Println("入mysql错误：", err)
				}
			}

			// college_years  大学毕业时间
			for _, year := range spec.CollegeYears {
				adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
					//ID:         0,
					AudienceID: audience.ID,
					FbName:     "college_years",
					FbType:     "education_statuses",
					//FbID:       interest.Id,
					Type:      constants.NumberOne,
					CreatedAt: audience.CreateTime.Unix(),
					UpdatedAt: audience.UpdateTime.Unix(),
					Label:     "大学就读年份",
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
				if err != nil {
					fmt.Println("入mysql错误：", err)
				}
				// 根据 audience.ID 将year插入ad_conf_audience.remark
			}

			// relationship_statuses 感情状况

			// income 收入

			// family_statuses 家庭状态

			// behaviors 行为

		}
	}
}
