/**
 * @author zhanghao
 * @date 2020/8/11
 * @description: 版位组 数据模型
 */
const BaseModel = require('@cc/cargo').BaseModel;

class CfgPositionModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'CfgPosition';
    }
    getSchema() {
        return {
            _id: Number, // 版位组id
            company_id: Number, //公司id
            name: String, // 版位组名称
            positionType: { //版位类型
                type: String,
                enum: ['autoPosition', 'editPosition'],
            },
            publisher_platforms: {  // 版位平台
                type: Array,
                default: undefined
            },
            facebook_positions: { //fb版位数组,
                type: Array,
                default: undefined
            },
            instagram_positions: { // nstagram 版位数组
                type: Array,
                default: undefined
            },
            messenger_positions: { // messenger 版位数组
                type: Array,
                default: undefined
            },
            audience_network_positions: { // audience_network 版位数组
                type: Array,
                default: undefined
            },
            containPosition: Array, // 所有版位的 对照数组信息
            create_user: String, // 创建人
            user_id: Number,
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
        data._id = await require("../IdGenModel").genId('cfg_position', 1400001);
        let o = new this.model(data);
        return o.save();
    }
}

module.exports = new CfgPositionModel();