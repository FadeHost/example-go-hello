// A hello-world web app in Go, small enough to read in one go.
//
// It listens on the port FadeHost gives it in PORT, and once it is up it
// asks itself for the page and prints the answer, so the console shows the
// app really is serving and not just starting.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello from %s, hosted on FadeHost.\n", runtime.Version())
	})

	go selfCheck(port)

	log.Printf("listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Ask ourselves for the page until the listener is up, then say what came
// back. Twenty tries at a quarter of a second is five seconds, which is far
// longer than a Go program needs to start.
func selfCheck(port string) {
	for i := 0; i < 20; i++ {
		time.Sleep(250 * time.Millisecond)
		resp, err := http.Get("http://127.0.0.1:" + port + "/")
		if err != nil {
			continue
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		log.Printf("self check: HTTP %d, %q", resp.StatusCode, strings.TrimSpace(string(body)))
		return
	}
	log.Printf("self check: nothing answered on port %s", port)
}
