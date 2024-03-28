package service

import "net/http"

func Init() {
	http.HandleFunc("/http/Register", new(UserImpl).Register)
}
