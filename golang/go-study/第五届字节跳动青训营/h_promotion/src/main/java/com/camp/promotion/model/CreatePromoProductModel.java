package com.camp.promotion.model;

import com.camp.promotion.convertet.Converter;
import com.camp.promotion.entity.HPromoProduct;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotNull;
import java.util.function.Function;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreatePromoProductModel implements Converter<CreatePromoProductModel, HPromoProduct> {

    /**
     * 活动promo id
     */
    private Long promoId;
    /**
     * 活动名称
     */
    private String promoName;
    /**
     * 活动spu id
     */
    private Long spuId;
    /**
     * 活动sku id
     */
    private Long skuId;
    /**
     * 数量
     */
    private Integer promoStock;
    /**
     * 价格
     */
    private Integer promoPrice;

    @Override
    public HPromoProduct convert(Function<CreatePromoProductModel, HPromoProduct> f) {
        return f.apply(this);
    }
}
