package org.halokid.controller;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.annotation.IgnoreResponseAdvice;
import org.halokid.service.IJWTService;
import org.halokid.vo.JwtToken;
import org.halokid.vo.UsernameAndPassword;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequestMapping("/authority")
public class AuthorityController {

  private final IJWTService ijwtService;


  public AuthorityController(IJWTService ijwtService){
    this.ijwtService=ijwtService;
  }

  /**
   * <h2>从授权中心获取Token （其实就是登陆功能），且返回信息中没有统一响应的包装</h2>
   * @param usernameAndPassword
   * @return
   * @throws Exception
   */
  @IgnoreResponseAdvice
  @PostMapping("/token")
  public JwtToken token(@RequestBody UsernameAndPassword usernameAndPassword)
      throws Exception{
    log.info("request to get token with param: [{}]", JSON.toJSONString(usernameAndPassword));
    return new JwtToken(ijwtService.generateToken(usernameAndPassword.getUsername()
        ,usernameAndPassword.getPassword()));

  }

  /**
   * <h2>注册用户并返回当前注册用户的Token ,即通过授权中心创建用户</h2>
   * @param usernameAndPassword
   * @return
   */
  @IgnoreResponseAdvice
  @PostMapping("/register")
  public JwtToken register(@RequestBody UsernameAndPassword usernameAndPassword)
      throws Exception{
    log.info("register user with param: [{}]",JSON.toJSONString(
        usernameAndPassword
    ));
    return new JwtToken(ijwtService.registerUserAndGenerateToken(usernameAndPassword));
  }
}



