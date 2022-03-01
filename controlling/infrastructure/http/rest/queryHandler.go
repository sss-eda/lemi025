package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sss-eda/lemi025/controlling"
)

// HandleQuery - Handle an incoming query
//   -> The responses for any control query will look the same.
//   -> The requests will look different depending on the usecase.
func HandleQuery[Request controlling.RequestType](
	handle func(func(*controlling.Response) error, *Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := Request{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if request.QueryName != Request.

		handle(func(response *controlling.Response) error {
			return json.NewEncoder(w).Encode(response)
		}, &request)
	}
}
