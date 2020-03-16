package controllers

import (
	web "github.com/kyleu/dbui/internal/app/web"
	"net/http"

	"emperror.dev/errors"

	"github.com/gorilla/mux"
	"github.com/kyleu/dbui/internal/gen/templates"
)

var _sandboxes = []string{"gallery", "testbed"}

func SandboxList(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		ctx.Title = "Sandboxes"
		ctx.Breadcrumbs = web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox")
		return templates.SandboxList(_sandboxes, ctx, w)
	})
}

func SandboxForm(w http.ResponseWriter, r *http.Request) {
	act(w, r, func(ctx web.RequestContext) (int, error) {
		key := mux.Vars(r)["key"]
		if key == "testbed" {
			return 0, errors.WithStack(errors.New("error!"))
		}
		ctx.Title = "[" + key + "] Sandbox"
		bc := web.Breadcrumb{Path: ctx.Route("sandbox.run", "key", key), Title: key}
		ctx.Breadcrumbs = append(web.BreadcrumbsSimple(ctx.Route("sandbox"), "sandbox"), bc)
		return templates.SandboxForm(key, ctx, w)
	})
}
