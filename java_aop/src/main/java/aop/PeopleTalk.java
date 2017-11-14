package aop;

/**
 * Created by wist on 2017/11/14.
 */

/**
 * 真实主题角色， 定义真实的对象
 */
public class PeopleTalk implements ITalk {
  public String username;
  public String age;

  public PeopleTalk(String username, String age) {
    this.username = username;
    this.age = age;
  }

  public void  talk(String msg) {
    System.out.println(msg + "你好， 我是 " + username + ", 我年龄是 " + age);
  }

  public String getName() {
    return username;
  }

  public void setName(String name) {
    this.username = name;
  }

  public String getAge() {
    return age;
  }

  public void setAge(String age) {
    this.age = age;
  }
}
