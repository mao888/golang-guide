package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HSpec)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
public class HSpec implements Serializable {
    private static final long serialVersionUID = -66475501811467207L;
    /**
     * id
     */
    private Long id;
    /**
     * 分类id,对应category表的主键
     */
    private Long categoryId;
    /**
     * 规格名称
     */
    private String specKey;
    /**
     * 类目状态0-废弃，1-正常
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

    public String getSpecKey() {
        return specKey;
    }

    public void setSpecKey(String specKey) {
        this.specKey = specKey;
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

