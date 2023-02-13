package com.camp.promotion.dao;

import com.camp.promotion.entity.HSpecDetail;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (HSpecDetail)表数据库访问层
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
@Mapper
public interface HSpecDetailDao {

    List<HSpecDetail> queryAll();
    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HSpecDetail queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hSpecDetail 查询条件
     * @param pageable    分页对象
     * @return 对象列表
     */
    List<HSpecDetail> queryAllByLimit(HSpecDetail hSpecDetail, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hSpecDetail 查询条件
     * @return 总行数
     */
    long count(HSpecDetail hSpecDetail);

    /**
     * 新增数据
     *
     * @param hSpecDetail 实例对象
     * @return 影响行数
     */
    int insert(HSpecDetail hSpecDetail);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HSpecDetail> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HSpecDetail> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HSpecDetail> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HSpecDetail> entities);

    /**
     * 修改数据
     *
     * @param hSpecDetail 实例对象
     * @return 影响行数
     */
    int update(HSpecDetail hSpecDetail);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

