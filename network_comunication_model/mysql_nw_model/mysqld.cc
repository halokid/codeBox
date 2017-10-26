
while (!abort_loop)
{
  retval = poll(fds, socket_count, -1);

  for (int i = 0; i < socket_count; ++i) {
    if (fds[i].revents & POLLIN ) {
      sock = pfs_fds[i];
      flags = fcntl(mysql_socket_getfd(sock), F_GETFL, 0);
      break;
    }
  }


  new_sock = mysql_socket_accept(key_socket_client_connection, sock, 
                  (struct sockaddr *)(&cAddr), &length );

  thd = new THD;

  create_new_thread(thd);
  
}


// create_new_thread最终会调用create_thread_to_handle_connection处理请求
void create_thread_to_handle_connection(THD *thd)
{
  mysql_mutex_assert_owner(&LOCK_thread_count);
   
  if (blocked_pthread_count >  wake_pthread) //说明ThreadCache中还有空闲线程
  {
    /* Wake up blocked pthread */
    waiting_thd_list->push_back(thd); //将thd放到队列，并通知线程拿走thd
    wake_pthread++;
    mysql_cond_signal(&COND_thread_cache);
  }
  else // 没有空闲线程了，创建新线程
  {
    mysql_thread_create(key_thread_one_connection,
                        &thd->real_id, &connection_attrib,
                        handle_one_connection, // 新线程入口函数为handle_one_connection
                        (void*) thd)
  }
   
  mysql_mutex_unlock(&LOCK_thread_count);
}


void do_handle_one_connection(THD *thd) {
  for (;;)
  {
    while (thd_is_connection_alive(thd)) {
      if (do_command(thd))
        break;
    }


    //当上一个连接的请求全都被处理完时， 在条件变量 COND_thread_cache 上等待
    if (blocked_pthread_count < max_blocked_pthreads && 
        !abort_loop && !kill_blocked_pthreads_flag)
        {
          mysql_cond_wait(&COND_thread_cache, &LOCK_thread_count);
          thd= waiting_thd_list->front(); // 从队列中取出第一个元素
          waiting_thd_list->pop_front(); 
        }  
  }
}




int ha_innobase::write_row (unchar* record) {
  innobase_srv_conc_enter_innodb(prebuilt->trx); //操作数据前，先控制并发
  error = row_insert_for_mysql((byte*)record, prebuilt); //操作数据
  innobase_srv_conc_exit_innodb(prebuilt->trx);
}


void innobase_srv_conc_enter_innodb(trx_t *trx) {
  if (srv_thread_concurrency) {
    if (trx->n_ticket_to_enter_innodb > 0) {    //线程里已有tickets
      --trx->n_ticket_to_enter_innodb;
    } else {
      srv_conc_enter_onnodb(trx);
    }
  }
}



/** FIXME:
下面这段代码的逻辑是:

下面三种就是 线程的 handle 模式
1. 每个连接一个线程处理 
2. 不用线程去处理调度， 把连接调度到其他的逻辑代码去
3， 把连接加入到 线程池 里面去


**/
if (thread_handing <= SCHEDULER_ONE_THREAD_PER_CONNECTION) {
  one_thread_per_connection_scheduler(thread_schedule, &max_connections,
                                       &connection_count);
  
}
else if (thread_handing == SCHEDULER_NO_THREADS) {
  one_thread_scheduler(thread_schedule);
} 
else {
  pool_of_threads_scheduler(thread_schedule, &max_connections, &connection_count);
}















