package art_need

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"github.com/mao888/mao-gutils/constants"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

func RunArtNeedTagRelation() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("artneeds")

	// 2、从mongo查询数据
	mArtNeeds := make([]*MArtNeeds, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mArtNeeds)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mArtNeeds)

	// 3、将mongo数据装入切片
	for _, need := range mArtNeeds {

		// 需求id >= 8000000 的都不要
		if need.ID >= 8000000 {
			continue
		}

		tag := strings.Split(need.Tag, "+")
		// 如果tag数组为1或0，则 "+" 后无字符串，跳过
		if len(tag) == 1 || len(tag) == 0 {
			continue
		}
		// 根据"+"后的字符串去mysql/dictionaries中查type为4，若没有则添加，若有，则拿到id
		dictionary := make([]*Dictionary, 0)
		err = db2.MySQLClientCruiser.Table("dictionaries").
			Where("label = ?", tag[1]).Find(&dictionary).Error
		if err != nil {
			fmt.Println("根据"+"后的字符串去mysql/dictionaries中查 错误：", err)
			return
		}
		// 若没有则添加
		if len(dictionary) == constants.NumberZero {
			dic := &Dictionary{
				//ID:        0,
				Label:     tag[1],
				Code:      tag[1],
				Type:      constants.NumberFour,
				Remark:    "",
				CreatedAt: time.Now().Unix(),
			}
			err = db2.MySQLClientCruiser.Table("dictionaries").Create(dic).Error
			if err != nil {
				fmt.Println("向dictionaries中添加标签 错误", err)
				return
			}
			// 并插入 art_need_tag_relations
			artNeedTagRelations := &ArtNeedTagRelation{
				//ID:     0,
				NeedID: need.ID,
				TagID:  dic.ID,
			}
			err = db2.MySQLClientCruiser.Table("art_need_tag_relations").Create(artNeedTagRelations).Error
			if err != nil {
				fmt.Println("插入 art_need_tag_relations 错误", err)
				return
			}
			continue
		}

		// 若有，则拿到id,并插入 art_need_tag_relations
		artNeedTagRelations := &ArtNeedTagRelation{
			//ID:     0,
			NeedID: need.ID,
			TagID:  dictionary[0].ID,
		}
		err = db2.MySQLClientCruiser.Table("art_need_tag_relations").Create(artNeedTagRelations).Error
		if err != nil {
			fmt.Println("插入 art_need_tag_relations 错误", err)
			return
		}
	}
}
