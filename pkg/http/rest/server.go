package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/sss-eda/lemi025"
)

type Server struct {
	*http.Server
	driver lemi025.Driver
}

func NewServer(
	driver lemi025.Driver,
) *Server {
	mux := http.DefaultServeMux

	// OK. This is a problem, because this should be
	server := Server{
		Server: &http.Server{
			Addr:           ":8080",
			Handler:        mux,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
		driver: driver,
	}

	mux.HandleFunc("/command", server.readConfig)

	return &server
}

func (server *Server) readConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)

	err := server.driver.ReadConfig(lemi025.ReadConfigCommand{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err2 := encoder.Encode(err)
		if err2 != nil {
			log.Println(err2)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println(err)
	}

	return
}

func (server *Server) readTime(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)

	err := server.driver.ReadTime(lemi025.ReadTimeCommand{})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err2 := encoder.Encode(err)
		if err2 != nil {
			log.Println(err2)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println(err)
	}

	return
}

func (server *Server) setTime(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	command := lemi025.SetTimeCommand{}

	err := decoder.Decode(&command)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err2 := encoder.Encode(err)
		if err2 != nil {
			log.Println(err2)
		}
		return
	}

	err = server.driver.SetTime(command)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err2 := encoder.Encode(err)
		if err2 != nil {
			log.Println(err2)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println(err)
	}

	return
}
