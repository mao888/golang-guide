package com.camp.promotion.dao;

import com.camp.promotion.entity.HPromo;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.data.domain.Pageable;

import java.util.List;

/**
 * (HPromo)表数据库访问层
 *
 * @author makejava
 * @since 2022-11-29 20:22:24
 */
@Mapper
public interface HPromoDao {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HPromo queryById(Long id);

    /**
     * 查询指定行数据
     *
     * @param hPromo   查询条件
     * @param pageable 分页对象
     * @return 对象列表
     */
    List<HPromo> queryAllByLimit(HPromo hPromo, @Param("pageable") Pageable pageable);

    /**
     * 统计总行数
     *
     * @param hPromo 查询条件
     * @return 总行数
     */
    long count(HPromo hPromo);

    /**
     * 新增数据
     *
     * @param hPromo 实例对象
     * @return 影响行数
     */
    int insert(HPromo hPromo);

    /**
     * 批量新增数据（MyBatis原生foreach方法）
     *
     * @param entities List<HPromo> 实例对象列表
     * @return 影响行数
     */
    int insertBatch(@Param("entities") List<HPromo> entities);

    /**
     * 批量新增或按主键更新数据（MyBatis原生foreach方法）
     *
     * @param entities List<HPromo> 实例对象列表
     * @return 影响行数
     * @throws org.springframework.jdbc.BadSqlGrammarException 入参是空List的时候会抛SQL语句错误的异常，请自行校验入参
     */
    int insertOrUpdateBatch(@Param("entities") List<HPromo> entities);

    /**
     * 修改数据
     *
     * @param hPromo 实例对象
     * @return 影响行数
     */
    int update(HPromo hPromo);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 影响行数
     */
    int deleteById(Long id);

}

