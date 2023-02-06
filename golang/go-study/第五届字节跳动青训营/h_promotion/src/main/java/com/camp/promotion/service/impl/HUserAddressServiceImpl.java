package com.camp.promotion.service.impl;

import com.camp.promotion.entity.HUserAddress;
import com.camp.promotion.dao.HUserAddressDao;
import com.camp.promotion.service.HUserAddressService;
import org.springframework.stereotype.Service;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;

import javax.annotation.Resource;

/**
 * (HUserAddress)表服务实现类
 *
 * @author makejava
 * @since 2022-12-04 18:44:32
 */
@Service("hUserAddressService")
public class HUserAddressServiceImpl implements HUserAddressService {
    @Resource
    private HUserAddressDao hUserAddressDao;

    /**
     * 通过ID查询单条数据
     *
     * @param id 主键
     * @return 实例对象
     */
    @Override
    public HUserAddress queryById(Long id) {
        return this.hUserAddressDao.queryById(id);
    }

    /**
     * 分页查询
     *
     * @param hUserAddress 筛选条件
     * @param pageRequest  分页对象
     * @return 查询结果
     */
    @Override
    public Page<HUserAddress> queryByPage(HUserAddress hUserAddress, PageRequest pageRequest) {
        long total = this.hUserAddressDao.count(hUserAddress);
        return new PageImpl<>(this.hUserAddressDao.queryAllByLimit(hUserAddress, pageRequest), pageRequest, total);
    }

    /**
     * 新增数据
     *
     * @param hUserAddress 实例对象
     * @return 实例对象
     */
    @Override
    public HUserAddress insert(HUserAddress hUserAddress) {
        this.hUserAddressDao.insert(hUserAddress);
        return hUserAddress;
    }

    /**
     * 修改数据
     *
     * @param hUserAddress 实例对象
     * @return 实例对象
     */
    @Override
    public HUserAddress update(HUserAddress hUserAddress) {
        this.hUserAddressDao.update(hUserAddress);
        return this.queryById(hUserAddress.getId());
    }

    /**
     * 通过主键删除数据
     *
     * @param id 主键
     * @return 是否成功
     */
    @Override
    public boolean deleteById(Long id) {
        return this.hUserAddressDao.deleteById(id) > 0;
    }
}
