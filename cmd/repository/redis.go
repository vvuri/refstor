package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"refstor/cmd/model"
)

type RedisRepo struct {
	Client *redis.Client
}

func imageIDKey(sid string) string {
	//sid := hex.EncodeToString(id)
	//sid := fmt.Sprintf("%X", id[10:])
	return fmt.Sprintf("image:%X", sid)
}

func (r *RedisRepo) Insert(ctx context.Context, image model.Image) error {
	data, err := json.Marshal(image)
	if err != nil {
		return fmt.Errorf("Fail to encode image: %w", err)
	}

	uid := uuid.New()
	sid := fmt.Sprintf("%X", uid[10:])
	key := imageIDKey(sid)

	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("Fail to set: %w", err)
	}
	return nil
}

var ErrNotExist = errors.New("image do not exist")

func (r *RedisRepo) FindByID(ctx context.Context, sid string) (model.Image, error) {
	key := imageIDKey(sid)
	value, err := r.Client.Get(ctx, key).Result()
	if erroer.Is(err, redis.Nil) {
		return model.Image{}, ErrNotExist
	} else if err != nil {
		return model.Image{}, fmt.Errorf("Fail to get by id: %w", err)
	}

	var image model.Image
	err = json.Unmarshal([]byte(value), &image)
	if err != nil {
		return model.Image{}, fmt.Errorf("Fail to decode image json: %w", err)
	}
	return image, nil
}

func (r *RedisRepo) DeleteByID(ctx context.Context, sid string) error {
	key := imageIDKey(sid)

	err := r.Client.Del(ctx, key).Err()
	if erroer.Is(err, redis.Nil) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("Fail to delete by id: %w", err)
	}

	returm nil
}

func (r *RedisRepo) Update(ctx context.Context, sid string) error {


}
