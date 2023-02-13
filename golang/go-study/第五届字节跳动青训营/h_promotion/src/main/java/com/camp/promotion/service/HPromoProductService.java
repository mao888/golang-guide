package com.camp.promotion.service;

import com.camp.promotion.entity.HPromoProduct;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (HPromoProduct)表服务接口
 *
 * @author makejava
 * @since 2022-12-03 11:01:32
 */
public interface HPromoProductService {

    HPromoProduct queryByIdAndSkuId(Long promoId, Long skuId, Long spuId);

    int decreaseStock(Long promoId, Long skuId, Integer quantity);

    int insertBatch(List<HPromoProduct> promoProducts);

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HPromoProduct queryById(Long id);

    /**
     * 分页查询
     *
     * @param hPromoProduct 筛选条件
     * @param pageRequest   分页对象
     * @return 查询结果
     */
    Page<HPromoProduct> queryByPage(HPromoProduct hPromoProduct, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hPromoProduct 实例对象
     * @return 实例对象
     */
    HPromoProduct insert(HPromoProduct hPromoProduct);

    /**
     * 修改数据
     *
     * @param hPromoProduct 实例对象
     * @return 实例对象
     */
    HPromoProduct update(HPromoProduct hPromoProduct);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
