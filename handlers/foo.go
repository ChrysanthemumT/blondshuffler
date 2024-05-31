package handlers

import "net/http"

func HandlerFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))
}
