/**
 * Created by r00x on 2018/4/7.
 */


import java.sql.Time;

/**
 * 线程锁， 锁住执行逻辑
 */


class MyThreadLock implements Runnable {
  private int ticket = 10;

  /**
   *
   */

  /**
   * // 如果程序像下面的逻辑， 那是什么都不会输出的， 为什么呢？？
   * // 因为假如是这个  synchronized 的逻辑 是走到了  for 循环 20 的逻辑上面去， 但是这个逻辑是要比多线程外面的逻辑是要慢的
   * // 所以程序就会执行外面的逻辑， 而没有等待 多线程 里面的这个逻辑跑完
   * // 上面的都是屁话
   * // synchronized 锁住了 for 循环里面的逻辑, 所以下面这个逻辑肯定是 等某一个 票贩子全部卖完之后， 才会处理其他票贩子的逻辑
   *
   * @Override public void run() {
   * synchronized (this) {     //表示程序逻辑上锁
   * //      System.out.println("根本就不会触发这个流程");
   * for (int x = 0; x < 20; x++) {
   * //        System.out.println("根本就不会触发这个流程 xxxxxxxxx ");
   * if (this.ticket > 0) {
   * try {
   * Thread.sleep(200);      //模拟网络延迟
   * } catch (InterruptedException e) {
   * e.printStackTrace();
   * }
   * System.out.println(Thread.currentThread().getName() + "卖票, ticket = " + this.ticket --);
   * }
   * }
   * }
   * //    try {
   * //      Thread.sleep(20000);
   * //    } catch (InterruptedException e) {
   * //      e.printStackTrace();
   * //    }
   * System.out.println("end here");
   * }
   **/

  /**
   * 下面这个程序是 线程安全的， 虽然可以保证变量的统一性， ticket， 但是速度会比 非线程安全的灰慢点， 在高并发的时候，这个性能
   * 差距是比较明显的， 所以为了处理高并发的系统， 如果可以不用线程安全， 就可以不用哈， 如果影响了程序的合理逻辑， 那么就应该用
   * 线程安全
   * 这种情况在 高并发的情况下~~ 也可能会造成死锁
   */
  @Override
  public void run() {
    for (int x = 0; x < 20; x++) {
      synchronized (this) {     //表示程序逻辑上锁
//      System.out.println("根本就不会触发这个流程");
//        System.out.println("根本就不会触发这个流程 xxxxxxxxx ");
        if (this.ticket > 0) {
          try {
            Thread.sleep(200);      //模拟网络延迟
          } catch (InterruptedException e) {
            e.printStackTrace();
          }
          System.out.println(Thread.currentThread().getName() + "卖票, ticket = " + this.ticket--);
        }
      }
    }
//    try {
//      Thread.sleep(20000);
//    } catch (InterruptedException e) {
//      e.printStackTrace();
//    }
    System.out.println("end here");
  }


}

public class xianchengsuo {
  public static void main(String[] args) {
    MyThreadLock mt = new MyThreadLock();
    Thread t1 = new Thread(mt, "票贩子A");
    Thread t2 = new Thread(mt, "票贩子B");
    Thread t3 = new Thread(mt, "票贩子C");

    t1.start();
    t2.start();
    t3.start();

  }
}






