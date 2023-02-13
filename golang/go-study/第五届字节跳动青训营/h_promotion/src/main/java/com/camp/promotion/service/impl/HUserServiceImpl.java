package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HUser;
import com.camp.promotion.dao.HUserDao;
import com.camp.promotion.service.HUserService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HUser)表服务实现类
 *
 * @author makejava
 * @since 2022-11-29 20:22:26
 */
@Service("hUserService")
public class HUserServiceImpl implements HUserService {
    @Resource
    private HUserDao hUserDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HUser queryById(Long id) {
        return this.hUserDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hUser       筛选条件
     * @param pageRequest 分页对象
     * @return 查询结果
     */
    @Override
    public Page<HUser> queryByPage(HUser hUser, PageRequest pageRequest) {
        long total = this.hUserDao.count(hUser);
        return new PageImpl<>(this.hUserDao.queryAllByLimit(hUser, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hUser 实例对象
     * @return 实例对象
     */
    @Override
    public HUser insert(HUser hUser) {
        this.hUserDao.insert(hUser);
        return hUser;
    }

    /**
     * 修改数据
     *
     * @param hUser 实例对象
     * @return 实例对象
     */
    @Override
    public HUser update(HUser hUser) {
        this.hUserDao.update(hUser);
        return this.queryById(hUser.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hUserDao.deleteById(id) > 0;
    }
}
