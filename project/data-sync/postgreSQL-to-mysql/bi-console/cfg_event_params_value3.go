package main

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
	"time"
	"unsafe"
)

// CfgEventParamsValuePG3 From PostgreSQL fotoabledb data_cfg.cfg_event_params_value
type CfgEventParamsValuePG3 struct {
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

// CfgEventParamsValue3 From MySQL bi_console cfg_event_params_value
type CfgEventParamsValue3 struct {
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

func FunCfgEventParamsValue3() {

	// 1、pg查数据
	cfgEventParamsValuePG := make([]*CfgEventParamsValuePG3, 0)
	err := db2.PostgreSQLClient.Table("data_cfg.cfg_event_params_value").
		Find(&cfgEventParamsValuePG).Error
	if err != nil {
		fmt.Println("RunCfgEventParamsValue PostgreSQLClient Find err:", err)
		return
	}
	// 切片总长度
	fmt.Println("RunCfgEventParamsValue PostgreSQLClient Find len(cfgEventParamsValuePG):", len(cfgEventParamsValuePG))

	// 获取切片的底层数组大小
	//// 使用反射获取切片的底层数组的大小
	//sliceHeader := reflect.SliceHeader{
	//	Data: uintptr(unsafe.Pointer(&cfgEventParamsValuePG[0])),
	//	Len:  len(cfgEventParamsValuePG),
	//	Cap:  cap(cfgEventParamsValuePG),
	//}
	// 计算底层数组的大小
	arraySize := int(unsafe.Sizeof(cfgEventParamsValuePG[0])) * cap(cfgEventParamsValuePG)
	arraySizeGB := bytesToGigabytes(arraySize) // 将字节转换为千兆字节（GB）
	fmt.Printf("切片的内存大小: %d 字节\n", arraySize)
	fmt.Printf("切片的内存大小: %f GB\n", arraySizeGB)

	if len(cfgEventParamsValuePG) == 0 {
		fmt.Println("No more data to migrate.")
		return
	}

	// 2、转换数据并存入MySQL
	//for i, v := range cfgEventParamsValuePG {
	//	cfgEventParamsValue := &CfgEventParamsValue3{
	//		AppID:       v.AppID,
	//		Params:      v.Params,
	//		ParamsValue: v.ParamsValue,
	//		ParamsLabel: v.ParamsLabel,
	//	}
	//	// 3、mysql存数据
	//	err = db2.MySQLClientBI.Table("cfg_event_params_value").Create(cfgEventParamsValue).Error
	//	if err != nil {
	//		fmt.Println("RunCfgEventParamsValue MySQLClientBI CreateInBatches err:", err)
	//		return
	//	}
	//	fmt.Println("第 ", i, " 条数据迁移完成")
	//}
	// 每次迁移的批次大小
	batchSize := 500
	totalRecords := len(cfgEventParamsValuePG)

	// 迭代批次
	for startIdx := 0; startIdx < totalRecords; startIdx += batchSize {
		endIdx := startIdx + batchSize
		if endIdx > totalRecords {
			endIdx = totalRecords
		}

		// 批次数据
		batchData := cfgEventParamsValuePG[startIdx:endIdx]

		// 转换数据并存入MySQL
		batchRecords := make([]*CfgEventParamsValue3, 0)
		for _, v := range batchData {
			cfgEventParamsValue := &CfgEventParamsValue3{
				AppID:       v.AppID,
				Params:      v.Params,
				ParamsValue: v.ParamsValue,
				ParamsLabel: v.ParamsLabel,
			}
			batchRecords = append(batchRecords, cfgEventParamsValue)
		}

		// 批量插入数据
		err := db2.MySQLClientBI.Table("cfg_event_params_value").CreateInBatches(batchRecords, len(batchRecords)).Error
		if err != nil {
			fmt.Println("RunCfgEventParamsValue MySQLClientBI CreateInBatches err:", err)
			return
		}
		// 打印批次信息
		fmt.Printf("Migrating records from offset %d to %d\n", startIdx, endIdx-1)
	}
}

func bytesToGigabytes(bytes int) float64 {
	gigabytes := float64(bytes) / (1024 * 1024 * 1024)
	return gigabytes
}

func main() {
	startTime := time.Now()

	FunCfgEventParamsValue3()

	elapsedTime := time.Since(startTime)
	fmt.Printf("Migration complete! Time elapsed: %s\n", elapsedTime)
	fmt.Println("Migration complete!")
}
