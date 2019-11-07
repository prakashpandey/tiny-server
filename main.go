package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 8181, "Port at which the server will run.")
	address := fmt.Sprintf(":%d", *port)
	dir := flag.String("dir", ".", "Default directory from which file wil be served.")
	fmt.Printf("Serving file from: %s directory.\n", *dir)
	fmt.Printf("Running server at: %s\n", address)
	if err := http.ListenAndServe(address, http.FileServer(http.Dir(*dir))); err != nil {
		panic(err)
	}
}
