package org.halokid.rocket;

import com.alibaba.fastjson.JSON;
import lombok.extern.apachecommons.CommonsLog;
import lombok.extern.slf4j.Slf4j;
import org.apache.rocketmq.spring.annotation.RocketMQMessageListener;
import org.apache.rocketmq.spring.core.RocketMQListener;
import org.halokid.vo.Message;

/**
 * <h1> 第一个 RocketMQ 消费者</h1>
 */
@Slf4j
@CommonsLog
@RocketMQMessageListener(
    topic = "halokid-study-rocketmq",
    consumerGroup = "halokid-springboot-rocketmq-tag-string"
)
public class RocketMQConsumerString implements RocketMQListener<String> {
  @Override
  public void onMessage(String message) {

    Message rocketMessage = JSON.parseObject(message, Message.class);
    log.info("consumer message in RocketMQConsumerString: {}",
        JSON.toJSONString(rocketMessage));
  }
}


