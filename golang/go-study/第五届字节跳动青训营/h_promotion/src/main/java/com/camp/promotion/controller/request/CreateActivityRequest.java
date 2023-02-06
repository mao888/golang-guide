package com.camp.promotion.controller.request;

import com.camp.promotion.convertet.Converter;
import com.camp.promotion.model.CreateActivityModel;
import lombok.Data;

import javax.validation.constraints.NotEmpty;
import javax.validation.constraints.NotNull;
import java.util.List;
import java.util.function.Function;

@Data
public class CreateActivityRequest implements Converter<CreateActivityRequest, CreateActivityModel> {

    /**
     * 活动名称
     */
    @NotNull(message = "promoName不能为null")
    private String promoName;
    /**
     * 开始时间
     */
    @NotNull(message = "开始时间不能为null")
    private Long startTime;
    /**
     * 结束时间
     */
    @NotNull(message = "结束时间不能为null")
    private Long endTime;
    /**
     * 参与秒杀的商品
     */
    @NotEmpty(message = "秒杀商品不能为空")
    private List<CreatePromoProductRequest> promoProducts;


    @Override
    public CreateActivityModel convert(Function<CreateActivityRequest, CreateActivityModel> f) {
        return f.apply(this);
    }
}
