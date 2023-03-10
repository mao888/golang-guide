const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

// 素材
class AdFlowModel extends BaseModel {

    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return "JenkinsBuildInfo";
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            ci_code: String, //Jenkins 任务编号
            adflow_id: Number,  // 广告投入信息id
            build_number: Number, // Jenkins Build 编号
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
        }
    }
}

module.exports = new AdFlowModel();