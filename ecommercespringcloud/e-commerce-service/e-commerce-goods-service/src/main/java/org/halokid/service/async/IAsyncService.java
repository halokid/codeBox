package org.halokid.service.async;

import org.halokid.goods.GoodsInfo;

import java.util.List;

public interface IAsyncService {

  // 异步将商品信息保存下来
  void asyncImportGoods(List<GoodsInfo> goodsInfos, String taskId);
}
