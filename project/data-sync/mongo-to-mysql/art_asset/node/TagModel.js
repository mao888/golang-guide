const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class TagModel extends BaseModel {

	constructor() {
		super('plat_console');
	}

	getName() {
		return 'Tag';
	}

	getSchema() {
		return {
            _id: String, //company_id + '-' + ttype
            tag_list: [{
                _id: Number,
                name: String,
                company_id: {
                    type: Number,
                    ref: 'Company'
                },
                ttype: { //类型：art_store美术库标签
                    type: String,
                    enum: ['art_store']
                },
                category_id: { //标签分类
                    type: Number,
                    ref: 'Category',
                    default: 0
                },
                create_time: {
                    type: Date,
                    default: Date.now
                },
                update_time: {
                    type: Date,
                    default: Date.now
                }
            }]
		};
	}

	// async save(data) {
	// 	data._id = await require('../IdGenModel').genId('tag_id');
	// 	data.sort = await require('../IdGenModel').genId('tag_sort_id', 101);
	// 	return super.save(data);
	// }
}

module.exports = new TagModel();