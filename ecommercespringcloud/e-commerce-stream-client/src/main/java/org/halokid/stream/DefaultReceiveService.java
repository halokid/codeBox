package org.halokid.stream;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.vo.HalokidMessage;
import org.springframework.cloud.stream.annotation.EnableBinding;
import org.springframework.cloud.stream.annotation.StreamListener;
import org.springframework.cloud.stream.messaging.Sink;

@Slf4j
@EnableBinding(Sink.class)
public class DefaultReceiveService {

  @StreamListener(Sink.INPUT)
  public void receiveMessage(Object payload) {

    log.info("in DefaultReceiveService consume message start");
    HalokidMessage message = JSON.parseObject(
      payload.toString(), HalokidMessage.class
    );

    log.info("in DefaultReceiveService consume message success: {}",
        JSON.toJSONString(message));
  }
}
