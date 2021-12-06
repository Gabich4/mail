package utils

import (
	"encoding/json"
	"net/http"
	"profile/models"
)

// SerializeReponseJSON serializes data into response structure.
func SerializeResponseJSON(w http.ResponseWriter, data interface{}) {
	var resp models.Response
	switch data := data.(type) {
	case error:
		resp.Error = data.Error()
	default:
		resp.Data = data
		resp.Success = true
	}

	_ = json.NewEncoder(w).Encode(resp)
}
