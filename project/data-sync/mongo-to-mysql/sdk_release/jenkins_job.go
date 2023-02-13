package sdk_release

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

// JenkinsJob mapped from table admin_console <jenkins_jobs>
type JenkinsJob struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	JenkinsName string `gorm:"column:jenkins_name;not null" json:"jenkins_name"`   // jenkins_name
	JenkinsURL  string `gorm:"column:jenkins_url;not null" json:"jenkins_url"`     // jenkins_url
	ChildSdkID  int32  `gorm:"column:child_sdk_id;not null" json:"child_sdk_id"`   // 子sdk项目id
	CreatedAt   int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"` // 创建时间
	UpdatedAt   int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"` // 更新时间
	IsDeleted   bool   `gorm:"column:is_deleted;not null" json:"is_deleted"`       // 是否删除
}

func RunJenkinsJobsChildSdk() {
	// 1、从mysql查询 jenkins_jobs
	jenkinsJob := make([]*JenkinsJob, 0)
	err := db2.MySQLClientAdmin.Table("jenkins_jobs").Find(&jenkinsJob).Error
	if err != nil {
		fmt.Println("从mysql查询 jenkins_jobs 错误：", err)
		return
	}

	// 2、从mysql查询 child_sdk
	childSdk := make([]*ChildSdk, 0)
	err = db2.MySQLClientAdmin.Table("child_sdk").Find(&childSdk).Error
	if err != nil {
		fmt.Println("从mysql查询 child_sdk 错误：", err)
		return
	}

	idMap := map[string]int32{}
	for _, sdk := range childSdk {
		idMap[sdk.Jenkins] = sdk.ID
	}
	fmt.Println("map", idMap)

	for _, job := range jenkinsJob {
		job.ChildSdkID = idMap[job.JenkinsName]
		err := db2.MySQLClientAdmin.Table("jenkins_jobs").Where("id = ?", job.ID).
			UpdateColumn("child_sdk_id", job.ChildSdkID).Error
		if err != nil {
			fmt.Println("mysql更新 child_sdk_id 错误", err)
			return
		}
	}
}
