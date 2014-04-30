package floodgate

import (
  "net/http"
  "net/http/httputil"
  "net/url"
)

type Floodgate struct {
  proxy *httputil.ReverseProxy
  holder *Holder
  counter *Counter
}

func NewFloodgate(target string) *Floodgate {
  fg := Floodgate{}

  url, _ := url.Parse(target)
  fg.proxy = httputil.NewSingleHostReverseProxy(url)

  fg.counter = NewCounter()
  go fg.counter.Run()

  fg.holder = NewHolder()
  go fg.holder.Run()

  return &fg
}

func (dh *Floodgate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  dh.counter.Count(1)
  dh.holder.Hold()

  dh.proxy.ServeHTTP(w, r)

  dh.counter.Count(-1)
}

func (dh *Floodgate) Hold() {
  dh.holder.StateChanger <- false
}

func (dh *Floodgate) Release() {
  dh.holder.StateChanger <- true
}

func main() {
  handler := NewFloodgate("http://localhost:81")

  http.Handle("/", handler)
  http.HandleFunc("/hold", func(http.ResponseWriter, *http.Request){ handler.Hold() })
  http.HandleFunc("/release", func(http.ResponseWriter, *http.Request){ handler.Release() })
  http.ListenAndServe(":8081", nil)
}
