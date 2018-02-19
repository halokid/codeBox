/**
 * Created by wist on 2018/2/19.
 * 多例范例， 但是现在已经可以优化成枚举的做法了
 * 实际上 枚举就是一种高级的多例设计？？？？
 */

/**
 * 多例的设计
 */
class Color {

  private static final Color RED = new Color("RED");
  private static final Color GREEN = new Color("GREEN");
  private static final Color BLUE = new Color("BLUE");

  private String title;

  private Color(String title) {
    this.title = title;
  }

  public static Color getInstance(int ch) {
    switch (ch) {
      case 0:
        return RED;
      case 1:
        return GREEN;
      case 2:
        return BLUE;
      default:
        return null;
    }
  }

  public String toString() {
    return this.title;
  }
}


/**
 * 枚举的设计
 */
enum Color2 {
  RED, GREEN, BLUE
}



public class DuoLi {
  public static void main(String[] args) {
    System.out.println(Color.getInstance(0));   //多例
    System.out.println(Color2.RED);             //枚举
  }
}





