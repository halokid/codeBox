
/**

我们讨论过在使用用Go channel时利用一种常用的模式，来创建一个二级channel系统，一个来queue job，另外一个来控制使用多少个worker来并发操作JobQueue。

想法是，以一个恒定速率并行上传到S3，既不会导致机器崩溃也不好产生S3的连接错误。这样我们选择了创建一个Job/Worker模式。对于那些熟悉Java、C#等语言的开发者，可以把这种模式想象成利用channel以golang的方式来实现了一个worker线程池，作为一种替代。

**/

var (
  MaxWorker = os.Getenv("MAX_WORKERS")
  MaxQueue = os.Getenv("MAX_QUEUE")
)


// TODO:  首先定义一个 Job 的结构体

// Job指要运行的任务
type Job struct {
  Payload Payload
}


// TODO: 定义一个储存Job 的channel,  这个就是一个 Job队列

// 一个可以发送工作请求的缓冲 channel
var JobQueue chan Job 


// TODO: 定义 Worker 的结构体

// Worker 表示要执行任务的 worker
type Worker struct {
  WorkerPool chan chan Job  
  JobChannel chan  Job
  quit  chan bool   
} 


// TODO:  定义一个函数，生成一个 Worker

// 处理上传到 S3 的实际逻辑
func NewWorker(workerPool chan chan Job) Worker {
  return Worker {
    workerPool:   workerPool,
    JobChannel:   make(chan Job),
    quit:         make(chan bool)
  }
}



// Start方法开启一个worker循环， 监听退出channel, 可按需停止这个循环
fucn (w Worker) Start() {
  go func() {
    for {
      // 将当前的 worker 注册到 worker 队列中
      w.WorkerPool <- w.JobChannel

      select {
      case job := <- w.JobChannel:
        // 此时我们接收到一个工作请求
        if err := job.Payload.UploadToS3(); err != nil {
          log.Errorf("Error uploading to S3: %s", err.Error())
        }

      case <- w.quit:
        // 此处接收一个停止信号
        return 
      }
    } 
  }()
}


// Stop方法控制 worker 停止监听工作请求
func (w Worker) Stop() {
  go func() {
    w.quit() <- true
  }()
}




// 我们修改了 Web 请求处理器，使之能够创建一个携带载荷信息的 Job 实例， 然后把它发到 JobQueue channel 中供 worker 消费
func payloadHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.WriteHeader(http.StatusMethodNotAllowed)
    return
  }

  // 把包体读入到一个字符串中， 进行json解析
  var content = &PayloadCollection{}
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // TODO: 事实上 content 只是一串 json 类型的数据格式

  // 遍历载荷， 逐个队列化用以上传到S3 
  for _, payload := range content.Payloads {

    // TODO:  JSON 类型的数据格式， 描述要做的 payload

    // 创建一个带有载荷的任务, Job实例
    work := Job{ Payload:  payload}  

    // 然后把它放到队列中
    JobQueue <- work
  }

  w.WriteHeader(http.StatusOK)
}



func main() {
  disptcher := NewDispatcher(MaxWorker)
  disptcher.Run() 
}


/**
证明自己对写程序的思路还是不够清晰， 很多东西都可以用方程式去理解， 所以自己可以尝试先用方程式去理解任何
东西，不行再去寻找其他的理解方式
**/
























