package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"manager/internal/common/config"
	"manager/internal/common/dto"
	"strconv"
)

type RequestStatusRepo struct {
	client *redis.Client
}

func NewRequestStatusRepo(ctx context.Context) (repo *RequestStatusRepo, err error) {
	configMap := config.GetRedisConfigMap()
	host, port := configMap["host"], configMap["port"]
	password := configMap["password"]
	db, err := strconv.Atoi(configMap["db"])
	if err != nil {
		return
	}

	repo = &RequestStatusRepo{client: redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})}

	_, err = repo.client.Ping(ctx).Result()
	if err != nil {
		_ = repo.client.Close()
	}

	return
}

func DestroyRequestStatusRepo(repo *RequestStatusRepo) (err error) {
	err = repo.client.Close()
	return
}

func (repo *RequestStatusRepo) Create(uuid string, requestStatus dto.RequestStatus, opts ...interface{}) (err error) {
	if len(opts) < 1 {
		err = fmt.Errorf("context not provided")
		return
	}

	ctx, ok := opts[0].(context.Context)
	if !ok {
		err = fmt.Errorf("the additional argument must be of type context.Context")
		return
	}

	requestStatusJson, err := json.Marshal(requestStatus)
	if err != nil {
		return
	}

	err = repo.client.Set(ctx, uuid, string(requestStatusJson), 0).Err()

	return
}

func (repo *RequestStatusRepo) Read(uuid string, opts ...interface{}) (requestStatus dto.RequestStatus, err error) {
	if len(opts) < 1 {
		err = fmt.Errorf("context not provided")
		return
	}

	ctx, ok := opts[0].(context.Context)
	if !ok {
		err = fmt.Errorf("the additional argument must be of type context.Context")
		return
	}

	requestStatusJson, err := repo.client.Get(ctx, uuid).Result()
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(requestStatusJson), &requestStatus)

	return
}

func (repo *RequestStatusRepo) Update(uuid string, requestStatus dto.RequestStatus, opts ...interface{}) (err error) {
	if len(opts) < 1 {
		err = fmt.Errorf("context not provided")
		return
	}

	ctx, ok := opts[0].(context.Context)
	if !ok {
		err = fmt.Errorf("the additional argument must be of type context.Context")
		return
	}

	requestStatusJson, err := json.Marshal(requestStatus)
	if err != nil {
		return
	}

	err = repo.client.Set(ctx, uuid, string(requestStatusJson), 0).Err()

	return
}

func (repo *RequestStatusRepo) Delete(uuid string, opts ...interface{}) (err error) {
	if len(opts) < 1 {
		err = fmt.Errorf("context not provided")
		return
	}

	ctx, ok := opts[0].(context.Context)
	if !ok {
		err = fmt.Errorf("the additional argument must be of type context.Context")
		return
	}

	err = repo.client.Del(ctx, uuid).Err()

	return
}
