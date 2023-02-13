package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HPromoProduct)实体类
 *
 * @author makejava
 * @since 2022-12-03 11:01:30
 */
public class HPromoProduct implements Serializable {
    private static final long serialVersionUID = -72847303198394301L;
    /**
     * id
     */
    private Long id;
    /**
     * 活动id
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
     * 活动库存
     */
    private Integer promoStock;
    /**
     * 价格,秒杀价
     */
    private Integer promoPrice;
    /**
     * 创建时间
     */
    private Date createTime;
    /**
     * 更新时间
     */
    private Date updateTime;


    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getPromoId() {
        return promoId;
    }

    public void setPromoId(Long promoId) {
        this.promoId = promoId;
    }

    public String getPromoName() {
        return promoName;
    }

    public void setPromoName(String promoName) {
        this.promoName = promoName;
    }

    public Long getSpuId() {
        return spuId;
    }

    public void setSpuId(Long spuId) {
        this.spuId = spuId;
    }

    public Long getSkuId() {
        return skuId;
    }

    public void setSkuId(Long skuId) {
        this.skuId = skuId;
    }

    public Integer getPromoStock() {
        return promoStock;
    }

    public void setPromoStock(Integer promoStock) {
        this.promoStock = promoStock;
    }

    public Integer getPromoPrice() {
        return promoPrice;
    }

    public void setPromoPrice(Integer promoPrice) {
        this.promoPrice = promoPrice;
    }

    public Date getCreateTime() {
        return createTime;
    }

    public void setCreateTime(Date createTime) {
        this.createTime = createTime;
    }

    public Date getUpdateTime() {
        return updateTime;
    }

    public void setUpdateTime(Date updateTime) {
        this.updateTime = updateTime;
    }

}

