package v1

import (
	"io/ioutil"
	"net/http"
	"profile/common"
	"profile/utils"

	"github.com/go-chi/chi"
)

// createReceiversOnUser godoc
// @Summary Create receivers list for user
// @Accept  plain
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /receivers/{user_id} [post]
func (a *App) createReceiversOnUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user_id")
	receivers := ""
	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	} else {
		receivers = string(bytes)
	}

	if err := a.Logic.Create(username, receivers); err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}
}

// readReceiversOnUser godoc
// @Summary Get receivers list by user
// @Produce json
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /receivers/{user_id} [get]
func (a *App) readReceiversOnUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user_id")
	receivers, err := a.Logic.Read(username)
	if err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusNotFound)
		utils.SerializeResponseJSON(w, err)
		return
	}

	utils.SerializeResponseJSON(w, receivers)
}

// updateReceiversOnUser godoc
// @Summary Update receivers list by user
// @Accept plain
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /receivers/{user_id} [put]
func (a *App) updateReceiversOnUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user_id")
	receivers := ""
	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	} else {
		receivers = string(bytes)
	}

	if err := a.Logic.Update(username, receivers); err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}
}

// deleteReceiversOnUser godoc
// @Summary Delete receivers list by user
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /receivers/{user_id} [delete]
func (a *App) deleteReceiversOnUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "user_id")
	if err := a.Logic.Delete(username); err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}
}

// readAllReceivers godoc
// @Summary Get receivers list for all users
// @Produce json
// @Success 200 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /receivers/ [get]
func (a *App) readAllReceivers(w http.ResponseWriter, r *http.Request) {
	allReceivers, err := a.Logic.ReadAll()
	if err != nil {
		common.Logger.Print(err)
		w.WriteHeader(http.StatusNotFound)
		utils.SerializeResponseJSON(w, err)
		return
	}

	utils.SerializeResponseJSON(w, allReceivers)
}
