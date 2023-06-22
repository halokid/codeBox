package org.halokid;


public class Main {

  /*
  n -->>> 0 1 2 3 4 5 6 7  8
          0 1 1 2 3 5 8 13 21
   */
  // O(2^n)
  public static int fib1(int n) {
    if (n <= 1) return n;
    return fib1(n - 1) + fib1(n - 2);
  }

  // O(n)
  public static int fib2(int n) {
    if (n <= 1) return n;

    int first = 0;
    int second = 1;
    for (int i = 0; i < n - 1; i++) {
      int sum = first + second;
      first = second;
      second = sum;
    }
    return second;
  }

  public static void main(String[] args) {
//    System.out.println(fib1(46));
//    System.out.println(fib2(46));

    int n = 46;
    TimeTool.check("fib1", new TimeTool.Task() {
      @Override
      public void execute() {
        System.out.println(fib1(n));
      }
    });

    TimeTool.check("fib2", new TimeTool.Task() {
      @Override
      public void execute() {
        System.out.println(fib2(n));
      }
    });
  }
}


