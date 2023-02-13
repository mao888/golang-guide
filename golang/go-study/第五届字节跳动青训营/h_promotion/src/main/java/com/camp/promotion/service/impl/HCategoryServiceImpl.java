package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HCategory;
import com.camp.promotion.dao.HCategoryDao;
import com.camp.promotion.service.HCategoryService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HCategory)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:21
 */
@Service("hCategoryService")
public class HCategoryServiceImpl implements HCategoryService {
    @Resource
    private HCategoryDao hCategoryDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HCategory queryById(Long id) {
        return this.hCategoryDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hCategory   筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HCategory> queryByPage(HCategory hCategory, PageRequest pageRequest) {
        long total = this.hCategoryDao.count(hCategory);
        return new PageImpl<>(this.hCategoryDao.queryAllByLimit(hCategory, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hCategory 实例对象
     * @return 实例对象
     */
    @Override
    public HCategory insert(HCategory hCategory) {
        this.hCategoryDao.insert(hCategory);
        return hCategory;
    }

    /**
     * 修改数据
     *
     * @param hCategory 实例对象
     * @return 实例对象
     */
    @Override
    public HCategory update(HCategory hCategory) {
        this.hCategoryDao.update(hCategory);
        return this.queryById(hCategory.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hCategoryDao.deleteById(id) > 0;
    }
}
