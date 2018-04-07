/**
 * Created by r00x on 2018/4/7.
 */

/**
 * 并发执行多线程卖票的程序
 */

class MyThreadpp implements Runnable {
  private int ticket = 10;      //这里是一共要卖出的票的总数
  @Override
  public void run() {
    for (int x = 0; x < 20; x++) {
      if (this.ticket > 0) {
        try {
          Thread.sleep(200);      //模拟网络延迟， 此程序会出现票的超卖问题
        } catch (InterruptedException e) {
          e.printStackTrace();
        }
        System.out.println(Thread.currentThread().getName() + "卖票, ticket = " + this.ticket--);
      }
    }
  }
}


public class maipiao {
  public static void main(String[] args) {
    MyThreadpp mt = new MyThreadpp();
    new Thread(mt, "票贩子A").start();
    new Thread(mt, "票贩子B").start();
    new Thread(mt, "票贩子C").start();
  }
}
