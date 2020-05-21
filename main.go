package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

var (
	dir  *string
	port *int
)

func init() {
	port = flag.Int("port", 8181, "Port at which the server will run.")
	dir = flag.String("dir", ".", "Default directory from which file wil be served.")
	flag.Parse()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	download := strings.TrimSpace(r.URL.Query().Get("download"))
	if download == "true" {
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", r.URL.Path))
	}
	http.FileServer(http.Dir(*dir)).ServeHTTP(w, r)
}

func main() {
	// mustInitFlags()
	address := fmt.Sprintf(":%d", *port)
	fmt.Printf("Serving file from: %s directory.\n", *dir)
	fmt.Printf("Running server at: %s\n", address)
	if err := http.ListenAndServe(address, http.HandlerFunc(rootHandler)); err != nil {
		panic(err)
	}
}
