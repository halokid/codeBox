package org.halokid.service;

import org.halokid.vo.UsernameAndPassword;

/*
JWT 相关服务接口定义
 */
public interface IJWTService {

  /*
  生成JWT Toekn 使用默认的超时时间
   */
  String generateToken(String username, String password) throws Exception;

  String generateToken(String username, String password, int expire) throws Exception;

  String registerUserAndGenerateToken(UsernameAndPassword usernameAndPassword);
}
