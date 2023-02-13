package com.camp.promotion.dao;

import com.camp.promotion.entity.HUserAddress;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (HUserAddress)表数据库访问层
 *
 * @author makejava
 * @since 2022-12-04 18:44:32
 */
@Mapper
public interface HUserAddressDao {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HUserAddress queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hUserAddress 查询条件
     * @param pageable     分页对象
     * @return 对象列表
     */
    List<HUserAddress> queryAllByLimit(HUserAddress hUserAddress, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hUserAddress 查询条件
     * @return 总行数
     */
    long count(HUserAddress hUserAddress);

    /**
     * 新增数据
     *
     * @param hUserAddress 实例对象
     * @return 影响行数
     */
    int insert(HUserAddress hUserAddress);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HUserAddress> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HUserAddress> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HUserAddress> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HUserAddress> entities);

    /**
     * 修改数据
     *
     * @param hUserAddress 实例对象
     * @return 影响行数
     */
    int update(HUserAddress hUserAddress);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

