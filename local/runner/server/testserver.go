package server

import (
	"fmt"
	"net/http"
	"time"
)

type TestOpts struct {
	Message string
	Fail    bool
	Delay   time.Duration
	Port    int
	Route   string
}

// TestServer starts a test server that will respond to requests with the given message.
// paramaters: TestOpts
// message: string
//
//	the message to respond with
//
// fail: bool
//
//	if true, the server will respond with a 500 error
//
// delay: time.Duration
//
//	the amount of time to wait before responding
//
// port: int
//
//	the port to listen on
//
// root: string
//
//	the root path to listen on
//
// returns:
//
//	error if the server fails to start
//
// example:
//
//	err := TestServer(TestOpts{
//	  "hello",
//	  false,
//	  time.Duration(time.Duration(10).Seconds()),
//	  8080,
//	})
func TestServer(opts TestOpts) error {
	http.HandleFunc(opts.Route, func(w http.ResponseWriter, r *http.Request) {
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
				http.Error(
					w,
					"Server failed to respond",
					http.StatusInternalServerError,
				)
			}
		}
	})

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

	return nil
}
