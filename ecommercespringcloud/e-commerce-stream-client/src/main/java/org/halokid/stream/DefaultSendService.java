package org.halokid.stream;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.vo.HalokidMessage;
import org.springframework.cloud.stream.annotation.EnableBinding;
import org.springframework.cloud.stream.messaging.Source;
import org.springframework.messaging.support.MessageBuilder;

/**
 * <h1> use default communication channel to realize message send</h1>
 */
@Slf4j
@EnableBinding(Source.class)
public class DefaultSendService {

  private final Source source;

  public DefaultSendService(Source source) {
    this.source = source;
  }

  /**
   * <h2> use default Output channel to send message</h2>
   */
  public void sendMessage(HalokidMessage message) {
    String _message = JSON.toJSONString(message);
    log.info("in DefaultSendService send message: {}", _message);

    // Spring Messaging, unitive message program model, is an important part of the Stream component
    source.output().send(MessageBuilder.withPayload(_message).build());
  }
}







