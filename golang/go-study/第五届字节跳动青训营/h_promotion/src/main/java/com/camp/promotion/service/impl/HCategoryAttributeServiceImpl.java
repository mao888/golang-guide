package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HCategoryAttribute;
import com.camp.promotion.dao.HCategoryAttributeDao;
import com.camp.promotion.service.HCategoryAttributeService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HCategoryAttribute)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:23
 */
@Service("hCategoryAttributeService")
public class HCategoryAttributeServiceImpl implements HCategoryAttributeService {
    @Resource
    private HCategoryAttributeDao hCategoryAttributeDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HCategoryAttribute queryById(Long id) {
        return this.hCategoryAttributeDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hCategoryAttribute 筛选条件
     * @param pageRequest        分页对象
     * @return 查询结果
     */
    @Override
    public Page<HCategoryAttribute> queryByPage(HCategoryAttribute hCategoryAttribute, PageRequest pageRequest) {
        long total = this.hCategoryAttributeDao.count(hCategoryAttribute);
        return new PageImpl<>(this.hCategoryAttributeDao.queryAllByLimit(hCategoryAttribute, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 实例对象
     */
    @Override
    public HCategoryAttribute insert(HCategoryAttribute hCategoryAttribute) {
        this.hCategoryAttributeDao.insert(hCategoryAttribute);
        return hCategoryAttribute;
    }

    /**
     * 修改数据
     *
     * @param hCategoryAttribute 实例对象
     * @return 实例对象
     */
    @Override
    public HCategoryAttribute update(HCategoryAttribute hCategoryAttribute) {
        this.hCategoryAttributeDao.update(hCategoryAttribute);
        return this.queryById(hCategoryAttribute.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hCategoryAttributeDao.deleteById(id) > 0;
    }
}
