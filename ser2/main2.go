package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	Favorite []string
	Msg      string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	user := User{[]string{"TWO", "Blue", "Yellow", "Black"}, message}

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
