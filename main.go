package main

import (
  "net/http"
  "flag"
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
  var (
    target_address  = flag.String("target", "http://localhost:81", "location of the service being protected")
    listen_address  = flag.String("listen", ":8081", "adddress for floodgate proxy")
    control_address = flag.String("control", ":8082", "address for the control entrypoints")
  )

  flag.Parse()

  defaultStream := Stream{*target_address, *listen_address, *control_address}

  RunFloodgate(defaultStream)

  // Wait forever
  <-make(chan struct{})
}
