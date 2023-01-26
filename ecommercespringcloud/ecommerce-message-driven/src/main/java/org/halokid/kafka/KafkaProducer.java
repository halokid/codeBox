package org.halokid.kafka;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.stereotype.Component;
import org.springframework.util.concurrent.ListenableFuture;

import java.util.concurrent.TimeUnit;

/**
 * <h1> kafka producer</h1>
 */
@Slf4j
@Component
public class KafkaProducer {

  private final KafkaTemplate<String, String> kafkaTemplate;

  public KafkaProducer(KafkaTemplate<String, String> kafkaTemplate) {
    this.kafkaTemplate = kafkaTemplate;
  }

  /**
   * <h2> send kafka message</h2>
   */
  public void sendMessage(String key, String value, String topic) {
    if (StringUtils.isBlank(value) || StringUtils.isBlank(topic)) {
      throw new IllegalArgumentException("value or topic is null or empty");
    }

    ListenableFuture<SendResult<String, String>> future = StringUtils.isBlank(key) ?
        kafkaTemplate.send(topic, value) : kafkaTemplate.send(topic, key, value);

    // 异步回调的方式获取通知
    future.addCallback(
        success -> {
          assert null != success && null != success.getRecordMetadata();
          // 发送到 kafka 的 topic
          String _topic = success.getRecordMetadata().topic();
          // 消息发送到的分区
          int partition = success.getRecordMetadata().partition();
          // 消息在分区内的 offset
          long offset = success.getRecordMetadata().offset();

          log.info("send kafka message success: [{}], [{}], [{}]", _topic, partition, offset);
        },
        failure -> {
          log.error("send kafka message failure: [{}], [{}], [{}]", key, value, topic);
        }
    );

    // 同步等待的方式获取通知
    try {
      SendResult<String, String> sendResult = future.get(5, TimeUnit.SECONDS);

      // 发送到 kafka 的 topic
      String _topic = sendResult.getRecordMetadata().topic();
      // 消息发送到的分区
      int partition = sendResult.getRecordMetadata().partition();
      // 消息在分区内的 offset
      long offset = sendResult.getRecordMetadata().offset();

      log.info("send kafka message success: [{}], [{}], [{}]", _topic, partition, offset);
    } catch (Exception ex) {
      log.error("send kafka message failure: {}, {}, {}", key, value, topic);
    }
  }
}







