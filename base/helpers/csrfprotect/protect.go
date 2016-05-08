package csrfprotect

import (
	"net/http"

	"github.com/justinas/nosurf"
)

var Inject, Verify func(http.Handler) http.Handler
var allowedMethods = []string{"GET", "HEAD", "OPTIONS", "TRACE"}

func init() {
	Inject = nosurf.NewPure

	Verify = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !contains(allowedMethods, r.Method) {
				var tokenRetrieved string

				if tkn := r.Header.Get("X-CSRF-Token"); tkn != "" {
					tokenRetrieved = tkn
				} /*else if tkn := r.FormValue("csrf_token"); tkn != "" {
					tokenRetrieved = tkn
				}*/

				tokenOriginal := nosurf.Token(r)

				if !nosurf.VerifyToken(tokenOriginal, tokenRetrieved) {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusBadRequest)
					return
				}
			}

			next.ServeHTTP(w, r)
		})
	}
}

func contains(vals []string, s string) bool {
	for _, v := range vals {
		if v == s {
			return true
		}
	}

	return false
}
