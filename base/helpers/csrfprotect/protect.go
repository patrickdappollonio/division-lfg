package csrfprotect

import (
	"net/http"

	"github.com/gorilla/csrf"
)

var Inject, Verify func(http.Handler) http.Handler

func init() {
	Verify = csrf.Protect(
		[]byte("32-byte-long-auth-key"),
		csrf.CookieName("__csrf_protect"),
	)

	Inject = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-CSRF-Token", csrf.Token(r))
			next.ServeHTTP(w, r)
		})
	}
}
