package com.camp.promotion.controller.request;

import com.camp.promotion.convertet.ConvertFunction;
import com.camp.promotion.convertet.Converter;
import com.camp.promotion.model.CreateOrderModel;
import lombok.Data;

import javax.validation.constraints.Min;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import java.util.function.Function;

@Data
public class CreateOrderRequest implements Converter<CreateOrderRequest, CreateOrderModel> {

    /**
     * 用户id
     */
    @NotNull(message = "用户id不能为null")
    private Long userId;
    /**
     * 活动id
     */
    @NotNull(message = "秒杀活动id不能为null")
    private Long promoId;
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
    @NotNull(message = "quantity不能为null")
    private Integer quantity;
    /**
     * 价格
     */
    @Min(0)
    @NotNull(message = "promoPrice不能为null")
    private Integer promoPrice;
    /**
     * 总价
     */
    @Min(0)
    @NotNull(message = "totalAmount不能为null")
    private Integer totalAmount;
    /**
     * 收货地址id
     */
    @NotNull(message = "addressId不能为null")
    private Long addressId;
    /**
     * 收货地址
     */
    @NotBlank(message = "shippingAddress不能为null")
    private String shippingAddress;


    @Override
    public CreateOrderModel convert(Function<CreateOrderRequest, CreateOrderModel> f) {
        return f.apply(this);
    }
}
