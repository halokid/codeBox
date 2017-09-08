/**
 * Created by r00xx on 2017/9/8.
 */


/**
 *  Test Case 类， 测试用例类
 */

import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class TestJunit {
  String message = "hey junit";

  //新建一个这个类的类型
  MessageUtil messageUtil = new MessageUtil(message);


  @Test
  public void testPrintMessage() {
    message = "new junit";

    assertEquals( message, messageUtil.printMessage() );
  }
}
