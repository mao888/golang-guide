package com.camp.promotion.service;

import com.camp.promotion.entity.HOrder;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (HOrder)表服务接口
 *
 * @author makejava
 * @since 2022-12-04 21:32:33
 */
public interface HOrderService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HOrder queryById(Long id);

    /**
     * 分页查询
     *
     * @param hOrder 筛选条件
     * @param pageRequest      分页对象
     * @return 查询结果
     */
    Page<HOrder> queryByPage(HOrder hOrder, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hOrder 实例对象
     * @return 实例对象
     */
    int insert(HOrder hOrder);

    /**
     * 新增数据
     *
     * @param hOrders 实例对象
     * @return 实例对象
     */
    int batchInsert(List<HOrder> hOrders);

    /**
     * 修改数据
     *
     * @param hOrder 实例对象
     * @return 实例对象
     */
    HOrder update(HOrder hOrder);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
