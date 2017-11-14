package aop;

/**
 * Created by wist on 2017/11/14.
 */

/**
 * 代理主题角色，内部包含对真实主题的引用， 并且提供和真实主题角色相同的接口
 */
public class TalkProxy implements ITalk {

  private ITalk talker;

  public TalkProxy(ITalk talker) {
    //super();
    this.talker = talker;
  }

  public void talk(String msg) {
    //引用自己内部的 talker 这个 ITalk 类的实例，然后执行自己的接口方法
    talker.talk(msg);
  }

  public void talk(String msg, String singname) {
    talker.talk(msg);
    sing(singname);
  }

  private void sing(String singname) {
    System.out.println("唱歌: " + singname);
  }
}
