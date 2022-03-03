package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sss-eda/lemi025/controlling"
)

// CommandRequest TODO
type CommandRequest[Command controlling.Commands] struct {
	Command Command `json:"command"`
}

// CommandSuccessResponse TODO
type CommandSuccessResponse struct{}

// CommandFailureResponse TODO
type CommandFailureResponse struct {
	Error error `json:"error"`
}

// HandleCommand TODO
func HandleCommand[Command controlling.Commands](
	handle func(*Command) error,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CommandRequest[Command]{}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			response, err2 := json.Marshal(CommandFailureResponse{
				Error: err,
			})
			if err2 != nil {
				log.Println(err2)
			}
			w.Write(response)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = handle(&request.Command)
		if err != nil {
			response, err2 := json.Marshal(CommandFailureResponse{
				Error: err,
			})
			if err2 != nil {
				log.Println(err2)
			}
			w.Write(response)
			w.WriteHeader(http.StatusInternalServerError)
		}

		response, err := json.Marshal(CommandSuccessResponse{})
		w.Write(response)
		w.WriteHeader(http.StatusOK)
	}
}
