/**
 * Created by r00x on 2018/4/7.
 */


class MyThread extends Thread {     // 是一个线程的主类
  private String title;
  public MyThread(String title) {
    this.title = title;
  }

  @Override
  public void run() {       //所有的线程从此开发执行
    for (int x = 0;  x < 10; x++) {
      System.out.println(this.title + ", x = " + x);
    }
  }
}


public class duoxiancheng {
  public static void main(String args[]) {
    MyThread m1 = new MyThread("thread A");
    MyThread m2 = new MyThread("thread B");
    MyThread m3 = new MyThread("thread C");

//    m1.run();
//    m2.run();
//    m3.run();

    m1.start();
//    m1.start();
    m2.start();
    m3.start();
  }
}
