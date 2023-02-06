package com.camp.promotion.convertet;

import com.camp.promotion.controller.request.CreateActivityRequest;
import com.camp.promotion.controller.request.CreateOrderRequest;
import com.camp.promotion.controller.request.CreatePromoProductRequest;
import com.camp.promotion.entity.*;
import com.camp.promotion.enums.PromoStatusEnum;
import com.camp.promotion.model.*;

import java.util.Date;
import java.util.List;
import java.util.stream.Collectors;

public class PromoConverter {

    public static HPromoSpuModel convertSpuModel(HProduct spu) {
        HPromoSpuModel spuModel = HPromoSpuModel.builder()
                .spuId(spu.getId())
                .categoryId(spu.getCategoryId())
                .shopId(spu.getShopId())
                .title(spu.getTitle())
                .subtitle(spu.getSubtitle())
                .mainImage(spu.getMainImage())
                .subImages(spu.getSubImages())
                .detail(spu.getDetail())
                .price(spu.getPrice())
                .categoryData(spu.getCategoryData())
                .specData(spu.getSpecData())
                .status(spu.getStatus())
                .build();

        return spuModel;
    }

    public static HPromoSkuModel convertSkuModel(HPromoProduct promoProduct, HSku sku) {
        HPromoSkuModel skuModel = HPromoSkuModel.builder()
                .spuId(promoProduct.getSpuId())
                .skuId(promoProduct.getSkuId())
                .promoStock(promoProduct.getPromoStock())
                .promoPrice(promoProduct.getPromoPrice())
                .price(sku.getPrice())
                .shopId(sku.getShopId())
                .specDetailIds(sku.getSpecDetailIds())
                .build();

        return skuModel;
    }

    public static HPromoProductModel convertProductModel(HPromo promo) {
        HPromoProductModel promoProductModel = HPromoProductModel.builder()
                .promoId(promo.getId())
                .promoName(promo.getPromoName())
                .startDate(promo.getStartDate())
                .endDate(promo.getEndDate())
                .status(promo.getStatus())
                .build();

        return promoProductModel;
    }

    public static CreateOrderModel convertCreateOrderModel(CreateOrderRequest createOrderRequest) {
        CreateOrderModel model = CreateOrderModel.builder()
                .userId(createOrderRequest.getUserId())
                .promoId(createOrderRequest.getPromoId())
                .skuId(createOrderRequest.getSkuId())
                .spuId(createOrderRequest.getSpuId())
                .quantity(createOrderRequest.getQuantity())
                .addressId(createOrderRequest.getAddressId())
                .shippingAddress(createOrderRequest.getShippingAddress())
                .build();

        return model;
    }

    public static HOrder initOrder(CreateOrderModel model) {
        HOrder order = new HOrder();
        order.setUserId(model.getUserId());
        order.setSkuId(model.getSkuId());
        order.setSpuId(model.getSpuId());
        order.setActivityId(model.getPromoId());
        order.setAddressId(model.getAddressId());
        order.setShippingAddress(model.getShippingAddress());
        return order;
    }

    public static OrderModel convertOrderModel(HOrder order) {
        OrderModel model = OrderModel.builder()
                .orderId(order.getId())
                .userId(order.getUserId())
                .spuId(order.getSpuId())
                .skuId(order.getSkuId())
                .shopId(order.getShopId())
                .payment(order.getPayment())
                .paymentType(order.getPaymentType())
                .postage(order.getPostage())
                .quantity(order.getQuantity())
                .status(order.getStatus())
                .addressId(order.getAddressId())
                .shippingAddress(order.getShippingAddress())
                .activityId(order.getActivityId())
                .build();

        return model;
    }

    public static CreateActivityModel convertCreateActivityModel(CreateActivityRequest createActivityRequest) {
        List<CreatePromoProductModel> promoProductModels = createActivityRequest.getPromoProducts().stream()
                .map(PromoConverter::convertCreateActivityModel).collect(Collectors.toList());
        CreateActivityModel model = CreateActivityModel.builder()
                .promoName(createActivityRequest.getPromoName())
                .startTime(createActivityRequest.getStartTime())
                .endTime(createActivityRequest.getEndTime())
                .promoProducts(promoProductModels)
                .build();
        return model;
    }

    public static CreatePromoProductModel convertCreateActivityModel(CreatePromoProductRequest createPromoProductRequest) {
        CreatePromoProductModel model = CreatePromoProductModel.builder()
                .spuId(createPromoProductRequest.getSpuId())
                .skuId(createPromoProductRequest.getSkuId())
                .promoStock(createPromoProductRequest.getPromoStock())
                .promoPrice(createPromoProductRequest.getPromoPrice())
                .build();

        return model;
    }

    public static HPromo convertHPromo(CreateActivityModel model) {
        Date now = new Date();
        HPromo promo = new HPromo();
        promo.setPromoName(model.getPromoName());
        promo.setStartDate(new Date(model.getStartTime()));
        promo.setEndDate(new Date(model.getEndTime()));
        promo.setStatus(PromoStatusEnum.ONLINE.getCode());
        promo.setCreateTime(now);
        promo.setUpdateTime(now);
        return promo;
    }

    public static HPromoProduct convertHPromoProduct(CreatePromoProductModel model) {
        Date now = new Date();
        HPromoProduct promoProduct = new HPromoProduct();
        promoProduct.setPromoId(model.getPromoId());
        promoProduct.setPromoName(model.getPromoName());
        promoProduct.setSpuId(model.getSpuId());
        promoProduct.setSkuId(model.getSkuId());
        promoProduct.setPromoPrice(model.getPromoPrice());
        promoProduct.setPromoStock(model.getPromoStock());
        promoProduct.setCreateTime(now);
        promoProduct.setUpdateTime(now);
        return promoProduct;
    }
}
