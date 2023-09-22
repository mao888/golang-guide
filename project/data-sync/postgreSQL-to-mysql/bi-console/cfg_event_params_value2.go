package main

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

// CfgEventParamsValuePG2 From PostgreSQL fotoabledb data_cfg.cfg_event_params_value
type CfgEventParamsValuePG2 struct {
	//   "app_id" varchar(2000),
	//  "params" varchar(2000),
	//  "params_value" varchar(2000),
	//  "params_label" varchar(2000),
	//  "person" varchar(2000),
	//  "remark" varchar(200),
	//  "create_time" timestamp(0),
	//  "update_time" timestamp(0)
	AppID       string `gorm:"column:app_id;NOT NULL;" json:"app_id"`
	Params      string `gorm:"column:params;NOT NULL;" json:"params"`
	ParamsValue string `gorm:"column:params_value;NOT NULL;" json:"params_value"`
	ParamsLabel string `gorm:"column:params_label;NOT NULL;" json:"params_label"`
	Person      string `gorm:"column:person;NOT NULL;" json:"person"`
	Remark      string `gorm:"column:remark;NOT NULL;" json:"remark"`
	CreateTime  string `gorm:"column:create_time;NOT NULL;" json:"create_time"`
	UpdateTime  string `gorm:"column:update_time;NOT NULL;" json:"update_time"`
}

// CfgEventParamsValue2 From MySQL bi_console cfg_event_params_value
type CfgEventParamsValue2 struct {
	//  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
	//  `app_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '应用id',
	//  `params` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数名称',
	//  `params_value` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数值',
	//  `params_label` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '参数值名称',
	//  `creator` int NOT NULL COMMENT '创建人id',
	//  `created_at` int NOT NULL DEFAULT '0' COMMENT '创建时间',
	//  `updated_at` int NOT NULL DEFAULT '0' COMMENT '更新时间',
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AppID       string `gorm:"column:app_id;NOT NULL;" json:"app_id"`
	Params      string `gorm:"column:params;NOT NULL;" json:"params"`
	ParamsValue string `gorm:"column:params_value;NOT NULL;" json:"params_value"`
	ParamsLabel string `gorm:"column:params_label;NOT NULL;" json:"params_label"`
	Creator     int32  `gorm:"column:creator;NOT NULL;" json:"creator"`
	CreatedAt   int32  `gorm:"column:created_at;NOT NULL;" json:"created_at"`
	UpdatedAt   int32  `gorm:"column:updated_at;NOT NULL;" json:"updated_at"`
}

func FunCfgEventParamsValue2() {
	batchSize := 5000 // 每批次处理的数据量

	offset := 0
	for {
		// 1、pg查数据
		cfgEventParamsValuePG := make([]*CfgEventParamsValuePG2, 0)
		err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").
			Offset(offset).
			Limit(batchSize).
			Find(&cfgEventParamsValuePG).Error
		if err != nil {
			fmt.Println("RunCfgEventParamsValue PostgreSQLClient Find err:", err)
			return
		}

		if len(cfgEventParamsValuePG) == 0 {
			fmt.Println("No more data to migrate.")
			return
		}

		// 2、转换数据并存入MySQL
		cfgEventParamsValue := make([]*CfgEventParamsValue2, 0)
		for _, v := range cfgEventParamsValuePG {
			cfgEventParamsValue = append(cfgEventParamsValue, &CfgEventParamsValue2{
				AppID:       v.AppID,
				Params:      v.Params,
				ParamsValue: v.ParamsValue,
				ParamsLabel: v.ParamsLabel,
				Creator:     0,
			})
		}

		// 3、mysql存数据
		err = db2.MySQLClientBI.Table("cfg_event_params_value").CreateInBatches(cfgEventParamsValue, batchSize).Error
		if err != nil {
			fmt.Println("RunCfgEventParamsValue MySQLClientBI CreateInBatches err:", err)
			return
		}

		offset += batchSize
		fmt.Printf("Successfully migrated records from offset %d to %d\n", offset-batchSize, offset-1)
	}
}

func main() {
	//for {
	FunCfgEventParamsValue2() // 使用协程执行数据迁移
	fmt.Println("Batch migration complete!")

	// 可以在这里加入一些延时，以防止过于频繁的查询
	//time.Sleep(time.Second * 5) // 5秒延时
	//}

	fmt.Println("Migration complete!")
}
