package com.camp.promotion.controller;

import com.camp.promotion.common.ResponseData;
import com.camp.promotion.entity.HCategory;
import com.camp.promotion.entity.HSpecDetail;
import com.camp.promotion.lock.RedisDistributedLock;
import com.camp.promotion.mq.OrderProducer;
import com.camp.promotion.service.HCategoryService;
import com.camp.promotion.service.HSpecDetailService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.annotation.Resource;
import java.util.List;

@Slf4j
@Controller
@ResponseBody
@RequestMapping("/api/v1/promo/product")
public class CategoryController {

    @Resource
    private HCategoryService hCategoryService;
    @Resource
    private HSpecDetailService hSpecDetailService;
    @Resource
    private OrderProducer orderProducer;

    @GetMapping("/list")
    public ResponseData<?> listCategory(@RequestParam Long id) {
        HCategory category = hCategoryService.queryById(id);
        return ResponseData.Success(category);
    }

    @GetMapping("/all")
    public ResponseData<?> queryAll() {
        List<HSpecDetail> specDetails = hSpecDetailService.queryAll();
        return ResponseData.Success(specDetails);
    }

    @GetMapping("/test")
    public ResponseData<?> testLock() {
        try (RedisDistributedLock lock = new RedisDistributedLock("test_key", 10000)) {
            if (lock.tryLock()) {
                return ResponseData.Success("ok");
            }
        } catch (Exception e) {
            throw new RuntimeException("lock fail");
        }

        return ResponseData.Success("false");
    }

    @GetMapping("/mq")
    public ResponseData<?> sendMsg() {
        orderProducer.send();
        return ResponseData.Success("ok");
    }
}
