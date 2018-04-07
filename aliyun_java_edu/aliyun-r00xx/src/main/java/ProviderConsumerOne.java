/**
 * Created by r00x on 2018/4/7.
 */

/**
 *  这个程序虽然可以解决数据的 锁住 设置 和 获取， 但是还是不能解决我们 消费者 交替取值的问题， 这个也是会重复的取值， 重复的设值
 *  两个线程不可以交替执行， 我们的目标是要两个线程可以交替执行
 */

class DataProviderOne implements Runnable {
  private DataOne data;
  public DataProviderOne (DataOne data) {
    this.data = data;
  }

  @Override
  public void run() {
    for (int x = 0; x < 50; x++) {
      if (x % 2 == 0) {
        this.data.set("老李", "是个好人");
      }
      else {
        this.data.set("民族败类", "老方B");
      }
    }
  }
}



class DataConsumerOne implements Runnable {
  private DataOne data;
  public DataConsumerOne(DataOne data) {
    this.data = data;
  }

  @Override
  public void run() {
    for (int x = 0; x < 50; x++) {
      this.data.get();
    }
  }
}



class DataOne {      //负责保存数据
  private String title;
  private String note;

  public synchronized void set(String title, String note) {
    this.title = title;
    try {
      Thread.sleep(100);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
    this.note = note;
  }

  public synchronized void get() {
    try {
      Thread.sleep(50);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
    System.out.println(this.title + "=" + this.note);
  }

}


public class ProviderConsumerOne {
  public static void main(String[] args) {
    DataOne data = new DataOne();
    new Thread(new DataProviderOne(data)).start();
    new Thread(new DataConsumerOne(data)).start();
  }
}
