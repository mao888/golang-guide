const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class ActiveLibraryModel extends BaseModel {

    constructor() {
        super('plat_console');
    }

    getName() {
        return 'ActiveLibrary';
    }

    getSchema() {
        return {
            _id: Number,
            company_id: {
                type: Number,
                ref: 'Company'
            },
            name: String,
            desc: String,
            size: String,
            category_id: {
                type: Number,
                ref: 'Category'
            },
            tag_arr: [{
                type: Number,
                ref: 'Tag'
            }], //标签，上限10
            thumb_url: String, 	//预览图
            url: String, 	//视频URL
            creator_id: Number,			//上传用户id
            author: Number, //  作者
            game_id: Number, //  游戏id
            ue_download_url: String, //UE下载地址
            maya_download_url: String, //Maya下载地址
            done_time: Date,  // 完成时间
            relation_artneeds: Array, // 关联需求
            deleted_time: Date, // 删除状态
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            }
        };
    }

    async save(data) {
        data._id = await require('../IdGenModel').genId('active_library', 2000000);
        return super.save(data);
    }

}

module.exports = new ActiveLibraryModel();