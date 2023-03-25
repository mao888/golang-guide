const BaseModel = require('@cc/cargo').BaseModel;

// 素材
class AdFlowModel extends BaseModel {

    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return "adflow";
    }

    getSchema() {
        return {
            _id: Number,
            fbid: String,
            type: String,  //创建类型，0，1， 2， （向目标广告组中新建广告， 复制目标广告组参数并新建广告组， 手动新建广告组）
            campaign_type: String, //广告系列类型1，2， （1：新建广告系列， 2：使用已有广告系列）
            store_url: String,
            main_page_id: String,
            instagram_id: String,
            account_id: String,
            campaign_datas: Array,
            success_campaign: Array,//[{name: ,fb_id:}]
            success_adset: Array,
            success_ad: Array,
            fail_campaign: Array,
            fail_adset: Array,
            fail_ad: Array,
            start_time: String, //now为立即执行 或指定一个具体的PST时间执行,投放排期
            end_time: String,//2020-09-02 11:11:11
            create_user: String,
            cfg_frame_id: String,  // 广告结构方案id （v3版 新增字段）
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            }
        }
    }

    async save(data) {
        data._id = await require("../IdGenModel").genId('ad_flow', 1900000);
        return super.save(data);
    }

}

module.exports = new AdFlowModel();