package com.camp.promotion.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;


@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class HPromoSpuModel implements Serializable {

    private static final long serialVersionUID = 258795279137834888L;
    /**
     * 商品id
     */
    private Long spuId;
    /**
     * 分类id,对应category表的主键
     */
    private Long categoryId;
    /**
     * 店铺id
     */
    private Long shopId;
    /**
     * 商品名称
     */
    private String title;
    /**
     * 商品副标题
     */
    private String subtitle;
    /**
     * 产品主图,url相对地址
     */
    private String mainImage;
    /**
     * 图片地址,json格式,扩展用
     */
    private String subImages;
    /**
     * 商品详情
     */
    private String detail;
    /**
     * 价格,吊牌价
     */
    private Integer price;
    /**
     * 类目属性信息，json
     */
    private String categoryData;
    /**
     * 规格信息，json
     */
    private String specData;
    /**
     * 商品状态.1-在售 2-下架 3-删除
     */
    private Integer status;

}
