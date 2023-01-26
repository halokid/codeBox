package org.halokid.advice;

import org.halokid.annotation.IgnoreResponseAdvice;
import org.halokid.vo.CommonResponse;
import org.springframework.core.MethodParameter;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.servlet.mvc.method.annotation.ResponseBodyAdvice;

/*
实现统一响应
 */
// TODO: 这里只要这样设置了这个注解（RestControllerAdvice), 就会自动配置所有的 org.halokid 包下的restController,
// TODO: 所有的restful  api 都会走下面的wrap方法去进行逻辑处理
@RestControllerAdvice(value = "org.halokid")
public class CommonResponseDataAdvice implements ResponseBodyAdvice<Object> {
  // 判断是否需要对响应进行处理
  @Override
  public boolean supports(MethodParameter methodParameter, Class<? extends HttpMessageConverter<?>> aClass) {
    if (methodParameter.getDeclaringClass().isAnnotationPresent(IgnoreResponseAdvice.class)) {
      return false;
    }
    if (methodParameter.getMethod().isAnnotationPresent(IgnoreResponseAdvice.class)) {
      return false;
    }
    return true;
  }

  //
  @Override
  public Object beforeBodyWrite(Object o, MethodParameter methodParameter, MediaType mediaType, Class<? extends HttpMessageConverter<?>> aClass, ServerHttpRequest serverHttpRequest, ServerHttpResponse serverHttpResponse) {

    // 定义最终的返回对象
    CommonResponse<Object> response = new CommonResponse<>(0, "");

    if (null == o) {
      return response;
    } else if (o instanceof CommonResponse) {
      response = (CommonResponse<Object>) o;
    } else {
      response.setData(o);
    }

    return response;
  }
}



