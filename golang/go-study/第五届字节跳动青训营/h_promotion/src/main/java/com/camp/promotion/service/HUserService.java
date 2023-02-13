package com.camp.promotion.service;

import com.camp.promotion.entity.HUser;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HUser)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:26
 */
public interface HUserService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HUser queryById(Long id);

    /**
     * 分页查询
     *
     * @param hUser       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HUser> queryByPage(HUser hUser, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hUser 实例对象
     * @return 实例对象
     */
    HUser insert(HUser hUser);

    /**
     * 修改数据
     *
     * @param hUser 实例对象
     * @return 实例对象
     */
    HUser update(HUser hUser);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
