package handlers

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/patrickdappollonio/division-lfg/base/helpers/render"
)

func GetHome(_ context.Context, w http.ResponseWriter, r *http.Request) {
	render.Template.HTML(w, http.StatusOK, "home", nil)
}
