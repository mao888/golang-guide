const BaseModel = require('@cc/cargo').BaseModel;

class AssetLogModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'assetlog';
    }

    getSchema() {
        return {
            asset_id: Number,
            asset_name: String,
            account_id: String,
            state: String,
            asset_type: String,
            asset_sign: String,
            create_time: {
                type: Date,
                default: Date.now
            }
        };
    }

}


module.exports = new AssetLogModel();