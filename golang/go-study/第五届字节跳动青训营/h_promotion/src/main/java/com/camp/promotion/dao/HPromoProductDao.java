package com.camp.promotion.dao;

import com.camp.promotion.entity.HPromoProduct;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (HPromoProduct)表数据库访问层
 *
 * @author makejava
 * @since 2022-12-03 11:01:30
 */
@Mapper
public interface HPromoProductDao {


    HPromoProduct queryByPromoIdAndSkuId(@Param("promoId") Long promoId,
                                    @Param("skuId") Long skuId,
                                    @Param("spuId") Long spuId);

    int decreaseStock(@Param("promoId") Long promoId,
                      @Param("skuId") Long skuId,
                      @Param("quantity") Integer quantity);


    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HPromoProduct queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hPromoProduct 查询条件
     * @param pageable      分页对象
     * @return 对象列表
     */
    List<HPromoProduct> queryAllByLimit(HPromoProduct hPromoProduct, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hPromoProduct 查询条件
     * @return 总行数
     */
    long count(HPromoProduct hPromoProduct);

    /**
     * 新增数据
     *
     * @param hPromoProduct 实例对象
     * @return 影响行数
     */
    int insert(HPromoProduct hPromoProduct);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HPromoProduct> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HPromoProduct> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HPromoProduct> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HPromoProduct> entities);

    /**
     * 修改数据
     *
     * @param hPromoProduct 实例对象
     * @return 影响行数
     */
    int update(HPromoProduct hPromoProduct);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

