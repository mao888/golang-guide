const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class DownloadRecordModel extends BaseModel {

	constructor() {
		super('plat_console');
	}

	getName() {
		return 'DownloadRecord';
	}

	getSchema() {
		return {
			_id: {
				type: String,
				default: shortid.generate,
			},
			asset_id: {		//下载的美术资源id
				type: Number,
				ref: 'ArtSource'
			},
			user_id: Number, 	//下载用户id
			user_name: String, 	//下载用户名
			department_id: String,	 //部门id
			department_name: String, //部门名称(不变更)
			create_time: {
				type: Date,
				default: Date.now
			}
		};
	}
}

module.exports = new DownloadRecordModel();