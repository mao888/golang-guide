package com.camp.promotion.limit;

import com.camp.promotion.common.BizException;
import com.camp.promotion.common.ResponseEnum;
import com.google.common.util.concurrent.RateLimiter;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.stereotype.Component;

import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;

@Slf4j
@Aspect
@Component
public class RateLimitAspect {

    private Map<String, RateLimiter> limiterMap = new HashMap<>();

    @Around("@annotation(com.camp.promotion.limit.RateLimit)")
    public Object around(ProceedingJoinPoint proceedingJoinPoint) throws Throwable {
        MethodSignature signature = (MethodSignature) proceedingJoinPoint.getSignature();
        Method method = signature.getMethod();

        RateLimit annotation = method.getAnnotation(RateLimit.class);
        if (annotation == null) {
            return proceedingJoinPoint.proceed();
        }

        String key = annotation.value();
        if (StringUtils.isBlank(key)) {
            String className = method.getDeclaringClass().getSimpleName();
            String methodName = method.getName();
            key = className + "_" + methodName;
        }
        log.info("rate limit key = {}", key);
        RateLimiter limiter = limiterMap.get(key);
        if (limiter == null) {
            limiter = RateLimiter.create(annotation.rate().getCount());
            limiterMap.put(key, limiter);
        }
        if (!limiter.tryAcquire()) {
            throw new BizException(ResponseEnum.OVER_RATE_LIMIT);
        }
        return proceedingJoinPoint.proceed();
    }
}
