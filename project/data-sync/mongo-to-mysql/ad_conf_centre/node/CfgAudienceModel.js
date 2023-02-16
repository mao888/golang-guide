/**
 * @author zhanghao
 * @date 2020/8/11
 * @description: 受众组 数据模型
 */
const BaseModel = require('@cc/cargo').BaseModel;

class CfgAudienceModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'CfgAudience';
    }
    getSchema() {
        return {
            _id: Number, // 受众组id
            company_id: Number, //公司id
            name: String, // 受众组名称
            account_id: String, // 账户id
            custom_audiences: {  // 包含受众
                type: Array,
                default: undefined
            },
            excluded_custom_audiences: { //排除受众,
                type: Array,
                default: undefined
            },
            flexible_spec: { // 包含细分定位、缩小细分定位
                type: Array,
                default: undefined
            },
            exclusions: { // 排除细分定位
                type: Object,
                default: undefined
            },
            targeting_optimization: { // 细分定位扩展优化
                type: String,
                default: undefined
            },
            verify_str: String, // 受众信息校验参数
            create_user: String, // 创建人
            user_id: Number, // 创建人id
            source_sys: {
                type: String,
                default: 'gia'
            },
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
    async save(data) {
        data._id = await require("../IdGenModel").genId('cfg_audience', 1100001);
        let o = new this.model(data);
        return o.save();
    }
}

module.exports = new CfgAudienceModel();