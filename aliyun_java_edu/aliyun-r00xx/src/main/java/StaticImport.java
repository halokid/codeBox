//import util.MyMath;
import util.MyMath;
import static util.MyMath.*;

/**
 * Created by wist on 2018/2/19.
 * 静态导入
 */
public class StaticImport {
  public static void main(String[] args) {
    System.out.println(MyMath.add(5, 6));
    System.out.println(MyMath.sub(5, 6));

    System.out.println(add(7, 8));
  }
}
