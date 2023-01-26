package org.halokid.service;


import brave.Tracer;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author r0x
 * @description <h1>使用代码更直观的看到 Sleuth 生成的相关跟踪信息</h1>
 * @date 2021/9/20 9:37
 */
@Slf4j
@Service
public class SleuthTraceInfoService {
  /**
   * brave.Tracer 跟踪对象
   */
  @Autowired
  private Tracer tracer;


  /**
   * <h2>打印当前的跟踪信息到日志中</h2>
   */
  public void logCurrentTraceInfo() {
    log.info("Sleuth trace id: [{}]", tracer.currentSpan().context().traceId());
    log.info("Sleuth span id: [{}]", tracer.currentSpan().context().spanId());
  }
}