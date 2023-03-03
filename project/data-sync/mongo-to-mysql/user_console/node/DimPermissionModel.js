const { BaseModel } = require('@cc/cargo');
const shortid = require('shortid');

class DimPermissionModel extends BaseModel {

    constructor() {
        super('rambler');
    }

    getName() {
        return 'DimPermission';
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate,
            },
            company_id: Number, // 公司id
            user_id: Number, //  用户id
            app_ids: { type: Array, default: undefined },
            account_ids: { type: Array, default: undefined },
            platforms: { type: Array, default: undefined },
            game_ids: { type: Array, default: undefined },
            delete_time: Date,
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
            child_system: {
                type: String,
                default: 'bi',
            },
        };
    }
}

module.exports = new DimPermissionModel();