package com.camp.promotion.entity;

import java.util.Date;
import java.io.Serializable;

/**
 * (HCategory)实体类
 *
 * @author makejava
 * @since 2022-11-29 20:22:17
 */
public class HCategory implements Serializable {
    private static final long serialVersionUID = -70258730050838158L;
    /**
     * id
     */
    private Long id;
    /**
     * 父类目id当id=0时说明是根节点,一级类目
     */
    private Long parentId;
    /**
     * 类目名称
     */
    private String name;
    /**
     * 类目级别1-一级,2-二级，3-三级
     */
    private Integer level;
    /**
     * 是否父类目0-否，1-是
     */
    private Integer beParent;
    /**
     * 类目状态0-废弃，1-正常
     */
    private Integer status;
    /**
     * 排序编号,同类展示顺序,数值相等则自然排序
     */
    private Integer sortOrder;
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

    public Long getParentId() {
        return parentId;
    }

    public void setParentId(Long parentId) {
        this.parentId = parentId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getLevel() {
        return level;
    }

    public void setLevel(Integer level) {
        this.level = level;
    }

    public Integer getBeParent() {
        return beParent;
    }

    public void setBeParent(Integer beParent) {
        this.beParent = beParent;
    }

    public Integer getStatus() {
        return status;
    }

    public void setStatus(Integer status) {
        this.status = status;
    }

    public Integer getSortOrder() {
        return sortOrder;
    }

    public void setSortOrder(Integer sortOrder) {
        this.sortOrder = sortOrder;
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

