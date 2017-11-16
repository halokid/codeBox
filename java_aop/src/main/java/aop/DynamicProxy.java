package aop;

/**
 * Created by r00xx on 2017/11/15.
 */

import java.lang.reflect.InvocationHandler;
import java.lang.reflect.Method;
import java.lang.reflect.Proxy;

/**
 * 动态代理
 */
public class DynamicProxy implements InvocationHandler {

  private Object target;

  public Object bind(Object target) {
    this.target = target;
    return Proxy.newProxyInstance(target.getClass().getClassLoader(),
                                  target.getClass().getInterfaces(), this);
  }


  @Override
  public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
    Object result = null;

    System.out.println("切面之前执行");

    result = method.invoke(target, args);

    System.out.println("切面之后执行");

    return result;
  }

}





