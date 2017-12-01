package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/foecum/digitalocean-client/client"
	"github.com/foecum/digitalocean-client/droplet"
	"github.com/foecum/digitalocean-client/user"
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
	user    *user.User
}

// New created a server mux
func (h *Handler) New(accessToken string) {
	h.Mux = http.NewServeMux()
	client := client.RegisterClient(accessToken)
	h.droplet = &droplet.Droplet{}
	h.user = &user.User{}

	h.droplet.Client = client
	h.user.Client = client

	h.Mux.HandleFunc("/list/droplets", h.getDropletOrRegionList)
	h.Mux.HandleFunc("/list/regions", h.getDropletOrRegionList)
	h.Mux.HandleFunc("/user/info", h.getUserInfo)
}

func (h *Handler) getUserInfo(w http.ResponseWriter, r *http.Request) {
	data, err := h.user.GetUserInfo()

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		data := err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not get droplets: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
	if err != nil {
		log.Printf("could not encode droplets: %v", err)
	}

}

func (h *Handler) getDropletOrRegionList(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error

	params := strings.Split(r.URL.Path, "/")
	switch params[len(params)-1] {
	case "droplets":
		data, err = h.droplet.GetDroplets()
	case "regions":
		data, err = h.droplet.GetRegions()
	}

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		data := err.Error()
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
		if err != nil {
			log.Printf("could not get droplets: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response{Data: data, Success: err == nil})
	if err != nil {
		log.Printf("could not encode droplets: %v", err)
	}
}
