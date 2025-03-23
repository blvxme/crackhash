package handling

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"manager/internal/service"
	"net/http"
)

func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Errorf("Invalid HTTP method: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	params := r.URL.Query()
	requestId := params.Get("requestId")
	if requestId == "" {
		log.Errorf("Invalid query parameter: requestId=\"%s\"\n", requestId)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	requestStatus, err := service.HandleStatusRequest(requestId, r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(requestStatus); err != nil {
		log.Errorf("Failed to encode response: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Infof("Sending response to %s: %v\n", r.RemoteAddr, requestStatus)
}
