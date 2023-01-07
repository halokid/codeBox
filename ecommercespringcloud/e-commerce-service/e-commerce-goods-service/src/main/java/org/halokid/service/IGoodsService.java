package org.halokid.service;

import org.halokid.common.TableId;
import org.halokid.goods.DeductGoodsInventory;
import org.halokid.goods.GoodsInfo;
import org.halokid.goods.SimpleGoodsInfo;
import org.halokid.vo.PageSimpleGoodsInfo;

import java.util.List;

public interface IGoodsService {

  // 根据TableId查询商品详细信息
  List<GoodsInfo> getGoodsInfoByTableId(TableId tableId);

  //  获取分页的商品信息
  PageSimpleGoodsInfo getSimpleGoodsInfoByPage(int page);

  // 根据TableId 查询商品简单信息
  List<SimpleGoodsInfo> getSimpleGoodsInfoByTableId(TableId tableId);

  // 扣减商品库存
  Boolean deductGoodsInventory(List<DeductGoodsInventory> deductGoodsInventories);
}
