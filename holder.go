package floodgate

import (
  "fmt"
)

type Holder struct {
  Req chan struct{}
  StateChanger chan bool
  Active bool
}

func NewHolder() *Holder {
  h := Holder{}

  h.Active       = true
  h.Req          = make(chan struct{})
  h.StateChanger = make(chan bool)

  return &h
}

func (h *Holder) Hold() {
  h.Req <- struct{}{}
}

func (h *Holder) UpdateState(state bool) {
  h.Active = state

  fmt.Printf("Changing state: %v\n", h.Active)
}

func (h *Holder) Run() {
  for {
    for h.Active {
      select {
        case <-h.Req:
        case s := <-h.StateChanger: h.UpdateState(s)
      }
    }

    h.UpdateState(<-h.StateChanger)
  }
}
