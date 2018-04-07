/**
 * Created by r00x on 2018/4/7.
 */

/**
 * 取得线程名称
 */
class MyThreadzz implements Runnable {

  @Override
  public void run() {
    for (int x = 0; x < 10; x++) {
      System.out.println(Thread.currentThread().getName() + ", x = " + x);
    }
  }
}


public class xianchengname {
  public static void main(String[] args) {
    MyThreadzz mt = new MyThreadzz();
    new Thread(mt).start();
    new Thread(mt).start();
    new Thread(mt, "有名线程").start();
  }
}
