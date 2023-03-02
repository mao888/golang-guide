package version_console

import (
	"fmt"
	db2 "github.com/mao888/golang-guide/project/data-sync/db"
)

func RunIsGrayTo2() {
	var version []*Version
	db2.MySQLClientVersion.Table("version").Where("is_gray = ?", 1).
		Where("gray_scale = ?", 100).Where("is_deleted = ?", 0).Find(&version)
	fmt.Println("len(version):", len(version))

	for i, v := range version {
		fmt.Println("version:", i)

		err := db2.MySQLClientVersion.Table("version").Where("id = ?", v.ID).
			UpdateColumn("is_gray", 0).Error
		if err != nil {
			fmt.Println("更新数据 错误：", err)
			return
		}
	}
}
