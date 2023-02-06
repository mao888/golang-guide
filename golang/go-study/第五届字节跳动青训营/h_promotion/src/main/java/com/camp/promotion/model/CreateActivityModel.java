package com.camp.promotion.model;

import com.camp.promotion.convertet.Converter;
import com.camp.promotion.entity.HPromo;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.function.Function;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class CreateActivityModel implements Converter<CreateActivityModel, HPromo> {

    /**
     * 活动名称
     */
    private String promoName;
    /**
     * 开始时间
     */
    private Long startTime;
    /**
     * 结束时间
     */
    private Long endTime;
    /**
     * 参与秒杀的商品
     */
    private List<CreatePromoProductModel> promoProducts;

    @Override
    public HPromo convert(Function<CreateActivityModel, HPromo> f) {
        return f.apply(this);
    }
}
