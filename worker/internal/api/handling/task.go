package handling

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"worker/internal/common/dto"
	"worker/internal/service"
)

func HandleTaskRequest(w http.ResponseWriter, r *http.Request) {
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

	var taskRequest dto.TaskRequest
	if err := json.NewDecoder(r.Body).Decode(&taskRequest); err != nil {
		log.Errorf("Failed to decode request body: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	go service.HandleTaskRequest(taskRequest)
}
