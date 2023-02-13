package com.camp.promotion.model;

import com.camp.promotion.convertet.ConvertFunction;
import com.camp.promotion.convertet.Converter;
import com.camp.promotion.entity.HOrder;
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
public class CreateOrderModel implements Converter<CreateOrderModel, HOrder> {
    /**
     * 用户id
     */
    private Long userId;
    /**
     * 活动id
     */
    private Long promoId;
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
    private Integer quantity;
    /**
     * 收货地址id
     */
    private Long addressId;
    /**
     * 收货地址
     */
    private String shippingAddress;

    @Override
    public HOrder convert(Function<CreateOrderModel, HOrder> f) {
        return f.apply(this);
    }
}
