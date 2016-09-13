public class HelloWorldServiceImpl implements HelloWorldService {
  @Override
  public String sayHello(String msg) {
    String result = "hello world" + msg;
    System.out.println(result);
    return result;
  }
}


