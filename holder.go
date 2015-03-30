package main

import (
  "log"
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

  if h.Active {
    log.Println("releasing traffic")
  } else {
    log.Println("holding traffic")
  }
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
