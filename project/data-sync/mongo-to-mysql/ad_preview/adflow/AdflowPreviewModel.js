/**
 * @author zhanghao
 * @date 2020/8/14
 * @description:  广告预览及创建配置信息
 */
const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class AdflowPreviewModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'AdflowPreview';
    }
    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            fbid: String,
            account_id: String,
            struct_cfg: Object,
            campaign_datas: [{
                name: String,
                budget: {
                    type: Number,
                    default: undefined
                },
                adset_datas: [{
                    name: String,
                    budget: {
                        type: Number,
                        default: undefined
                    },
                    cost: {
                        type: Number,
                        default: undefined
                    },
                    targeting_data: Array,
                    ad_datas: Array
                }]
            }],
            create_user: String, // 创建人
            update_time: {
                type: Date,
                default: Date.now
            },
            create_time: {
                type: Date,
                default: Date.now
            },
        };
    }
}

module.exports = new AdflowPreviewModel();