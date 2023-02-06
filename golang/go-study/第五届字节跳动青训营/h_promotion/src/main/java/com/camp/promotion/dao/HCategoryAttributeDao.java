package com.camp.promotion.dao;

import com.camp.promotion.entity.HCategoryAttribute;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (HCategoryAttribute)表数据库访问层
 *
 * @author makejava
 * @since 2022-11-29 20:22:22
 */
@Mapper
public interface HCategoryAttributeDao {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HCategoryAttribute queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hCategoryAttribute 查询条件
     * @param pageable           分页对象
     * @return 对象列表
     */
    List<HCategoryAttribute> queryAllByLimit(HCategoryAttribute hCategoryAttribute, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hCategoryAttribute 查询条件
     * @return 总行数
     */
    long count(HCategoryAttribute hCategoryAttribute);

    /**
     * 新增数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 影响行数
     */
    int insert(HCategoryAttribute hCategoryAttribute);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HCategoryAttribute> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HCategoryAttribute> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HCategoryAttribute> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HCategoryAttribute> entities);

    /**
     * 修改数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 影响行数
     */
    int update(HCategoryAttribute hCategoryAttribute);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

