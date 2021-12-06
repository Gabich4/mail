package v1

import (
	"encoding/json"
	"net/http"
	"profile/models"
	"profile/utils"
)

func receiveStatus(w http.ResponseWriter, r *http.Request) {
	var sendRequestStatus models.SendRequestStatus
	err := json.NewDecoder(r.Body).Decode(&sendRequestStatus)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}
	utils.SerializeResponseJSON(w, sendRequestStatus)
}
