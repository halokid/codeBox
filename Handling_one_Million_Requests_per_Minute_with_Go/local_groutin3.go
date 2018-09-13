
/**

我们讨论过在使用用Go channel时利用一种常用的模式，来创建一个二级channel系统，一个来queue job，另外一个来控制使用多少个worker来并发操作JobQueue。

想法是，以一个恒定速率并行上传到S3，既不会导致机器崩溃也不好产生S3的连接错误。这样我们选择了创建一个Job/Worker模式。对于那些熟悉Java、C#等语言的开发者，可以把这种模式想象成利用channel以golang的方式来实现了一个worker线程池，作为一种替代。

**/

var (
  MaxWorker = os.Getenv("MAX_WORKERS")
  MaxQueue = os.Getenv("MAX_QUEUE")
)



