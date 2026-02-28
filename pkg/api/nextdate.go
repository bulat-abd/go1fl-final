package api

import (
	"net/http"
	"time"

	"github.com/bulat-abd/go1fl-final/pkg/datecalc"
	log "github.com/sirupsen/logrus"
)

func nextDateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		nowstr := req.FormValue("now")
		now, _ := time.Parse("20060102", nowstr)
		date := req.FormValue("date")
		repeat := req.FormValue("repeat")
		if repeat == "" {
			JsonError(w, "No repeat value provided", http.StatusBadRequest)
			return
		}
		result, err := datecalc.NextDate(now, date, repeat)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(result))
		if err != nil {
			log.Error("error in nextDateHandler while writing response:", err)
		}
	}
}
