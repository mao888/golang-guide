package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HOrder;
import com.camp.promotion.dao.HOrderDao;
import com.camp.promotion.service.HOrderService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (HOrder)表服务实现类
 *
 * @author makejava
 * @since 2022-12-04 21:32:33
 */
@Service("hOrderService")
public class HOrderServiceImpl implements HOrderService {
    @Resource
    private HOrderDao hOrderDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HOrder queryById(Long id) {
        return this.hOrderDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hOrder 筛选条件
     * @param pageRequest      分页对象
     * @return 查询结果
     */
    @Override
    public Page<HOrder> queryByPage(HOrder hOrder, PageRequest pageRequest) {
        long total = this.hOrderDao.count(hOrder);
        return new PageImpl<>(this.hOrderDao.queryAllByLimit(hOrder, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hOrder 实例对象
     * @return 实例对象
     */
    @Override
    public int insert(HOrder hOrder) {
        return this.hOrderDao.insert(hOrder);
    }

    /**
     * 新增数据
     *
     * @param hOrders 实例对象
     * @return 实例对象
     */
    @Override
    public int batchInsert(List<HOrder> hOrders) {
        return this.hOrderDao.insertBatch(hOrders);
    }

    /**
     * 修改数据
     *
     * @param hOrder 实例对象
     * @return 实例对象
     */
    @Override
    public HOrder update(HOrder hOrder) {
        this.hOrderDao.update(hOrder);
        return this.queryById(hOrder.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hOrderDao.deleteById(id) > 0;
    }
}
