package v1

import (
	"bodyshop/common"
	"bodyshop/models"
	"bodyshop/utils"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// send godoc
// @Summary Send message to mailsender
// @Description Construct message by template and params and send it to mailsender for receivers
// @Accept  json
// @Produce  json
// @Param incomingRequest body models.IncomingRequest true "New Incoming Request"
// @Success 200 {object} utils.response
// @Failure 500 {object} utils.response
// @Router /send [post]
func (a *App) send(w http.ResponseWriter, r *http.Request) {
	var inRequest models.IncomingRequest
	if err := json.NewDecoder(r.Body).Decode(&inRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}

	template, err := a.Logic.FormMessage(
		inRequest.TemplateId,
		inRequest.Parameters,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	// forming request to mailsender
	var mailSenderRequest models.MailsenderRequest
	mailSenderRequest.Receivers = inRequest.Receivers
	mailSenderRequest.Message = template

	payload, err := json.Marshal(mailSenderRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	// sending request to mailsender
	resp, err := a.MailsenderClient.Post(
		common.ServiceConfig.MailsenderConnection+"/api/v1/send",
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		// TODO: error for bad communication within services
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, errors.New("bad response from mailsender"))
		return
	}

	w.WriteHeader(http.StatusOK)
	utils.SerializeResponseJSON(w, "successfully send message")
}
