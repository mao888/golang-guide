package main

func main() {
	// mongo_to_mysql.RunGame()

	//
	// 将mongo数据 plat_console/projects 迁移到mysql中的 admin_console/sdk_project表
	//sdk_release.RunSdkProject()

	// 将mongo数据 plat_console/projectversions 迁移到mysql中的 admin_console/sdk_release_record表
	//sdk_release.RunSdkReleaseRecord()

	// 更新 sdk_release_record表中 version_number 字段为 可排序字段 version_ordinal
	//sdk_release.VersionNumberToOrdinal()

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
	//remote_config.RunRemoteConfig()

	// 四、版本检查数据迁移
	// Mongo/plat_console/environments 数据迁移到 ARK version_console/env
	//version_console.RunEnvAndVersion()

	//Mongo/plat_console/gamelanguageconfs 数据迁移到 ARK version_console/language
	//version_console.RunLanguage()

	// Mongo/plat_console/whitelists 数据迁移到 ARK version_console/whitelist
	// version_console.RunWhitelist()

	// Mongo/plat_console/versions 数据迁移到 ARK version_console/version
	//version_console.RunVersion()  废弃

	// 更新 is_gray
	//version_console.RunIsGrayTo2()

	// 五、bi数据工具数据迁移
	// Mongo/rambler/thirdadvconfigs 数据迁移到 ARK admin_console/bi_data
	//bi_data.RunBiData()

	// 六、美术需求数据迁移
	// 1、美术需求主表
	// Mongo/plat_console/artneeds 数据迁移到 ARK cruiser_console/art_needs
	//art_need.RunArtNeeds()

	// 2、美术需求标签多对多关联表
	// Mongo/plat_console/artneeds.tag 数据迁移到 ARK cruiser_console/art_need_tag_relations
	//art_need.RunArtNeedTagRelation()

	// 3、美术需求创意负责人多对多关联表
	// Mongo/plat_console/artneeds.creative_user 数据迁移到 ARK cruiser_console/art_need_person_relations
	//art_need.RunArtNeedPersonRelation()

	// 4、美术需求关联需求多对多关联表
	// Mongo/plat_console/artneeds.relatedList 数据迁移到 ARK cruiser_console/art_need_relations
	//art_need.RunArtNeedRelation()

	// 5、美术需求素材尺寸多对多关联表
	// Mongo/plat_console/artneeds.size 数据迁移到 ARK cruiser_console/art_need_material_size_relations
	//art_need.RunArtNeedMaterialSizeRelation()

	// 6、美术需求语种多对多关联表
	// Mongo/plat_console/artneeds.language 数据迁移到 ARK cruiser_console/art_need_language_relations
	//art_need.RunArtNeedLanguageRelation()

	// 7、美术需求子任务表
	// Mongo/plat_console/artneeds.design_user 数据迁移到 ARK cruiser_console/art_tasks
	//art_need.RunArtTask()

	// 8、美术需求附件-终稿
	// Mongo/plat_console/artattachments 数据迁移到 ARK cruiser_console/art_attachments
	//art_need.RunArtAttachment()

	// 美术需求附件-终稿 迁移未完成需求 的附件-终稿件
	//art_need.RunArtAttachment2()

	// 9、美术需求默认描述表
	// Mongo/plat_console/artneeds.main_desc 数据迁移到 ARK cruiser_console/base_desc_template
	//art_need.RunBaseDescTemplate()

	// 10、美术需求主表-更新补充说明
	// Mongo/plat_console/artneedlogs 数据迁移到 ARK cruiser_console/art_needs.extra_desc
	//art_need.RunSupplyDesc()

	// 11、更新 artattachments.url 域名
	//art_need.RunUrlReplace()

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
	//ad_conf_centre.RunAdConfScheme() // 210

	// 广告投放-配置中心-版位组
	// Mongo/cruiser_console_v2/cfgpositions -> ARK cruiser_console/ad_conf_position
	//ad_conf_centre.RunAdConfPosition() // 20

	// 广告投放-配置中心-国家组表
	// Mongo/cruiser_console_v2/cfgcountries -> ARK cruiser_console/ad_conf_country
	//ad_conf_centre.RunAdConfCountry() // 214

	// 广告投放-配置中心-文案表
	// Mongo/cruiser_console_v2/adtexts -> ARK cruiser_console/ad_conf_copywriting
	//ad_conf_centre.RunAdConfCopywriting() // 336

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

	// 九、广告投放-素材中心
	// 原始时间统计：73 + 61 + 134 + 185 = 453min = 7.55h
	// 时间统计开：4个协程，共 2.5h

	// 1、广告素材主表（它的ID会社交关联到广告素材tag，尺寸，语言，负责人等关联表) （4.3万 73min）
	// Mongo/assetcenters -> ARK cruiser_console/ad_material
	//var wg sync.WaitGroup
	//wg.Add(4)
	//go func() {
	//defer wg.Done()
	//ad_material.RunAdMaterial()
	//}()

	// 2、广告素材语言关联表-语言表多对多关联表 （4.3万 61min）
	// Mongo/assetcenters.asset_language -> ARK cruiser_console/ad_material_language_relations
	//go func() {
	// defer wg.Done()
	// ad_material.RunAdMaterialLanguageRelation()
	//}()

	// 3、广告素材人员关联表-人员表多对多关联表 （11万 134min）
	// Mongo/assetcenters.creative_user/design_user -> ARK cruiser_console/ad_material_person_relations
	//go func() {
	// defer wg.Done()
	// ad_material.RunAdMaterialPersonRelation()
	//}()

	// 4、广告素材 上传同步 返回对照表          (17万 185 min)
	// Mongo/assetcenters.media_list -> ARK cruiser_console/ad_material_sync_success
	//go func() {
	// defer wg.Done()
	// ad_material.RunAdMaterialSyncSuccess()
	//}()
	//wg.Wait()

	// 5、更新 广告素材主表 url域名
	//ad_material.RunUrlReplace()

	// ad_material mysql to mysql
	//ad_material.RunAdMaterialMysqlToMysql()

	// ad_material_sync_success 洗数据 success_id
	//ad_material.RunAdMaterialSyncSuccess()

	// 十、数据权限
	// BI数据权限
	// Mongo/rambler/dimpermissions.child_system:"bi" -> ARK user_console/policy user_console/policy_resource user_console/user_perm
	//user_console.RunUserPermPoliceResourceBi()
	//user_console.RunUserPermPoliceResourceBiApp()
	//user_console.RunUserPermPoliceResourceBiGame()

	// 广告投放数据权限
	// Mongo/rambler/dimpermissions.child_system:"cruiser_v2,pandora,art_needs" -> ARK user_console/policy user_console/policy_resource user_console/user_perm
	//user_console.RunUserPermPoliceResourceAdAppAndGame()

	// 十一、创建广告

	// 十二、BI-配置管理-自定义参数配置
	// PostgreSQL/data_cfg.cfg_event_params_value -> MySQL/bi_console.cfg_event_params_value
}
