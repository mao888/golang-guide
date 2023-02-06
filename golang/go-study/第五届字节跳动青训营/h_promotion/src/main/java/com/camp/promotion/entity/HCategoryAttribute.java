package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HCategoryAttribute)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:22
 */
public class HCategoryAttribute implements Serializable {
    private static final long serialVersionUID = 749884106007017319L;
    /**
     * id
     */
    private Long id;
    /**
     * 类目id
     */
    private Long categoryId;
    /**
     * 类目属性名称
     */
    private String attributeKey;
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

    public String getAttributeKey() {
        return attributeKey;
    }

    public void setAttributeKey(String attributeKey) {
        this.attributeKey = attributeKey;
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

