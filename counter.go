package floodgate

import (
  "fmt"
  "time"
)

type Counter struct {
  Req chan int
}

func NewCounter() *Counter {
  c := Counter{}
  c.Req = make(chan int)

  return &c
}

func (c *Counter) Count(v int) {
  c.Req <- v
}

func (c *Counter) Run() {
  ticker  := time.NewTicker(time.Millisecond * 1000)
  queued  := 0
  handled := 0

  for {
    select {
      case d := <-c.Req: queued += d; if d < 0 { handled += -d }
      case <-ticker.C: c.register()
    }
  }
}

func (c *Counter) register() {
  fmt.Printf("Queued requests: %v, handled %v\n", queued, handled)
}
