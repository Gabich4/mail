package v1

import (
	"html/template"
	"io"
	"net/http"
	"profile/common"
	"profile/utils"
	"text/template/parse"
)

// uploadTemplate godoc
// @Summary Upload new template
// @Accept  plain
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /upload_template [post]
func uploadTemplate(w http.ResponseWriter, r *http.Request) {

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		common.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	} else if len(bytes) == 0 {
		common.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		utils.SerializeResponseJSON(w, err)
		return
	}

	// TODO: to logic
	tmpl := template.Must(template.New("").Parse(string(bytes)))
	sl := make([]string, 0, len(tmpl.Tree.Root.Nodes))
	m := make(map[string]bool, len(tmpl.Tree.Root.Nodes))
	for _, node := range tmpl.Tree.Root.Nodes {
		if node.Type() == parse.NodeAction {
			ns := node.String()
			if _, exist := m[ns]; !exist {
				m[ns] = true
				sl = append(sl, ns)
			}
		}
	}

	utils.SerializeResponseJSON(w, sl)
}
