package main

import (
  "context"
  "log"
  "os"
  "sync"
)

type APIConnection struct {}

func Open() *APIConnection {
  return &APIConnection{}
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
  log.Println("ReadFile...")
  return nil
}

func (a *APIConnection) ResolvedAddress(ctx context.Context) error {
  log.Println("ResolveAddress...")
  return nil
}


func main() {

  defer log.Printf("Done.")

  log.SetOutput(os.Stdout)
  log.SetFlags(log.Ltime | log.LUTC)

  apiConnection := Open()
  var wg sync.WaitGroup
  wg.Add(20)

  for i := 0; i < 10; i++ {
    go func() {
      defer wg.Done()
      err := apiConnection.ReadFile(context.Background())
      if err != nil {
        log.Printf("cannot ReadFile: %v", err)
      }
    }()

  }
  //log.Printf("ReadFile")

  for i := 0; i < 10; i++ {
    go func() {
      defer wg.Done()
      err := apiConnection.ResolvedAddress(context.Background())
      if err != nil {
        log.Printf("cannot ResolveAddress: %v", err)

      }
    }()
  }
  //log.Printf("ResolveAddress")

  wg.Wait()
}





