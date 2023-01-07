package org.halokid.service;

import lombok.extern.slf4j.Slf4j;
import org.halokid.common.TableId;
import org.halokid.goods.GoodsInfo;
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
  }
}



