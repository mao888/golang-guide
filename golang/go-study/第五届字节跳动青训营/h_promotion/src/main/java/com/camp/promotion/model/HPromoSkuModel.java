package com.camp.promotion.model;

import com.camp.promotion.convertet.ConvertFunction;
import com.camp.promotion.convertet.Converter;
import com.camp.promotion.convertet.MultiConverter;
import com.camp.promotion.entity.HPromoProduct;
import com.camp.promotion.entity.HSku;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.function.Function;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
@JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
public class HPromoSkuModel implements Serializable, MultiConverter<HPromoProduct, HSku, HPromoSkuModel> {

    private static final long serialVersionUID = 5555514090859900792L;
    private HPromoProduct promoProduct;
    private HSku sku;

    public HPromoSkuModel(HPromoProduct promoProduct, HSku sku) {
        this.promoProduct = promoProduct;
        this.sku = sku;
    }

    /**
     * 活动spu id
     */
    private Long spuId;
    /**
     * 活动sku id
     */
    private Long skuId;
    /**
     * 活动库存
     */
    private Integer promoStock;
    /**
     * 价格,秒杀价
     */
    private Integer promoPrice;
    /**
     * 店铺id
     */
    private Long shopId;
    /**
     * 规格信息
     */
    private String specDetailIds;
    /**
     * 价格,销售价
     */
    private Integer price;

    public HPromoSkuModel convert(ConvertFunction<HPromoProduct, HSku, HPromoSkuModel> f) {
        return f.apply(this.promoProduct, this.sku);
    }
}
