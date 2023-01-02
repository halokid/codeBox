package org.halokid.entity;

import com.alibaba.fastjson.JSON;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.halokid.constant.BrandCategory;
import org.halokid.constant.GoodsCategory;
import org.halokid.constant.GoodsStatus;
import org.halokid.converter.BrandCategoryConverter;
import org.halokid.converter.GoodsCategoryConverter;
import org.halokid.converter.GoodsStatusConverter;
import org.halokid.goods.GoodsInfo;
import org.halokid.goods.SimpleGoodsInfo;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

/**
 * <h1>商品表实体类定义</h1>
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
@Entity
@EntityListeners(AuditingEntityListener.class)
@Table(name = "t_ecommerce_goods")
public class EcommerceGoods {

  /**自增主键*/
  @Id
  @GeneratedValue(strategy =GenerationType.IDENTITY)
  @Column(name = "id",nullable = false)
  private Long id;

  /**商品类型*/
  @Column(name = "goods_category",nullable = false)
  @Convert(converter = GoodsCategoryConverter.class)
  private GoodsCategory goodsCategory;

  /** 品牌分类*/
  @Column(name = "brand_category",nullable = false)
  @Convert(converter = BrandCategoryConverter.class)
  private BrandCategory brandCategory;

  /**商品名称*/
  @Column(name="goods_name",nullable = false)
  private String goodsName;


  /**商品图片*/
  @Column(name = "goods_pic",nullable = false)
  private String goodsPic;

  /**商品描述信息*/
  @Column(name ="goods_description",nullable = false)
  private String goodsDescription;

  /**商品状态*/
  @Column(name = "goods_status",nullable = false)
  @Convert(converter = GoodsStatusConverter.class)
  private GoodsStatus goodsStatus;

  /**商品价格  单位：分、厘*/
  @Column(name="price",nullable = false)
  private Integer price;

  /**总供应量*/
  @Column(name="supply",nullable = false)
  private Long supply;

  /**库存*/
  @Column(name = "inventory",nullable = false)
  private Long inventory;

  /**商品属性，json 字符串存储*/
  @Column(name="goods_property",nullable = false)
  private String goodsProperty;

  @CreatedDate
  @Column(name = "create_time",nullable = false)
  private Date createTime;

  @LastModifiedDate
  @Column(name = "update_time",nullable = false)
  private Date updateTime;

  /**
   * <h2>将 GoodsInfo 转成实体对象</h2>
   * @param goodsInfo
   * @return
   */
  public static EcommerceGoods to(GoodsInfo goodsInfo){
    EcommerceGoods ecommerceGoods= new EcommerceGoods();

    ecommerceGoods.setGoodsCategory(GoodsCategory.of(goodsInfo.getGoodsCategory()));
    ecommerceGoods.setBrandCategory(BrandCategory.of(goodsInfo.getBrandCategory()));
    ecommerceGoods.setGoodsName(goodsInfo.getGoodsName());
    ecommerceGoods.setGoodsPic(goodsInfo.getGoodsPic());
    ecommerceGoods.setGoodsDescription(goodsInfo.getGoodsDescription());
    ecommerceGoods.setGoodsStatus(GoodsStatus.ONLINE); //这里增加审核 流程
    ecommerceGoods.setPrice(goodsInfo.getPrice());
    ecommerceGoods.setSupply(goodsInfo.getSupply());
    ecommerceGoods.setInventory(goodsInfo.getSupply());
    ecommerceGoods.setGoodsProperty(JSON.toJSONString(goodsInfo.getGoodsProperty()));
    return ecommerceGoods;
  }

  /**
   * <h2>将实体对象转成 GoodsInfo对象</h2>
   * @return
   */
  public GoodsInfo toGoodsInfo(){
    GoodsInfo goodsInfo = new GoodsInfo();
    goodsInfo.setId(this.id);
    goodsInfo.setBrandCategory(this.brandCategory.getCode());
    goodsInfo.setGoodsName(this.goodsName);
    goodsInfo.setGoodsCategory(this.goodsCategory.getCode());
    goodsInfo.setGoodsPic(this.goodsPic);
    goodsInfo.setGoodsDescription(this.goodsDescription);
    goodsInfo.setGoodsStatus(this.goodsStatus.getStatus());
    goodsInfo.setPrice(this.price);
    goodsInfo.setGoodsProperty(JSON.parseObject(this.goodsProperty, GoodsInfo.GoodsProperty.class));
    goodsInfo.setSupply(this.supply);
    goodsInfo.setInventory(this.inventory);
    goodsInfo.setCreateTime(this.createTime);
    goodsInfo.setUpdateTime(this.updateTime);
    return  goodsInfo;
  }

  /**
   * <h2>将实体对象转成SimpleGoodsInfo对象</h2>
   * @return
   */
  public SimpleGoodsInfo toSimple(){
    SimpleGoodsInfo goodsInfo = new SimpleGoodsInfo();
    goodsInfo.setId(this.id);
    goodsInfo.setGoodsName(this.goodsName);
    goodsInfo.setGoodsPic(this.goodsPic);
    goodsInfo.setPrice(this.price);
    return goodsInfo;
  }
}











