const BaseModel = require('@cc/cargo').BaseModel;
const shortid = require('shortid');

class FBAdsetModel extends BaseModel {
    constructor() {
        super('cruiser_console_v2');
    }

    getName() {
        return 'FBAdset';
    }

    getSchema() {
        return {
            _id: {
                type: String,
                default: shortid.generate
            },
            status: { //状态
                type: String,
                enum: ['ACTIVE', 'PAUSED', 'DELETED', 'ARCHIVED']
            },
            adset_id: String,
            adset_name: String,
            campaign_id: String,
            account_id: String,
            created_by_cruiser: {
                type: Boolean,
                default: true
            }, //是否为自动化创建
            clone_template_open: { //克隆模板是否开启
                type: Boolean,
                default: false
            },
            create_user: String,//创建人
            user_id: Number,
            //受众
            targeting: {
                custom_audiences: {
                    type: Array,
                    default: undefined
                }, //自定义受众
                excluded_custom_audiences: {
                    type: Array,
                    default: undefined
                }, //排除自定义受众

                connections: {
                    type: Array,
                    default: undefined
                }, //关系 与以下内容建立关系网络的用户
                excluded_connections: {
                    type: Array,
                    default: undefined
                }, //排除与以下内容建立关系网络的用户
                friends_of_connections: {
                    type: Array,
                    default: undefined
                }, //与以下内容建立关系网络的用户的好友
                app_install_state: String, //not_installed 设置关系时加入该属性

                geo_locations: {
                    location_types: {
                        type: Array,
                        default: undefined
                    },
                    countries: {
                        type: Array,
                        default: undefined
                    },
                    country_groups: {
                        type: Array,
                        default: undefined
                    }
                },
                age_min: Number,
                age_max: Number,
                genders: {
                    type: Array,
                    default: undefined
                },
                locales: {
                    type: Array,
                    default: undefined
                }, //语言
                flexible_spec: {
                    type: Array,
                    default: undefined
                }, //细分定位 包含至少符合一项条件的用户
                exclusions: Object, //细分定位 排除符合以下至少一项条件的用户
                targeting_optimization: { //expansion_all-当能够以更低的安装费用提高安装量时，系统就会扩展兴趣范围
                    type: String,
                    enum: ['expansion_all', 'none']
                },

                user_os: {
                    type: Array,
                    default: undefined
                },
                user_device: {
                    type: Array,
                    default: undefined
                }, //包含的设备
                excluded_user_device: {
                    type: Array,
                    default: undefined
                }, //被排除的设备
                wireless_carrier: {
                    type: Array,
                    default: undefined
                }, //仅在连接wifi时 ["Wifi"]

                //-----------版位---------
                publisher_platforms: {
                    type: Array,
                    default: undefined
                },//enum: ['facebook', 'instagram', 'messenger', 'audience_network']
                facebook_positions: {
                    type: Array,
                    default: undefined
                },//enum: ['feed', 'right_hand_column', 'instant_article', 'marketplace', 'suggested_video', 'instream_video', 'story']
                instagram_positions: {
                    type: Array,
                    default: undefined
                },//enum: ['stream', 'story']
                messenger_positions: {
                    type: Array,
                    default: undefined
                },//enum: ['messenger_home', 'sponsored_messages', 'story']
                audience_network_positions: {
                    type: Array,
                    default: undefined
                },//enum: ['classic', 'instream_video', 'rewarded_video']
                //-----------版位---------
            },

            //---------未开启cbo时特有字段-----
            bid_strategy: { //竞价策略，比campaign层级少两个
                type: String,
                enum: ['LOWEST_COST_WITHOUT_CAP', 'LOWEST_COST_WITH_BID_CAP', 'COST_CAP', 'TARGET_COST', 'LOWEST_COST_WITH_MIN_ROAS']
            },
            daily_budget: Number,//单日预算
            pacing_type: {
                type: Array,
                default: undefined
            },//投放时段或类型说明字段 enum: ['standard', 'no_pacing', 'day_parting']

            //---------未开启cbo时特有字段-----

            //---------开启cbo时字段----------
            daily_min_spend_target: Number, //开启cbo后广告组限额下限-每日预算
            daily_spend_cap: Number, //开启cbo后广告组限额上限-每日预算
            lifetime_min_spend_target: Number, //开启cbo后广告组限额下限-总预算
            lifetime_spend_cap: Number, //开启cbo后广告组限额上限-总预算
            bid_constraints: { //roas竞价策略值
                roas_average_floor: Number
            },
            start_time: String,//cbo开启时，广告系列选择总预算，广告组排期选择开始和结束时间
            end_time: String,
            //---------开启cbo时字  段----------

            attribution_spec: {
                type: Array,
                default: undefined
            }, //转化时间窗
            bid_amount: Number,
            optimization_goal: {//优化方式
                type: String,
                enum: ['APP_INSTALLS', 'VALUE', 'OFFSITE_CONVERSIONS', 'LINK_CLICKS', 'DERIVED_EVENTS', 'LANDING_PAGE_VIEWS']
            },
            billing_event: {
                type: String,
                enum: ['APP_INSTALLS', 'IMPRESSIONS', 'LINK_CLICKS']
            }, //计费方式
            positionType: { //自定义字段
                type: String,
                enum: ['autoPosition', 'editPosition']
            },
            promoted_object: {
                application_id: String,
                object_store_url: String,
                custom_event_type: String, //当优化目标为OFFSITE_CONVERSIONS optimize_event_value值会填充到custom_event_type属性
            },//缺少时报错找不到推广对象
            create_time: {
                type: Date,
                default: Date.now
            },
            create_type: String,
            company_id: Number,
        }
    }
}

module.exports = new FBAdsetModel();