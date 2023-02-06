package com.camp.promotion.service;

import com.camp.promotion.entity.HSpec;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HSpec)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
public interface HSpecService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HSpec queryById(Long id);

    /**
     * 分页查询
     *
     * @param hSpec       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HSpec> queryByPage(HSpec hSpec, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hSpec 实例对象
     * @return 实例对象
     */
    HSpec insert(HSpec hSpec);

    /**
     * 修改数据
     *
     * @param hSpec 实例对象
     * @return 实例对象
     */
    HSpec update(HSpec hSpec);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
