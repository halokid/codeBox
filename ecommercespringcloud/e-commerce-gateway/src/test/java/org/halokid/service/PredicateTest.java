package org.halokid.service;

import lombok.extern.slf4j.Slf4j;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.Arrays;
import java.util.List;
import java.util.function.Predicate;

/**
 * java8 predicate 使用方法与思想
 */
@Slf4j
@SpringBootTest
@RunWith(SpringRunner.class)
public class PredicateTest {

  public static List<String> MICRO_SERVICE = Arrays.asList(
      "nacos","authority","gateway","ribbon","feign","hystrix",
      "e-commerce"
  );

  /**
   * <h2>test 方法主要用与参数符不符合规则，返回值是 boolean</h2>
   */
  @Test
  public void testPredicateTest(){
    Predicate<String> letterLengthLimit = s->s.length()>5;
    MICRO_SERVICE.stream().filter(letterLengthLimit).forEach(
        System.out::println);
  }

  /**
   * <h2> and 方法 等同于逻辑与 &&，存在短路特性，需要所有的条件都满足</h2>
   */
  @Test
  public void testPredicateAnd(){
    Predicate<String> letterLengthLimit = s->s.length()>5;
    Predicate<String> letterStartWith = s->s.startsWith("gate");

    MICRO_SERVICE.stream().filter(letterLengthLimit.and(letterStartWith)).forEach(
        System.out::println
    );
  }

  /**
   * <h2>or 方法  等同于逻辑或 || ,多个条件只要满足一个即可</h2>
   */
  @Test
  public void testPredicateOr(){
    Predicate<String> letterLengthLimit = s->s.length()>5;
    Predicate<String> letterStartWith = s->s.startsWith("gate");

    MICRO_SERVICE.stream().filter(letterLengthLimit.or(letterStartWith)).forEach(
        System.out::println
    );
  }

  /**
   * <h2>negate 等同于我们的逻辑非 ！</h2>
   */
  @Test
  public void testPredicateNegate(){
    Predicate<String> letterStartWith = s->s.startsWith("gate");

    MICRO_SERVICE.stream().filter(letterStartWith.negate()).forEach(
        System.out::println
    );
  }

  /**
   *  <h2>isEqual 类似于 equals(),区别在与：先判断对象是否为NULL,
   *  不为null 再使用equals 进行比较</h2>
   */
  @Test
  public void testPredicateIsEqual(){
    Predicate<String> equalGateway = s->Predicate.isEqual("gateway").test(s);
    MICRO_SERVICE.stream().filter(equalGateway).forEach(
        System.out::println
    );
  }
}









