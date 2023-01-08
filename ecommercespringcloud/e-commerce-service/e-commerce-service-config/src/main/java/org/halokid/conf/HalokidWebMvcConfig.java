package org.halokid.conf;

import org.halokid.filter.LoginUserInfoInterceptor;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurationSupport;

/**
 * <h1>Web Mvc 配置</h1>
 * 配置拦截器不会生效 ，要做到mvc  配置里面
 */
@Configuration    // TODO: 这个注解就是在spring运行期间，让这个class自动配置进去整个运行时的作用
public class HalokidWebMvcConfig extends WebMvcConfigurationSupport {

  /**
   * <h2> 添加拦截器配置</h2>
   * @param registry
   */
  @Override
  protected void addInterceptors(InterceptorRegistry registry) {
    // 添加用户身份统一登录拦截的拦截器
    registry.addInterceptor(new LoginUserInfoInterceptor())
        .addPathPatterns("/**").order(0);
  }

  /**
   * <h2>让MVC 加载 Swagger 的静态资源</h2>
   * @param registry
   */
  @Override
  protected void addResourceHandlers(ResourceHandlerRegistry registry) {

//        registry.addResourceHandler("/**")
//                .addResourceLocations("classpath:/static/");
//        registry.addResourceHandler("swagger-ui.html")
//                .addResourceLocations("classpath:/META-INF/resource");
//        registry.addResourceHandler("doc.html")
//                .addResourceLocations("classpath:/META-INF/resource");
//        registry.addResourceHandler("/webjars/**")
//                .addResourceLocations("classpath:/META-INF/resource/webjars");

    registry.addResourceHandler("/**").
        addResourceLocations("classpath:/static/");
    registry.addResourceHandler("swagger-ui.html")
        .addResourceLocations("classpath:/META-INF/resources/");
    registry.addResourceHandler("doc.html")
        .addResourceLocations("classpath:/META-INF/resources/");
    registry.addResourceHandler("/webjars/**")
        .addResourceLocations("classpath:/META-INF/resources/webjars/");

    super.addResourceHandlers(registry);
  }
}



