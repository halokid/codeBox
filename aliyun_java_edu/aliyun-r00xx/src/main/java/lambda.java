/**
 * Created by r00x on 2018/4/7.
 */

@FunctionalInterface
//如果要调用 lambda 的话， 则该接口类必须只有一个方法， 不然就会报错
interface IMessagey {
  public void print();      //这是一个接口， 接口类方法必须有子类
//  public void print2();
}


public class lambda {
  public static void main (String[] args) {
    //匿名函数式编程
//    IMessagey msg = () -> System.out.print("hello world");  //这一句的 ()  实际上就等同接口类  print 方法的 ()
    //下面是多行的情况
    IMessagey msg = () -> {
      System.out.print("hello world");
      System.out.print("hello world");
      System.out.print("hello world");
    };
    msg.print();
  }
}
