package com.camp.promotion.service;

import com.camp.promotion.model.HPromoProductModel;
import com.github.benmanes.caffeine.cache.CacheLoader;
import com.github.benmanes.caffeine.cache.Caffeine;
import com.github.benmanes.caffeine.cache.LoadingCache;
import lombok.extern.slf4j.Slf4j;
import org.checkerframework.checker.nullness.qual.Nullable;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.concurrent.TimeUnit;

@Slf4j
@Service
public class LocalCacheService {

    @Resource
    private RedisService redisService;

//    private LoadingCache<String, HPromoProductModel> GuavaCache;

    private LoadingCache<String, HPromoProductModel> caffeineCache;

    @PostConstruct
    public void init() {
//        GuavaCache = CacheBuilder.newBuilder()
//                .initialCapacity(10)
//                .maximumSize(100)
//                .expireAfterWrite(3, TimeUnit.SECONDS)
//                .build(new CacheLoader<>() {
//                    @Override
//                    public HPromoProductModel load(String key) {
//                        return redisService.getVal(key);
//                    }
//                });

        caffeineCache = Caffeine.newBuilder()
                .initialCapacity(10)
                .maximumSize(100)
                .expireAfterWrite(5,TimeUnit.SECONDS)
                .build(new CacheLoader<>() {
                    @Override
                    public @Nullable HPromoProductModel load(String key) {
                        return (HPromoProductModel) redisService.getVal(key);
                    }
                });
    }

    public HPromoProductModel getVal(String key) {

        HPromoProductModel model = null;
        try {
            model = this.caffeineCache.get(key);
        } catch (Exception e) {
            return null;
        }
        log.info("get HPromoProductModel from local cache, key = {}, res is null {}", key, model == null);
        return model;
    }
}
