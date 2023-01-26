package org.halokid.controller;

import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.halokid.kafka.KafkaProducer;
import org.halokid.vo.Message;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Slf4j
@RestController
@RequestMapping("/kafka")
public class KafkaController {

  private final ObjectMapper mapper;
  private final KafkaProducer kafkaProducer;

  public KafkaController(ObjectMapper mapper, KafkaTemplate kafkaTemplate, KafkaProducer kafkaProducer) {
    this.mapper = mapper;
    this.kafkaProducer = kafkaProducer;
  }

  /**
   * <h2>发送 kafka 消息</h2>
   */
  @GetMapping("/send-message")
  public void sendMessage(@RequestParam(required = false) String key,
                          @RequestParam String topic) throws Exception {
    Message message = new Message(
       1,
       "halokid-study-ecommerece"
    );
    kafkaProducer.sendMessage(key, mapper.writeValueAsString(message), topic);
  }
}

