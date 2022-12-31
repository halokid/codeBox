package org.halokid.config;

import com.alibaba.fastjson.JSON;
import com.alibaba.nacos.api.NacosFactory;
import com.alibaba.nacos.api.config.ConfigService;
import com.alibaba.nacos.api.config.listener.Listener;
import com.alibaba.nacos.api.exception.NacosException;
import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.gateway.route.RouteDefinition;
import org.springframework.context.annotation.DependsOn;
import org.springframework.stereotype.Component;
import org.springframework.util.CollectionUtils;

import javax.annotation.PostConstruct;
import java.util.List;
import java.util.Properties;
import java.util.concurrent.Executor;

/**
 * <h1> 通过nacos 下发动态路由配置，监听Nacos 中路由配置变更</h1>
 */
@Slf4j
@Component
@DependsOn({"gatewayConfig"})   // 依赖GatewayConfig 加载完成后再加载
public class DynamicRouteServiceImplByNacos {
  /**
   * Nacos 配置服务
   */
  private ConfigService configService;

  private final DynamicRouteServiceImpl dynamicRouteService;

  public DynamicRouteServiceImplByNacos(DynamicRouteServiceImpl dynamicRouteService) {
    this.dynamicRouteService = dynamicRouteService;
  }

  /**
   * <h2>Bean 在容器中构造完成就会执行 init 方法</h2>
   */
  @PostConstruct
  public void init() {
    log.info("gateway route init....");
    try {
      // 初始化 Nacos 配置客户端
      configService = initConfigService();
      if (null == configService) {
        log.error("init config service fail");
        return;
      }
      // 通过Nacos Config 并指定路由配置去获取路由配置
      String configInfo = configService.getConfig(
          GatewayConfig.NACOS_ROUTE_DATA_ID,
          GatewayConfig.NACOS_ROUTE_GROUP,
          GatewayConfig.DEFAULT_TIMEOUT
      );
      log.info("get current gateway config: [{}]", configInfo);
      List<RouteDefinition> definitionList = JSON.parseArray(configInfo, RouteDefinition.class);

      if (!CollectionUtils.isEmpty(definitionList)) {
        for (RouteDefinition definition : definitionList) {
          log.info("init gateway config: [{}]", definition.toString());
          dynamicRouteService.addRouteDefinition(definition);
        }
      }
      //
    } catch (Exception ex) {
      log.error("gateway route init hao some error:[{}]", ex.getMessage(), ex);
    }
    // 设置监听器
    dynamicRouteByNacosListener(GatewayConfig.NACOS_ROUTE_DATA_ID,
        GatewayConfig.NACOS_ROUTE_GROUP);
  }


  /**
   * <h2>初始化 Nacos Config</h2>
   *
   * @return
   */
  private ConfigService initConfigService() {
    try {
      Properties properties = new Properties();
      log.info("GatewayConfig.NACOS_SERVER_ADDR :{}", GatewayConfig.NACOS_SERVER_ADDR);
      properties.setProperty("serverAddr", GatewayConfig.NACOS_SERVER_ADDR);
      properties.setProperty("namespace", GatewayConfig.NACOS_NAMESPACE);
      return configService = NacosFactory.createConfigService(properties);
    } catch (Exception ex) {
      log.error("init gateway nacos config error: [{}]", ex.getMessage(), ex);
      return null;
    }
  }

  /**
   * <h2>监听 Nacos 下发的动态路由配置</h2>
   *
   * @param dataId
   * @param group
   */
  private void dynamicRouteByNacosListener(String dataId, String group) {
    try {
      // 给Nacos Config 客户端增加一个监听器
      configService.addListener(dataId, group, new Listener() {

        /**
         * <h2> 自己提供线程池执行操作</h2>
         * @return
         */
        @Override
        public Executor getExecutor() {
          return null;
        }

        /**
         * <h2>监听器收到配置更新</h2>
         * @param configInfo  Nacos 中最新的配置定义
         */
        @Override
        public void receiveConfigInfo(String configInfo) {
          log.info("start to update config: [{}]", configInfo);

          List<RouteDefinition> definitionList = JSON.parseArray(configInfo, RouteDefinition.class);
          log.info("update route: [{}]", definitionList.stream().toString());
          dynamicRouteService.updateList(definitionList);
        }
      });
    } catch (NacosException ex) {
      log.error("dynamic update gateway config error: [{}]", ex.getErrCode(), ex);
    }
  }
}

