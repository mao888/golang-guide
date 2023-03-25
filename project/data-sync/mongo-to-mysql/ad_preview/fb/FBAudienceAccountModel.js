const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class FBAudienceAccountModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'FBAudienceAccount';
    }
    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            audience_id: String, //受众id
            audience_type: String, //受众类型
            app_id: String, // 目标应用，
            beta_id: String,
            app_name: String, //应用id
            account_id: String, // 目标账户id
            account_name: String, // 账户名称
            create_user: String, // 添加人
            auto_type: {   // 是否自动更新， true  开启， false：关闭
                type: Boolean,
                default: false
            },
            request_state: {  // 1代表更新完成， 0代表正在更新中， 2 还未更新
                type: String,
                default: '2'
            },
            fb_audience_id: {
                type: String,
                default: ''
            },// fb 创建的后的受众id
            last_update_num: {
                type: Number,
                default: 0
            }, // 最新更新条数,
            update_err_msg: {
                type: String,
                default: ''
            }, // 更新错误信息,
            errObj: {
                type: Object,
                default: {}
            },
            deleted: {    // 是否删除
                type: Date
            },
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
            company_id: Number
        };
    }
}

module.exports = new FBAudienceAccountModel();