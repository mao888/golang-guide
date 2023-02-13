package com.camp.promotion.controller;

import com.alibaba.fastjson2.JSON;
import com.camp.promotion.common.BizException;
import com.camp.promotion.common.Constant;
import com.camp.promotion.common.ResponseData;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.controller.request.CreateActivityRequest;
import com.camp.promotion.controller.request.CreateOrderRequest;
import com.camp.promotion.convertet.PromoConverter;
import com.camp.promotion.limit.RateEnum;
import com.camp.promotion.limit.RateLimit;
import com.camp.promotion.model.*;
import com.camp.promotion.service.*;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.CollectionUtils;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;

@Slf4j
@RestController
@RequestMapping("/api/v1/promo/product")
public class PromotionController {

    @Resource
    private CacheService cacheService;
    @Resource
    private HPromotionService hPromotionService;
    @Resource
    private HPromoOrderService hPromoOrderService;
    @Resource
    private RiskManagementService riskManagementService;

    @PostMapping("/activity")
    @RateLimit(rate = RateEnum.RATE_100_PER_SECONDS)
    public ResponseData<?> createPromoActivity(@RequestBody CreateActivityRequest createActivityRequest) {

        Long start = createActivityRequest.getStartTime();
        Long end = createActivityRequest.getEndTime();
        if (start > end || System.currentTimeMillis() > start) {
            throw new BizException(ResponseEnum.ILLEGAL_ARGUMENT);
        }
        if (CollectionUtils.isEmpty(createActivityRequest.getPromoProducts())) {
            throw new BizException(ResponseEnum.ILLEGAL_ARGUMENT);
        }

        CreateActivityModel model = createActivityRequest.convert(PromoConverter::convertCreateActivityModel);
        List<CreatePromoProductModel> promoProductModels = hPromotionService.CreatePromoActivity(model);

        return ResponseData.Success(promoProductModels);
    }

    @GetMapping("/detail")
    @RateLimit(rate = RateEnum.RATE_5000_PER_SECONDS)
    public ResponseData<?> getProductDetail(@RequestParam(name = "promoId") Long promoId,
                                            @RequestParam(name = "skuId") Long skuId,
                                            @RequestParam(name = "spuId") Long spuId) {

        if (promoId <= 0 || skuId <= 0 || spuId <= 0) {
            throw new BizException(ResponseEnum.ILLEGAL_ARGUMENT);
        }

        String key = Constant.generatePromoProductKey(promoId, skuId, spuId);
        HPromoProductModel promoProductModel = cacheService.getVal(key);
        if (promoProductModel == null) {
            promoProductModel = hPromotionService.getPromotionProductDetail(promoId, skuId, spuId);
        }
        if (promoProductModel == null || !promoProductModel.checkPromo()) {
            log.error("promo activity not exist, params = {}", JSON.toJSONString(promoProductModel));
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }

        log.info("get promo product detail success, res = {}", promoProductModel.getPromoId());
        return ResponseData.Success(promoProductModel);
    }

    @PostMapping("/order")
    @RateLimit(rate = RateEnum.RATE_1000_PER_SECONDS)
    public ResponseData<?> createOrder(@RequestBody CreateOrderRequest createOrderRequest) {

        // 登录验证 mock
        if (createOrderRequest.getUserId() == null) {
            throw new BizException(ResponseEnum.USER_NOT_LOGIN);
        }

        // risk management
        if (!riskManagementService.riskManagement(createOrderRequest.getUserId())) {
            throw new BizException(ResponseEnum.USER_UNDER_RISK);
        }

        CreateOrderModel model = createOrderRequest.convert(PromoConverter::convertCreateOrderModel);
        OrderModel order = hPromoOrderService.createOrder(model);
        return ResponseData.Success(order);
    }
}
