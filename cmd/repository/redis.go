package repository

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"refstor/cmd/model"
)

type RedisRepo struct {
	Client *redis.Client
}

func imageIDKey(id UUID) string {
	fmt.Println(id)
	sid := hex.EncodeToString(id)
	fmt.Println(sid)
	return fmt.Sprintf("image:%w", sid)
}

func (r *RedisRepo) Insert(ctx context.Context, image model.Image) error {
	data, err := json.Marshal(image)
	if err != nil {
		return fmt.Errorf("Fail to encode image: %w", err)
	}
	key := imageIDKey(uuid.New())
	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("Fail to set: %w", err)
	}
	return nil
}
