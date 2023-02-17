const BaseModel = require('@cc/cargo').BaseModel;

class AssetSyncNoticeModel extends BaseModel {
    constructor() {
        super('plat_console');
    }

    getName() {
        return 'assetsyncnotice';
    }

    getSchema() {
        return {
            _id: Number,
            target: String,   // 发送目标标识 规则: 类型_操作用户_公司_需求_任务总量_素材数量_账户数量_时间戳
            status: String,   // 状态 noStarted | finish
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date
            }
        };
    }


    async save(data) {
        data._id = await require('../IdGenModel').genId('asset_sync_noice');
        return super.save(data);
    }

}


module.exports = new AssetSyncNoticeModel();