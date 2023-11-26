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
func TestServer(opts TestOpts) error {

	fmt.Println("server started")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got request")
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

func main() {
	TestServer(TestOpts{
		"hello",
		false,
		time.Duration(2 * time.Second),
		8085,
	})
}
