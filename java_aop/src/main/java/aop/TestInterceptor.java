package aop;

/**
 * Created by r00xx on 2017/11/16.
 */
public class TestInterceptor {

  public static void main(String[] args) {
    PeopleTalkCglib peopleTalk = (PeopleTalkCglib) new CglibProxy().getInstance(new PeopleTalkCglib());

    peopleTalk.talk("业务方法");
//    peopleTalk.spreak("业务方法");
  }
}
