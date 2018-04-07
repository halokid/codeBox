import java.util.concurrent.ExecutionException;
import java.util.concurrent.FutureTask;

/**
 * Created by r00x on 2018/4/7.
 */

/**
 * 这种多线程的处理形式， 主要是为了返回我们的处理结果， 想想多线程还可以返回处理结果，这个是比较酷的
 */

class MyThreadxx implements java.util.concurrent.Callable<String> {

  @Override
  public String call() throws Exception {
    for (int x = 0; x < 20; x++) {
      System.out.println("卖票, x = " + x);
    }
    return "票卖完了， 下次吧···";
  }
}

public class ThreadCallable {
  public static void main(String[] args) throws ExecutionException, InterruptedException {
    FutureTask<String> task = new FutureTask<String>(new MyThreadxx());
    new Thread(task).start();
    System.out.println(task.get());
  }
}
