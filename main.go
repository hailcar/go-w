package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", Timer)
	http.ListenAndServe(":8080", nil)
}
func Timer(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	w.Write([]byte(now.Format("2006-01-02 15:04:05")))
}
