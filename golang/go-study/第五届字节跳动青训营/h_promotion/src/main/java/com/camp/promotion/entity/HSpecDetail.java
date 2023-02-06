package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HSpecDetail)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
public class HSpecDetail implements Serializable {
    private static final long serialVersionUID = 148396565009656432L;
    /**
     * id
     */
    private Long id;
    /**
     * 规格
     */
    private Long specId;
    /**
     * 规格属性名称
     */
    private String specKey;
    /**
     * 规格属性值
     */
    private String specValue;
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

    public Long getSpecId() {
        return specId;
    }

    public void setSpecId(Long specId) {
        this.specId = specId;
    }

    public String getSpecKey() {
        return specKey;
    }

    public void setSpecKey(String specKey) {
        this.specKey = specKey;
    }

    public String getSpecValue() {
        return specValue;
    }

    public void setSpecValue(String specValue) {
        this.specValue = specValue;
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

