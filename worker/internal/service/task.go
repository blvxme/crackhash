package service

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"worker/internal/common/config"
	"worker/internal/common/dto"
	"worker/pkg/seqgen"
)

func HandleTaskRequest(taskRequest dto.TaskRequest) {
	alphabetLen := len(taskRequest.Alphabet)

	totalNWords := int64(0)
	for i := 1; i <= taskRequest.MaxLength; i++ {
		totalNWords += pow(alphabetLen, i)
	}

	offset := int64(taskRequest.PartNumber) * (totalNWords / int64(taskRequest.PartCount))

	nWords := int64(0)
	if taskRequest.PartNumber < taskRequest.PartCount-1 {
		nWords = totalNWords / int64(taskRequest.PartCount)
	} else {
		nWords = totalNWords - int64(taskRequest.PartCount-1)*(totalNWords/int64(taskRequest.PartCount))
	}

	sg := seqgen.New(taskRequest.Alphabet, taskRequest.MaxLength, offset, nWords)

	data := make([]string, 0)

	for i := int64(0); i < nWords; i++ {
		nextWord, err := sg.Next()
		if err != nil {
			log.Errorf("Failed to generate next word: %v\n", err)
			break
		}

		hash := md5.Sum([]byte(nextWord))
		hashStr := hex.EncodeToString(hash[:])

		if taskRequest.Hash == hashStr {
			log.Infof("Found matching word: %s (hash: %s)\n", nextWord, hashStr)
			data = append(data, nextWord)
		}
	}

	log.Infof("The word search is complete. Words found: %d (%v)\n", len(data), data)

	if len(data) > 0 {
		taskResponse := dto.TaskResponse{RequestId: taskRequest.RequestId, Data: data}
		sendTaskResponse(taskResponse)
	}
}

func pow(a, b int) (result int64) {
	result = 1
	for i := 1; i <= b; i++ {
		result *= int64(a)
	}

	return
}

func sendTaskResponse(taskResponse dto.TaskResponse) {
	jsonResponse, err := json.Marshal(taskResponse)
	if err != nil {
		log.Errorf("Failed to encode response: %v\n", err)
		return
	}

	addr := fmt.Sprintf(
		"http://%s:%s/internal/api/manager/hash/crack/request",
		config.GetManagerHost(), config.GetManagerPort(),
	)

	req, err := http.NewRequest(http.MethodPatch, addr, bytes.NewBuffer(jsonResponse))
	if err != nil {
		log.Errorf("Failed to create request: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Failed to send request: %v\n", err)
		return
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Errorf("Failed to close response body: %v\n", err)
		}
	}()

	log.Infof("Response status from %s: %d\n", addr, resp.StatusCode)
}
