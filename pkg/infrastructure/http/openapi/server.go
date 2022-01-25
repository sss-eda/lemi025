package openapi

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/sss-eda/lemi025/pkg/application"
)

// Server TODO
type Server struct {
	service    application.Server
	wsUpgrader websocket.Upgrader
}

// NewServer TODO
func NewServer(
	service application.Server,
) (*Server, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	server := Server{
		service:    service,
		wsUpgrader: upgrader,
	}

	return &server, nil
}

// Serve TODO
func (server *Server) Serve(
	port uint64,
) error {
	router := mux.NewRouter().PathPrefix("/lemi025").Subrouter()

	router.HandleFunc("/commands/readConfig", server.readConfig)
	router.HandleFunc("/commands/readTime", server.readTime)
	router.HandleFunc("/events/configRead", server.configRead)
	// router.HandleFunc("/events/timeRead", server.timeRead)

	return http.ListenAndServe(strconv.FormatUint(port, 10), router)
}

func (server *Server) readConfig(
	w http.ResponseWriter,
	r *http.Request,
) {
	request :=

		server.service.ReadConfig(request.ID)

	return
}

func (server *Server) configRead(
	w http.ResponseWriter,
	r *http.Request,
) {
	ws, err := server.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Print(messageType)
		fmt.Println(p)
	}
}

func (server *Server) readTime(
	w http.ResponseWriter,
	r *http.Request,
) {
	return
}
