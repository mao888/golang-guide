package com.camp.promotion;

import com.camp.promotion.util.ApplicationContextUtil;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.ConfigurableApplicationContext;

/**
 * author Ryan Cui
 */
@SpringBootApplication
public class HPromotionApplication {

    public static void main(String[] args) {
        ConfigurableApplicationContext applicationContext = SpringApplication.run(HPromotionApplication.class, args);
        ApplicationContextUtil.setApplicationContext(applicationContext);
    }

}
