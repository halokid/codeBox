package org.halokid.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/**
 * <h1> 登录用户信息</h1>
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class LoginUserInfo {

  /**  用户id */
  private Long id;

  /** 用户名 */
  private String username;
}