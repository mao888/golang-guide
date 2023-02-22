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
	//version_console.RunVersion()  废弃

	// 五、bi数据工具数据迁移
	// Mongo/rambler/thirdadvconfigs 数据迁移到 ARK admin_console/bi_data
	// bi_data.RunBiData()

	// 六、美术需求数据迁移
	// 美术需求主表
	// Mongo/plat_console/artneeds 数据迁移到 ARK cruiser_console/art_needs
	//art_need.RunArtNeeds()

	// 美术需求标签多对多关联表
	// Mongo/plat_console/artneeds.tag 数据迁移到 ARK cruiser_console/art_need_tag_relations
	//art_need.RunArtNeedTagRelation()

	// 美术需求创意负责人多对多关联表
	// Mongo/plat_console/artneeds.creative_user 数据迁移到 ARK cruiser_console/art_need_person_relations
	//art_need.RunArtNeedPersonRelation()

	// 美术需求关联需求多对多关联表
	// Mongo/plat_console/artneeds.relatedList 数据迁移到 ARK cruiser_console/art_need_relations
	//art_need.RunArtNeedRelation()

	// 美术需求素材尺寸多对多关联表
	// Mongo/plat_console/artneeds.size 数据迁移到 ARK cruiser_console/art_need_material_size_relations
	//art_need.RunArtNeedMaterialSizeRelation()

	// 美术需求语种多对多关联表
	// Mongo/plat_console/artneeds.language 数据迁移到 ARK cruiser_console/art_need_language_relations
	//art_need.RunArtNeedLanguageRelation()

	// 美术需求子任务表
	// Mongo/plat_console/artneeds.design_user 数据迁移到 ARK cruiser_console/art_tasks
	//art_need.RunArtTask()

	// 美术需求附件-终稿
	// Mongo/plat_console/artattachments 数据迁移到 ARK cruiser_console/art_attachments
	//art_need.RunArtAttachment()

	// 美术需求默认描述表
	// Mongo/plat_console/artneeds.main_desc 数据迁移到 ARK cruiser_console/base_desc_template
	//art_need.RunBaseDescTemplate()

	// 美术需求主表-更新补充说明
	// Mongo/plat_console/artneedlogs 数据迁移到 ARK cruiser_console/art_needs.extra_desc
	//art_need.RunSupplyDesc()

	// 七、广告投放-配置中心数据迁移
	// 广告配置中心-受众组
	// Mongo/cruiser_console_v2/cfgaudiences -> ARK cruiser_console/ad_conf_audience
	//ad_conf_centre.RunAdConfAudience()

	// 广告配置中心-受众组 排除包含受众关联表
	// Mongo/cruiser_console_v2/cfgaudiences.custom_audiences/cfgaudiences.excluded_custom_audiences -> ARK cruiser_console/ad_conf_audience_include_relations
	//ad_conf_centre.RunAdConfAudienceIncludeRelation()

	// 广告配置中心-受众组细分定位关联表
	// Mongo/cruiser_console_v2/cfgaudiences.flexible_spec/cfgaudiences.exclusions -> ARK cruiser_console/ad_conf_audience_isegmentation_relations
	//ad_conf_centre.RunAdConfAudienceIsegmentationRelation()

	// 广告投放-配置中心-结构方案表
	// Mongo/cruiser_console_v2/cfgframes -> ARK cruiser_console/ad_conf_scheme
	//ad_conf_centre.RunAdConfScheme()

	// 广告投放-配置中心-版位组
	// Mongo/cruiser_console_v2/cfgpositions -> ARK cruiser_console/ad_conf_position
	//ad_conf_centre.RunAdConfPosition()

	// 广告投放-配置中心-国家组表
	// Mongo/cruiser_console_v2/cfgcountries -> ARK cruiser_console/ad_conf_country
	//ad_conf_centre.RunAdConfCountry()

	// 八、资产库（内部）、动作库（内部）
	// 美术资产库表，该表保护（美术资产，动作资产，音乐资产等)
	// Mongo/artsources -> ARK cruiser_console/art_asset
	//art_asset.RunArtAsset1() // 美术资产

	// 美术资产库表，该表保护（美术资产，动作资产，音乐资产等)
	// Mongo/activelibraries -> ARK cruiser_console/art_asset
	//art_asset.RunArtAsset2() // 动作资产

	// 资产标签表
	// Mongo/tags -> ARK cruiser_console/art_asset_tags
	//art_asset.RunArtAssetTag()

	// 美术资产-标签表多对多关联表
	// Mongo/artsources/activelibraries -> ARK cruiser_console/art_asset_tag_relations
	//art_asset.RunArtAssetTagRelation()

	// 美术需求-资产多对多关联表
	// Mongo/artsources/activelibraries -> ARK cruiser_console/art_need_asset_relations
	//art_asset.RunArtNeedAssetRelation()
}
