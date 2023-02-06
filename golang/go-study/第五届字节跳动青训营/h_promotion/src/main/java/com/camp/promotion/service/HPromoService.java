package com.camp.promotion.service;

import com.camp.promotion.entity.HPromo;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

/**
 * (HPromo)表服务接口
 *
 * @author makejava
 * @since 2022-12-03 10:59:54
 */
public interface HPromoService {

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HPromo queryById(Long id);

    /**
     * 分页查询
     *
     * @param hPromo      筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HPromo> queryByPage(HPromo hPromo, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hPromo 实例对象
     * @return 实例对象
     */
    int insert(HPromo hPromo);

    /**
     * 修改数据
     *
     * @param hPromo 实例对象
     * @return 实例对象
     */
    HPromo update(HPromo hPromo);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
