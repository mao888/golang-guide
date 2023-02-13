package com.camp.promotion.model;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class OrderModel {

    /**
     * 订单号
     */
    private Long orderId;
    /**
     * 用户id
     */
    private Long userId;
    /**
     * 活动spu id
     */
    private Long spuId;
    /**
     * 活动sku id
     */
    private Long skuId;
    /**
     * 店铺id
     */
    private Long shopId;
    /**
     * 实际付款金额,单位是元,保留两位小数
     */
    private Integer payment;
    /**
     * 支付类型,1-在线支付
     */
    private Integer paymentType;
    /**
     * 运费,单位是元
     */
    private Integer postage;
    /**
     * 商品数量
     */
    private Integer quantity;
    /**
     * 订单状态:0-已取消-1-未付款，2-已付款，3-已发货，4-交易成功，5-交易关闭
     */
    private Integer status;
    /**
     * 地址id
     */
    private Long addressId;
    /**
     * 收货地址
     */
    private String shippingAddress;
    /**
     * 活动id
     */
    private Long activityId;

}
