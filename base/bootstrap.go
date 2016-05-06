package base

import (
	"github.com/patrickdappollonio/division-lfg/base/handlers"
	"github.com/patrickdappollonio/division-lfg/base/helpers/csrfprotect"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	Router.Use(
		// canonical.Enforce(config.DomainName, config.ShouldUseSSL),
		middleware.Logger,
		middleware.Recoverer,
		csrfprotect.Inject,
		csrfprotect.Verify,
	)

	Router.Get("/", handlers.GetHome)
	Router.Post("/test", handlers.Test)
}
