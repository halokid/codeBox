package org.halokid.service.async;

import lombok.extern.slf4j.Slf4j;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.halokid.constant.AsyncTaskStatusEnum;
import org.halokid.vo.AsyncTaskInfo;
import org.springframework.stereotype.Component;

import java.util.Date;

/**
 * <h2>异步任务执行监控切面</h2>
 */
@Slf4j
@Aspect
@Component
public class AsyncTaskMonitor {
  /**注入异步任务管理器*/
  private final AsyncTaskManager asyncTaskManager;

  public AsyncTaskMonitor(AsyncTaskManager asyncTaskManager) {
    this.asyncTaskManager = asyncTaskManager;
  }

  /**
   * <h2>异步任务执行的环绕切面</h2>
   * TODO: 环绕切面让我们可以在方法执行之前和执行之后做一些"额外"的操作
   * @param proceedingJoinPoint
   * @return
   */
  // TODO: 这个是异步执行的关键1， 定义了这个注解，当 AsyncServiceImpl 里面的异步任务调用的时候，就会触发这个func， 这个就是执行
  // TODO: 异步任务的触发入口，可以在这里管理异步任务的状态
  @Around("execution(* org.halokid.service.async.AsyncServiceImpl.*(..))")
  public Object taskHandle(ProceedingJoinPoint proceedingJoinPoint){
    // 获取taskId ，调用异步任务传入的第二个参数
    String taskId = proceedingJoinPoint.getArgs()[1].toString();
    // 获取任务信息，在提交任务的时候就已经放入到容器中了
    AsyncTaskInfo taskInfo = asyncTaskManager.getTaskInfo(taskId);
    log.info("AsyncTaskMonitor is monitoring async task:[{}]",taskId);
    taskInfo.setStatus(AsyncTaskStatusEnum.RUNNING);
    asyncTaskManager.setTaskInfo(taskInfo);//设置为运行状态
    AsyncTaskStatusEnum status;
    Object result;
    try{
      //执行异步任务
      result = proceedingJoinPoint.proceed();
      status = AsyncTaskStatusEnum.SUCCESS;
    }catch (Throwable ex){
      // 异步任务出现异常
      result=null;
      status=AsyncTaskStatusEnum.FAILED;
      log.error("AsyncTaskMoitor: async task[{}] is failed,Error Info:[{}]",
          taskId,ex.getMessage(),ex);

    }

    // 设置异步任务其他的信息，再次重新放入到容器中
    taskInfo.setEndTime(new Date());
    taskInfo.setStatus(status);
    taskInfo.setTotalTime(String.valueOf(
        taskInfo.getEndTime().getTime()-taskInfo.getStartTime().getTime()
    ));
    asyncTaskManager.setTaskInfo(taskInfo);
    return result;
  }
}

