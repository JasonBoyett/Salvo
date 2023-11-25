package testserver

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TestOpts struct {
  Message string
  Fail bool
  Delay time.Duration
  Port int
}



// TestServer starts a test server that will respond to requests with the given message.
// paramaters: TestOpts
// message: string
//  the message to respond with
// fail: bool
//  if true, the server will respond with a 500 error
// delay: time.Duration
//  the amount of time to wait before responding
// port: int
//  the port to listen on
// returns:
//  a function that will stop the server
//
// example:
// kill, err := TestServer(TestOpts{
//   "hello", 
//   false, 
//   time.Duration(time.Duration(10).Seconds()), 
//   8080,
// })
func TestServer(opts TestOpts) (func() error, error) {

  ch := make(chan error, 1)

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

    time.Sleep(opts.Delay)

    if opts.Fail {
      http.Error(
        w, 
        "Server intentionally failed", 
        http.StatusInternalServerError,
      )
    } else {
      _, err := w.Write([]byte(opts.Message))
      if err != nil {
        ch <- err
      }
    }    

  })

  server := &http.Server{Addr: fmt.Sprintf(":%d", opts.Port)}
  
  go func() {
    if err := server.ListenAndServe(); err != nil {
      ch <- err
      close(ch)
    }
  }()
  
  kill := func() error {
    if err := server.Shutdown(context.TODO()); err != nil {
      return err
    }
    return nil
  }

  err := <- ch

  return kill, err
}
