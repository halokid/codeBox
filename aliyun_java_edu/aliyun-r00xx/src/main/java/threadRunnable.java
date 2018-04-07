/**
 * Created by r00x on 2018/4/7.
 */


class MyThreadx implements Runnable {
  private int ticket = 10;

  @Override
  public void run() {
    for (int x = 0; x < 10; x++) {
      if (this.ticket > 0) {
        System.out.println("卖票, ticket = " + this.ticket -- );
      }
    }
  }
}


public class threadRunnable {
  public static  void main(String[] args) {
    MyThreadx mt =  new MyThreadx();
    new Thread(mt).start();
    new Thread(mt).start();
    new Thread(mt).start();
  }
}
