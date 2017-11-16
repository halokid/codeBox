package aop;

/**
 * Created by wist on 2017/11/14.
 */

/**
 * 代理主题角色，内部包含对真实主题的引用， 并且提供和真实主题角色相同的接口
 *
 * 代理类的要点就是， 每一个业务类都要有一个属于自己的代理类
 */
public class TalkProxy implements ITalk {

  private ITalk talker;

  public TalkProxy(ITalk talker) {
    //super();
    this.talker = talker;
  }

  public void talk(String msg) {
    //引用自己内部的 talker 这个 ITalk 类的实例，然后执行自己的接口方法
    System.out.println("这个是代理类自己重写的方法， 跑玩这个方法，假如代理的是 PeopleTalk 类的实例，其实还是会跑people.talk 的");
    talker.talk(msg);
  }

  /**
   * 代理类自己添加的一个 切面， 就是凡是经过代理类执行的业务，都要经过这个切面的逻辑处理
   * @param msg
   * @param singname
   */
  public void talk(String msg, String singname) {
    talker.talk(msg);
    sing(singname);
  }

  private void sing(String singname) {
    System.out.println("唱歌: " + singname);
  }
}
