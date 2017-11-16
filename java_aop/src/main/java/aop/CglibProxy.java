package aop;

import jdk.nashorn.internal.runtime.regexp.joni.EncodingHelper;
import org.mockito.cglib.proxy.Enhancer;
import org.mockito.cglib.proxy.MethodInterceptor;
import org.mockito.cglib.proxy.MethodProxy;

import java.lang.reflect.Method;

/**
 * Created by r00xx on 2017/11/16.
 */
public class CglibProxy implements MethodInterceptor {

  private Object target;

  public Object getInstance(Object target) {
    this.target = target;
    Enhancer enhancer = new Enhancer();
    enhancer.setSuperclass(this.target.getClass());
    //回调方法
    enhancer.setCallback(this);
    //创建代理对象
    return enhancer.create();
  }

  @Override
  public Object intercept(Object proxy, Method method, Object[] args, MethodProxy methodProxy) throws Throwable {
    Object result = null;
    System.out.println("事务开始");
    result = methodProxy.invokeSuper(proxy, args);
    System.out.println("事务结束");
    return result;
  }
}




