package main

import version_console "github.com/mao888/golang-guide/project/data-sync/mongo-to-mysql/version-console"

func main() {
	// mongo_to_mysql.RunGame()

	// 将mongo数据 plat_console/projects 迁移到mysql中的 admin_console/sdk_project表
	// mongo_to_mysql.RunSdkProject()

	// 将mongo数据 plat_console/projectversions 迁移到mysql中的 admin_console/sdk_release_record表
	// mongo_to_mysql.RunSdkReleaseRecord()

	// 更新 sdk_release_record表中 version_number 字段为 可排序字段 version_ordinal
	// mongo_to_mysql.VersionNumberToOrdinal()

	// 将mongo数据 plat_console/sdks 迁移到mysql中的 admin_console/child_sdk表
	// mongo_to_mysql.RunChildSdk()

	// 将mongo数据 plat_console/sdkversions 迁移到mysql中的 admin_console/child_sdk_release_record表
	// mongo_to_mysql.RunChildSdkReleaseRecord()

	// 从admin_console/child_sdk.id 同步到 admin_console/jenkins_jobs.child_sdk_id 字段
	// mongo_to_mysql.RunJenkinsJobsChildSdk()

	// 二、gm-system 数据迁移到 ARK application_console

	// 将 gm-system/m_activities 入库 application_console/activities
	// mysql_to_mysql.RunActivity()

	// 将 gm-system/m_command 入库 application_console/command
	// mysql_to_mysql.RunCommand()

	// 将 gm-system/m_command_project_team gm-system/m_project_team_player plat_console/gameconfigs入库 application_console/gm_config
	// mysql_to_mysql.RunGmConfig()
	// mysql_to_mysql.RunMGameConfigs()

	// 将 gm-system/m_mail_goods 入库 application_console/mail_goods
	// mysql_to_mysql.RunMailGoods()

	// 将 gm-system/m_mails 入库 application_console/mails
	// mysql_to_mysql.RunMails()

	// 三、远程配置数据迁移
	// Mongo/app_console/remote_config 数据迁移到 ARK application_console/remote_config
	// remote_config.RunRemoteConfig()

	// 四、版本检查数据迁移
	// Mongo/plat_console/environments 数据迁移到 ARK version_console/env
	version_console.RunEnvAndVersion()

	// Mongo/plat_console/gamelanguageconfs 数据迁移到 ARK version_console/language
	//version_console.RunLanguage()

	// Mongo/plat_console/whitelists 数据迁移到 ARK version_console/whitelist
	// version_console.RunWhitelist()

	// Mongo/plat_console/versions 数据迁移到 ARK version_console/version
	// version_console.RunVersion()

	// 五、bi数据工具数据迁移
	// Mongo/rambler/thirdadvconfigs 数据迁移到 ARK admin_console/bi_data
	// bi_data.RunBiData()
}
