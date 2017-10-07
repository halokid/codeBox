

def get_manager():
    class WorkManager(BaseManager):
        pass
    # 由于只是从共享网络中获取，因此只需要注册名字即可
    WorkManager.register('get_jobid_queue')
    WorkManager.register('get_tofs_list')
    WorkManager.register('get_pCOs')
    WorkManager.register('get_pO2s')
    # 这里的地址和验证码要与服务端相同才可以进行数据共享
    manager = WorkManager(address=(ADDR, PORT), authkey=AUTHKEY)
    return manager



if "__main__" == __name__:
    manager = get_manager()
    print "work manager connect to {}:{}...".format(ADDR, PORT)
    # 将客户端本地的manager连接到相应的服务端manager
    manager.connect()
    # 获取共享的结果收集列表
    shared_tofs_list = manager.get_tofs_list()
    # 获取共享的任务队列
    shared_jobid_queue = manager.get_jobid_queue()
    # 从服务端获取计算参数
    pCOs = manager.get_pCOs()
    shared_pO2s = manager.get_pO2s()
    # 创建进程池在本地计算机进行多核并行
    pool = Pool()
    while 1:
        try:
            indices = shared_jobid_queue.get_nowait()
            pO2s = [shared_pO2s[i] for i in indices]
            print "Run {}".format(str(pO2s))
            tofs_2d = pool.map(task, pO2s)
            # Update shared tofs list.
            for idx, tofs_1d in zip(indices, tofs_2d):
                shared_tofs_list[idx] = tofs_1d
        # 直到将任务队列中的任务全部取完，结束任务进程
        except Queue.Empty:
            break




