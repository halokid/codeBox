package org.halokid.conf;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.client.RestTemplate;

@Configuration
public class GatewayBeanConf {

  @Bean
  public RestTemplate restTemplate(){
    return new RestTemplate();
  }
}
