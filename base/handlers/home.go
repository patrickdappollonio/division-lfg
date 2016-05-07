package handlers

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/patrickdappollonio/division-lfg/base/helpers/render"
	"github.com/patrickdappollonio/division-lfg/base/models/binding"
)

func GetHome(_ context.Context, w http.ResponseWriter, r *http.Request) {
	bnd := binding.GetDefault(r)
	render.Template.HTML(w, http.StatusOK, "home", bnd)
}
