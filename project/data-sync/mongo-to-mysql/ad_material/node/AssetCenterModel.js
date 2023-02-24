const BaseModel = require('@cc/cargo').BaseModel;

// 素材
class AssetCenterModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return "assetcenter";
    }

    getSchema() {
        return {
            _id: Number,
            asset_type: { //素材类型
                type: String,
                enum: ['image', 'video'],
            },
            asset_md5: String, //md5
            asset_name: String,//素材名
            asset_url: String, //素材下载地址
            asset_thumbnail: String, //素材缩略图
            asset_width: Number, //素材宽
            asset_height: Number, //素材高
            asset_duration: {
                type: Number,
                default: 0
            }, //素材时长
            asset_language: Number,
            tag: String,//标签
            asset_size: String,//尺寸
            fbid: String, //
            media_list: Array, //该素材在每个账户中的对应ID [{account_id:"123", media_id:"678",creative_id:""}]
            create_user: String,
            creative_user: String,//创意人员id
            design_user: String,//设计人员id
            asset_url_info: String,//包含完整域名的url
            asset_thumbnail_url: String,////包含完整域名的url
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
            delete_time: Date //假删除对应字段
        }
    }

    async save(data) {
        data._id = await require("../IdGenModel").genId('cruiser_assets', 1000000);
        return super.save(data);
    }

    async insertMany(arr) {
        for (let item of arr) {
            item._id = await require("../IdGenModel").genId('cruiser_assets', 1000000);
        }
        return super.insertMany(arr);
    }
}

module.exports = new AssetCenterModel();