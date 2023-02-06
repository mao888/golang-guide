package com.camp.promotion.service;

import com.alibaba.fastjson2.JSON;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.io.ClassPathResource;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.data.redis.core.script.DefaultRedisScript;
import org.springframework.scripting.support.ResourceScriptSource;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import javax.annotation.Resource;
import java.util.List;
import java.util.concurrent.TimeUnit;

@Slf4j
@Service
public class RedisService {

    private DefaultRedisScript<Long> redisScript;

    @PostConstruct
    public void init() {
        redisScript = new DefaultRedisScript<>();
        redisScript.setScriptSource(new ResourceScriptSource(new ClassPathResource("decrease-stock.lua")));
        redisScript.setResultType(Long.class);
    }

    @Resource
    private RedisTemplate<String, Object> redisTemplate;

    public Object getVal(String key) {
        return redisTemplate.opsForValue().get(key);
    }

    public void setVal(String key, Object val) {
        redisTemplate.opsForValue().set(key, JSON.toJSONString(val));
    }

    public void setVal(String key, Object val, long timeout, TimeUnit unit) {
        redisTemplate.opsForValue().set(key, val, timeout, unit);
    }

    public boolean setNX(String key, String value, long timeout, TimeUnit unit) {
       return redisTemplate.opsForValue().setIfAbsent(key, value, timeout, unit);
    }

    public Long eval(DefaultRedisScript<Long> redisScript, List<String> keys, List<String> values) {
        return redisTemplate.execute(redisScript, keys, values.toArray());
    }

    public Long decreaseStock(String... keys) {
        return redisTemplate.execute(this.redisScript, List.of(keys));
    }

    public Long increment(String key) {
        return redisTemplate.opsForValue().increment(key);
    }

    public Long decrement(String key) {
        return redisTemplate.opsForValue().decrement(key);
    }
}
