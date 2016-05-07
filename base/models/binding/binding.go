package binding

import (
	"net/http"

	"github.com/justinas/nosurf"
)

type KV map[string]interface{}

func GetDefault(r *http.Request) KV {
	return KV{
		"CSRFToken": nosurf.Token(r),
	}
}
