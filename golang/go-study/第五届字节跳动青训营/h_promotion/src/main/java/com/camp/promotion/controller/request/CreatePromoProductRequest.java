package com.camp.promotion.controller.request;

import lombok.Data;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;

@Data
public class CreatePromoProductRequest {

    /**
     * 活动spu id
     */
    @NotNull(message = "spuId不能为null")
    private Long spuId;
    /**
     * 活动sku id
     */
    @NotNull(message = "skuId不能为null")
    private Long skuId;
    /**
     * 数量
     */
    @Min(1)
    @NotNull(message = "promoStock不能为null")
    private Integer promoStock;
    /**
     * 价格
     */
    @Min(0)
    @NotNull(message = "promoPrice不能为null")
    private Integer promoPrice;
}
