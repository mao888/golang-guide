package ad_material

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/ad_material/bean"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"sync"
)

// RunAdMaterialPersonRelation 广告素材人员关联表-人员表多对多关联表
func RunAdMaterialPersonRelation(wg sync.WaitGroup) {
	defer wg.Done()
	// 1、建立连接
	db := db2.MongoClient.Database("cruiser_console_v2")
	coll := db.Collection("assetcenters")

	// 2、从mongo查询数据
	mAssetCenter := make([]*bean.MAssetCenter, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mAssetCenter)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mAssetCenter)

	// 3、将mongo数据装入切片
	//artAsset := make([]*bean.ArtAsset, 0)
	for _, center := range mAssetCenter {

		// CreativeUser
		creativeUser := strings.Split(center.CreativeUser, "+")
		for _, s := range creativeUser {
			// PersonID
			// 根据简称查询员工id
			user := make([]*bean.User, 0)
			err = db2.MySQLClientUser.Table("user").
				Where("abbreviation = ?", s).Find(&user).Error
			if err != nil {
				fmt.Println("根据简称查询员工id 错误：", err)
				return
			}
			var personId int32
			if len(user) == 0 || s == constants.EmptyString {
				personId = 1000
			} else {
				personId = user[0].ID
			}

			adMaterialPersonRelation := &bean.AdMaterialPersonRelation{
				//ID:         0,
				MaterialID: center.Id,
				PersonID:   personId,
				Type:       constants.NumberOne,
			}
			// 4、将装有mongo数据的切片入库
			err = db2.MySQLClientCruiser.Table("ad_material_person_relations").Create(adMaterialPersonRelation).Error
			if err != nil {
				fmt.Println("入mysql/ad_material_person_relations 错误：", err)
				return
			}
		}

		// DesignUser
		designUser := strings.Split(center.DesignUser, "+")
		for _, s := range designUser {
			// PersonID
			// 根据简称查询员工id
			user := make([]*bean.User, 0)
			err = db2.MySQLClientUser.Table("user").
				Where("abbreviation = ?", s).Find(&user).Error
			if err != nil {
				fmt.Println("根据简称查询员工id 错误：", err)
				return
			}
			var personId int32
			if len(user) == 0 || s == constants.EmptyString {
				personId = 1000
			} else {
				personId = user[0].ID
			}

			adMaterialPersonRelation := &bean.AdMaterialPersonRelation{
				//ID:         0,
				MaterialID: center.Id,
				PersonID:   personId,
				Type:       constants.NumberTwo,
			}
			// 4、将装有mongo数据的切片入库
			err = db2.MySQLClientCruiser.Table("ad_material_person_relations").Create(adMaterialPersonRelation).Error
			if err != nil {
				fmt.Println("入mysql/ad_material_person_relations 错误：", err)
				return
			}
		}
	}
}
