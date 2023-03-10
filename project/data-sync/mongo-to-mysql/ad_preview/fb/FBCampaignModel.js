const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class FBCampaignModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'FBCampaign'
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            campaign_id: String,
            ad_target: String, //  广告目标字段， 1：应用安装量，2: 转化量,
            campaign_name: String,
            cfg_frame_id: String, // 广告结构方案id （v3版 新增字段）
            audience_name: String, //受众名称
            custom_suf: String, //自定义后缀
            account_id: String,
            create_user: String, //创建人
            user_id: Number,
            ai_level_cfg: String,//ai广告模板层级配置
            ai_ruler_id: {
                type: Number,
                ref: 'aiassetruler'
            },
            ai_ruler_open: Boolean,
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
            created_by_cruiser: {
                type: Boolean,
                default: true
            }, //是否为自动化创建

            status: { //状态
                type: String,
                enum: ['ACTIVE', 'PAUSED', 'DELETED', 'ARCHIVED']
            },
            buying_type: { //购买类型
                type: String,
                enum: ['AUCTION', 'RESERVED']
            },
            objective: { //营销目标
                type: String,
                enum: ['APP_INSTALLS', 'CONVERSIONS']
            },

            //CBO特有属性
            daily_budget: Number,
            lifetime_budget: Number,
            bid_strategy: { //竞价策略
                type: String,
                enum: ['LOWEST_COST_WITHOUT_CAP', 'LOWEST_COST_WITH_BID_CAP', 'COST_CAP', 'TARGET_COST', 'LOWEST_COST_WITH_MIN_ROAS']
            },
            pacing_type: {
                type: Array,
                default: undefined
            }, //投放时段或类型说明字段 enum: ['standard', 'no_pacing', 'day_parting']
            create_type: String,
            company_id: Number,
        }
    }
}

module.exports = new FBCampaignModel();
