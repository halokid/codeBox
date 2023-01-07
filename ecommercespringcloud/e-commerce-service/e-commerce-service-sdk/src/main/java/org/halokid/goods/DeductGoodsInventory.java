package org.halokid.goods;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 */
@ApiModel(description = "扣减商品库存对象")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class DeductGoodsInventory {

  @ApiModelProperty(value = "商品主键 id")
  private Long goodsId;

  @ApiModelProperty(value = "扣减的个数")
  private Integer count;
}


