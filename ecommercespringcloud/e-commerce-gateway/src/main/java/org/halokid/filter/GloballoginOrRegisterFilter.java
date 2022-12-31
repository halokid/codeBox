package org.halokid.filter;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.collections4.ResettableIterator;
import org.halokid.constant.CommonConstant;
import org.halokid.contant.GatewayConstant;
import org.halokid.util.TokenParseUtil;
import org.halokid.vo.JwtToken;
import org.halokid.vo.LoginUserInfo;
import org.halokid.vo.UsernameAndPassword;
import org.springframework.cloud.client.ServiceInstance;
import org.springframework.cloud.client.loadbalancer.LoadBalancerClient;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.cloud.gateway.filter.GlobalFilter;
import org.springframework.core.Ordered;
import org.springframework.core.io.buffer.DataBuffer;
import org.springframework.core.io.buffer.DataBufferUtils;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.server.reactive.ServerHttpRequest;
import org.springframework.http.server.reactive.ServerHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.nio.CharBuffer;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.atomic.AtomicReference;

/**
 * todo: 这个是全局过滤器，不用在factory文件夹里面实现factory方法
 */
@Slf4j
@Component
public class GloballoginOrRegisterFilter  implements GlobalFilter, Ordered {

  private final LoadBalancerClient loadBalancerClient;

  private final RestTemplate restTemplate;

  public GloballoginOrRegisterFilter(LoadBalancerClient loadBalancerClient, RestTemplate restTemplate) {
    this.loadBalancerClient = loadBalancerClient;
    this.restTemplate = restTemplate;
  }

  /**
   * <h2>登录、注册、鉴权</h2>
   * 1. 如果的登录或注册，则去授权中心拿到Token 并返回给客户端
   * 2. 如果是访问其他的服务，则鉴权，没有权限返回 401
   * @param exchange
   * @param chain
   * @return
   */
  @Override
  public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
    ServerHttpRequest request = exchange.getRequest();
    ServerHttpResponse response = exchange.getResponse();

    // 1.如果是登录
    if(request.getURI().getPath().contains(GatewayConstant.LOGIN_URI)){
      // 去授权中心拿 token
      String token = getTokenFromAuthorityCenter(
          request, GatewayConstant.AUTHORITY_CENTER_TOKEN_URL_FORMAT
      );
      // header 中不能设置null
      response.getHeaders().add(
          CommonConstant.JWT_USER_INFO_KEY,
          null == token ? "null" : token);
      response.setStatusCode(HttpStatus.OK);
      return response.setComplete();
    }

    // 2. 如果是注册
    if(request.getURI().getPath().contains(GatewayConstant.REGISTER_URI)){
      // 去授权中心拿 token: 先创建用户，再返回Token
      String token = getTokenFromAuthorityCenter(
          request,GatewayConstant.AUTHORITY_CENTER_REGISTER_URL_FORMAT
      );
      // header 中不能设置null
      response.getHeaders().add(
          CommonConstant.JWT_USER_INFO_KEY,
          null == token ? "null" : token);
      response.setStatusCode(HttpStatus.OK);
      return response.setComplete();
    }

    // 3. 访问其他的服务，则鉴权，校验是否能够从Token 中解析出用户信息
    HttpHeaders headers = request.getHeaders();
    String token = headers.getFirst(CommonConstant.JWT_USER_INFO_KEY);
    log.info("get token from header: {}", token);
    LoginUserInfo loginUserInfo = null;
    try {
      loginUserInfo = TokenParseUtil.parseUserInfoFromToken(token);
    }catch (Exception ex){
      log.error("parse user info from token error: [{}]",ex.getMessage());
    }
    // 获取不到登录用户信息，返回401
    if(null==loginUserInfo){
      response.setStatusCode(HttpStatus.UNAUTHORIZED);
      return response.setComplete();
    }

    // 解析通过，则方行
    log.info("--pass--");
    return chain.filter(exchange);
  }

  @Override
  public int getOrder() {
    return HIGHEST_PRECEDENCE+2;
  }

  private String getTokenFromAuthorityCenter(ServerHttpRequest request, String uriFormat) {
    ServiceInstance serviceInstance = loadBalancerClient.choose(
        CommonConstant.AUTHORITY_CENTER_SERVICE_ID
    );

    log.info("Nacos Client info: {}, {}, {}", serviceInstance.getServiceId(),
        serviceInstance.getInstanceId(), JSON.toJSONString(serviceInstance.getMetadata()));

    String requestUrl = String.format(uriFormat, serviceInstance.getHost(), serviceInstance.getPort());

    UsernameAndPassword requestBody = JSON.parseObject(
      parseBodyFromRquest(request), UsernameAndPassword.class
    );

    log.info("login request url and body: {}, {}", requestUrl, JSON.toJSONString(requestBody));

    HttpHeaders headers = new HttpHeaders();
    headers.setContentType(MediaType.APPLICATION_JSON);
    JwtToken token = restTemplate.postForObject(requestUrl,
        new HttpEntity<>(JSON.toJSONString(requestBody), headers),
        JwtToken.class);

    if (null != token) {
      return token.getToken();
    }

    return null;
  }

  private String parseBodyFromRquest(ServerHttpRequest request) {
    Flux<DataBuffer> body = request.getBody();
    AtomicReference<String>  bodyRef = new AtomicReference<>();

    body.subscribe(buffer -> {
      CharBuffer charBuffer = StandardCharsets.UTF_8.decode(buffer.asByteBuffer());

      DataBufferUtils.release(buffer);
      bodyRef.set(charBuffer.toString());
    });

    return bodyRef.get();
  }
}
