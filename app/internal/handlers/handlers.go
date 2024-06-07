package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func URLParam(r *http.Request, name string) string {
	ctx := r.Context()

	params := ctx.Value("params").(map[string]string)
	return params[name]
}

func URLParamInt(r *http.Request, name string) (int, error) {
	ctx := r.Context()

	params := ctx.Value("params").(map[string]string)
	param, found := params[name]
	if !found {
		return 0, fmt.Errorf("param %s not found", name)
	}
	i, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return i, nil
}
