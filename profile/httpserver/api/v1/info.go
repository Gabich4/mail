package v1

import (
	"io"
	"io/ioutil"
	"net/http"
	"profile/common"
	"profile/utils"
)

// info godoc
// @Summary Get info about current user
// @Description Get info about current user from auth by token
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /i [get]
func info(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8000/me", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}
	// copy cookies
	utils.CopyCookies(*r, req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		common.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			common.Logger.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.SerializeResponseJSON(w, err)
		return
	}

	w.WriteHeader(resp.StatusCode)
	utils.SerializeResponseJSON(w, body)
}
