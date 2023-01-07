package main

import "github.com/mao888/golang-guide/project/mongo-to-mysql/internal"

func main() {
	//internal.RunGame()

	// 将mongo数据迁移到mysql中的sdk_project表
	//internal.RunSdkProject()

	// 将mongo数据迁移到mysql中的sdk_release_record表
	//internal.RunSdkReleaseRecord()

	// 将mongo数据迁移到mysql中的child_sdk表
	internal.RunChildSdk()
}
