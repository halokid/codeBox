package org.halokid.controller;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.rocket.RocketMQProducer;
import org.halokid.vo.Message;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * <h1>SpringBoot 集成 RocketMQ</h1>
 * */
@Slf4j
@RestController
@RequestMapping("/rocket-mq")
public class RocketMQController {

  private static final Message RocketMQMessage = new Message(
      1,
      "Halokid-Study-RocketMQ-In-SpringBoot"
  );

  private final RocketMQProducer rocketMQProducer;

  public RocketMQController(RocketMQProducer rocketMQProducer) {
    this.rocketMQProducer = rocketMQProducer;
  }

  @GetMapping("/message-with-value")
  public void sendMessageWithValue() {
    rocketMQProducer.sendMessageWithValue(JSON.toJSONString(RocketMQMessage));
  }

  @GetMapping("/message-with-key")
  public void sendMessageWithKey() {
    rocketMQProducer.sendMessageWithKey("Halokid", JSON.toJSONString(RocketMQMessage));
  }

  @GetMapping("/message-with-tag")
  public void sendMessageWithTag() {
    rocketMQProducer.sendMessageWithTag("halokid",
        JSON.toJSONString(RocketMQMessage));
  }

  @GetMapping("/message-with-all")
  public void sendMessageWithAll() {
    rocketMQProducer.sendMessageWithAll("Halokid", "halokid",
        JSON.toJSONString(RocketMQMessage));
  }
}




