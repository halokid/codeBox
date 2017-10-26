

struct scheduler_functions {
  uint max_threads;
  bool (*int)(void);
  bool (*int_new_connection_thread)(void);

  void (*add_connection) (THD *thd);  //当有新的连接到来时，会回调此函数
  void (*thd_wait_begin) (THD *thd, int wait_type);
  void (*the_wait_end) (THD *thd); 
  void (*post_kill_notification) (THD *thd);

  bool (*end_thread) (THD *thd, bool cache_thread);
  void (*end)(void);
};



