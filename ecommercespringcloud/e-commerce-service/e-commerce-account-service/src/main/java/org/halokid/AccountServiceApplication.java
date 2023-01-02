package org.halokid;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Import;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

/**
 * <h1>用户账户微服务启动入口</h1>
 * 127.0.0.1:8003/e-commerce-account-service/doc.html
 * http://127.0.0.1:8003/e-commerce-account-service/swagger-ui.html
 */
@EnableJpaAuditing
@EnableDiscoveryClient
@SpringBootApplication
//@Import(DataSourceProxyAutoConfiguration.class)
public class AccountServiceApplication {

  public static void main(String[] args) {
    SpringApplication.run(AccountServiceApplication.class,args);
  }
}