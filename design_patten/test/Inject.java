/**
 * 如果在 Class A 中，有 Class B 的实例，则称 Class A 对 Class B 有一个依赖。例如下面类 Human 中用到一个 Father 对象，我们就说类 Human 对类 Father 有一个依赖。
 */

public class Human {

  Father father;

  public  Human() {
    father = new Father();
  }
}


/**
 *
 *
 * 仔细看这段代码我们会发现存在一些问题：
 (1). 如果现在要改变 father 生成方式，如需要用new Father(String name)初始化 father，需要修改 Human 代码；
 (2). 如果想测试不同 Father 对象对 Human 的影响很困难，因为 father 的初始化被写死在了 Human 的构造函数中；
 (3). 如果new Father()过程非常缓慢，单测时我们希望用已经初始化好的 father 对象 Mock 掉这个过程也很困难。
 */




// 依赖注入的写法
public class Human {

  Father father;

  public Human (Father father) {
    this.father = father;
  }
}



// java中的依赖注入，注解是最常见的

public  class Human {

  @Inject Father father;

  public  Human() {

  }
}


/**
 * 上面这段代码看起来很神奇：只是增加了一个注解，Father 对象就能自动注入了？这个注入过程是怎么完成的？

 实质上，如果你只是写了一个 @Inject 注解，Father 并不会被自动注入。你还需要使用一个依赖注入框架，并进行简单的配置。现在 Java 语言中较流行的依赖注入框架有 Google Guice、Spring 等，而在 Android 上比较流行的有 RoboGuice、Dagger 等。其中 Dagger 是我现在正在项目中使用的。如果感兴趣，你可以到 Dagger 实现原理解析 了解更多依赖注入和 Dagger 实现原理相关信息。
 */






//------------------- 清晰的 依赖注入代码例子  ------------------------

public class A {

  private B b;

  public A(B b) {
    this.b = b;
  }

  public void myMethod() {
    b.m();
  }
}




//通常我们这样调用， 所以 A 是 依赖 B 的
A a = new A(new B());
a.myMethod();


// 而依赖注入就是解决这样的 高耦合问题
//而通过依赖注入容器或框架，A对B的依赖无需我们编写代码时赋予，只要我们从一个工厂或容器中获取A的实例，这个工厂或容器是依赖注入框架提供的，它会在里面偷偷地将B的实例注入到A中

 // factory.getA() 方法已自动把依赖的类的注入了，我们不必关心每个类的依赖关系
A a = Factory.getA();
a.myMethod();



class Factory {

  public A getA() {
   A a = new A(new B());
   return a;
  }
}














































































