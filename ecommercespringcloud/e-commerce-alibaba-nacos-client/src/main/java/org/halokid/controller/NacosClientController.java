package org.halokid.controller;

import org.halokid.config.ProjectConfig;
import org.halokid.service.NacosClientService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

/**
 * <h1>nacos client controller</h1>
 */
@Slf4j
@RestController
@RequestMapping("/nacos-client")
public class NacosClientController {
  @Autowired
  private NacosClientService nacosClientService;

  @Autowired
  private ProjectConfig projectConfig;


  /**
   * <h2>根据 service id 获取服务所有的实例信息</h2>
   */
  @GetMapping("/service-instance")
  public List<ServiceInstance> logNacosClientInfo(
      @RequestParam(defaultValue = "e-commerce-nacos-client") String serviceId) {

    log.info("coming in log nacos client info: [{}]", serviceId);
    return nacosClientService.getNacosClientInfo(serviceId);
  }

  /**
   * <h2>动态获取 Nacos 中的配置信息</h2>
   */
  @GetMapping("/project-config")
  public ProjectConfig getProjectConfig() {
    return projectConfig;
  }
}
