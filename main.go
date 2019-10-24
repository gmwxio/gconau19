package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/wxio/okmod"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s %T", time.Now(), okmod.HW{})
}

func main() {
	http.HandleFunc("/", greet)
	li, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("listening on http://%s:%d\n", "localhost", li.Addr().(*net.TCPAddr).Port)
	err = http.Serve(li, nil)
	log.Fatalf("%v", err)
}
