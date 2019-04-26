package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	user := User{[]string{"ONE", "Pink", "Blue", "Yellow"}, message}

	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func ConnectServer2(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	url := "http://godockertwo:8000/ " + message
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func main() {
	// http.HandleFunc("/", sayHello)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	panic(err)
	// }
	http.HandleFunc("/", ConnectServer2)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
