package hello;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.concurrent.atomic.AtomicLong;

/**
 * Created by r00xx<82049406@qq.com> on 2016/12/28.
 */

@RestController
public class GreetController {
  private static final String template = "Hello, %s!";
  private final AtomicLong counter = new AtomicLong();

  @RequestMapping("/greeting")
  public Greeting greeting(@RequestParam(value="name", defaultValue = "World") String name) {
    return  new Greeting(counter.incrementAndGet(), String.format(template, name));
  }

}



