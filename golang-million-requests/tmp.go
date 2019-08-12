package main 


const maxWorkers = 5

type job struct {
  name            string 
  duration        time.Duration 
}


func doWork(id int, j job) {
  time.Sleep(j.duration)
  fmt.Printf("do work %n, %s\n", id, j.name) 
}



func main() {
  jobs := make(chan job)

  wg := &sync.WaitGroup{}

  wg.Add(maxWorkers)
  for i := 1; i <= maxWorkers; i++ {
    go func (i int) {
      defer wg.Done()

      for j := range jobs {
        doWork(i, j)
      } 
    }(i)
  } 


  iCk := 0
  for i := 0; i < 100; i++ {
    name := fmt.Sprintf("job-%d", i)
    duration := time.Duration(rand.Intn(1000)) * time.Millisecond

    jobs <- job{name: name,  duration:  duration}

    iCk++
  }

  close(jobs)

  wg.Wait()
}