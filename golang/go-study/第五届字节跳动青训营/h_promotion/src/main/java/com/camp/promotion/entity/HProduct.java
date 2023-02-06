package com.camp.promotion.entity;

import com.camp.promotion.convertet.ConvertFunction;
import com.camp.promotion.convertet.Converter;
import com.camp.promotion.model.HPromoSpuModel;

import java.util.Date;
import java.io.Serializable;
import java.util.function.Function;

/**
 * (HProduct)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:24
 */
public class HProduct implements Serializable, Converter<HProduct, HPromoSpuModel> {
    private static final long serialVersionUID = -39693500405111662L;
    /**
     * 商品id
     */
    private Long id;
    /**
     * 分类id,对应category表的主键
     */
    private Long categoryId;
    /**
     * 店铺id
     */
    private Long shopId;
    /**
     * 商品名称
     */
    private String title;
    /**
     * 商品副标题
     */
    private String subtitle;
    /**
     * 产品主图,url相对地址
     */
    private String mainImage;
    /**
     * 图片地址,json格式,扩展用
     */
    private String subImages;
    /**
     * 商品详情
     */
    private String detail;
    /**
     * 价格,吊牌价
     */
    private Integer price;
    /**
     * 类目属性信息，json
     */
    private String categoryData;
    /**
     * 规格信息，json
     */
    private String specData;
    /**
     * 商品状态.1-在售 2-下架 3-删除
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

    public Long getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(Long categoryId) {
        this.categoryId = categoryId;
    }

    public Long getShopId() {
        return shopId;
    }

    public void setShopId(Long shopId) {
        this.shopId = shopId;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getSubtitle() {
        return subtitle;
    }

    public void setSubtitle(String subtitle) {
        this.subtitle = subtitle;
    }

    public String getMainImage() {
        return mainImage;
    }

    public void setMainImage(String mainImage) {
        this.mainImage = mainImage;
    }

    public String getSubImages() {
        return subImages;
    }

    public void setSubImages(String subImages) {
        this.subImages = subImages;
    }

    public String getDetail() {
        return detail;
    }

    public void setDetail(String detail) {
        this.detail = detail;
    }

    public Integer getPrice() {
        return price;
    }

    public void setPrice(Integer price) {
        this.price = price;
    }

    public String getCategoryData() {
        return categoryData;
    }

    public void setCategoryData(String categoryData) {
        this.categoryData = categoryData;
    }

    public String getSpecData() {
        return specData;
    }

    public void setSpecData(String specData) {
        this.specData = specData;
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

    @Override
    public HPromoSpuModel convert(Function<HProduct, HPromoSpuModel> f) {
        return f.apply(this);
    }

}

