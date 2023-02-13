package com.camp.promotion.service;

import com.camp.promotion.model.CreateOrderModel;
import com.camp.promotion.model.OrderModel;

public interface HPromoOrderService {

    OrderModel createOrder(CreateOrderModel createOrderModel);
}
