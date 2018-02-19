import com.sun.xml.internal.ws.message.stream.StreamAttachment;

/**
 * Created by wist on 2018/2/19.
 * 范型，占位符， 避免了向下转型的错误
 */

class PointPlus <T> {
  private T x;
  private String y;

  public T getX() {
    return x;
  }

  public void setX(T x) {
    this.x = x;
  }

  public void setY(String y) {
    this.y = y;
  }

  public String getY() {
    return y;
  }
}


public class FanXingPlus {
  public static void main(String[] args) {
    PointPlus<String> p = new PointPlus<String>();
    p.setX("你好");
    p.setY("世界");

    String x = p.getX();
    String y = p.getY();

    System.out.println(x);
    System.out.println(y);
  }
}



