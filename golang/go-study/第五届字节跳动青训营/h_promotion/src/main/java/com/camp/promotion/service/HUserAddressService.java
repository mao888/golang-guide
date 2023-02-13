package com.camp.promotion.service;

import com.camp.promotion.entity.HUserAddress;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HUserAddress)表服务接口
 *
 * @author makejava
 * @since 2022-12-04 18:44:32
 */
public interface HUserAddressService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HUserAddress queryById(Long id);

    /**
     * 分页查询
     *
     * @param hUserAddress 筛选条件
     * @param pageRequest  分页对象
     * @return 查询结果
     */
    Page<HUserAddress> queryByPage(HUserAddress hUserAddress, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hUserAddress 实例对象
     * @return 实例对象
     */
    HUserAddress insert(HUserAddress hUserAddress);

    /**
     * 修改数据
     *
     * @param hUserAddress 实例对象
     * @return 实例对象
     */
    HUserAddress update(HUserAddress hUserAddress);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
