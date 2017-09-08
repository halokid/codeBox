/**
 * Created by r00xx on 2017/9/8.
 */

/**
 *   Test Case， 测试用例的运行类
 */


import org.junit.runner.JUnitCore;
import org.junit.runner.Result;
import org.junit.runner.notification.Failure;

public class TestRunner {

  // 主执行方法
  public static void main(String[] args) {
    Result result = JUnitCore.runClasses(TestJunit.class);

    // 循环输出测试的错误信息
    for (Failure failure : result.getFailures()) {
      System.out.println(failure.toString());
    }

    System.out.println(result.wasSuccessful());
  }
}
