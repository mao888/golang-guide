package com.camp.promotion.model;

import com.camp.promotion.common.BizException;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.enums.PromoStatusEnum;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;
import java.util.Date;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class HPromoProductModel implements Serializable {

    private static final long serialVersionUID = -81907624022631674L;
    /**
     * 活动id
     */
    private Long promoId;
    /**
     * 活动名称
     */
    private String promoName;
    /**
     * 开始时间
     */
    private Date startDate;
    /**
     * 结束时间
     */
    private Date endDate;
    /**
     * 活动状态.0-创建，1-上线中，2-已下线
     */
    private Integer status;
    /**
     * 秒杀spu信息
     */
    private HPromoSpuModel spuModel;
    /**
     * 秒杀sku信息
     */
    private HPromoSkuModel skuModel;

    public boolean checkPromo() {
        Date now = new Date();
        if (this.getStatus() != PromoStatusEnum.ONLINE.getCode()
                || this.getEndDate().before(now) || this.getStartDate().after(now)) {
            return false;
        }
        return true;
    }
}
