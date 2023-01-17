package version_console

import (
	"context"
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Env mapped from table version_console <env>
type Env struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`  // 主键
	AppID     int32  `gorm:"column:app_id;not null" json:"app_id"`               // 应用id
	Type      int32  `gorm:"column:type;not null" json:"type"`                   // 环境类型 0未知 1生产 2测试 3自定义 4开发 5预发
	Name      string `gorm:"column:name;not null" json:"name"`                   // 环境名称
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	IsDeleted int32  `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除(0否1是)
}

func RunEnv() {
	// 1、建立连接
	db := db2.MongoClient.Database("plat_console")
	coll := db.Collection("environments")

	// 2、从mongo查询数据
	mEnvironment := make([]*MEnvironment, 0)
	err := coll.Find(context.TODO(), bson.M{}).All(&mEnvironment)
	if err != nil {
		fmt.Println("Mongo查询错误：", err)
		return
	}
	fmt.Println(mEnvironment)

	// 3、将mongo数据装入切片
	envs := make([]*Env, 0)
	for _, environment := range mEnvironment {
		fmt.Println(environment.AppID)
		// Type
		t := -1
		if environment.EnvID == 0 {
			t = 4
		} else if environment.EnvID == 1 {
			t = 2
		} else if environment.EnvID == 2 {
			t = 5
		} else if environment.EnvID == 3 {
			t = 1
		} else {
			t = 3
		}
		// IsDeleted
		isDeleted := 0
		if environment.DeleteTime != nil {
			isDeleted = 1
		}
		env := &Env{
			//ID:        0,
			AppID:     int32(environment.AppID),
			Type:      int32(t),
			Name:      environment.Name,
			UpdatedAt: environment.UpdateTime.Unix(),
			CreatedAt: environment.CreateTime.Unix(),
			IsDeleted: int32(isDeleted),
		}
		envs = append(envs, env)
	}
	// 4、将装有mongo数据的切片入库
	err = db2.MySQLClientVersion.Table("env").CreateInBatches(envs, len(envs)).Error
	if err != nil {
		fmt.Println("入mysql错误：", err)
	}
}
