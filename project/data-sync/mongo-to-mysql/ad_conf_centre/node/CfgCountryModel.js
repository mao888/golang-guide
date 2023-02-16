/**
 * @author zhanghao
 * @date 2020/8/11
 * @description: 国家组 数据模型
 */
const BaseModel = require('@cc/cargo').BaseModel;

class CfgCountryModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'CfgCountry';
    }
    getSchema() {
        return {
            _id: Number, //国家组id
            company_id: Number, //公司id
            name: String, //国家组名称
            fbid: String, //fbId gia迁移后废弃
            game_id: Number, // 游戏id
            geo_locations: {  // 包含国家，
                type: Object,
                default: undefined
            },
            excluded_geo_locations: { //排除国家,
                type: Object,
                default: undefined
            },
            contain_countries: { // 包含国家 的中文名称数组
                type: Array,
                default: undefined
            },
            excluded_countries: { // 排除国家的中文名称数据
                type: Array,
                default: undefined
            },
            create_user: String, // 创建人
            user_id: Number, //创建人id
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
        data._id = await require("../IdGenModel").genId('cfg_country', 1200001);
        let o = new this.model(data);
        return o.save();
    }
}

module.exports = new CfgCountryModel();