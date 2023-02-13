package com.camp.promotion.service;

import com.camp.promotion.common.BizException;
import com.camp.promotion.common.Constant;
import com.camp.promotion.common.ResponseData;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.lock.RedisDistributedLock;
import com.camp.promotion.model.HPromoProductModel;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.concurrent.TimeUnit;

@Slf4j
@Service
public class CacheService {

    @Resource
    private RedisService redisService;
    @Resource
    private LocalCacheService localCacheService;
    @Resource
    private HPromotionService hPromotionService;

    public HPromoProductModel getVal(String key) {
        HPromoProductModel model = localCacheService.getVal(key);

        String[] elements = key.split("_");
        if (elements.length != 5) {
            log.error("promo activity cache invalid key = {}", key);
            throw new BizException(ResponseEnum.INVALID_CACHE_KEY);
        }
        var promoId = Long.parseLong(elements[2]);
        var skuId = Long.parseLong(elements[3]);
        var spuId = Long.parseLong(elements[4]);
        if (model == null) {
            try (RedisDistributedLock lock = new RedisDistributedLock("lock_" + key, 10000)) {
                if (lock.tryLock()) {
                    model = hPromotionService.getPromotionProductDetail(promoId, skuId, spuId);
                    redisService.setVal(key, model, 3L, TimeUnit.MINUTES);
                } else {
                    // 获取分布式锁失败
                    boolean isCached = false;
                    while (!isCached) {
                        Thread.sleep(20);
                        model = localCacheService.getVal(key);
                        if (model != null) {
                            isCached = true;
                        }
                    }
                }
            } catch (Exception e) {
                log.error("lock fail, e=", e);
                throw new RuntimeException("lock fail");
            }
        } else {
            String stockKey = Constant.generatePromoStockKey(promoId, skuId, spuId);
            Integer stock = (Integer) redisService.getVal(stockKey);
            if (stock != null) {
                model.getSkuModel().setPromoStock(stock);
            }
        }

        return model;
    }

    public Long decreaseStock(String... keys) {
        return redisService.decreaseStock(keys);
    }

    public Long increaseStock(String key) {
        return redisService.increment(key);
    }

    public void setStock(String key, Integer val) {
        redisService.setVal(key, val);
    }
}
