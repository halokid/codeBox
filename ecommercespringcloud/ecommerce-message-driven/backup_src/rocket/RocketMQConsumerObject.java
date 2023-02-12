package org.halokid.rocket;

import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.spring.annotation.RocketMQMessageListener;
import org.apache.rocketmq.spring.core.RocketMQListener;
import org.halokid.vo.Message;
import org.springframework.stereotype.Component;
import com.alibaba.fastjson.JSON;

/**
 * <h1>第四个, RocketMQ 消费者, 指定消费带有 tag 的消息, 且消费的是 Java Pojo</h1>
 * */
@Slf4j
@Component
@RocketMQMessageListener(
    topic = "imooc-study-rocketmq",
    consumerGroup = "qinyi-springboot-rocketmq-tag-object",
    selectorExpression = "qinyi"    // 根据 tag 做过滤
)
public class RocketMQConsumerObject implements RocketMQListener<Message> {

  @Override
  public void onMessage(Message message) {

    log.info("consume message in RocketMQConsumerObject: [{}]",
        JSON.toJSONString(message));
    // do something
  }
}









