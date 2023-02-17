const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class ArtSourceModel extends BaseModel {

    constructor() {
        super('plat_console');
    }

    getName() {
        return 'ArtSource';
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
            category_id: {
                type: Number,
                ref: 'Category'
            },
            tag_arr: Array,  	//标签，上限10
            thumb_arr: Array, 	//预览图，上限20
            // source_arr: Array,	//资源，上限20  废弃
            download_count: {	//下载次数
                type: Number,
                default: 0
            },
            creator_id: Number,			//上传用户id
            // support_id: String,			//废弃   供应部门id
            // support_department: String,	//废弃  供应部门
            author: Number, //  作者
            game_id: Number, //  游戏id
            download_url: String, //UE下载地址
            maya_download_url: String, //Maya下载地址
            done_time: Date,  // 完成时间 
            relation_artneeds: Array, // 关联需求
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date
            }
        };
    }

    async save(data) {
        data._id = await require('../IdGenModel').genId('art_source');
        return super.save(data);
    }

}

module.exports = new ArtSourceModel();