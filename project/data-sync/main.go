package main

import "github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql"

func main() {
	//data-sync.RunGame()

	// 将mongo数据迁移到mysql中的sdk_project表
	//data-sync.RunSdkProject()

	// 将mongo数据迁移到mysql中的sdk_release_record表
	//data-sync.RunSdkReleaseRecord()

	// 将mongo数据迁移到mysql中的child_sdk表
	//data-sync.RunChildSdk()

	// 将mongo数据迁移到mysql中的child_sdk_release_record表
	mongo_to_mysql.RunChildSdkReleaseRecord()
}
