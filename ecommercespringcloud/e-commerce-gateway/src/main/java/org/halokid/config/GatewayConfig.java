package org.halokid.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Configuration;

/**
 * <h1>配置类，读取Nacos 相关的配置项，用于配置监听器</h1>
 */
@Configuration
public class GatewayConfig {

  /** 读取配置的超时时间*/
  public static final long DEFAULT_TIMEOUT = 30000;

  /** Nacos 服务器地址*/
  public static String NACOS_SERVER_ADDR;

  /** 命名空间*/
  public static String NACOS_NAMESPACE;

  /** data-id*/
  public static String NACOS_ROUTE_DATA_ID;
  /** 分组id*/
  public static String NACOS_ROUTE_GROUP;

  @Value("${spring.cloud.nacos.discovery.server-addr}")
  public void setNacosServerAddr(String nacosServerAddr){
    NACOS_SERVER_ADDR=nacosServerAddr;
  }

  @Value("${spring.cloud.nacos.discovery.namespace}")
  public void setNacosNamespace(String nacosNamespace){
    NACOS_NAMESPACE = nacosNamespace;
  }

  @Value("${nacos.gateway.route.config.data-id}")
  public void setNacosRouteDataId(String nacosRouteDataId){
    NACOS_ROUTE_DATA_ID = nacosRouteDataId;
  }

  @Value("${nacos.gateway.route.config.group}")
  public void setNacosRouteGroup(String nacosRouteGroup){
    NACOS_ROUTE_GROUP = nacosRouteGroup;
  }
}



