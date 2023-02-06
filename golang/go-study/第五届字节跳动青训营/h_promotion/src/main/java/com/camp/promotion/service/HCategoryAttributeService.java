package com.camp.promotion.service;

import com.camp.promotion.entity.HCategoryAttribute;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HCategoryAttribute)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:23
 */
public interface HCategoryAttributeService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HCategoryAttribute queryById(Long id);

    /**
     * 分页查询
     *
     * @param hCategoryAttribute 筛选条件
     * @param pageRequest        分页对象
     * @return 查询结果
     */
    Page<HCategoryAttribute> queryByPage(HCategoryAttribute hCategoryAttribute, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 实例对象
     */
    HCategoryAttribute insert(HCategoryAttribute hCategoryAttribute);

    /**
     * 修改数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 实例对象
     */
    HCategoryAttribute update(HCategoryAttribute hCategoryAttribute);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
