const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class CategoryModel extends BaseModel {

	constructor() {
		super('plat_console');
	}

	getName() {
		return 'Category';
	}

	getSchema() {
		return {
			_id: Number,
			company_id: Number, //绑定公司id，若为0则为通用型
			parent_id: { 		//父节点id，若为0则为一级节点
				type: Number,
				default: 0
			},
			ctype: {	//art_store美术库；art_tag美术库标签；
				type: String,
				enum: ['art_store', 'art_tag']
			},
			level: {
				type: Number,
				default: 1
			},
			name: {
				type: String,
				default: '',
			},
			enable: { //是否被禁用
				type: Boolean,
				default: true
			},
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
		data._id = await require('../IdGenModel').genId('category_id');
		return super.save(data);
	}
}

module.exports = new CategoryModel();