/**
 * Created by r00xx on 2017/9/8.
 */


/**
 * 实际的逻辑代码类
 */

public class MessageUtil {

  private String message;

  //constructor
  public MessageUtil(String message) {
    this.message = message;
  }

  // prints the message
  public String printMessage() {
    System.out.println(message);
    return message;
  }
}
