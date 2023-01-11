### gm-system 数据迁移到 ARK application_console
### 实施描述
- 1、建立MongoDb连接 
- 2、从mongo查询数据 
- 3、将从mongo中查出的games.id(int)作为key, games.game_id(string)作为value,存入map 
- 4、从mysql查询 gm-system 相关表数据
- 5、将 gm-system/table 入库 application_console/table

| 数据库/表（Mysql、Mongo）               | 迁移方式 | 目标数据库/表（Mysql）                               |
|----------------------------------|------|----------------------------------------------|
| gm-system/Mysql                  |      | application_console/Mysql                    |
| gm-system/m_activities           | go脚本 | application_console/activities              |
| gm-system/m_command              | go脚本 | application_console/ command      |
| gm-system/m_command_project_team | go脚本 | application_console/gm_config                |
| gm-system/m_project_team_player  | go脚本 | application_console/gm_config |
| (Mongo)plat_console/gameconfigs  | go脚本 | application_console/gm_config |
| gm-system/m_mail_goods           | go脚本 | application_console/ mail_goods|
| gm-system/m_mails                | go脚本 | application_console/mails |
| gm-system/m_command_option       | 直接迁移 | application_console/command_option |
| gm-system/command_template       | 直接迁移 | application_console/command_template |
| gm-system/m_mail_languages       | 直接迁移 | application_console/ mail_languages|