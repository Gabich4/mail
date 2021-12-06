package v1

import (
	"bytes"
	"encoding/json"
	"errors"
	"mailsender/common"
	"mailsender/models"
	"mailsender/utils"

	"net/http"
)

// Send godoc
// @Summary Inserts message in repository and sends status to profile
// @Description Inserts message with waiting status to the connected repository,
// @Description sends it to profile service, updates status with one randomly chosen,
// @Description finally, sends status to profile
// @Accept json
// @Produce json
// @Param incomingRequest body models.IncomingSendRequest true "Incoming Request"
// @Success 200 {object} models.Response
// @Failure 500 {object} models.Response
// @Failure default {object} models.Response
// @Router /send [post]
func (a *App) send(w http.ResponseWriter, r *http.Request) {
	// validation of incoming request data
	var inRequest models.IncomingSendRequest
	if err := json.NewDecoder(r.Body).Decode(&inRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// inserting message to collection
	id, err := a.Logic.InsertMessage(inRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// forming post request
	var outRequest models.OutgoingStatusRequest
	outRequest.Message = inRequest.Message
	outRequest.Receivers = inRequest.Receivers
	outRequest.Status = models.WaitingStatus
	outJson, err := json.Marshal(outRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	// sending post-request to profile
	resp, err := a.ProfileClient.Post(
		common.ServiceConfig.ProfileConnection+"/status",
		"application/json",
		bytes.NewBuffer(outJson),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// check that response status is ok
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, errors.New("profile response status is not OK"))
		return
	}

	status := utils.GetRandomStatus()
	if err := a.Logic.UpdateMessage(id, status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	outRequest.Status = status
	outJson, err = json.Marshal(outRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	resp, err = a.ProfileClient.Post(
		common.ServiceConfig.ProfileConnection+"/status",
		"application/json",
		bytes.NewBuffer(outJson),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, errors.New("profile response status is not OK"))
		return
	}

	utils.SerializeResponseJSON(w, "successfully inserted message")
}
