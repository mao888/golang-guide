/**
 * @author zhanghao
 * @date 2021/02/10
 * @description: 应用参数数据模型
 */
const BaseModel = require('@cc/cargo').BaseModel;

class CfgAppModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }
    getName() {
        return 'CfgApp';
    }
    getSchema() {
        return {
            _id: Number,
            company_id: Number,
            game_id: Number,
            git_address: String,
            onelink_id: String,
            conversion_params: [{  // 转化量参数
                token: String,  // 转化量广告token
                fbh5_domain_name: String, // f5 广告域名配置
                ci_task: String,  //  任务编号
                pixel: String,   //  像素代码
                pixel_id: String,
            }],
            user_id: Number, // 创建人id
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
        data._id = await require("../IdGenModel").genId('cfg_game', 210001);
        let o = new this.model(data);
        return o.save();
    }
}

module.exports = new CfgAppModel();