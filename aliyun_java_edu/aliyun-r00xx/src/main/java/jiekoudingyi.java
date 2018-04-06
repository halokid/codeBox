/**
 * Created by r00x on 2018/4/6.
 */


interface IMessagex {
  public default void fun() {     //追加了普通方法， 有方法体了
    System.out.print("have fun");
  }

  public static IMessagex getInstance() {
    return new MessageImplx();
  }


  public void print();

}

class MessageImplx implements IMessagex {

  @Override
  public void print() {
    System.out.print("hello world");
  }
}


public class jiekoudingyi {
  public static void main(String[] args) {
//        IMessagex msg = new MessageImplx();       //这种写法也可以
    IMessagex msg = IMessagex.getInstance();
    msg.print();
    msg.fun();
  }
}
