const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class CloudUrlModel extends BaseModel {

    constructor() {
        super('plat_console');
    }

    getName() {
        return 'CloudUrl';
    }

    getSchema() {
        return {
            _id: Number,
            company_id: {
                type: Number,
                ref: 'Company'
            },
            asset_id: Number, 	  //关联到第三方资源类型id，如美术资源库id
            cloud_type: {		  //资源存储的CDN类型，默认Akamai
                type: String,
                default: 'Akamai'
            },
            use_type: {			 // 资源用途：art_store_source美术资源库源文件；art_store_thumb美术资源库预览图
                type: String,
                enum: ['art_store_source', 'art_store_thumb', 'player_feedback']
            },
            source_type: { 		// 资源类型video、image、zip、file
                type: String,
                enum: ['video', 'image', 'zip', 'file'],
                default: 'file'
            },
            name: String,		 // 资源名称
            cloud_dir: String,   // 云地址路径
            url: String,		 // 资源地址
            thumbnail_url: String,		 // 资源缩略图地址
            size: Number, 		 // 资源大小
            suffix: String, 	 // 资源后缀名
            create_time: { 		 // 创建时间
                type: Date,
                default: Date.now
            }
        };
    }

    async save(data) {
        data._id = await require('../IdGenModel').genId('cloudurl_id');
        return super.save(data);
    }

}

module.exports = new CloudUrlModel();