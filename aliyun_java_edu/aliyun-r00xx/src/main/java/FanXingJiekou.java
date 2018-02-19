/**
 * Created by wist on 2018/2/19.
 * 泛型接口
 * 泛型接口的两种形式， 一种是延续泛型的参数， 另外一种是定义好泛型的参数类型
 */

interface IMessage<T> {
  public void print(T t);
}


class MessageImpl<T> implements IMessage<T> {

  public void print(T t) {
    System.out.println(t);
  }
}


class MessageImplTwo implements IMessage<String> {

  public void print(String s) {
    System.out.println(s);
  }
}

public class FanXingJiekou {
  public static void main(String[] args) {
    IMessage<String> msg = new MessageImpl<String>();
    msg.print("hello world");

    IMessage<String> msg2 = new MessageImplTwo();
    msg2.print("xxxxx");
  }
}
