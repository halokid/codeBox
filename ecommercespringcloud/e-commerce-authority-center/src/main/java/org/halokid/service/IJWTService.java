package org.halokid.service;

import org.halokid.vo.UsernameAndPassword;

/*
JWT 相关服务接口定义
 */
public interface IJWTService {

  /**
   * <h2>生成 JWT Token,使用默认的超时时间</h2>
   * @param username
   * @param password
   * @return
   */
  String generateToken(String username,String password) throws Exception;

  /**
   * <h2>生成指定超时时间Token 单位是天</h2>
   * @param username
   * @param password
   * @param expire
   * @return
   * @throws Exception
   */
  String generateToken(String username,String password,int expire) throws Exception;


  /**
   * <h2>注册用户 同时生成Token返回</h2>
   * @param usernameAndPassword
   * @return
   * @throws Exception
   */
  String registerUserAndGenerateToken(UsernameAndPassword usernameAndPassword) throws Exception;
}
