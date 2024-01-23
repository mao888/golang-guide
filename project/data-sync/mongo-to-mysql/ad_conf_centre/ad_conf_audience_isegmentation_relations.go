package ad_conf_centre

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_conf_centre/bean"
	"github.com/mao888/mao-gutils/constants"
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
		if len(audience.FlexibleSpec) != constants.NumberZero {
			for _, spec := range audience.FlexibleSpec {

				if len(spec.EducationStatuses) != constants.NumberZero {
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
							Label:      bean.EducationStatusesMapChinese[status],
						}
						// 4、将装有mongo数据的切片入库
						err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
						if err != nil {
							fmt.Println("入mysql错误：", err)
						}
					}
				}

				// interests 兴趣
				if len(spec.Interests) != constants.NumberZero {
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
				}

				// college_years  大学毕业时间
				if len(spec.CollegeYears) != constants.NumberZero {
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
					// 根据 audience.ID 将year更新到ad_conf_audience.remark
					err = db2.MySQLClientCruiser.Table("ad_conf_audience").Where("id = ?", audience.ID).
						UpdateColumn("remark", fmt.Sprintf("%d-%d", spec.CollegeYears[0], spec.CollegeYears[len(spec.CollegeYears)-1])).Error
					if err != nil {
						fmt.Println("更新数据 错误：", err)
						return
					}
				}

				// relationship_statuses 感情状况
				if len(spec.RelationshipStatuses) != constants.NumberZero {
					for _, status := range spec.RelationshipStatuses {
						adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
							//ID:         0,
							AudienceID: audience.ID,
							FbName:     bean.RelationshipStatusesMap[status],
							FbType:     "relationship_statuses",
							FbID:       string(status),
							Type:       constants.NumberOne,
							CreatedAt:  audience.CreateTime.Unix(),
							UpdatedAt:  audience.UpdateTime.Unix(),
							Label:      bean.RelationshipStatusesMap[status],
						}
						// 4、将装有mongo数据的切片入库
						err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
						if err != nil {
							fmt.Println("入mysql错误：", err)
						}
					}
				}

				// income 收入
				if len(spec.Income) != constants.NumberZero {
					for _, s := range spec.Income {
						adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
							//ID:         0,
							AudienceID: audience.ID,
							FbName:     s.Name,
							FbType:     "income",
							FbID:       s.Id,
							Type:       constants.NumberOne,
							CreatedAt:  audience.CreateTime.Unix(),
							UpdatedAt:  audience.UpdateTime.Unix(),
							Label:      s.Name,
						}
						// 4、将装有mongo数据的切片入库
						err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
						if err != nil {
							fmt.Println("入mysql错误：", err)
						}
					}
				}

				// family_statuses 家庭状态
				if len(spec.FamilyStatuses) != constants.NumberZero {
					for _, s := range spec.FamilyStatuses {
						adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
							//ID:         0,
							AudienceID: audience.ID,
							FbName:     s.Name,
							FbType:     "family_statuses",
							FbID:       s.Id,
							Type:       constants.NumberOne,
							CreatedAt:  audience.CreateTime.Unix(),
							UpdatedAt:  audience.UpdateTime.Unix(),
							Label:      s.Name,
						}
						// 4、将装有mongo数据的切片入库
						err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
						if err != nil {
							fmt.Println("入mysql错误：", err)
						}
					}
				}

				// behaviors 行为
				if len(spec.Behaviors) != constants.NumberZero {
					for _, s := range spec.Behaviors {
						adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
							//ID:         0,
							AudienceID: audience.ID,
							FbName:     s.Name,
							FbType:     "behaviors",
							FbID:       s.Id,
							Type:       constants.NumberOne,
							CreatedAt:  audience.CreateTime.Unix(),
							UpdatedAt:  audience.UpdateTime.Unix(),
							Label:      s.Name,
						}
						// 4、将装有mongo数据的切片入库
						err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
						if err != nil {
							fmt.Println("入mysql错误：", err)
						}
					}
				}
			}
		}

		// 排除细分定位
		if audience.Exclusions != nil {
			// education_statuses 教育程度
			if len(audience.Exclusions.EducationStatuses) != constants.NumberZero {
				for _, status := range audience.Exclusions.EducationStatuses {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     bean.EducationStatusesMap[status],
						FbType:     "education_statuses",
						FbID:       string(status),
						Type:       constants.NumberTwo,
						CreatedAt:  audience.CreateTime.Unix(),
						UpdatedAt:  audience.UpdateTime.Unix(),
						Label:      bean.EducationStatusesMapChinese[status],
					}
					// 4、将装有mongo数据的切片入库
					err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
					if err != nil {
						fmt.Println("入mysql错误：", err)
					}
				}
			}

			// interests 兴趣
			if len(audience.Exclusions.Interests) != constants.NumberZero {
				for _, interest := range audience.Exclusions.Interests {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     interest.Name,
						FbType:     "interests",
						FbID:       interest.Id,
						Type:       constants.NumberTwo,
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
			}

			// college_years  大学毕业时间
			if len(audience.Exclusions.CollegeYears) != constants.NumberZero {
				adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
					//ID:         0,
					AudienceID: audience.ID,
					FbName:     "college_years",
					FbType:     "education_statuses",
					//FbID:       interest.Id,
					Type:      constants.NumberTwo,
					CreatedAt: audience.CreateTime.Unix(),
					UpdatedAt: audience.UpdateTime.Unix(),
					Label:     "大学就读年份",
				}
				// 4、将装有mongo数据的切片入库
				err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
				if err != nil {
					fmt.Println("入mysql错误：", err)
				}
				// 根据 audience.ID 将year更新到ad_conf_audience.remark
				err = db2.MySQLClientCruiser.Table("ad_conf_audience").Where("id = ?", audience.ID).
					UpdateColumn("remark", fmt.Sprintf("%d-%d", audience.Exclusions.CollegeYears[0], audience.Exclusions.CollegeYears[len(audience.Exclusions.CollegeYears)-1])).Error
				if err != nil {
					fmt.Println("更新数据 错误：", err)
					return
				}
			}

			// relationship_statuses 感情状况
			if len(audience.Exclusions.RelationshipStatuses) != constants.NumberZero {
				for _, status := range audience.Exclusions.RelationshipStatuses {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     bean.RelationshipStatusesMap[status],
						FbType:     "relationship_statuses",
						FbID:       string(status),
						Type:       constants.NumberTwo,
						CreatedAt:  audience.CreateTime.Unix(),
						UpdatedAt:  audience.UpdateTime.Unix(),
						Label:      bean.RelationshipStatusesMap[status],
					}
					// 4、将装有mongo数据的切片入库
					err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
					if err != nil {
						fmt.Println("入mysql错误：", err)
					}
				}
			}

			// income 收入
			if len(audience.Exclusions.Income) != constants.NumberZero {
				for _, s := range audience.Exclusions.Income {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     s.Name,
						FbType:     "income",
						FbID:       s.Id,
						Type:       constants.NumberTwo,
						CreatedAt:  audience.CreateTime.Unix(),
						UpdatedAt:  audience.UpdateTime.Unix(),
						Label:      s.Name,
					}
					// 4、将装有mongo数据的切片入库
					err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
					if err != nil {
						fmt.Println("入mysql错误：", err)
					}
				}
			}

			// family_statuses 家庭状态
			if len(audience.Exclusions.FamilyStatuses) != constants.NumberZero {
				for _, s := range audience.Exclusions.FamilyStatuses {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     s.Name,
						FbType:     "family_statuses",
						FbID:       s.Id,
						Type:       constants.NumberTwo,
						CreatedAt:  audience.CreateTime.Unix(),
						UpdatedAt:  audience.UpdateTime.Unix(),
						Label:      s.Name,
					}
					// 4、将装有mongo数据的切片入库
					err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
					if err != nil {
						fmt.Println("入mysql错误：", err)
					}
				}
			}

			// behaviors 行为
			if len(audience.Exclusions.Behaviors) != constants.NumberZero {
				for _, s := range audience.Exclusions.Behaviors {
					adConfAudienceIsegmentationRelation := &bean.AdConfAudienceIsegmentationRelation{
						//ID:         0,
						AudienceID: audience.ID,
						FbName:     s.Name,
						FbType:     "behaviors",
						FbID:       s.Id,
						Type:       constants.NumberOne,
						CreatedAt:  audience.CreateTime.Unix(),
						UpdatedAt:  audience.UpdateTime.Unix(),
						Label:      s.Name,
					}
					// 4、将装有mongo数据的切片入库
					err = db2.MySQLClientCruiser.Table("ad_conf_audience_isegmentation_relations").Create(adConfAudienceIsegmentationRelation).Error
					if err != nil {
						fmt.Println("入mysql错误：", err)
					}
				}
			}
		}
	}

}
