public class RPCProxyClient implements java.lang.reflect.InvocationHandler {
  private Object obj;
  public RPCProxyClient(Object obj) {
    this.obj = obj;
  }
  
  //得到被代理对象
  public static Object getProxy(Object obj) {
    return java.lang.reflect.Proxy.newProxyInstance(obj.getClass().getClassLoader(), obj.getClass().getInterfaces(), new RPCProxyClient(obj) );
  }
  
  
  //调用此方法执行
  public Object invoke(Object proxy, Method method, Object[] args) throws Throwable {
    //结果参数
    Object result = new Object();
    //.. 执行通信相关逻辑
    //... 逻辑代码
    return result;
  }
}



/**
下面为调用此远程 RPC 的逻辑代码
**/
public class Test {
  public static void main(String[] args) {
    HelloWorldService helloWorldService = (HelloWorldService)RPCProxyClient.getProxy(HelloWorldService.class);
    helloWorldService.sayHello("test");
  }
}



/**

整个怎么理解呢？？
就是 每一个 tread 都是一个 hashmap 的元素， 而这个 requestID 就是这个哈希map的key，
当thread发起的时候，先线程锁定这个 thread， 服务端收到 requestID 之后， 处理完，
发包回去给客户端，客户端收到回复的 requestID 的时候， 根据这个ID 解锁对应的 thread, 
再唤醒等待的线程， 告诉这些线程，现在已经可以发起连接了。。

**/


//消息里呆 requestID 的原理
public Object get() {     //旋锁
  synchronized(this) {    // 是否有结果
    while (!isDone) {     // 没结果是释放锁， 让当前现成处于等待状态
      wait();
    }
  }
}



private void setDone(Response res) {
  this.res = res;
  isDone = true;
  synchronized (this) {   //获取锁， 因为前面 wait() 已经释放了 callback 的锁了
    notifyAll();
  }
}
















