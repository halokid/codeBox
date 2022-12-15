package org.halokid.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.io.Serializable;

/*
通用响应对象定义
example:
{
  "code": 0,
  "message":  "",
  "data":  {}
}
 */

@Data
@NoArgsConstructor
@AllArgsConstructor
public class CommonResponse<T> implements Serializable {

  // 错误码
  private Integer code;

  // 错误消息
  private String message;

  // 泛型返回数据
  private T Data;

  public CommonResponse(Integer code, String message) {
    this.code = code;
    this.message = message;
  }
}















