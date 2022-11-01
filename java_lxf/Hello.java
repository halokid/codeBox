import java.lang.reflect.Array;
import java.util.Arrays;

public class Hello {
  public static void main(String[] args) {
    System.out.println("Hello, world!");
    c1();
    c2();

    // ------------------------------------------
    Outer outer = new Outer("Nested"); // 实例化一个Outer
    Outer.Inner inner = outer.new Inner(); // 实例化一个Inner
    inner.hello();
  }

  private static void c1() {
    char c1 = 'A';
    char c2 = '中';

    int age = 25;
    String s = "aeg is " + age;
    System.out.println(s);

    // ---------------------------------------------
    String x = "hello";
    String t = x;
    x = "world";
    System.out.println(t);
  }

  private static void c2() {
    int[] ns = new int[3];
    ns[0] = 1;
    ns[1] = 2;
    ns[2] = 3;
    System.out.println(ns);
    System.out.println(Arrays.toString(ns));
    System.out.println(ns.length);

    int[][] magicSquare =
      {
        {16, 3, 2, 13},
        {5, 10, 11, 8},
        {9, 6, 7, 3}
      };
    System.out.println(Arrays.toString(magicSquare));
    for (int[] a:magicSquare) {
      System.out.println(Arrays.toString(a));
    }

    // -------------------------------------------
    String[] names = {"ABC", "XYZ", "zoo"};
    String s = names[1];
    names[1] = "cat";
    System.out.println(s);
  }
}

class Outer {
  private String name;

  Outer(String name) {
    this.name = name;
  }

  class Inner {
    void hello() {
      System.out.println("Hello, " + Outer.this.name);
    }
  }

  void asyncHello() {
    Runnable r = new Runnable() {   // todo: 匿名类
      @Override
      public void run() {
        System.out.println("Hello, " + Outer.this.name);
      }
    };
    new Thread(r).start();
  }
}







