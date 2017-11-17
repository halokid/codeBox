package aop;

/**
 * Created by wist on 2017/11/14.
 */

/**
 * 抽象主题角色， 声明了真实主题和代理主题的共同接口
 * 这个是属于业务类
 */
public interface ITalk {
  public void talk(String msg);
}
