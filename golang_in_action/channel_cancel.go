package golang_in_action

func isCanceled(cancelChan chan struct{}) bool {
  select {
  //case i := <-cancelChan:
  case <-cancelChan:
   //fmt.Println("isCanceled读取到", cancelChan)
    //fmt.Println("isCanceled读取到", i)
    return true
  default:
    return false
  }
}

func cancel2(cancelChan chan struct{}) {
  cancelChan <- struct{}{}
}

func cancel1(cancelChan chan struct{}) {
  close(cancelChan)
}