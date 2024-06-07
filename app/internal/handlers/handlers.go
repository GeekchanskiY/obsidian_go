package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/lib/pq"
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

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	// TODO: Divide error logs to DEBUG and non-DEBUG for outcoming data
	log.Println(err)
	if err, ok := err.(*pq.Error); ok {
		log.Printf("psql error: %s, %s \n", err.Code, err.Message)
		if err.Code == "23503" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Cant delete this item, it contains depended object"))
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Database Error"))
			return
		}
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Error"))
}
