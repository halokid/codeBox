package org.halokid.partition;

import com.alibaba.fastjson.JSON;
import lombok.extern.slf4j.Slf4j;
import org.halokid.vo.HalokidMessage;
import org.springframework.cloud.stream.binder.PartitionKeyExtractorStrategy;
import org.springframework.cloud.stream.binder.PartitionSelectorStrategy;
import org.springframework.messaging.Message;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class halokidPartitionSelectorStrategy implements PartitionSelectorStrategy {

  @Override
  public int selectPartition(Object key, int partitionCount) {

    int partition = key.toString().hashCode() % partitionCount;
    log.info("SrpingCloud Stream Selector Key: {}, {}, {}", key.toString(), partitionCount, partition);
    return partition;
  }
}



