package com.camp.promotion.limit;

import java.lang.annotation.*;

@Documented
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface RateLimit {

    String value() default "";

    RateEnum rate() default RateEnum.RATE_100_PER_SECONDS;
}
