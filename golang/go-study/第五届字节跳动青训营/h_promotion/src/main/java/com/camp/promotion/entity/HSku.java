package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HSku)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
public class HSku implements Serializable {
    private static final long serialVersionUID = -80891613527625211L;
    /**
     * id
     */
    private Long id;
    /**
     * spu_id
     */
    private Long spuId;
    /**
     * 店铺id
     */
    private Long shopId;
    /**
     * 规格信息
     */
    private String specDetailIds;
    /**
     * 库存
     */
    private Integer stock;
    /**
     * 价格,销售价
     */
    private Integer price;
    /**
     * 商品状态.0-无需，1-有效
     */
    private Integer status;
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

    public Long getSpuId() {
        return spuId;
    }

    public void setSpuId(Long spuId) {
        this.spuId = spuId;
    }

    public Long getShopId() {
        return shopId;
    }

    public void setShopId(Long shopId) {
        this.shopId = shopId;
    }

    public String getSpecDetailIds() {
        return specDetailIds;
    }

    public void setSpecDetailIds(String specDetailIds) {
        this.specDetailIds = specDetailIds;
    }

    public Integer getStock() {
        return stock;
    }

    public void setStock(Integer stock) {
        this.stock = stock;
    }

    public Integer getPrice() {
        return price;
    }

    public void setPrice(Integer price) {
        this.price = price;
    }

    public Integer getStatus() {
        return status;
    }

    public void setStatus(Integer status) {
        this.status = status;
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

