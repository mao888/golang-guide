const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

// 素材
class AdTextModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return "adtext";
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            fbid: String,
            en_text: String,
            translation: [{
                lang: Number,
                text: String
            }],
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            }
        }
    }
}

module.exports = new AdTextModel();