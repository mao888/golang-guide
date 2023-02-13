package com.camp.promotion.service;

import com.camp.promotion.entity.HCategory;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HCategory)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:20
 */
public interface HCategoryService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HCategory queryById(Long id);

    /**
     * 分页查询
     *
     * @param hCategory   筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HCategory> queryByPage(HCategory hCategory, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hCategory 实例对象
     * @return 实例对象
     */
    HCategory insert(HCategory hCategory);

    /**
     * 修改数据
     *
     * @param hCategory 实例对象
     * @return 实例对象
     */
    HCategory update(HCategory hCategory);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
