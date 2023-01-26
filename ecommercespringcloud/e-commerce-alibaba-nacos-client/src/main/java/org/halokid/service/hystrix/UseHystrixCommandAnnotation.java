package org.halokid.service.hystrix;

import org.halokid.service.NacosClientService;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixCommand;
import com.netflix.hystrix.contrib.javanica.annotation.HystrixProperty;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;

/**
 * @author r0x
 * @description <h1>UseHystrixCommandAnnotation</h1>
 * @date 2021/9/23 9:32
 */
@Slf4j
@Service
public class UseHystrixCommandAnnotation {
  private final NacosClientService nacosClientService;

  public UseHystrixCommandAnnotation(NacosClientService nacosClientService) {
    this.nacosClientService = nacosClientService;
  }

  @HystrixCommand(
      // 用于对 Hystrix 命令进行分组, 分组之后便于统计展示于仪表盘、上传报告和预警等等
      // 内部进行度量统计时候的分组标识, 数据上报和统计的最小维度就是 groupKey
      groupKey = "NacosClientService",
      // HystrixCommand 的名字, 默认是当前类的名字, 主要方便 Hystrix 进行监控、报警等
      commandKey = "NacosClientService",
      // 舱壁模式
      threadPoolKey = "NacosClientService",
      // 后备模式
      fallbackMethod = "getNacosClientInfoFallback",
      // 断路器模式
      commandProperties = {
          // 超时时间, 单位毫秒, 超时进 fallback
          @HystrixProperty(name = "execution.isolation.thread.timeoutInMilliseconds", value = "1500"),
          // 判断熔断的最少请求数, 默认是10; 只有在一定时间内请求数量达到该值, 才会进行成功率的计算
          @HystrixProperty(name = "circuitBreaker.requestVolumeThreshold", value = "10"),
          // 熔断的阈值默认值 50, 表示在一定时间内有50%的请求处理失败, 会触发熔断
          @HystrixProperty(name = "circuitBreaker.errorThresholdPercentage", value = "10"),
      },
      // 舱壁模式
      threadPoolProperties = {
          @HystrixProperty(name = "coreSize", value = "30"),
          @HystrixProperty(name = "maxQueueSize", value = "101"),
          @HystrixProperty(name = "keepAliveTimeMinutes", value = "2"),
          @HystrixProperty(name = "queueSizeRejectionThreshold", value = "15"),
          // 在时间窗口中, 收集统计信息的次数; 在 1440ms 的窗口中一共统计 12 次
          @HystrixProperty(name = "metrics.rollingStats.numBuckets", value = "12"),
          // 时间窗口, 从监听到第一次失败开始计时
          @HystrixProperty(name = "metrics.rollingStats.timeInMilliseconds", value = "1440")
      }
  )
  public List<ServiceInstance> getNacosClientInfo(String serviceId) {
    log.info("use hystrix command annotation to get nacos client info: [{}], [{}]",
        serviceId, Thread.currentThread().getName());
    return nacosClientService.getNacosClientInfo(serviceId);
  }

  /**
   * <h2>getNacosClientInfo 的兜底策略 - Hystrix 后备模式</h2>
   */
  public List<ServiceInstance> getNacosClientInfoFallback(String serviceId) {

    log.warn("can not get nacos client, trigger hystrix fallback: [{}], [{}]",
        serviceId, Thread.currentThread().getName());
    return Collections.emptyList();
  }
}
