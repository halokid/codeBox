/**
 * Created by r00x on 2018/4/7.
 */


class DataProvider implements Runnable {
  private Data data;
  public DataProvider (Data data) {
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



class DataConsumer implements Runnable {
  private Data data;
  public DataConsumer(Data data) {
    this.data = data;
  }

  @Override
  public void run() {
    for (int x = 0; x < 50; x++) {
      this.data.get();
    }
  }
}



class Data {      //负责保存数据
  private String title;
  private String note;
  // flag = true:  表示还没消费完呢， 不允许生产， 要等待消费完才能允许生产， 避免重复生产， 消费完之后  flag 的值为 false
  // flag = false: 表示还没生产完呢， 不允许消费， 要等待生产完才能允许消费， 生产完之后 flag的值为 true
  private boolean flag = false;    //初始化为 false， 表示不能get， 因为一开始还没有set的时候， 消费者取得的值是为 null的， 要等set完之后， flag才是 true

  public synchronized void get() {
    // 如果这个时候是不允许消费者取走的话， 那么就等待到 允许消费者取走
    if (flag == false) {   // 还没有生产呢， 所以要 wait， 等待生产完才能消费
      try {
        super.wait();   // 如果 flag 为flase 的时候， hold住逻辑， 不能取值， 因为还没有生产完呢
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
    }
    try {
      Thread.sleep(50);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
    System.out.println(this.title + "=" + this.note);
//    this.flag = true;    //表示等待完消费者取走之后， 允许生产了
    this.flag = false;    //消费完了， 重新设置 flag 的值为 false， 表示不能再消费了啊， 必须要等待生产（也就是flag的值为 true）才可以再消费啊
    super.notify();     //唤醒等待生产的线程， 让生产线程去生产
  }


  public synchronized void set(String title, String note) {
    if (this.flag == true) {     //还没消费完呢， 不允许重复生产哦， 要等生一个消费完（flag值为false），才可以生产哦
      try {
        super.wait();       // hold 住逻辑， 等待吧
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
    }
    this.title = title;
    try {
      Thread.sleep(100);
    } catch (InterruptedException e) {
      e.printStackTrace();
    }
    this.note = note;
//    this.flag = false;         //生产完成之后， 设置flag的值表示生产完毕， 允许取走了
    this.flag = true;         // 生产完了哦， 设置 flag 的值为true， 告诉程序不允许再生产了， 需要消费完（设置flag值为flase）才可以再生产
    super.notify();           //唤醒 等待消费 的线程
  }



}


public class ProviderConsumer {
  public static void main(String[] args) {
    Data data = new Data();
    new Thread(new DataProvider(data)).start();
    new Thread(new DataConsumer(data)).start();
  }
}



//fixme: 面试题 sleep() 和  wait() 的区别, sleep sh8  tread 定义的方法， 到了一定的时候，线程会自动唤醒，wait 是 object的方法， 如果要唤醒， 必须调用 notify， notifyall等方法才能唤醒
