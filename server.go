package main

import (
	"fmt"
	"log"
	"net/http"
)

func accountVerification(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	fmt.Println("Starting server in port 443")
	http.HandleFunc("/v1/accounts", accountVerification)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
