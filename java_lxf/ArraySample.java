import java.util.Arrays;

public class ArraySample {
  public static void main(String[] args) {
    int[] ns = {1, 4, 9, 16, 25};
    for (int i = 0; i < ns.length; i++) {
      int n = ns[i];
      System.out.println(n);
    }

    System.out.println("------------------------------------");
    for (int n : ns) {
      System.out.println(n);
    }

    System.out.println("------------------------------------");
    System.out.println(Arrays.toString(ns));

    System.out.println("------------------------------------");
    bubbingSort();

    c1();

    // ----------------------------------------------
    for (String arg: args) {
      System.out.println(arg);
      if ("-version".equals(arg)) {
        System.out.println("v 1.0");
        break;
      }
    }
  }

  // todo: 冒泡排序的特点是，每一轮循环后，最大的一个数被交换到末尾，因此，下一轮循环就可以“刨除”最后的数，每一轮循环都比上一轮循环的结束位置靠前一位。
  private static void bubbingSort() {
    int[] ns = { 28, 12, 89, 73, 65, 18, 96, 50, 8, 36 };
    System.out.printf("before sort -->>> %s\n", Arrays.toString(ns));

    // todo: from small to large
    for (int i = 0; i < ns.length - 1; i++) {
      for (int j = 0; j < ns.length - 1 - i; j++) {
        if (ns[j] > ns[j + 1]) {
          int tmp = ns[j];
          ns[j] = ns[j + 1];
          ns[j + 1] = tmp;
        }
      }
    }

    System.out.printf("after sort -->>> %s\n", Arrays.toString(ns));

    // todo: build-in sort function
    int[] ns2 = { 28, 12, 89, 73, 65, 18, 96, 50, 8, 36 };
    Arrays.sort(ns2);
    System.out.printf("build-in sort -->>> %s\n", Arrays.toString(ns2));
  }

  private static void c1() {
    int[][] ns = {
      {1, 2, 3, 4},
      {5, 6, 7, 8}
    };
    System.out.println(ns.length);
  }

}







