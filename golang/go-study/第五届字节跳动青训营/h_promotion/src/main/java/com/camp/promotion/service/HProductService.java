package com.camp.promotion.service;

import com.camp.promotion.entity.HProduct;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (HProduct)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:24
 */
public interface HProductService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HProduct queryById(Long id);

    List<HProduct> queryByIdList(List<Long> ids);

    /**
     * 分页查询
     *
     * @param hProduct    筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HProduct> queryByPage(HProduct hProduct, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hProduct 实例对象
     * @return 实例对象
     */
    HProduct insert(HProduct hProduct);

    /**
     * 修改数据
     *
     * @param hProduct 实例对象
     * @return 实例对象
     */
    HProduct update(HProduct hProduct);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
