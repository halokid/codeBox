/**
 * Created by r00x on 2018/4/7.
 */


class DataProviderxx implements Runnable {
  private Dataxx data;
  public DataProviderxx (Dataxx data) {
    this.data = data;
  }

  @Override
  public void run() {
    for (int x = 0; x < 50; x++) {
      if (x % 2 == 0) {
        this.data.setTitle("老李");
        try {
          Thread.sleep(1000);
        } catch (InterruptedException e) {
          e.printStackTrace();
        }
        this.data.setNote("是个好人");
      }
      else {
        this.data.setTitle("民族败类");
        try {
          Thread.sleep(1000);
        } catch (InterruptedException e) {
          e.printStackTrace();
        }
        this.data.setNote("老方B");
      }
    }
  }
}


class DataConsumerxx implements Runnable {
  private Dataxx data;
  public DataConsumerxx(Dataxx data) {
    this.data = data;
  }

  @Override
  public void run() {
    for (int x = 0; x < 50; x++) {
      try {
        Thread.sleep(1000);
      } catch (InterruptedException e) {
        e.printStackTrace();
      }
      System.out.println(this.data.getTitle() + ":" + this.data.getNote());
    }
  }
}



class Dataxx {      //负责保存数据
  private String title;
  private String note;

  public void setNote(String note) {
    this.note = note;
  }

  public void setTitle(String title) {
    this.title = title;
  }

  public String getNote() {
    return note;
  }

  public String getTitle() {
    return title;
  }
}


public class ProviderConsumerWrong {
  public static void main(String[] args) {
    Data data = new Data();
    new Thread(new DataProvider(data)).start();
    new Thread(new DataConsumer(data)).start();
  }
}
