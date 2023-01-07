package org.halokid.service.async;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.CollectionUtils;
import org.apache.commons.collections4.IterableUtils;
import org.apache.commons.lang3.time.StopWatch;
import org.halokid.constant.GoodsConstant;
import org.halokid.dao.EcommerceGoodsDao;
import org.halokid.entity.EcommerceGoods;
import org.halokid.goods.GoodsInfo;
import org.halokid.goods.SimpleGoodsInfo;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.*;
import java.util.concurrent.TimeUnit;
import java.util.stream.Collectors;


/**
 * <h1>异步服务接口实现</h1>
 */
@Slf4j
@Service
@Transactional
public class AsyncServiceImpl implements IAsyncService{
  private final EcommerceGoodsDao ecommerceGoodsDao;
  private final StringRedisTemplate redisTemplate;

  public AsyncServiceImpl(EcommerceGoodsDao ecommerceGoodsDao,
                          StringRedisTemplate redisTemplate) {
    this.ecommerceGoodsDao = ecommerceGoodsDao;
    this.redisTemplate = redisTemplate;
  }

  /**
   * <h2>异步任务需要加上注解，并指定使用的线程池</h2>
   * 异步任务处理两件事：
   * 1. 将商品信息保存到数据表
   * 2. 更新商品缓存
   * @param goodsInfos
   * @param taskId
   */
  // TODO: 这个是异步执行的关键2， 定义了这个注解，下面的func就可以异步执行，但是触发执行是在 AsyncTaskMonitor 里面触发的
  @Async("getAsyncExecutor")
  @Override
  public void asyncImportGoods(List<GoodsInfo> goodsInfos, String taskId) {
    log.info("async task running taskId: [{}]",taskId);
    StopWatch watch= StopWatch.createStarted();

    // 1. 如果是goodsInfo 中存在重复的商品，不保存；直接返回，记录错误日志
    // 请求数据是否合法的标记
    boolean isIllegal = false;
    // 将商品信息字段 joint在一起，用来判断是否存在重复
    Set<String> goodsJointInfos = new HashSet<>(goodsInfos.size());
    // 过滤出来的，可以入库的商品信息（规则按照自己的业务需求自定义即可）
    List<GoodsInfo> filteredGoodsInfo= new ArrayList<>(goodsInfos.size());

    //走一遍循环，过滤非法参数与判定当前请求是否合法
    for(GoodsInfo goods: goodsInfos){
      // 基本条件不满足的，直接过滤器
      if(goods.getPrice()<=0||goods.getSupply()<=0){
        log.info("good info is invalid:[{}]", JSON.toJSONString(goods));
        continue;
      }
      // 组合商品信息
      String jointInfo = String.format(
          "%s,%s,%s",
          goods.getBrandCategory(),
          goods.getBrandCategory(),
          goods.getGoodsName()
      );
      if(goodsJointInfos.contains(jointInfo)){
        isIllegal=true;
      }
      // 加入到两个容器中
      goodsJointInfos.add(jointInfo);
      filteredGoodsInfo.add(goods);
    }

    // 如果存在重复商品或者是没有需要入库的商品，直接打印日志返回
    if(isIllegal || CollectionUtils.isEmpty(filteredGoodsInfo)){
      watch.stop();
      log.warn("import nothing:[{}]",JSON.toJSONString(filteredGoodsInfo));
      log.info("check and import goods done: [{}ms]",watch.getTime(TimeUnit.MILLISECONDS));
      return;
    }

    List<EcommerceGoods> ecommerceGoods = filteredGoodsInfo.stream()
        .map(EcommerceGoods::to)
        .collect(Collectors.toList());
    List<EcommerceGoods> targetGoods = new ArrayList<>(ecommerceGoods.size());

    //2. 保存goodsInfo之前先判断下是否存在重复商品
    ecommerceGoods.forEach(g -> {
      // limit 1
      if(null != ecommerceGoodsDao.findFirstByGoodsCategoryAndBrandCategoryAndGoodsName(
          g.getGoodsCategory(), g.getBrandCategory(), g.getGoodsName()
      ).orElse(null)){
        return;
      }

      targetGoods.add(g);
    });

    //商品信息入库
    List<EcommerceGoods> saveGoods = IterableUtils.toList(
        ecommerceGoodsDao.saveAll(targetGoods)
    );

    //TODO 将入库商品信息同步到Redis中
    log.info("save goods info to db and redis:[{}]",saveGoods.size());
    watch.stop();
    log.info("check and import goods success:[{}]",watch.getTime(TimeUnit.MILLISECONDS));
    saveNewGoodsInfoToRedis(saveGoods);
  }

  /**
   * <h2>将保存到数据表中的数据缓存到Redis中</h2>
   * dict: key-><id, SimpleGoodsInfo(json)>
   * @param saveGoods
   */
  private void saveNewGoodsInfoToRedis(List<EcommerceGoods> saveGoods){
    // 由于Redis 是内存存储，只存储简单商品信息
    List<SimpleGoodsInfo> simpleGoodsInfos=saveGoods.stream()
        .map(EcommerceGoods::toSimple)
        .collect(Collectors.toList());

    Map<String,String> id2JsonObject = new HashMap<>(simpleGoodsInfos.size());
    simpleGoodsInfos.forEach(
        g->id2JsonObject.put(g.getId().toString(), JSON.toJSONString(g))
    );
    // 保存到Redis中
    redisTemplate.opsForHash().putAll(
        GoodsConstant.ECOMMERCE_GOODS_DICT_KEY,
        id2JsonObject
    );
  }
}


