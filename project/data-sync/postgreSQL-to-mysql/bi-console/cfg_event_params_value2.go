package main

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"log"
	"time"
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

const batchSize3 = 200

func MigrateData() error {
	// 开启事务
	tx := db2.MySQLClientBI.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 发生错误时回滚事务
		}
	}()

	// 获取总记录数
	var totalRecords int64
	if err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").Count(&totalRecords).Error; err != nil {
		return err
	}

	offset := 0
	for offset < int(totalRecords) {
		// 查询数据
		cfgEventParamsValuePG := make([]*CfgEventParamsValuePG2, 0)
		if err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").
			Limit(batchSize3).Offset(offset).Find(&cfgEventParamsValuePG).Error; err != nil {
			return err
		}

		// 转换数据
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

		log.Printf("Migrating records from offset %d to %d", offset, offset+batchSize3-1) // 记录每批次迁移数据的起止

		// 插入数据
		if err := tx.Table("cfg_event_params_value").
			CreateInBatches(cfgEventParamsValue, batchSize3).Error; err != nil {
			return err
		}

		offset += batchSize3
		log.Printf("Successfully migrated records from offset %d to %d", offset-batchSize3, offset-1)
	}

	// 提交事务
	tx.Commit()
	return nil
}

func main() {
	startTime := time.Now()

	if err := MigrateData(); err != nil {
		log.Printf("Data migration failed: %v", err)
		return
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Migration complete! Time elapsed: %s\n", elapsedTime)
}
