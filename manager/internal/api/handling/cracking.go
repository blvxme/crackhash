package handling

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"manager/internal/common/dto"
	"manager/internal/service"
	"net/http"
)

func HandleCrackingRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Errorf("Invalid HTTP method: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Errorf("Failed to close request body: %v\n", err)
		}
	}()

	crackingRequest := dto.CrackingRequest{}
	if err := json.NewDecoder(r.Body).Decode(&crackingRequest); err != nil {
		log.Errorf("Failed to decode request body: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	uuid, err := service.HandleCrackingRequest(crackingRequest, r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	crackingResponse := dto.CrackingResponse{RequestId: uuid}
	if err = json.NewEncoder(w).Encode(crackingResponse); err != nil {
		log.Errorf("Failed to encode response: %v\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	log.Infof("Sending response to %s: %v\n", r.RemoteAddr, crackingResponse)
}
