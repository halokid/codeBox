package org.halokid.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.discovery.DiscoveryClient;
import org.springframework.stereotype.Service;

import java.util.List;

@Slf4j
@Service
public class NacoClientService {
  private final DiscoveryClient discoveryClient;

  public NacoClientService(DiscoveryClient discoveryClient) {
    this.discoveryClient = discoveryClient;
  }

  public List<ServiceInstance> getNacosClientInfo(String serviceId) {
    log.info("request nacos client to get service instance info: {}", serviceId);
    return discoveryClient.getInstances(serviceId);
  }
}



