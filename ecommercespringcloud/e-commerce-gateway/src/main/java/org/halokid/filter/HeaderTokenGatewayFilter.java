package org.halokid.filter;

import org.springframework.cloud.gateway.filter.GatewayFilter;
import org.springframework.cloud.gateway.filter.GatewayFilterChain;
import org.springframework.core.Ordered;
import org.springframework.http.HttpStatus;
import org.springframework.web.server.ServerWebExchange;
import reactor.core.publisher.Mono;

/**
 * todo: 这个是局部过滤器，要在factory文件夹里面实现factory方法
 * request header carry token authentication filter
 */
public class HeaderTokenGatewayFilter implements GatewayFilter, Ordered {
  @Override
  public Mono<Void> filter(ServerWebExchange exchange, GatewayFilterChain chain) {
    // find key is `token` from HTTP header
    String name = exchange.getRequest().getHeaders().getFirst("token");
    if ("halokid".equals(name)) {
      // todo: the request conntinue process
      return chain.filter(exchange);
    }

    // remark this request no permission, and end the request
    exchange.getResponse().setStatusCode(HttpStatus.UNAUTHORIZED);
    return exchange.getResponse().setComplete();
  }

  @Override
  public int getOrder() {
    return HIGHEST_PRECEDENCE;
  }
}



