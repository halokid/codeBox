package org.halokid.controller;

import lombok.extern.slf4j.Slf4j;
import org.halokid.service.NacoClientService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

/*
nacos client controller
 */
@Slf4j
@RestController
@RequestMapping("/nacos-client")
public class NacosClientController {
  @Autowired
  private NacoClientService nacoClientService;

  public NacosClientController(NacoClientService nacoClientService) {
    this.nacoClientService = nacoClientService;
  }

  // 根据 service id 获取服务所有的实例信息
  @GetMapping("/service-instance")
  public List<ServiceInstance> logNacosClientInfo(@RequestParam(defaultValue = "e-commerce-nacos-client") String serviceId) {
    log.info("coming in log nacos client info: {}", serviceId);
    return nacoClientService.getNacosClientInfo(serviceId);
  }
}




