package handlers

import (
	"net/http"
)

func URLParam(r *http.Request, name string) string {
	ctx := r.Context()

	params := ctx.Value("params").(map[string]string)
	return params[name]
}
