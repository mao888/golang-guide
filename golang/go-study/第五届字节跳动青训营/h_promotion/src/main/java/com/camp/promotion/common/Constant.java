package com.camp.promotion.common;

public class Constant {

    public static final int PROMO_QUANTITY = 1;
    public static final int POSTAGE = 0;

    public static final String PROMO_PREFIX = "promo_product_";
    public static final String PROMO_STOCK_PREFIX = "promo_product_";

    public static String generatePromoProductKey (Long promoId, Long skuId, Long spuId) {
        return "promo_product_" + promoId + "_" + skuId + "_" + spuId;
    }

    public static String generatePromoStockKey (Long promoId, Long skuId, Long spuId) {
        return "promo_product_stock_" + promoId + "_" + skuId + "_" + spuId;
    }
}

