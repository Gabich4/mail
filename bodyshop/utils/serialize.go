package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// SerializeReponseJSON serializes data into response structure.
func SerializeResponseJSON(w http.ResponseWriter, data interface{}) {
	var resp response
	switch data := data.(type) {
	case error:
		resp.Error = data.Error()
	default:
		resp.Data = data
		resp.Success = true
	}

	_ = json.NewEncoder(w).Encode(resp)
}
