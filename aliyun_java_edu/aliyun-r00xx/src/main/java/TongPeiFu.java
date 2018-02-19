/**
 * Created by wist on 2018/2/19.
 * 通配符
 *  ? extends, 设置泛型上限, 比如 ? extends int, 表示只能设置 integer 类型或其子类
 *  ? super, 设置泛型下限, 比如 ? super string, 表示只能设置 string 类型或其父类
 */

class Message<T> {
  private T note;
  public void setNote(T note) {
    this.note = note;
  }

  public T getNote() {
    return note;
  }
}

public class TongPeiFu {
  public static void main(String[] args) {
    Message<String>  msg = new Message<String>();
//    msg.setNote("99");
    msg.setNote("99");
    fun(msg);
  }

  public static void fun(Message<? super String> temp) {
    System.out.println(temp.getNote());
  }
}
