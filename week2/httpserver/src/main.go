package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	var VERSION string
	VERSION = os.Getenv("VERSION")
	// fmt.Println(VERSION)
	writer.Header().Add("VERSION", VERSION)
	for tag, v := range request.Header {
		for _, s := range v {
			writer.Header().Add(tag, s)
		}
	}
	writer.WriteHeader(200)
	// for tag, v := range writer.Header() {
	// 	fmt.Println(tag, v)
	// }
	fmt.Printf("source ip is: %s \n", request.Host)
	fmt.Printf("status is: %s \n", "200")
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8080", nil)
}
