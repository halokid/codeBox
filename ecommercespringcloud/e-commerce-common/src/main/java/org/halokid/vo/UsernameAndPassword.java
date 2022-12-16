package org.halokid.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

/*
用户名和密码
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class UsernameAndPassword {

  private String username;

  private String password;


}
