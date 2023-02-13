package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HSpec;
import com.camp.promotion.dao.HSpecDao;
import com.camp.promotion.service.HSpecService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HSpec)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:25
 */
@Service("hSpecService")
public class HSpecServiceImpl implements HSpecService {
    @Resource
    private HSpecDao hSpecDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HSpec queryById(Long id) {
        return this.hSpecDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hSpec       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HSpec> queryByPage(HSpec hSpec, PageRequest pageRequest) {
        long total = this.hSpecDao.count(hSpec);
        return new PageImpl<>(this.hSpecDao.queryAllByLimit(hSpec, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hSpec 实例对象
     * @return 实例对象
     */
    @Override
    public HSpec insert(HSpec hSpec) {
        this.hSpecDao.insert(hSpec);
        return hSpec;
    }

    /**
     * 修改数据
     *
     * @param hSpec 实例对象
     * @return 实例对象
     */
    @Override
    public HSpec update(HSpec hSpec) {
        this.hSpecDao.update(hSpec);
        return this.queryById(hSpec.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hSpecDao.deleteById(id) > 0;
    }
}
