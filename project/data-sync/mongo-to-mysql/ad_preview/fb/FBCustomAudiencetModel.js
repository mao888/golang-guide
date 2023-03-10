const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class FBCustomAudiencetModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'FBCustomAudience';
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            audience_name: String, //受众名称
            audience_desc: String, //受众描述
            source_app: String, // 来源应用Id
            source_app_name: String,// 来源应用名称
            app_eng_name: String, // 共同应用名称
            platform_type: String, // 投放平台限制条件
            platforms: Array, // 投放平台值
            country_type: String, // 国家限制条件
            countries: Array, // 国家值
            optimize_type: String,  // 优化方式限制条件
            optimizes: Array, // 优化方式值
            audience_type: {
                type: String,
                default: 'custom',
            }, // custom 自定义受众
            user_types: Array, //用户类型
            // day_num: Number, // N 日内
            // income_type: String, // 收费类型
            // income_value: Number, // 收入值
            power_apps: Array, // 受众授权应用
            create_user: String, // 创建人,
            audString: String, // 受众简化对象字符串
            auto_type: {    // 是否开启
                type: Boolean,
                default: false,
            },
            kochava_id: String, // kochavaId
            count_users: Number, // 受众用户量
            deleted: {
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

module.exports = new FBCustomAudiencetModel();