package com.camp.promotion.lock;

import com.camp.promotion.service.RedisService;
import com.camp.promotion.util.ApplicationContextUtil;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.io.ClassPathResource;
import org.springframework.data.redis.core.script.DefaultRedisScript;
import org.springframework.scripting.support.ResourceScriptSource;

import java.util.List;
import java.util.UUID;
import java.util.concurrent.TimeUnit;

@Slf4j
public class RedisDistributedLock extends AbstractDistributedLock {

    private final RedisService redisService;

    private final String key;

    private final String value;

    private static final DefaultRedisScript<Long> redisScript;

    private static final String EXPECTED = ", expectedValue=";

    static {
        redisScript = new DefaultRedisScript<>();
        redisScript.setScriptSource(new ResourceScriptSource(new ClassPathResource("unlock.lua")));
        redisScript.setResultType(Long.class);
    }


    /**
     * @param redisService   redis
     * @param key            锁键
     * @param survivalMillis 锁最大存活时间
     */
    public RedisDistributedLock(RedisService redisService, String key, long survivalMillis) {
        this(redisService, key);
        this.survivalMillis = survivalMillis;
    }

    /**
     * @param redisService redis
     * @param key   锁键
     */
    public RedisDistributedLock(RedisService redisService, String key) {
        this.redisService = redisService;
        this.key = key;
        this.value = UUID.randomUUID().toString().replace("-", "");
    }

    /**
     * @param key 锁键
     */
    public RedisDistributedLock(String key) {
        this(ApplicationContextUtil.getBean(RedisService.class), key);
    }

    /**
     * @param key            锁键
     * @param survivalMillis 锁最大存活时间
     */
    public RedisDistributedLock(String key, long survivalMillis) {
        this(ApplicationContextUtil.getBean(RedisService.class), key, survivalMillis);
    }

    /**
     * @param name bean名称
     * @param key           锁键
     */
    public RedisDistributedLock(String name, String key) {
        this((RedisService) ApplicationContextUtil.getBean(name), key);
    }

    /**
     * @param name bean名称
     * @param key            锁键
     * @param survivalMillis 锁最大存活时间
     */
    public RedisDistributedLock(String name, String key, long survivalMillis) {
        this((RedisService) ApplicationContextUtil.getBean(name), key, survivalMillis);
    }

    @Override
    public boolean tryLock() {
        boolean lockResult = redisService.setNX(key, value, survivalMillis, TimeUnit.MILLISECONDS);

        if (lockResult) {
            isLock.set(true);
        }
        return lockResult;
    }



    @Override
    public void unlock() {
        Long result;
        try {
            result = redisService.eval(redisScript, List.of(key), List.of(value));
        } catch (Exception e) {
            String msg = "unlock redis key failed, executing unlock scripts failed, key=" + key + EXPECTED + value;
            throw new IllegalStateException(msg, e);
        }
        if (null == result) {
            String msg = "unlock redis key failed, unlock scripts' result is null, key=" + key + EXPECTED + value;
            throw new IllegalStateException(msg);
        }
        if (1 != result) {
            String msg = "unlock redis key failed, unlock scripts' result not 1, result=" + result + ", key=" + key + EXPECTED + value;
            throw new IllegalStateException(msg);
        }
        isLock.set(false);
    }

}
