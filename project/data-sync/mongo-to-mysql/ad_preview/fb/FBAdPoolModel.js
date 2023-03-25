const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');
const CustomEnum = require('../../utils/Enum');

class FBAdPoolModel extends BaseModel {

    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'FBAdPool';
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            aijob_model_id: Number,
            account_id: {
                type: String,
            },
            campaign_id: {
                type: String,
                ref: 'FBCampaign'
            },
            adset_id: {
                type: String,
                ref: 'FBAdset'
            },
            ad_id: {
                type: String,
                ref: 'FBAd'
            },
            asset_id: { //素材ID
                type: Number,
                ref: 'FBAsset'
            },
            status: { //广告当前所处池状态码
                type: String,
                enum: Object.values(CustomEnum.BID_POOL)
            },
            is_test: Boolean, //是否为泛投测试广告，即冷启动池中的广告
            clone_status: { //测试广告是否测试通过，如通过此字段作为次日克隆任务的触发点
                type: String,
                enum: ['ready', 'done']
            },
            priority: {     //排队开启优先级
                type: Number,
                default: 0
            },
            test_pass: Boolean,   //素材是否跳过测试阶段的广告
            reopen_times: Number, //复开次数
            last_bid_time: Date,  //上次调价时间
            acc_bid_num: Number,  //调价次数
            delete_time: Date,
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            }
        };
    }

}

module.exports = new FBAdPoolModel();