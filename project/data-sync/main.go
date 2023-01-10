package main

import mysql_to_mysql "github.com/mao888/golang-guide/project/data-sync/mysql-to-mysql"

func main() {
	//mongo_to_mysql.RunGame()

	// 将mongo数据迁移到mysql中的sdk_project表
	//mongo_to_mysql.RunSdkProject()

	// 将mongo数据迁移到mysql中的sdk_release_record表
	// mongo_to_mysql.RunSdkReleaseRecord()

	// 更新 sdk_release_record表中 version_number 字段为 可排序字段 version_ordinal
	// mongo_to_mysql.VersionNumberToOrdinal()

	// 将mongo数据迁移到mysql中的child_sdk表
	//data-sync.RunChildSdk()

	// 将mongo数据迁移到mysql中的child_sdk_release_record表
	//mongo_to_mysql.RunChildSdkReleaseRecord()

	// 将 gm-system/m_activities 入库 application_console/activities
	//mysql_to_mysql.RunActivity()

	// 将 gm-system/m_command 入库 application_console/command
	//mysql_to_mysql.RunCommand()

	// 将 gm-system/m_command_project_team gm-system/m_project_team_player 入库 application_console/gm_config
	// mysql_to_mysql.RunGmConfig()

	// 将 gm-system/m_mail_goods 入库 application_console/mail_goods
	mysql_to_mysql.RunMailGoods()
}
