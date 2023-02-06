package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HPromoProduct;
import com.camp.promotion.dao.HPromoProductDao;
import com.camp.promotion.service.HPromoProductService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (HPromoProduct)表服务实现类
 *
 * @author makejava
 * @since 2022-12-03 11:01:32
 */
@Service("hPromoProductService")
public class HPromoProductServiceImpl implements HPromoProductService {
    @Resource
    private HPromoProductDao hPromoProductDao;

    @Override
    public HPromoProduct queryByIdAndSkuId(Long promoId, Long skuId, Long spuId) {
        return hPromoProductDao.queryByPromoIdAndSkuId(promoId, skuId, spuId);
    }

    @Override
    public int decreaseStock(Long promoId, Long skuId, Integer quantity) {
        return hPromoProductDao.decreaseStock(promoId, skuId, quantity);
    }

    @Override
    public int insertBatch(List<HPromoProduct> promoProducts) {
        return hPromoProductDao.insertBatch(promoProducts);
    }

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HPromoProduct queryById(Long id) {
        return this.hPromoProductDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hPromoProduct 筛选条件
     * @param pageRequest   分页对象
     * @return 查询结果
     */
    @Override
    public Page<HPromoProduct> queryByPage(HPromoProduct hPromoProduct, PageRequest pageRequest) {
        long total = this.hPromoProductDao.count(hPromoProduct);
        return new PageImpl<>(this.hPromoProductDao.queryAllByLimit(hPromoProduct, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hPromoProduct 实例对象
     * @return 实例对象
     */
    @Override
    public HPromoProduct insert(HPromoProduct hPromoProduct) {
        this.hPromoProductDao.insert(hPromoProduct);
        return hPromoProduct;
    }

    /**
     * 修改数据
     *
     * @param hPromoProduct 实例对象
     * @return 实例对象
     */
    @Override
    public HPromoProduct update(HPromoProduct hPromoProduct) {
        this.hPromoProductDao.update(hPromoProduct);
        return this.queryById(hPromoProduct.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hPromoProductDao.deleteById(id) > 0;
    }
}
