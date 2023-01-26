package org.halokid.config;

import lombok.Data;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.cloud.client.loadbalancer.LoadBalanced;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;

/**
 * <h1>数据配置绑定</h1>
 */
@Data
@Component
@ConfigurationProperties(prefix = "project")
public class ProjectConfig {

  private String name;
  private String org;
  private String version;
  private String author;
}
