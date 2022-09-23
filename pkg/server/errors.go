package server

import (
	"encoding/json"
	"github.com/ketch-com/go-ketch-forwarder/pkg/types"
	"log"
	"net/http"
)

func WriteError(w http.ResponseWriter, metadata *types.Metadata, statusCode int, status string, message string) {
	var err error
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := new(types.Response)
	response.ApiVersion = types.ApiVersion
	response.Kind = types.ErrorKind
	response.Metadata = metadata
	response.Error, err = json.Marshal(&types.ErrorBody{
		Code:    statusCode,
		Status:  status,
		Message: message,
	})
	if err != nil {
		log.Println("failed to write error response")
		return
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println("failed to write error response")
		return
	}
}
