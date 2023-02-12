package org.halokid.rocket;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.client.producer.SendCallback;
import org.apache.rocketmq.client.producer.SendResult;
import org.apache.rocketmq.spring.core.RocketMQTemplate;
import org.apache.rocketmq.spring.support.RocketMQHeaders;
import org.springframework.messaging.Message;
import org.springframework.messaging.support.MessageBuilder;
import org.springframework.stereotype.Component;

/**
 * <h1> 通过 RocketMQProducer 发送消息</h1>
 */
@Slf4j
@Component
public class RocketMQProducer {

  // 类似 kafka 中的 topic, 默认的读写队列都是 4 个
  private static final String TOPIC = "halokid-study-rocketmq";

  // Rocket 客户端
  private final RocketMQTemplate rocketMQTemplate;

  public RocketMQProducer(RocketMQTemplate rocketMQTemplate) {
    this.rocketMQTemplate = rocketMQTemplate;
  }

  /**
   * <h2> 使用同步的方式发送消息， 不指定 key 和 tag</h2>
   */
  public void sendMessageWithValue(String value) {
    // 随机选择一个 Topic 的 Message Queue 发送消息
    SendResult sendResult = rocketMQTemplate.syncSend(TOPIC, value);
    log.info("sendMessageWithValue result: {}", JSON.toJSONString(sendResult));

    // 按照顺序发送消息
    SendResult sendResultOrderly = rocketMQTemplate.syncSendOrderly(
        TOPIC, value, "Halokid"
    );
    log.info("sendMessageWithValue orderly result: {}", JSON.toJSONString(sendResultOrderly));
  }

  /**
   * <h2> 使用异步的方式发送消息， 指定 key</h2>
   */
  public void sendMessageWithKey(String key, String value) {
    Message<String> message = MessageBuilder.withPayload(value)
        .setHeader(RocketMQHeaders.KEYS, key).build();

    // 异步发送消息， 并设定回调
    rocketMQTemplate.asyncSend(TOPIC, message, new SendCallback() {
      @Override
      public void onSuccess(SendResult sendResult) {
        log.info("sendMessageWithKey success result: {}",
            JSON.toJSONString(sendResult));
      }

      @Override
      public void onException(Throwable e) {
        log.error("sendMessageWithKey failure result: {}",
            e.getMessage(), e);
      }
    });
  }

  /**
   * <h2> 使用同步的方式发送消息， 带有tag， 且发送的是 Java Pojo</h2>
   */
  public void sendMessageWithTag(String tag, String value) {

    org.halokid.vo.Message message = JSON.parseObject(value, org.halokid.vo.Message.class);
    SendResult sendResult = rocketMQTemplate.syncSend(
        String.format("%s:%s", TOPIC, tag),
        message
    );
    log.info("sendMessageWithTag result: {}", JSON.toJSONString(sendResult));
  }

  /**
   * <h2> 使用同步的方式发送消息， 带有 key 和 tag</h2>
   */
  public void sendMessageWithAll(String key, String tag, String value) {

    Message<String> message = MessageBuilder.withPayload(value)
        .setHeader(RocketMQHeaders.KEYS, key).build();
    SendResult sendResult = rocketMQTemplate.syncSend(
        String.format("%s:%s", TOPIC, tag),
        message
    );
    log.info("sendMessageWithAll result: {}", JSON.toJSONString(sendResult));
  }
}