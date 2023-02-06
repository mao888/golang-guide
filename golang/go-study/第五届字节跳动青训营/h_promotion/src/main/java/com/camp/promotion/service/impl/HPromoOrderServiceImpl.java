package com.camp.promotion.service.impl;

import com.alibaba.fastjson2.JSON;
import com.camp.promotion.common.BizException;
import com.camp.promotion.common.Constant;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.convertet.PromoConverter;
import com.camp.promotion.entity.HOrder;
import com.camp.promotion.enums.OrderStatusEnum;
import com.camp.promotion.enums.PaymentTypeEnum;
import com.camp.promotion.model.CreateOrderModel;
import com.camp.promotion.model.HPromoProductModel;
import com.camp.promotion.model.OrderModel;
import com.camp.promotion.mq.OrderProducer;
import com.camp.promotion.service.*;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.util.Date;

@Slf4j
@Service
public class HPromoOrderServiceImpl implements HPromoOrderService {

    @Resource
    private HOrderService hOrderService;
    @Resource
    private HPromotionService hPromotionService;
    @Resource
    private HPromoProductService hPromoProductService;
    @Resource
    private CacheService cacheService;
    @Resource
    private IdGenerateService idGenerateService;
    @Resource
    private OrderProducer orderProducer;

    @Override
    @Transactional(rollbackFor = Exception.class)
    public OrderModel createOrder(CreateOrderModel createOrderModel) {

        // 订单初始化
        HOrder order = createOrderModel.convert(PromoConverter::initOrder);

        // 基础参数
        Long promoId = createOrderModel.getPromoId();
        Long skuId = createOrderModel.getSkuId();
        Long spuId = createOrderModel.getSpuId();

        // 校验活动和商品状态
        String key = Constant.generatePromoProductKey(promoId, skuId, spuId);
        HPromoProductModel promoProductModel = cacheService.getVal(key);
        if (promoProductModel == null) {
            promoProductModel = hPromotionService.getPromotionProductDetail(promoId, skuId, spuId);
        }
        if (promoProductModel == null || !promoProductModel.checkPromo()) {
            log.error("promo activity not exist, params = {}", JSON.toJSONString(createOrderModel));
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }
        order.setShopId(promoProductModel.getSpuModel().getShopId());


        long decrease = cacheService.decreaseStock(Constant.generatePromoStockKey(promoId, skuId, spuId));
        if (decrease <= 0) {
            log.error("promo activity product out of stock, params = {}", JSON.toJSONString(createOrderModel));
            throw new BizException(ResponseEnum.OUT_OF_STOCK);
        }

        // 秒杀价格从后端获取
        Integer promoPrice = promoProductModel.getSkuModel().getPromoPrice();
        // 数量固定为1
        order.setQuantity(Constant.PROMO_QUANTITY);
        order.setPayment(promoPrice * Constant.PROMO_QUANTITY);
        order.setPaymentType(PaymentTypeEnum.ONLINE.getCode());
        order.setPostage(Constant.POSTAGE);
        order.setStatus(OrderStatusEnum.NON_PAYMENT.getCode());

        Date now = new Date();
        order.setCreateTime(now);
        order.setUpdateTime(now);

        // 订单号生成
        long orderId = idGenerateService.generateId();
        order.setId(orderId);

        boolean sendResult = orderProducer.send(order);
        if (!sendResult) {
            cacheService.increaseStock(Constant.generatePromoStockKey(promoId, skuId, spuId));
            throw new BizException(ResponseEnum.CREATE_ORDER_FAIL);
        }

        // 返回数据
        OrderModel model = order.convert(PromoConverter::convertOrderModel);
        return model;
    }
}
