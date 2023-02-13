package com.camp.promotion.entity;

import com.camp.promotion.convertet.ConvertFunction;
import com.camp.promotion.convertet.Converter;
import com.camp.promotion.model.HPromoProductModel;

import java.util.Date;
import java.io.Serializable;
import java.util.function.Function;

/**
 * (HPromo)实体类
 *
 * @author makejava
 * @since 2022-12-03 10:59:51
 */
public class HPromo implements Serializable, Converter<HPromo, HPromoProductModel> {
    private static final long serialVersionUID = 765691746086102277L;
    /**
     * id
     */
    private Long id;
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
     * 商品状态.0-创建，1-上线中，2-已下线
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

    public String getPromoName() {
        return promoName;
    }

    public void setPromoName(String promoName) {
        this.promoName = promoName;
    }

    public Date getStartDate() {
        return startDate;
    }

    public void setStartDate(Date startDate) {
        this.startDate = startDate;
    }

    public Date getEndDate() {
        return endDate;
    }

    public void setEndDate(Date endDate) {
        this.endDate = endDate;
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
    public HPromoProductModel convert(Function<HPromo, HPromoProductModel> f) {
        return f.apply(this);
    }

}

