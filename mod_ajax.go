package main

import (
	"net/http"
)

func ajax(w http.ResponseWriter, r *http.Request) {
}

func init() {
	http.HandleFunc("/ajax", ajax)
}
