### SDK发版管理 数据迁移到 ARK  admin_console
- 1、建立MongoDb连接 
- 2、从mongo查询数据 
- 3、将mongo数据装入切片 
- 4、将装有mongo数据的切片入mysql库

| 数据库/表（MongoDb）         | 迁移方式 | 目标数据库/表（Mysql）                       |
| ---------------------------- | -------- | -------------------------------------------- |
| SDK发版管理/MongoDb          |          | admin_console/mysql                    |
| plat_console/projects        | go脚本   | admin_console/sdk_project              |
| plat_console/projectversions | go脚本   | admin_console/sdk_release_record       |
| plat_console/sdks            | go脚本   | admin_console/child_sdk                |
| plat_console/sdkversions     | go脚本   | admin_console/child_sdk_release_record |
