/**
 * Created by wist on 2018/2/19.
 * 泛型方法
 */
public class FanXingFangfa {

  public static void main(String[] args) {
    Integer data[] = fun(1, 2, 3, 4);
    for (int i : data) {
      System.out.println(i);
    }
  }


  public static  <T> T[] fun(T ... args) {
    return args;
  }

}
