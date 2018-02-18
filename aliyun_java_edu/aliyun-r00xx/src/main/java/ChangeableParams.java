/**
 * Created by wist on 2018/2/18.
 */
public class ChangeableParams {
  public static int add(int ... data) {
    int sum = 0;
//    for (int x = 0; x < data.length; x++) {
    for (int aData : data) {
      sum += aData;
    }
    return sum;
  }

  public static void main(String[] args) {
    System.out.println("开始");
    System.out.println(add());
    System.out.println(add(1, 2, 3, 4, 5));
    System.out.println(add(new int[] {1, 2, 3, 4, 5}));
    System.out.println(add(new int[] {1, 2, 3, 4, 5, 6, 7, 8}));
  }
}


