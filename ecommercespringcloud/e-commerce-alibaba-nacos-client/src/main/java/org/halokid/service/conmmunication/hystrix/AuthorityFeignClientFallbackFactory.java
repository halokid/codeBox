package org.halokid.service.conmmunication.hystrix;

import org.halokid.service.conmmunication.AuthorityFeignClient;
import org.halokid.vo.JwtToken;
import org.halokid.vo.UsernameAndPassword;
import feign.hystrix.FallbackFactory;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

/**
 * <h1>OpenFeign 集成 Hystrix 的另一种模式</h1>
 */
@Slf4j
@Component
public class AuthorityFeignClientFallbackFactory
    implements FallbackFactory<AuthorityFeignClient> {

  @Override
  public AuthorityFeignClient create(Throwable throwable) {

    log.warn("authority feign client get token by feign request error " +
        "(Hystrix FallbackFactory): [{}]", throwable.getMessage(), throwable);

    return new AuthorityFeignClient() {

      @Override
      public JwtToken getTokenByFeign(UsernameAndPassword usernameAndPassword) {
        return new JwtToken("qinyi-factory");
      }
    };
  }
}
