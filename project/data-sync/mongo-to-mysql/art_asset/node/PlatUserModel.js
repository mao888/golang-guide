const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class PlatUserModel extends BaseModel {
    constructor() {
        super('plat_console');
    }

    getName() {
        return 'PlatUser';
    }

    getSchema() {
        return {
            _id: Number,
            name: {
                type: String,
                default: ''
            },//昵称
            username: String,//用户姓名
            password: String,
            email: String,
            phone: {
                type: String,
                default: ''
            },
            avatar: {
                type: String,
                default: ''
            },
            role: Array, //存放账号级别角色
            enable: { //该用户是否被激活
                type: Boolean,
                default: true
            },
            user_tag: { //账户类型
                type: Number,
                enum: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10], //1管理员大权限
                default: 0 //0无
            },
            token: String,
            token_expire_time: Date,
            comments: { //备注
                type: String,
                default: ''
            },
            create_time: {
                type: Date,
                default: Date.now
            },
            update_time: {
                type: Date,
                default: Date.now
            },
            login_time: { //最后登录时间
                type: Date,
                default: Date.now
            },
            maintain_status: { //维护状态
                type: Boolean,
                default: false
            },
            guider_step: { //新手引导
                type: Number,
                default: 0
            },
            access_system: [String], //可访问的系统
            default_company: Number, //当前选中公司
        };
    }

    async save(data) {
        data._id = await require('../IdGenModel').genId('user_id');
        return super.save(data);
    }

}

module.exports = new PlatUserModel();