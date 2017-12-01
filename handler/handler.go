package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/foecum/digitalocean-client/droplet"
)

/*
	// h.Mux.HandleFunc("/list/droplet", sayHello)
	// h.Mux.HandleFunc("/rename/droplet", sayHello)
	// h.Mux.HandleFunc("/reboot/droplet", sayHello)
	// h.Mux.HandleFunc("/poweroff/droplet", sayHello)
	// h.Mux.HandleFunc("/rebuild/droplet", sayHello)
	// h.Mux.HandleFunc("/password/reset/droplet", sayHello)
	// h.Mux.HandleFunc("/delete/droplet", sayHello)
	// h.Mux.HandleFunc("/resize/droplet", sayHello)
	// h.Mux.HandleFunc("/user/info", sayHello)
*/

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"response"`
}

// Handler ...
type Handler struct {
	Mux     *http.ServeMux
	droplet *droplet.Droplet
}

// New created a server mux
func (h *Handler) New(accessToken string) {
	h.Mux = http.NewServeMux()

	h.droplet = &droplet.Droplet{}
	h.droplet.RegisterClient(accessToken)

	h.Mux.HandleFunc("/create/droplet", h.createDroplet)
	h.Mux.HandleFunc("/list/regions", h.getRegions)
}

func (h *Handler) createDroplet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *Handler) getRegions(w http.ResponseWriter, r *http.Request) {
	regions, err := h.droplet.GetRegions()

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		data := err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not get regions: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{Data: regions, Success: err == nil})
	if err != nil {
		log.Printf("could not encode regions: %v", err)
	}
}
