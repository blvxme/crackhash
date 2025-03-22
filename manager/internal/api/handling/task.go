package handling

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"manager/internal/common/dto"
	"manager/internal/service"
	"net/http"
)

func HandleTaskResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		log.Errorf("Invalid HTTP method: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.Errorf("Failed to close request body: %v\n", err)
		}
	}()

	var taskResponse dto.TaskResponse
	if err := json.NewDecoder(r.Body).Decode(&taskResponse); err != nil {
		log.Errorf("Failed to decode request body: %v\n", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := service.HandleTaskResponse(taskResponse, r.Context()); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
