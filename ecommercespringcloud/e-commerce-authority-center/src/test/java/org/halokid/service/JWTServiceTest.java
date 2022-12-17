package org.halokid.service;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.util.TokenParseUtil;
import org.halokid.vo.LoginUserInfo;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

/**
 * JWT相关服务测试类
 */
@Slf4j
@SpringBootTest
@RunWith(SpringRunner.class)
public class JWTServiceTest {

  @Autowired
  private IJWTService ijwtService;

  @Test
  public void testGenerateAndPareToken()throws Exception{
    String jwtToKen = ijwtService.generateToken(
        "halokid",
        "e10adc3949ba59abbe56e057f20f883e"
    );
    log.info("jwt token is: [{}]",jwtToKen);

    LoginUserInfo userInfo = TokenParseUtil.parseUserInfoFromToken(jwtToKen);
    log.info("parse token: [{}]", JSON.toJSONString(userInfo));
  }

}


