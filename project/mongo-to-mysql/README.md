### SDK发版管理 数据迁移到 ARK

| 数据库/表（MongoDb）         | 迁移方式 | 目标数据库/表（Mysql）                       |
| ---------------------------- | -------- | -------------------------------------------- |
| SDK发版管理/MongoDb          |          | application_console/mysql                    |
| plat_console/projects        | go脚本   | application_console/sdk_project              |
| plat_console/projectversions | go脚本   | application_console/sdk_release_record       |
| plat_console/sdks            | go脚本   | application_console/child_sdk                |
| plat_console/sdkversions     | go脚本   | application_console/child_sdk_release_record |
