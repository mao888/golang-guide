package main

func main() {
	// mongo_to_mysql.RunGame()

	// 将mongo数据 plat_console/projects 迁移到mysql中的 admin_console/sdk_project表
	//sdk_release.RunSdkProject()

	// 将mongo数据 plat_console/projectversions 迁移到mysql中的 admin_console/sdk_release_record表
	//sdk_release.RunSdkReleaseRecord()

	// 更新 sdk_release_record表中 version_number 字段为 可排序字段 version_ordinal
	// mongo_to_mysql.VersionNumberToOrdinal()

	// 将mongo数据 plat_console/sdks 迁移到mysql中的 admin_console/child_sdk表
	//sdk_release.RunChildSdk()

	// 将mongo数据 plat_console/sdkversions 迁移到mysql中的 admin_console/child_sdk_release_record表
	//sdk_release.RunChildSdkReleaseRecord()

	// 从admin_console/child_sdk.id 同步到 admin_console/jenkins_jobs.child_sdk_id 字段
	//sdk_release.RunJenkinsJobsChildSdk()

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
	//version_console.RunEnvAndVersion()

	// Mongo/plat_console/gamelanguageconfs 数据迁移到 ARK version_console/language
	//version_console.RunLanguage()

	// Mongo/plat_console/whitelists 数据迁移到 ARK version_console/whitelist
	// version_console.RunWhitelist()

	// Mongo/plat_console/versions 数据迁移到 ARK version_console/version
	// version_console.RunVersion()

	// 五、bi数据工具数据迁移
	// Mongo/rambler/thirdadvconfigs 数据迁移到 ARK admin_console/bi_data
	// bi_data.RunBiData()

	// 六、美术需求数据迁移
	// Mongo/plat_console/artneeds 数据迁移到 ARK cruiser_console/art_needs
	//art_need.RunArtNeeds()

	// Mongo/plat_console/artneeds.tag 数据迁移到 ARK cruiser_console/art_need_tag_relations
	//art_need.RunArtNeedTagRelation()

	// Mongo/plat_console/artneeds.creative_user 数据迁移到 ARK cruiser_console/art_need_person_relations
	// art_need.RunArtNeedPersonRelation()

	// Mongo/plat_console/artneeds.creative_user 数据迁移到 ARK cruiser_console/art_need_person_relations
	//art_need.RunArtNeedRelation()

	// Mongo/plat_console/artneeds.size 数据迁移到 ARK cruiser_console/art_need_material_size_relations
	//art_need.RunArtNeedMaterialSizeRelation()

	// Mongo/plat_console/artneeds.language 数据迁移到 ARK cruiser_console/art_need_language_relations
	//art_need.RunArtNeedLanguageRelation()

	// Mongo/plat_console/artneeds.design_user 数据迁移到 ARK cruiser_console/art_need_language_relations
	//art_need.RunArtTask()

	// Mongo/plat_console/artattachments 数据迁移到 ARK cruiser_console/art_attachments
	//art_need.RunArtAttachment()
}
