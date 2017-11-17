package aop;

/**
 * Created by r00xx on 2017/11/16.
 */
public class TestDynamic {
  public static void main(String[] args) {
    ITalk iTalk = (ITalk) new DynamicProxy().bind(new PeopleTalk("jxx", "20"));
    iTalk.talk("业务说明");
  }
}
