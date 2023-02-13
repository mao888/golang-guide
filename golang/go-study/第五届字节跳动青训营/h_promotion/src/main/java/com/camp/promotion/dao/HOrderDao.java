package com.camp.promotion.dao;

import com.camp.promotion.entity.HOrder;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;
import java.util.List;

/**
 * (HOrder)表数据库访问层
 *
 * @author makejava
 * @since 2022-12-04 21:32:30
 */
@Mapper
public interface HOrderDao {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HOrder queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hOrder 查询条件
     * @param pageable         分页对象
     * @return 对象列表
     */
    List<HOrder> queryAllByLimit(HOrder hOrder, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hOrder 查询条件
     * @return 总行数
     */
    long count(HOrder hOrder);

    /**
     * 新增数据
     *
     * @param hOrder 实例对象
     * @return 影响行数
     */
    int insert(HOrder hOrder);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HOrder> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HOrder> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HOrder> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HOrder> entities);

    /**
     * 修改数据
     *
     * @param hOrder 实例对象
     * @return 影响行数
     */
    int update(HOrder hOrder);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

