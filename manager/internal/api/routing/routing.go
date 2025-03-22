package routing

import (
	log "github.com/sirupsen/logrus"
	"manager/internal/api/handling"
	"net/http"
)

func SetUpRouting() (handler http.Handler) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hash/crack", handling.HandleCrackingRequest)
	mux.HandleFunc("/api/hash/status", handling.HandleStatusRequest)
	mux.HandleFunc("/internal/api/manager/hash/crack/request", handling.HandleTaskResponse)

	handler = wrapWithLogging(mux)

	return
}

func Run(handler http.Handler) (err error) {
	log.Infof("Starting server at port 8080\n")

	err = http.ListenAndServe(":8080", handler)

	return
}

func wrapWithLogging(next http.Handler) (handler http.Handler) {
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("New request: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})

	return
}
