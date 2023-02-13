package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HUserAddress)实体类
 *
 * @author makejava
 * @since 2022-12-04 18:44:32
 */
public class HUserAddress implements Serializable {
    private static final long serialVersionUID = -53419156607548980L;
    /**
     * 用户表id
     */
    private Long id;
    /**
     * user id
     */
    private Long userId;
    /**
     * 地址信息
     */
    private String address;
    /**
     * 状态.0-废弃 1-有效
     */
    private Integer status;
    /**
     * 是否默认.0-no 1-yes
     */
    private Integer beDefault;
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

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public Integer getStatus() {
        return status;
    }

    public void setStatus(Integer status) {
        this.status = status;
    }

    public Integer getBeDefault() {
        return beDefault;
    }

    public void setBeDefault(Integer beDefault) {
        this.beDefault = beDefault;
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

