package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"manager/internal/common/dto"
	"manager/internal/data/repo/redis"
)

func HandleStatusRequest(requestId string, ctx context.Context) (requestStatus dto.RequestStatus, err error) {
	requestStatusRepo, err := redis.NewRequestStatusRepo(ctx)
	if err != nil {
		log.Errorf("Failed to create RequestStatusRepo: %v\n", err)
		return
	}

	defer func() {
		if err = redis.DestroyRequestStatusRepo(requestStatusRepo); err != nil {
			log.Errorf("Failed to destroy RequestStatusRepo: %v\n", err)
		}
	}()

	requestStatus, err = requestStatusRepo.Read(requestId, ctx)
	if err != nil {
		log.Errorf("Failed to read request status: %v\n", err)
		return
	}

	return
}
