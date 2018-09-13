
/**

第二种方法是经典的处理大请求的方法， 把HTTP请求先放进队列里面去，通过控制队列的容量， 使之不会产生很多HTTP请求
同时处理的情况， 所以web处理程序不会卡死，

常见的一种做法是，举例来说， 我们先把要处理的HTTP任务放在redis队列里面去， 队列里面储存我们要处理HTTP请求的
一些数据， 然后我们后台有个进程去实际读取redis队列， 进行逻辑处理

但是这个例子是用golang 的 channel 来做队列，实际上跟上面的情况是一样的

**/


// 定义channel 队列的容量最多为100
MAX_QUEUE = 10000

var Queue = chan Payload

func init() {
  Queue = make(chan Payload, MAX_QUEUE)
}


func payloadHandler(w http.ResponseWriter, r *http.Request) {

  // 将包体读入到一个字符串进行json解析
  var content = &PayloadCollection()
  err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)) . Decode(&content)
  if err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  // 遍历每个有效载荷， 逐个排队之后上传S3
  for _, payload := range content.Payloads {
    // 把 payload 入队列, 以前的做法是入redis队列， 现在是入golang 的channel队列，其实原理上是一样的
    Queue <- payload
  }  
}


// 接下来， 我们再从队列中取job，然后处理它们
func StartProcessor() {
  for {
    select {
    case job := <-Queue:
      job.payload.UploadToS3()   // 仍然不够好
    }
  }
}


func main() {
  // 写进 channel, 满10000 就等待不写进， StartProcessor() 不断处理， 然后有空间了， 又不断写进
  payloadHandler()     

   // 循环处理 channel， 越处理， channel就能腾出空间继续写进 
  StartProcessor()     
}


/**
这种方法肯定是不好的， 因为限制了channle队列最大的容量为 10000， 意思就是当 channel塞满 10000 个请求的时候，
就暂时不把请求写进队列， 然后请求就在那里 hold住， StartProcessor 方法不断的处理 channel 里面的请求，一边
处理， 那边HTTP请求也就同时的把新的请求写进channel,这样的话，只能保证我们的程序不会奔溃，但是响应还是很慢的， 
实际上对于请求的客户端来说， 这种方案跟刚才第一种方案是一样的， 同样都是延时很大， 因为请求太大了， 但是对于
服务端处理程序来说， 这个会好一点， 因为服务端不用一下子接受那么多的请求， 而是只需针对 channel 队列里面的
10000条请求来处理就可以了。

所以这个方案是不及格的

那么问题来了， 在没有更多的抗压资源的情况下， 通过优化代码可以提供程序的性能吗？？

这里问题是关键是， 上传到S3 这一步逻辑， 我们是需要等待的， 如果没有这一步逻辑， StartProcessor() 消费 channel
数据的时候， 就会快速很多，，，，， 所以下一个优化版本的思路是什么？ 就是把 上传到S3 这一个逻辑也异步队列化？
那是不是可以认为， 实际上程序的优化有一个共同的大思路就是， 把耗时的同步程序能异步队列话就队列化呢？？ 哈哈
但是异步队列化的时候要注意原子性操作。
**/













