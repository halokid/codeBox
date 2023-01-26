package org.halokid.controller;

import org.halokid.service.SleuthTraceInfoService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * <h1>打印跟踪信息</h1>
 */
@Slf4j
@RestController
@RequestMapping("/sleuth")
public class SleuthTraceInfoController {
  @Autowired
  private SleuthTraceInfoService traceInfoService;

  /**
   * <h2>打印日志跟踪信息</h2>
   */
  @GetMapping("/trace-info")
  public void logCurrentTraceInfo() {
    traceInfoService.logCurrentTraceInfo();
  }
}
