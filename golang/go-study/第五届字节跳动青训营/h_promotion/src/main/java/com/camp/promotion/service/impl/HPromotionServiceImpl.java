package com.camp.promotion.service.impl;

import com.camp.promotion.common.BizException;
import com.camp.promotion.common.Constant;
import com.camp.promotion.common.ResponseEnum;
import com.camp.promotion.convertet.PromoConverter;
import com.camp.promotion.entity.*;
import com.camp.promotion.enums.ProductStatusEnum;
import com.camp.promotion.enums.PromoStatusEnum;
import com.camp.promotion.model.*;
import com.camp.promotion.service.*;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.CollectionUtils;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@Slf4j
@Service
public class HPromotionServiceImpl implements HPromotionService {

    @Resource
    private HPromoService hPromoService;
    @Resource
    private HPromoProductService hPromoProductService;
    @Resource
    private HSkuService hSkuService;
    @Resource
    private HProductService hProductService;
    @Resource
    private RedisService redisService;
    @Resource
    private IdGenerateService idGenerateService;


    @Override
    @Transactional(rollbackFor = Exception.class)
    public List<CreatePromoProductModel> CreatePromoActivity(CreateActivityModel createActivityModel) {

        HPromo promo = createActivityModel.convert(PromoConverter::convertHPromo);
        Long promoId = idGenerateService.generateId();
        promo.setId(promoId);
        int insertPromo = hPromoService.insert(promo);
        if (insertPromo < 1) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        List<CreatePromoProductModel> promoProducts = createActivityModel.getPromoProducts();
        Set<Long> spuIds = promoProducts.stream()
                .map(CreatePromoProductModel::getSpuId).collect(Collectors.toSet());
        if (CollectionUtils.isEmpty(spuIds)) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        // 获取spu并检验状态
        List<HProduct> products = hProductService.queryByIdList(List.copyOf(spuIds));
        if (CollectionUtils.isEmpty(products)) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        products = products.stream().filter(spu -> spu.getStatus() == ProductStatusEnum.ONLINE.getCode()).collect(Collectors.toList());
        if (CollectionUtils.isEmpty(products)) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        // 过滤掉下架的spu
        List<Long> finalSpuIds = products.stream().map(HProduct::getId).collect(Collectors.toList());
        promoProducts = promoProducts.stream().map(p -> {
            p.setPromoId(promo.getId());
            p.setPromoName(promo.getPromoName());
            return p;
        }).filter(p -> finalSpuIds.contains(p.getSpuId())).collect(Collectors.toList());

        // 从原库存扣除活动库存
        List<CreatePromoProductModel> successList = new ArrayList<>(promoProducts.size());
        for (CreatePromoProductModel model : promoProducts) {
            int decreaseStock = this.hSkuService.decreaseStock(model.getSkuId(), model.getPromoStock());

            if (decreaseStock == 1) {
                this.redisService.setVal(Constant.generatePromoStockKey(promoId, model.getSkuId(), model.getSpuId()), model.getPromoStock());
                successList.add(model);
            }
        }
        if (CollectionUtils.isEmpty(successList)) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        List<HPromoProduct> promoProductList = successList.stream().map(s -> {
            Long promoProductId = idGenerateService.generateId();
            HPromoProduct promoProduct = s.convert(PromoConverter::convertHPromoProduct);
            promoProduct.setId(promoProductId);
            return promoProduct;
        }).collect(Collectors.toList());


        int res = this.hPromoProductService.insertBatch(promoProductList);
        if (res < promoProductList.size()) {
            throw new BizException(ResponseEnum.CREATE_ACTIVITY_FAIL);
        }

        return successList;
    }

    @Override
    public HPromoProductModel getPromotionProductDetail(Long promoId, Long skuId, Long spuId) {

        // promotion activity
        HPromo promo = hPromoService.queryById(promoId);
        if (promo == null) {
            log.error("promo activity not exist, promoId = {}", promoId);
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }
        Date now = new Date();
        if (promo.getStatus() != PromoStatusEnum.ONLINE.getCode()
                || promo.getEndDate().before(now) || promo.getStartDate().after(now)) {
            log.error("promo activity not exist, promoId = {}", promoId);
            throw new BizException(ResponseEnum.ACTIVITY_OFFLINE);
        }

        // promotion activity product
        HPromoProduct hPromoProduct = hPromoProductService.queryByIdAndSkuId(promoId, skuId, spuId);
        if (hPromoProduct == null) {
            log.error("promo activity not exist, promoId = {}, skuId = {}, spuId = {}", promoId, skuId, spuId);
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }

        // spu info
        HProduct spu = hProductService.queryById(spuId);
        if (spu == null || spu.getStatus() != ProductStatusEnum.ONLINE.getCode()) {
            log.error("promo activity spu not exist, promoId = {}, skuId = {}, spuId = {}", promoId, skuId, spuId);
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }

        // sku info
        HSku sku = hSkuService.queryById(skuId);
        if (sku == null) {
            log.error("promo activity sku not exist, promoId = {}, skuId = {}, spuId = {}", promoId, skuId, spuId);
            throw new BizException(ResponseEnum.ACTIVITY_NOT_EXIST);
        }

        HPromoSpuModel spuModel = spu.convert(PromoConverter::convertSpuModel);

        HPromoSkuModel skuModel = new HPromoSkuModel(hPromoProduct, sku);
        skuModel = skuModel.convert(PromoConverter::convertSkuModel);

        HPromoProductModel promoProductModel = promo.convert(PromoConverter::convertProductModel);
        promoProductModel.setSpuModel(spuModel);
        promoProductModel.setSkuModel(skuModel);

        return promoProductModel;
    }
}
