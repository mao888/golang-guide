CREATE TABLE `h_category`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `parent_id`   bigint                                 DEFAULT '0' COMMENT '父类目id当id=0时说明是根节点,一级类目',
    `name`        varchar(50) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '类目名称',
    `level`       tinyint                                DEFAULT '0' COMMENT '类目级别1-一级,2-二级，3-三级',
    `be_parent`   tinyint                                DEFAULT '0' COMMENT '是否父类目0-否，1-是',
    `status`      tinyint                                DEFAULT '0' COMMENT '类目状态0-废弃，1-正常',
    `sort_order`  int                                    DEFAULT NULL COMMENT '排序编号,同类展示顺序,数值相等则自然排序',
    `create_time` timestamp NOT NULL                     DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL                     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name_parent` (`parent_id`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_category_attribute`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `category_id`   bigint                                  DEFAULT '0' COMMENT '类目id',
    `attribute_key` varchar(100) COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '类目属性名称',
    `status`        tinyint   NOT NULL                      DEFAULT '0' COMMENT '类目状态0-废弃，1-正常',
    `create_time`   timestamp NOT NULL                      DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   timestamp NOT NULL                      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY             `idx_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_order`
(
    `id`               bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
    `user_id`          bigint                                                        NOT NULL DEFAULT '0' COMMENT '用户id',
    `spu_id`           bigint                                                        NOT NULL DEFAULT '0' COMMENT '活动spu id',
    `sku_id`           bigint                                                        NOT NULL DEFAULT '0' COMMENT '活动sku id',
    `shop_id`          bigint                                                        NOT NULL DEFAULT '0' COMMENT '店铺id',
    `payment`          int                                                           NOT NULL DEFAULT '0' COMMENT '实际付款金额,单位是元,保留两位小数',
    `payment_type`     tinyint                                                       NOT NULL DEFAULT '0' COMMENT '支付类型,1-在线支付',
    `postage`          int                                                           NOT NULL DEFAULT '0' COMMENT '运费,单位是元',
    `quantity`         int                                                           NOT NULL DEFAULT '0' COMMENT '商品数量',
    `address_id`       bigint                                                        NOT NULL DEFAULT '0' COMMENT '地址id',
    `shipping_address` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '收货地址',
    `status`           tinyint                                                       NOT NULL DEFAULT '0' COMMENT '订单状态:0-已取消-1-未付款，2-已付款，3-已发货，4-交易成功，5-交易关闭',
    `payment_time`     timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '支付时间',
    `send_time`        timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发货时间',
    `end_time`         timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易完成时间',
    `close_time`       timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '交易关闭时间',
    `activity_id`      bigint                                                        NOT NULL DEFAULT '0' COMMENT '活动id',
    `create_time`      timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`      timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_product`
(
    `id`            bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `category_id`   bigint                                  NOT NULL DEFAULT '0' COMMENT '分类id,对应category表的主键',
    `shop_id`       bigint                                  NOT NULL DEFAULT '0' COMMENT '店铺id',
    `title`         varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品名称',
    `subtitle`      varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商品副标题',
    `main_image`    varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '产品主图,url相对地址',
    `sub_images`    text COLLATE utf8mb4_general_ci COMMENT '图片地址,json格式,扩展用',
    `detail`        text COLLATE utf8mb4_general_ci COMMENT '商品详情',
    `price`         int                                     NOT NULL DEFAULT '0' COMMENT '价格,吊牌价',
    `category_data` varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类目属性信息，json',
    `spec_data`     varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格信息，json',
    `status`        tinyint                                 NOT NULL DEFAULT '0' COMMENT '商品状态.1-在售 2-下架 3-删除',
    `create_time`   timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`   timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY             `idx_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_promo`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `promo_name`  varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动名称',
    `start_date`  timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
    `end_date`    timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
    `status`      tinyint                                 NOT NULL DEFAULT '0' COMMENT '商品状态.0-创建，1-上线中，2-已下线',
    `create_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_promo` (`promo_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_promo_product`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `promo_id`    bigint                                  NOT NULL DEFAULT '0' COMMENT '活动id',
    `promo_name`  varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '活动名称',
    `spu_id`      bigint                                  NOT NULL DEFAULT '0' COMMENT '活动spu id',
    `sku_id`      bigint                                  NOT NULL DEFAULT '0' COMMENT '活动sku id',
    `promo_stock` int                                     NOT NULL DEFAULT '0' COMMENT '活动库存',
    `promo_price` int                                     NOT NULL DEFAULT '0' COMMENT '价格,秒杀价',
    `create_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_promo_sku` (`promo_id`,`sku_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_sku`
(
    `id`              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `spu_id`          bigint                                  NOT NULL DEFAULT '0' COMMENT 'spu_id',
    `shop_id`         bigint                                  NOT NULL DEFAULT '0' COMMENT '店铺id',
    `spec_detail_ids` varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格信息',
    `stock`           int                                     NOT NULL DEFAULT '0' COMMENT '库存',
    `price`           int                                     NOT NULL DEFAULT '0' COMMENT '价格,销售价',
    `status`          tinyint                                 NOT NULL DEFAULT '0' COMMENT '商品状态.0-无需，1-有效',
    `create_time`     timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`     timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY               `idx_spu` (`spu_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_spec`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `category_id` bigint                                  NOT NULL DEFAULT '0' COMMENT '分类id,对应category表的主键',
    `spec_key`    varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格名称',
    `status`      tinyint                                 NOT NULL DEFAULT '0' COMMENT '类目状态0-废弃，1-正常',
    `create_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_category` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_spec_detail`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `spec_id`     bigint                                  NOT NULL DEFAULT '0' COMMENT '规格',
    `spec_key`    varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格属性名称',
    `spec_value`  varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规格属性值',
    `status`      tinyint                                 NOT NULL DEFAULT '0' COMMENT '类目状态0-废弃，1-正常',
    `create_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY           `idx_spec` (`spec_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_user`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户表id',
    `username`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
    `password`    varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
    `email`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone`       varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
    `create_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `h_user_address`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户表id',
    `user_id`     bigint                                                        NOT NULL DEFAULT '0' COMMENT 'user id',
    `address`     varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '地址信息',
    `status`      tinyint                                                       NOT NULL DEFAULT '0' COMMENT '状态.0-废弃 1-有效',
    `be_default`  tinyint                                                       NOT NULL DEFAULT '0' COMMENT '是否默认.0-no 1-yes',
    `create_time` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp                                                     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;