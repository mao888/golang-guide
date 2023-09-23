package main

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"gorm.io/gorm"
	"log"
	"sync"
)

// CfgEventParamsValuePG4 From PostgreSQL fotoabledb data_cfg.cfg_event_params_value
type CfgEventParamsValuePG4 struct {
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

// CfgEventParamsValue4 From MySQL bi_console cfg_event_params_value
type CfgEventParamsValue4 struct {
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

var wg2 sync.WaitGroup

const batchSize2 = 500
const maxGoroutines2 = 2 // 手动设置最大并发协程数量

type Result struct {
	Offset  int
	Success bool
}

func FunCfgEventParamsValue4(offset int, resultChan chan<- Result) {
	defer wg2.Done()

	// 1、pg查数据
	cfgEventParamsValuePG := make([]*CfgEventParamsValuePG4, 0)
	err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").
		Limit(batchSize2).Offset(offset).Find(&cfgEventParamsValuePG).Error
	if err != nil {
		log.Printf("RunCfgEventParamsValue PostgreSQLClient Find err at offset %d: %v", offset, err)
		resultChan <- Result{Offset: offset, Success: false}
		return
	}
	if len(cfgEventParamsValuePG) == 0 {
		log.Printf("RunCfgEventParamsValue PostgreSQLClient Find len(cfgEventParamsValuePG) == 0 at offset %d", offset)
		resultChan <- Result{Offset: offset, Success: true}
		return
	}

	// 2、转换数据
	cfgEventParamsValue := make([]*CfgEventParamsValue4, 0)
	for _, v := range cfgEventParamsValuePG {
		cfgEventParamsValue = append(cfgEventParamsValue, &CfgEventParamsValue4{
			AppID:       v.AppID,
			Params:      v.Params,
			ParamsValue: v.ParamsValue,
			ParamsLabel: v.ParamsLabel,
			Creator:     0,
		})
	}

	log.Printf("Migrating records from offset %d to %d", offset, offset+batchSize2-1) // 记录每批次迁移数据的起止

	// 3、mysql存数据
	err = db2.MySQLClientBI.Table("cfg_event_params_value").Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(cfgEventParamsValue, batchSize2).Error
	})
	if err != nil {
		log.Printf("Error inserting batch into MySQL from offset %d to %d: %v", offset, offset+batchSize2-1, err)
		resultChan <- Result{Offset: offset, Success: false}
		return
	}
	log.Printf("Successfully migrated records from offset %d to %d", offset, offset+batchSize2-1)
	resultChan <- Result{Offset: offset, Success: true}
}

func main() {
	// 获取总记录数
	var totalRecords int64
	db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").Count(&totalRecords)

	resultChan := make(chan Result, maxGoroutines2)

	offset := 0
	for offset < int(totalRecords) {
		wg2.Add(1)
		go FunCfgEventParamsValue4(offset, resultChan) // 使用协程执行数据迁移
		offset += batchSize2
	}

	// 等待所有协程完成
	go func() {
		wg2.Wait()
		close(resultChan)
	}()

	// 处理结果
	successCount := 0
	failureCount := 0
	for result := range resultChan {
		if result.Success {
			successCount++
		} else {
			failureCount++
		}
	}

	fmt.Printf("Migration complete! Successful: %d, Failed: %d\n", successCount, failureCount)
}
