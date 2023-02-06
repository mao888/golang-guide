package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HSku;
import com.camp.promotion.dao.HSkuDao;
import com.camp.promotion.service.HSkuService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (HSku)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
@Service("hSkuService")
public class HSkuServiceImpl implements HSkuService {
    @Resource
    private HSkuDao hSkuDao;


    @Override
    public int decreaseStock(Long id, Integer stock) {
        return this.hSkuDao.decreaseStock(id, stock);
    }

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HSku queryById(Long id) {
        return this.hSkuDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hSku        筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HSku> queryByPage(HSku hSku, PageRequest pageRequest) {
        long total = this.hSkuDao.count(hSku);
        return new PageImpl<>(this.hSkuDao.queryAllByLimit(hSku, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hSku 实例对象
     * @return 实例对象
     */
    @Override
    public HSku insert(HSku hSku) {
        this.hSkuDao.insert(hSku);
        return hSku;
    }

    /**
     * 修改数据
     *
     * @param hSku 实例对象
     * @return 实例对象
     */
    @Override
    public HSku update(HSku hSku) {
        this.hSkuDao.update(hSku);
        return this.queryById(hSku.getId());
    }

    @Override
    public int batchUpdate(List<HSku> hSkus) {
        return this.hSkuDao.batchUpdate(hSkus);
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hSkuDao.deleteById(id) > 0;
    }
}
