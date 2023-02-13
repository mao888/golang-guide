package com.camp.promotion.service;

import com.camp.promotion.entity.HSpecDetail;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (HSpecDetail)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:26
 */
public interface HSpecDetailService {

    List<HSpecDetail> queryAll();
    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HSpecDetail queryById(Long id);

    /**
     * 分页查询
     *
     * @param hSpecDetail 筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HSpecDetail> queryByPage(HSpecDetail hSpecDetail, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hSpecDetail 实例对象
     * @return 实例对象
     */
    HSpecDetail insert(HSpecDetail hSpecDetail);

    /**
     * 修改数据
     *
     * @param hSpecDetail 实例对象
     * @return 实例对象
     */
    HSpecDetail update(HSpecDetail hSpecDetail);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
