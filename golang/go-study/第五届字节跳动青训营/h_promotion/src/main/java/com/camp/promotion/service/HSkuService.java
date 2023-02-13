package com.camp.promotion.service;

import com.camp.promotion.entity.HSku;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;

import java.util.List;

/**
 * (HSku)表服务接口
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
public interface HSkuService {

    int decreaseStock(Long id, Integer stock);

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    HSku queryById(Long id);

    /**
     * 分页查询
     *
     * @param hSku        筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    Page<HSku> queryByPage(HSku hSku, PageRequest pageRequest);

    /**
     * 新增数据
     *
     * @param hSku 实例对象
     * @return 实例对象
     */
    HSku insert(HSku hSku);

    /**
     * 修改数据
     *
     * @param hSku 实例对象
     * @return 实例对象
     */
    HSku update(HSku hSku);

    /**
     * 批量修改数据
     *
     * @param hSkus 实例对象
     * @return 实例对象
     */
    int batchUpdate(List<HSku> hSkus);

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    boolean deleteById(Long id);

}
