package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HSpecDetail;
import com.camp.promotion.dao.HSpecDetailDao;
import com.camp.promotion.service.HSpecDetailService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;
import java.util.List;

/**
 * (HSpecDetail)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:26
 */
@Service("hSpecDetailService")
public class HSpecDetailServiceImpl implements HSpecDetailService {
    @Resource
    private HSpecDetailDao hSpecDetailDao;


    public List<HSpecDetail> queryAll() {
        return this.hSpecDetailDao.queryAll();
    }
    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HSpecDetail queryById(Long id) {
        return this.hSpecDetailDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hSpecDetail 筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HSpecDetail> queryByPage(HSpecDetail hSpecDetail, PageRequest pageRequest) {
        long total = this.hSpecDetailDao.count(hSpecDetail);
        return new PageImpl<>(this.hSpecDetailDao.queryAllByLimit(hSpecDetail, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hSpecDetail 实例对象
     * @return 实例对象
     */
    @Override
    public HSpecDetail insert(HSpecDetail hSpecDetail) {
        this.hSpecDetailDao.insert(hSpecDetail);
        return hSpecDetail;
    }

    /**
     * 修改数据
     *
     * @param hSpecDetail 实例对象
     * @return 实例对象
     */
    @Override
    public HSpecDetail update(HSpecDetail hSpecDetail) {
        this.hSpecDetailDao.update(hSpecDetail);
        return this.queryById(hSpecDetail.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hSpecDetailDao.deleteById(id) > 0;
    }
}
