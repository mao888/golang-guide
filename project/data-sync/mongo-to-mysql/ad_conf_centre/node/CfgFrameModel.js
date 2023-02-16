/**
 * @author zhanghao
 * @date 2020/8/12
 * @description: 结构方案数据模型
 */
const BaseModel = require('@cc/cargo').BaseModel;

class CfgFrameModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'CfgFrame';
    }
    getSchema() {
        return {
            _id: Number, // 结构方案id
            company_id: Number, //公司id
            name: String, // 结构方案名称
            campaign_dims: {  // campaign划分维度
                type: Array, //['countries', 'audiences', 'positions', 'age', 'sex', 'language', 'materials', 'tag', 'adtype'],
                default: undefined
            },
            adset_dims: {  // adset划分维度
                type: Array,
                default: undefined
            },
            campaign_limit: Number, // campaign数量上限
            adset_limit: Number, // adset数量上限
            is_cbo: Boolean, // 是否开启cbo
            optimization_goal: String, // 优化目标
            bid_strategy: String, // 竞价策略
            attribution_spec: Object, // 转化窗口
            custom_event_type: String, //AEO的14个应用事件
            billing_event: {
                type: String,
                default: 'IMPRESSIONS'
            }, //计费方式
            pacing_type: {
                type: Array,
                default: ['standard']
            },
            budget_limit: Number, // 单日预算上限
            verify_str: String,  // 结构方案信息校验参数
            create_user: String, // 创建人
            user_id: Number, // 创建人
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
        data._id = await require("../IdGenModel").genId('cfg_frame', 1300001);
        let o = new this.model(data);
        return o.save();
    }
}

module.exports = new CfgFrameModel();