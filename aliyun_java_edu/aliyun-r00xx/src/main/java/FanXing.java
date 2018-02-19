/**
 * Created by wist on 2018/2/19.
 * 范型， 向下兼容的时候是不安全的， 因为输入的接收者不知道object具体是什么类型
 * 平常写程序的时候， 尽量避免向下转型
 */
class Point {
  private Object x;
  private Object y;
  public Object getX() {
    return x;
  }

  public void setX(Object x) {
    this.x = x;
  }

  public Object getY() {
    return y;
  }

  public void setY(Object y) {
    this.y = y;
  }
}


public class FanXing {
  public static  void main(String[] args) {
    Point p = new Point();

    /**
    p.setX(10);
    p.setY(20);

    int x = (Integer) p.getX();
    int y = (Integer) p.getY();

    System.out.println(x);
    System.out.println(y);
    **/

    p.setX("你好");
    p.setY("世界");

    String x = (String) p.getX();
    System.out.println(x);

    String y = (String) p.getY();
    System.out.println(y);

  }
}

