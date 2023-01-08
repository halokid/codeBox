package org.halokid.service;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.common.TableId;
import org.halokid.goods.DeductGoodsInventory;
import org.halokid.goods.GoodsInfo;
import org.halokid.goods.SimpleGoodsInfo;
import org.halokid.vo.PageSimpleGoodsInfo;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 * <h1>商品微服务功能测试</h1>
 */
@Slf4j
@SpringBootTest
@RunWith(SpringRunner.class)
public class GoodsServiceTest {

  @Autowired
  private IGoodsService goodsService;

  @Test
  public void testGetGoodsInfoByTableId() {
    List<Long> ids = Arrays.asList(1L, 2L);
    List<TableId.Id>  tIds = ids.stream()
        .map(TableId.Id::new).collect(Collectors.toList());

    List<GoodsInfo> goodsInfos = goodsService.getGoodsInfoByTableId(new TableId(tIds));
    log.info("test get goods info by table id: {}", JSON.toJSONString(goodsInfos));
  }

  @Test
  public void testGetSimpleGoodsInfoByPage() {
    PageSimpleGoodsInfo pageSimpleGoodsInfo = goodsService.getSimpleGoodsInfoByPage(1);
    log.info("test get simple goods info by page: {}", JSON.toJSONString(pageSimpleGoodsInfo));
  }

  @Test
  public void testGetSimpleGoodsInfoByTableId() {
//    List<Long> ids = Arrays.asList(1L, 2L, 3L);
    List<Long> ids = Arrays.asList(1L, 2L);
    List<TableId.Id> tIds = ids.stream()
        .map(TableId.Id::new).collect(Collectors.toList());

    List<SimpleGoodsInfo> simpleGoodsInfos = goodsService.getSimpleGoodsInfoByTableId(
        new TableId(tIds)
    );
    log.info("test get simple goods info by tableId: {}", JSON.toJSONString(simpleGoodsInfos));
  }

  @Test
  public void testDeductGoodsInventory() {
    List<DeductGoodsInventory> deductGoodsInventories = Arrays.asList(
      new DeductGoodsInventory(1L, 100),
      new DeductGoodsInventory(2L, 66)
    );
    log.info("test deduct goods inventory: {}",
        goodsService.deductGoodsInventory(deductGoodsInventories));
  }
}






