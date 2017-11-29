package handler

import (
	"fmt"
	"net/http"
)

// New created a server mux
func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create/droplet", sayHello)
	mux.HandleFunc("/list/droplet", sayHello)
	mux.HandleFunc("/rename/droplet", sayHello)
	mux.HandleFunc("/reboot/droplet", sayHello)
	mux.HandleFunc("/poweroff/droplet", sayHello)
	mux.HandleFunc("/rebuild/droplet", sayHello)
	mux.HandleFunc("/password/reset/droplet", sayHello)
	mux.HandleFunc("/delete/droplet", sayHello)
	mux.HandleFunc("/resize/droplet", sayHello)
	mux.HandleFunc("/user/info", sayHello)
	mux.HandleFunc("/list/regions", sayHello)

	return mux
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
