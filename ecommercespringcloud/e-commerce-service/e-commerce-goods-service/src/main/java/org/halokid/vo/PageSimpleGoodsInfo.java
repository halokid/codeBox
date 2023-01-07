package org.halokid.vo;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.halokid.goods.SimpleGoodsInfo;

import java.util.List;

@ApiModel(description = "分页商品信息对象")
@Data
@NoArgsConstructor
@AllArgsConstructor
public class PageSimpleGoodsInfo {

  @ApiModelProperty(value = "分页简单商品信息")
  private List<SimpleGoodsInfo> simpleGoodsInfoList;

  @ApiModelProperty(value = "是否有更多的商品(分页)")
  private Boolean hasMore;
}

