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
//
// Returns
//	- error if the server fails to start
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
