package main

import "context"

func Stream(ctx context.Context, out chan <- Value) error {
  for {
    v, err := DoSomething(ctx)

    if err != nil {
      return err
    }

    select {
    case <- ctx.Done():
      return ctx.Err()

    case out <- v:
      // do something
    }
  }
}
