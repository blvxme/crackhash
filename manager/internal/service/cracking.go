package service

import (
	"context"
	gUuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"manager/internal/common/dto"
	"manager/internal/data/repo/redis"
	"time"
)

func HandleCrackingRequest(crackingRequest dto.CrackingRequest, ctx context.Context) (uuid string, err error) {
	uuid = gUuid.NewString()
	log.Infof("Generated UUID: %s\n", uuid)

	requestStatus := dto.RequestStatus{Status: dto.StatusInProgress, Data: nil}

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

	if err = requestStatusRepo.Create(uuid, requestStatus, ctx); err != nil {
		log.Errorf("Failed to save request status: %v\n", err)
		return
	}

	_ = time.AfterFunc(1*time.Hour, func() {
		ok, err := revokeTask(uuid)
		if err != nil {
			log.Errorf("Failed to revoke request with ID %s: %v\n", uuid, err)
		} else if ok {
			log.Infof("The request with ID %s has been assigned the status ERROR\n", uuid)
		}
	})

	go assignTasks(crackingRequest, uuid)

	return
}

func revokeTask(uuid string) (ok bool, err error) {
	ctx := context.Background()

	requestStatusRepo, err := redis.NewRequestStatusRepo(ctx)
	if err != nil {
		return
	}

	defer func() {
		if err = redis.DestroyRequestStatusRepo(requestStatusRepo); err != nil {
			return
		}
	}()

	requestStatus, err := requestStatusRepo.Read(uuid)
	if err != nil {
		return
	}

	if requestStatus.Status != dto.StatusInProgress {
		return
	}

	requestStatus.Status = dto.StatusError
	if err = requestStatusRepo.Update(uuid, requestStatus); err != nil {
		return
	}

	ok = true

	return
}
