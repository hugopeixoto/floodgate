package main

import (
  "net/http"
)

type Stream struct {
  TargetAddress  string
  ListenAddress  string
  ControlAddress string
}

func RunFloodgate(stream Stream) {
  handler := NewFloodgate(stream.TargetAddress)

  http.HandleFunc("/hold", func(http.ResponseWriter, *http.Request){ handler.Hold() })
  http.HandleFunc("/release", func(http.ResponseWriter, *http.Request){ handler.Release() })

  go http.ListenAndServe(stream.ControlAddress, nil)
  go http.ListenAndServe(stream.ListenAddress, handler)
}

func main() {
  defaultStream := Stream{"http://localhost:81", ":8081", ":8082"}

  RunFloodgate(defaultStream)

  // Wait forever
  <-make(chan struct{})
}
