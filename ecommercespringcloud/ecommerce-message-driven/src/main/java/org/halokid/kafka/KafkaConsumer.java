package org.halokid.kafka;

import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.halokid.vo.Message;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.util.Optional;

/**
 * <h1> kafka consumer</h1>
 */
@Slf4j
@Component
public class KafkaConsumer {

  private final ObjectMapper mapper;

  public KafkaConsumer(ObjectMapper mapper) {
    this.mapper = mapper;
  }

  /**
   * <h2>监听 Kafka 消息并消费</h2>
   */
  @KafkaListener(topics = {"halokid-springboot"}, groupId = "halokid-springboot-kafka-01")
  public void listener01(ConsumerRecord<String, String> record) throws Exception {

    String key = record.key();
    String value = record.value();

    Message kafkaMessage = mapper.readValue(value, Message.class);
    log.info("in listener01 consume kafka message: {}, {}", key, mapper.writeValueAsString(kafkaMessage));
  }

  /**
   * <h2>监听 Kafka 消息并消费</h2>
   * TODO: 假如不知道 ConsumerRecord 的 key 和 value 是什么类型
   */
  @KafkaListener(topics = {"halokid-springboot"}, groupId = "halokid-springboot-kafka-02")
  public void listener02(ConsumerRecord<?, ?> record) throws Exception {

    Optional<?> _kafkaMessage = Optional.ofNullable(record.value());
    if (_kafkaMessage.isPresent()) {    // TODO: 如果kafka的消息存在的话
      Object message = _kafkaMessage.get();
      Message kafkaMessage = mapper.readValue(message.toString(), Message.class);
      log.info("in listener02 consume kafka message: {}", mapper.writeValueAsString(kafkaMessage));
    }
  }
}





