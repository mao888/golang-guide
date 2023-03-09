/**
 * @author zhanghao
 * @date 2020/9/16
 * @description:
 */

const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class AdFlowCfgModel extends BaseModel {

    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return "adflowcfg";
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            o_id: String,
            fbid: String,
            campaign_id: String, //  广告系列id
            adset_id: String,  // 广告组id
            cfg_frame_id: String, // 结构方案id
            cfg_country_ids: Array, // 国家组id
            cfg_audience_ids: Array, // 受众组id
            cfg_position_ids: Array, // 版位组id
            ages: Array, // 年龄
            genders: Array,// 性别,
            contrys: Array,
            language_groups: Array, // 年龄组 [{ name: '国家组名称', languages: ['en', 'zh'] }]
            strategys: Array, // 优化方式
            user_os: String,  //  设备系统
            user_device: Array, // 包含设备
            excluded_user_device: Array, // 排除设备
            ver_min: Number, // 最低版本号
            ver_max: Number, // 最高版本号
            is_wifi: Boolean,  // 链接wifi
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

}

module.exports = new AdFlowCfgModel();