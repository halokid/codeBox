

func DoSomething(ch chan int) {
  ch <- 1
  var i = <-ch
}




static async void DoSomething<T> (T block) where T : ISourceBlock<int>, ITargeBlock<int> {
  await block.SendAsync(1);
  var i = await block.ReceiveAsync();
}



// in golang, base on groutine & channel programming model normal like this:
func (ch chan int) {
  for {     //dead loop
    var msg = <-ch
    
    Proccess(msg)     //del the msg
  }
}


