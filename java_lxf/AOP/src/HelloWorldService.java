
// todo:
// 感觉java的 IOC, AOP编程其实就是 装饰器， 把一些常用的组件， 流程初始化， 连接句柄等常用
// 的操作wrap起来, 减少代码的耦合， 重复运行， 交叉生命周期（这个容易引起bug）等情况


public class HelloWorldService implements IHelloWorldService {

  @Override
  public void sayHello() {
    System.out.println("hello! Spring AOP sample...");
  }
}


