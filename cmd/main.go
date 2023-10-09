package main

import (
	"log"
	"net/http"
)

const (
	Port = ":8080"
)

func main() {
	http.HandleFunc("/hello-go", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello from Go http server!"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("can't write to response writer %v\n", err)
			return
		}
		log.Println("message sent to the client")
	})

	log.Printf("App started and listening on %s\n", Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		log.Fatalf("can't listen and serve on %s\n", Port)
	}
}
