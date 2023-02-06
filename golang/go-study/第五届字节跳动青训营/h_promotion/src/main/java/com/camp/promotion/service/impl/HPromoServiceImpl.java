package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HPromo;
import com.camp.promotion.dao.HPromoDao;
import com.camp.promotion.service.HPromoService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HPromo)表服务实现类
 *
 * @author makejava
 * @since 2022-12-03 10:59:55
 */
@Service("hPromoService")
public class HPromoServiceImpl implements HPromoService {
    @Resource
    private HPromoDao hPromoDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HPromo queryById(Long id) {
        return this.hPromoDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hPromo      筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HPromo> queryByPage(HPromo hPromo, PageRequest pageRequest) {
        long total = this.hPromoDao.count(hPromo);
        return new PageImpl<>(this.hPromoDao.queryAllByLimit(hPromo, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hPromo 实例对象
     * @return 实例对象
     */
    @Override
    public int insert(HPromo hPromo) {
        return this.hPromoDao.insert(hPromo);
    }

    /**
     * 修改数据
     *
     * @param hPromo 实例对象
     * @return 实例对象
     */
    @Override
    public HPromo update(HPromo hPromo) {
        this.hPromoDao.update(hPromo);
        return this.queryById(hPromo.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hPromoDao.deleteById(id) > 0;
    }
}
