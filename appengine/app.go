package appengine

import (
	"net/http"

	"github.com/patrickdappollonio/division-lfg/base"
)

func init() {
	http.Handle("/", base.Router)
}
