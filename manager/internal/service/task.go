package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"manager/internal/common/config"
	"manager/internal/common/dto"
	"manager/internal/data/repo/redis"
	"net/http"
	"strconv"
)

func HandleTaskResponse(taskResponse dto.TaskResponse, ctx context.Context) (err error) {
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

	requestStatus, err := requestStatusRepo.Read(taskResponse.RequestId, ctx)
	if err != nil {
		log.Errorf("Failed to read request status: %v\n", err)
		return
	}

	requestStatus.Status = dto.StatusReady
	requestStatus.Data = append(requestStatus.Data, taskResponse.Data...)
	if err = requestStatusRepo.Update(taskResponse.RequestId, requestStatus, ctx); err != nil {
		log.Errorf("Failed to update request status: %v\n", err)
	}

	return
}

func assignTasks(crackingRequest dto.CrackingRequest, uuid string) {
	nWorkersStr := config.GetNWorkers()
	nWorkers, err := strconv.Atoi(nWorkersStr)
	if err != nil {
		log.Panicf("Invalid number of workers (%s): %v\n", nWorkersStr, err)
	}
	partCount := nWorkers

	alphabet := config.GetAlphabet()

	for i := 0; i < partCount; i++ {
		taskRequest := dto.TaskRequest{
			RequestId:  uuid,
			Alphabet:   alphabet,
			Hash:       crackingRequest.Hash,
			MaxLength:  crackingRequest.MaxLength,
			PartNumber: i,
			PartCount:  partCount,
		}

		if err = sendTaskRequest(taskRequest, i+1); err != nil {
			log.Errorf("Failed to send task request: %v\n", err)
		}
	}
}

func sendTaskRequest(taskRequest dto.TaskRequest, workerId int) (err error) {
	addr := fmt.Sprintf(
		"http://%s-%d:%s/internal/api/worker/hash/crack/task",
		config.GetWorkerHost(), workerId,
		config.GetWorkerPort(),
	)
	log.Infof("Sending task request to %s\n", addr)

	taskRequestJson, err := json.Marshal(taskRequest)
	if err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, addr, bytes.NewBuffer(taskRequestJson))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Errorf("Failed to close response body: %v\n", err)
		}
	}()

	log.Infof("Response status from %s: %d\n", addr, resp.StatusCode)

	return
}
