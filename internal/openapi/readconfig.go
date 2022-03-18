package openapi

import (
	"context"
	"net/http"
)

// ReadConfigCommandHandlerFunc TODO
func ReadConfigCommandHandlerFunc(
	ctx context.Context,
	command ReadConfigCommand,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func 