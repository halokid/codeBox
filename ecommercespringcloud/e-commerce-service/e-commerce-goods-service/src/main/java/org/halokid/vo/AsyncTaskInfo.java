package org.halokid.vo;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.halokid.constant.AsyncTaskStatusEnum;

import java.util.Date;

/**
 * <h1>异步任务执行信息</h1>
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class AsyncTaskInfo {
  /**异步任务 id*/
  private String taskId;

  private AsyncTaskStatusEnum status;

  /**异步任务开始时间 */
  private Date startTime;
  /**异步任务结束时间*/
  private Date endTime;
  /**异步任务总耗时*/
  private String totalTime;
}



