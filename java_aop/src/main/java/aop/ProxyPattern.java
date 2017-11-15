package aop;

/**
 * Created by r00xx on 2017/11/15.
 */

/**
 * 代理测试类， 使用代理
 */
public class ProxyPattern {
  public static void main(String[] args) {
    //不需要执行额外方法的
    ITalk people = new PeopleTalk("AOP", "18");
    people.talk("no proxy test");
    System.out.println("----------------------------------");

    //需要执行额外的方法的（切面）
    TalkProxy talker = new TalkProxy(people);
    talker.talk("proxy test", "代理");
  }
}
