package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HProduct;
import com.camp.promotion.dao.HProductDao;
import com.camp.promotion.service.HProductService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (HProduct)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:24
 */
@Service("hProductService")
public class HProductServiceImpl implements HProductService {
    @Resource
    private HProductDao hProductDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HProduct queryById(Long id) {
        return this.hProductDao.queryById(id);
    }

    @Override
    public List<HProduct> queryByIdList(List<Long> ids) {
        return this.hProductDao.queryByIdList(ids);
    }

    /**
     * 分页查询
     *
     * @param hProduct    筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HProduct> queryByPage(HProduct hProduct, PageRequest pageRequest) {
        long total = this.hProductDao.count(hProduct);
        return new PageImpl<>(this.hProductDao.queryAllByLimit(hProduct, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hProduct 实例对象
     * @return 实例对象
     */
    @Override
    public HProduct insert(HProduct hProduct) {
        this.hProductDao.insert(hProduct);
        return hProduct;
    }

    /**
     * 修改数据
     *
     * @param hProduct 实例对象
     * @return 实例对象
     */
    @Override
    public HProduct update(HProduct hProduct) {
        this.hProductDao.update(hProduct);
        return this.queryById(hProduct.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hProductDao.deleteById(id) > 0;
    }
}
