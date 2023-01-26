package org.halokid.controller;

import com.alibaba.fastjson.JSON;
import org.halokid.service.NacosClientService;
import org.halokid.service.hystrix.*;
import org.halokid.service.hystrix.request_merge.NacosClientCollapseCommand;
import lombok.extern.slf4j.Slf4j;
import org.halokid.service.hystrix.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;
import rx.Observable;
import rx.Observer;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.concurrent.Future;

/**
 * <h1>Hystrix Controller</h1>
 */
@Slf4j
@RestController
@RequestMapping("/hystrix")
public class HystrixController {
  @Autowired
  private UseHystrixCommandAnnotation hystrixCommandAnnotation;
  @Autowired
  private NacosClientService nacosClientService;
  @Autowired
  private CacheHystrixCommandAnnotation cacheHystrixCommandAnnotation;

  @GetMapping("/hystrix-command-annotation")
  public List<ServiceInstance> getNacosClientInfoUseAnnotation(
      @RequestParam String serviceId) {
    log.info("request nacos client info use annotation: [{}], [{}]",
        serviceId, Thread.currentThread().getName());
    return hystrixCommandAnnotation.getNacosClientInfo(serviceId);
  }

  @GetMapping("/simple-hystrix-command")
  public List<ServiceInstance> getServiceInstanceByServiceId(
      @RequestParam String serviceId) throws Exception {

    // 第一种方式
    List<ServiceInstance> serviceInstances01 = new NacosClientHystrixCommand(
        nacosClientService, serviceId
    ).execute();    // 同步阻塞，它会创建一个新线程执行NacosClientHystrixCommand中的run
    log.info("use execute to get service instances: [{}], [{}]",
        JSON.toJSONString(serviceInstances01), Thread.currentThread().getName());

    // 第二种方式
    List<ServiceInstance> serviceInstances02;
    Future<List<ServiceInstance>> future = new NacosClientHystrixCommand(
        nacosClientService, serviceId
    ).queue();      // 异步非阻塞
    // 这里可以做一些别的事, 需要的时候再去拿结果
    serviceInstances02 = future.get();
    log.info("use queue to get service instances: [{}], [{}]",
        JSON.toJSONString(serviceInstances02), Thread.currentThread().getName());

    // 第三种方式
    Observable<List<ServiceInstance>> observable = new NacosClientHystrixCommand(
        nacosClientService, serviceId
    ).observe();        // 热响应调用
    List<ServiceInstance> serviceInstances03 = observable.toBlocking().single();
    log.info("use observe to get service instances: [{}], [{}]",
        JSON.toJSONString(serviceInstances03), Thread.currentThread().getName());

    // 第四种方式
    Observable<List<ServiceInstance>> toObservable = new NacosClientHystrixCommand(
        nacosClientService, serviceId
    ).toObservable();        // 异步冷响应调用
    List<ServiceInstance> serviceInstances04 = toObservable.toBlocking().single();
    log.info("use toObservable to get service instances: [{}], [{}]",
        JSON.toJSONString(serviceInstances04), Thread.currentThread().getName());

    // execute = queue + get
    return serviceInstances01;
  }

  @GetMapping("/hystrix-observable-command")
  public List<ServiceInstance> getServiceInstancesByServiceIdObservable(
      @RequestParam String serviceId) {

    List<String> serviceIds = Arrays.asList(serviceId, serviceId, serviceId);
    List<List<ServiceInstance>> result = new ArrayList<>(serviceIds.size());

    NacosClientHystrixObservableCommand observableCommand =
        new NacosClientHystrixObservableCommand(nacosClientService, serviceIds);

    // 异步执行命令
    Observable<List<ServiceInstance>> observe = observableCommand.observe();

    // 注册获取结果
    observe.subscribe(
        new Observer<List<ServiceInstance>>() {
          // 执行 onNext 之后再去执行 onCompleted
          @Override
          public void onCompleted() {
            log.info("all tasks is complete: [{}], [{}]",
                serviceId, Thread.currentThread().getName());
          }

          @Override
          public void onError(Throwable e) {
            e.printStackTrace();
          }

          @Override
          public void onNext(List<ServiceInstance> instances) {
            result.add(instances);
          }
        }
    );

    log.info("observable command result is : [{}], [{}]",
        JSON.toJSONString(result), Thread.currentThread().getName());
    return result.get(0);
  }

  @GetMapping("/cache-hystrix-command")
  public void cacheHystrixCommand(@RequestParam String serviceId) {

    // 使用缓存 Command, 发起两次请求
    CacheHystrixCommand command1 = new CacheHystrixCommand(
        nacosClientService, serviceId
    );
    CacheHystrixCommand command2 = new CacheHystrixCommand(
        nacosClientService, serviceId
    );

    List<ServiceInstance> result01 = command1.execute();
    List<ServiceInstance> result02 = command2.execute();
    log.info("result01, result02: [{}], [{}]",
        JSON.toJSONString(result01), JSON.toJSONString(result02));

    // 清除缓存
    CacheHystrixCommand.flushRequestCache(serviceId);

    // 使用缓存 Command, 发起两次请求
    CacheHystrixCommand command3 = new CacheHystrixCommand(
        nacosClientService, serviceId
    );
    CacheHystrixCommand command4 = new CacheHystrixCommand(
        nacosClientService, serviceId
    );

    List<ServiceInstance> result03 = command3.execute();
    List<ServiceInstance> result04 = command4.execute();
    log.info("result03, result04: [{}], [{}]",
        JSON.toJSONString(result03), JSON.toJSONString(result04));
  }

  @GetMapping("/cache-annotation-01")
  public List<ServiceInstance> useCacheByAnnotation01(@RequestParam String serviceId) {

    log.info("use cache by annotation01(controller) to get nacos client info: [{}]",
        serviceId);

    List<ServiceInstance> result01 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation01(serviceId);
    List<ServiceInstance> result02 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation01(serviceId);

    // 清除掉缓存
    cacheHystrixCommandAnnotation.flushCacheByAnnotation01(serviceId);

    List<ServiceInstance> result03 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation01(serviceId);
    // 这里有第四次调用
    return cacheHystrixCommandAnnotation.useCacheByAnnotation01(serviceId);
  }

  @GetMapping("/cache-annotation-02")
  public List<ServiceInstance> useCacheByAnnotation02(@RequestParam String serviceId) {

    log.info("use cache by annotation02(controller) to get nacos client info: [{}]",
        serviceId);

    List<ServiceInstance> result01 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation02(serviceId);
    List<ServiceInstance> result02 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation02(serviceId);

    // 清除掉缓存
    cacheHystrixCommandAnnotation.flushCacheByAnnotation02(serviceId);

    List<ServiceInstance> result03 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation02(serviceId);
    // 这里有第四次调用
    return cacheHystrixCommandAnnotation.useCacheByAnnotation02(serviceId);
  }

  @GetMapping("/cache-annotation-03")
  public List<ServiceInstance> useCacheByAnnotation03(@RequestParam String serviceId) {

    log.info("use cache by annotation03(controller) to get nacos client info: [{}]",
        serviceId);

    List<ServiceInstance> result01 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation03(serviceId);
    List<ServiceInstance> result02 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation03(serviceId);

    // 清除掉缓存
    cacheHystrixCommandAnnotation.flushCacheByAnnotation03(serviceId);

    List<ServiceInstance> result03 =
        cacheHystrixCommandAnnotation.useCacheByAnnotation03(serviceId);
    // 这里有第四次调用
    return cacheHystrixCommandAnnotation.useCacheByAnnotation03(serviceId);
  }

  /**
   * <h2>编程方式实现请求合并</h2>
   */
  @GetMapping("/request-merge")
  public void requestMerge() throws Exception {

    // 前三个请求会被合并
    NacosClientCollapseCommand collapseCommand01 = new NacosClientCollapseCommand(
        nacosClientService, "e-commerce-nacos-client1");
    NacosClientCollapseCommand collapseCommand02 = new NacosClientCollapseCommand(
        nacosClientService, "e-commerce-nacos-client2");
    NacosClientCollapseCommand collapseCommand03 = new NacosClientCollapseCommand(
        nacosClientService, "e-commerce-nacos-client3");

    Future<List<ServiceInstance>> future01 = collapseCommand01.queue();
    Future<List<ServiceInstance>> future02 = collapseCommand02.queue();
    Future<List<ServiceInstance>> future03 = collapseCommand03.queue();

    future01.get();
    future02.get();
    future03.get();

    Thread.sleep(2000);

    // 过了合并的时间窗口, 第四个请求单独发起
    NacosClientCollapseCommand collapseCommand04 = new NacosClientCollapseCommand(
        nacosClientService, "e-commerce-nacos-client4");
    Future<List<ServiceInstance>> future04 = collapseCommand04.queue();
    future04.get();
  }

  /**
   * <h2>注解的方式实现请求合并</h2>
   */
  @GetMapping("/request-merge-annotation")
  public void requestMergeAnnotation() throws Exception {

    Future<List<ServiceInstance>> future01 = nacosClientService.findNacosClientInfo(
        "e-commerce-nacos-client1"
    );
    Future<List<ServiceInstance>> future02 = nacosClientService.findNacosClientInfo(
        "e-commerce-nacos-client2"
    );
    Future<List<ServiceInstance>> future03 = nacosClientService.findNacosClientInfo(
        "e-commerce-nacos-client3"
    );

    future01.get();
    future02.get();
    future03.get();

    Thread.sleep(2000);

    Future<List<ServiceInstance>> future04 = nacosClientService.findNacosClientInfo(
        "e-commerce-nacos-client4"
    );
    future04.get();
  }
}