package controllers

import (
	"net/http"

	"github.com/kyleu/dbui/internal/app/util"
	template "github.com/kyleu/dbui/internal/gen/templates"
)

func Health(w http.ResponseWriter, r *http.Request) {
	act(w, r, "Routes", func(ctx util.RequestContext) (int, error) {
		_, _ = w.Write([]byte("OK"))
		return 0, nil
	})
}

func Modules(w http.ResponseWriter, r *http.Request) {
	act(w, r, "Routes", func(ctx util.RequestContext) (int, error) {
		return template.ModulesList(ctx, w)
	})
}

func Routes(w http.ResponseWriter, r *http.Request) {
	act(w, r, "Routes", func(ctx util.RequestContext) (int, error) {
		return template.RoutesList(ctx, w)
	})
}
