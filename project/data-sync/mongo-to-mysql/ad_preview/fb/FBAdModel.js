const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class FBAdModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'FBAd';
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            status: { //状态
                type: String,
                enum: ['ACTIVE', 'PAUSED', 'DELETED', 'ARCHIVED']
            },
            ad_id: String,
            ad_name: String,
            adset_id: String,
            campaign_id: String,
            account_id: String,
            publish_type: Boolean, //使用主页而不是应用名称作为广告发布身份
            ad_text: String,
            ad_title: String,
            perform_type: String,
            deep_extend_url: String,
            create_user: String, //创建人
            user_id: Number,
            created_by_cruiser: {
                type: Boolean,
                default: true
            }, //是否为自动化创建
            asset_info: {
                media_id: String, //已上传到账户的素材
                file_type: String, //素材类型
                file_name: String, //素材名称
                file_id: Number, //素材在系统中的id
            },
            create_time: {
                type: Date,
                default: Date.now
            },
            create_type: String,
            company_id: Number,
        }
    }
}

module.exports = new FBAdModel();