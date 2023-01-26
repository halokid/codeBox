package org.halokid.service.conmmunication;

import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

/**
 * <h1>使用 Ribbon 之前的配置, 增强 RestTemplate</h1>
 */
@Component
public class RibbonConfig {

  /**
   * <h2>注入 RestTemplate</h2>
   */
  @Bean
  @LoadBalanced
  public RestTemplate restTemplate() {
    return new RestTemplate();
  }
}