package main

import (
  "context"
  "golang.org/x/time/rate"
  "sort"
)

type RateLimiter interface {
  Wait(ctx context.Context) error
  Limit() rate.Limit
}

type multiLimiter struct {
  limiters []RateLimiter
}

func MultiLimiter(limiters ...RateLimiter) *multiLimiter {
  byLimit := func(i, j int) bool {
    return limiters[i].Limit() < limiters[j].Limit()
  }
  sort.Slice(limiters, byLimit)
  return &multiLimiter{limiters:  limiters}
}

func (l *multiLimiter) Wait(ctx context.Context) error {
  for _, l := range l.limiters {
    if err := l.Wait(ctx); err != nil {
      return err
    }
  }
  return nil
}

func (l *multiLimiter) Limit() rate.Limit {
  return l.limiters[0].Limit()
}








