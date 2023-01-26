package org.halokid.partition;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.vo.HalokidMessage;
import org.springframework.cloud.stream.binder.PartitionKeyExtractorStrategy;
import org.springframework.messaging.Message;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class halokidPartitionKeyExtractorStrategy implements PartitionKeyExtractorStrategy {

  @Override
  public Object extractKey(Message<?> message) {

    HalokidMessage halokidMessage = JSON.parseObject(
      message.getPayload().toString(), HalokidMessage.class
    );

    String key = halokidMessage.getProjectName();
    log.info("SrpingCloud Stream Parition Key: {}", key);
    return key;
  }
}



