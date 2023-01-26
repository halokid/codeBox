package org.halokid.rocket;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.common.message.MessageExt;
import org.apache.rocketmq.spring.annotation.RocketMQMessageListener;
import org.apache.rocketmq.spring.core.RocketMQListener;
import org.springframework.stereotype.Component;

/**
 * <h1>第三个 RocketMQ 消费者, </h1>
 * */
@Slf4j
@Component
@RocketMQMessageListener(
    topic = "halokid-study-rocketmq",
    consumerGroup = "halokid-springboot-rocketmq-message-ext"
)
public class RocketMQConsumerMessageExt implements RocketMQListener<MessageExt> {

  @Override
  public void onMessage(MessageExt message) {

    String value = new String(message.getBody());
    log.info("consume message in RocketMQConsumerMessageExt: [{}], [{}]",
        message.getKeys(), value);
    log.info("MessageExt: [{}]", JSON.toJSONString(message));   // 会慢一些
  }
}
