package org.halokid.advice;

import lombok.extern.slf4j.Slf4j;
import org.halokid.vo.CommonResponse;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import javax.servlet.http.HttpServletRequest;

/*
全局异常捕获处理
 */
@Slf4j
@RestControllerAdvice
public class GlobalExceptionAdvice {

  @ExceptionHandler(value = Exception.class)
  public CommonResponse<String> handlerCommerceException(HttpServletRequest req, Exception ex) {
    CommonResponse<String> response = new CommonResponse<>(-1, "business error");
    response.setData(ex.getMessage());
    log.error("commerce service has error: {}", ex.getMessage(), ex);
    return response;
  }
}


