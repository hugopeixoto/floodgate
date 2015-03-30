package main

import (
  "log"
  "time"
)

type Counter struct {
  Req     chan int
  queued  int
  handled int
}

func NewCounter() *Counter {
  c := Counter{}

  c.Req = make(chan int)
  c.queued  = 0
  c.handled = 0

  return &c
}

func (c *Counter) Count(v int) {
  c.Req <- v
}

func (c *Counter) Run() {
  ticker  := time.NewTicker(time.Millisecond * 1000)

  for {
    select {
      case d := <-c.Req: c.queued += d; if d < 0 { c.handled += -d }
      case <-ticker.C: c.register()
    }
  }
}

func (c *Counter) register() {
  log.Printf("queued requests: %v, handled %v", c.queued, c.handled)
}
