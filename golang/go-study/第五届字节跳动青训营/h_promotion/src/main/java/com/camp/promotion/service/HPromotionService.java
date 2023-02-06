package com.camp.promotion.service;

import com.camp.promotion.model.CreateActivityModel;
import com.camp.promotion.model.CreatePromoProductModel;
import com.camp.promotion.model.HPromoProductModel;

import java.util.List;

public interface HPromotionService {

    List<CreatePromoProductModel> CreatePromoActivity(CreateActivityModel createActivityModel);

    HPromoProductModel getPromotionProductDetail(Long promoId, Long skuId, Long spuId);
}
