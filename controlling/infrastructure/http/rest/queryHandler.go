package rest

import (
	"net/http"

	"github.com/sss-eda/lemi025/controlling/application/readconfig"
	"github.com/sss-eda/lemi025/controlling/application/readtime"
	"github.com/sss-eda/lemi025/controlling/application/settime"
)

// Request TODO
type Request struct {
	Type    string
	Message []byte
}

// // Response TODO
// type Response struct {
// 	Error error
// }

// HandleQueries - Handle an incoming query
//   -> The responses for any control query will look the same.
//   -> The requests will look different depending on the usecase.
func HandleQueries(
	readConfig readconfig.HandlerFunc,
	readTime readtime.HandlerFunc,
	setTime settime.HandlerFunc,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.
	}
}
