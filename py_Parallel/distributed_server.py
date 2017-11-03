

# 管理分布式协作manager对象
def get_manager():
    '''创建服务端manager对象.
    '''
    # 自定义manager类
    class JobManager(BaseManager):
        pass
    # 创建任务队列，并将此数据对象共享在网络中
    jobid_queue = Queue()
    JobManager.register('get_jobid_queue', callable=lambda: jobid_queue)
    # 创建列表代理类，并将其共享再网络中
    tofs = [None] * N
    JobManager.register(
        'get_tofs_list', callable=lambda: tofs, proxytype=ListProxy)
    # 将分压参数共享到网络中
    JobManager.register('get_pCOs', callable=lambda: pCOs, proxytype=ListProxy)
    JobManager.register('get_pO2s', callable=lambda: pCOs, proxytype=ListProxy)
    # 创建manager对象并返回
    manager = JobManager(address=(ADDR, PORT), authkey=AUTHKEY)
    return manager


# 进行任务分配
def fill_jobid_queue(manager, nclient):
    indices = range(N)
    interval = N / nclient
    jobid_queue = manager.get_jobid_queue()
    start = 0
    for i in range(nclient):
        jobid_queue.put(indices[start: start + interval])
        start += interval
    if N % nclient > 0:
        jobid_queue.put(indices[start:])


def run_server():
    # 获取manager
    manager = get_manager()
    print "Start manager at {}:{}...".format(ADDR, PORT)
    # 创建一个子进程来启动manager
    manager.start()
    # 填充任务队列
    fill_jobid_queue(manager, NNODE)
    shared_job_queue = manager.get_jobid_queue()
    shared_tofs_list = manager.get_tofs_list()
    queue_size = shared_job_queue.qsize()
    # 循环进行监听，直到结果列表被填满
    while None in shared_tofs_list:
        if shared_job_queue.qsize() < queue_size:
            queue_size = shared_job_queue.qsize()
            print "Job picked..."
    return manager
