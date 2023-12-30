package functions

import (
	"net/http"

	"github.com/richardboase/npgpublic/models"
	"github.com/richardboase/npgpublic/sdk/cloudfunc"
	"github.com/richardboase/npgpublic/sdk/common"
	"github.com/richardboase/npgpublic/utils"
)

// api-user
func UserEntrypoint(w http.ResponseWriter, r *http.Request) {

	if cloudfunc.HandleCORS(w, r, "*") {
		return
	}

	app := common.NewApp()
	app.UseGCP("{{.ProjectID}}")
	app.UseGCPFirestore("{{.DatabaseID}}")

	function, err := cloudfunc.QueryParam(r, "function")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	userID, err := cloudfunc.QueryParam(r, "id")
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	user := &models.User{}
	if err := utils.GetDocument(app, userID, user); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	switch r.Method {

	case "GET":

		switch function {

		case "username":

			if err := cloudfunc.ServeJSON(w, user.Username); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		case "object":

			if err := cloudfunc.ServeJSON(w, user); err != nil {
				cloudfunc.HttpError(w, err, http.StatusInternalServerError)
				return
			}

			return

		}
	}
}
