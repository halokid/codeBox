package org.halokid.service.hystrix;

import org.halokid.service.NacosClientService;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixCommand;
import com.netflix.hystrix.contrib.javanica.cache.annotation.CacheKey;
import com.netflix.hystrix.contrib.javanica.cache.annotation.CacheRemove;
import com.netflix.hystrix.contrib.javanica.cache.annotation.CacheResult;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * <h1>使用注解方式开启 Hystrix 请求缓存</h1>
 */
@Slf4j
@Service
public class CacheHystrixCommandAnnotation {

  private final NacosClientService nacosClientService;

  public CacheHystrixCommandAnnotation(NacosClientService nacosClientService) {
    this.nacosClientService = nacosClientService;
  }

  // 第一种 Hystrix Cache 注解的使用方法
  @CacheResult(cacheKeyMethod = "getCacheKey")
  @HystrixCommand(commandKey = "CacheHystrixCommandAnnotation")
  public List<ServiceInstance> useCacheByAnnotation01(String serviceId) {
    log.info("use cache01 to get nacos client info: [{}]", serviceId);
    return nacosClientService.getNacosClientInfo(serviceId);
  }

  @CacheRemove(commandKey = "CacheHystrixCommandAnnotation",
      cacheKeyMethod = "getCacheKey")
  @HystrixCommand
  public void flushCacheByAnnotation01(String cacheId) {
    log.info("flush hystrix cache key: [{}]", cacheId);
  }

  //返回值必须是String类型不然会失效或者报错
  public String getCacheKey(String cacheId) {
    return cacheId;
  }

  /**
   * TODO: 这种是最常用的
   *  第二种 Hystrix Cache 注解的使用方法，@CacheResult不再指定key的获取方法
   *  用@CahceKey注解标注要作为key的方法参数
   */
  @CacheResult
  @HystrixCommand(commandKey = "CacheHystrixCommandAnnotation")
  public List<ServiceInstance> useCacheByAnnotation02(@CacheKey String serviceId) {

    log.info("use cache02 to get nacos client info: [{}]", serviceId);
    return nacosClientService.getNacosClientInfo(serviceId);
  }

  @CacheRemove(commandKey = "CacheHystrixCommandAnnotation")
  @HystrixCommand
  public void flushCacheByAnnotation02(@CacheKey String cacheId) {
    log.info("flush hystrix cache key: [{}]", cacheId);
  }

  /**
   * 第三种 Hystrix Cache 注解的使用方法：没有指定@CacheResult注解的cacheKeyMethod
   * 也没有用@CacheKey指定key值，此时使用默认值，即参数列表中所有的参数
   */
  @CacheResult
  @HystrixCommand(commandKey = "CacheHystrixCommandAnnotation")
  public List<ServiceInstance> useCacheByAnnotation03(String serviceId) {

    log.info("use cache03 to get nacos client info: [{}]", serviceId);
    return nacosClientService.getNacosClientInfo(serviceId);
  }

  @CacheRemove(commandKey = "CacheHystrixCommandAnnotation")
  @HystrixCommand
  public void flushCacheByAnnotation03(String cacheId) {
    log.info("flush hystrix cache key: [{}]", cacheId);
  }
}