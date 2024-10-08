package image

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	Client *redis.Client
}

func imageIDKey(sid string) string {
	//sid := hex.EncodeToString(id)
	//sid := fmt.Sprintf("%X", id[10:])
	return fmt.Sprintf("image:%X", sid)
}

func (r *RedisRepo) Insert(ctx context.Context, image ImageLink) error {
	data, err := json.Marshal(image)
	if err != nil {
		return fmt.Errorf("Fail to encode image: %w", err)
	}

	uid := uuid.New()
	sid := fmt.Sprintf("%X", uid[10:])
	key := imageIDKey(sid)

	txn := r.Client.TxPipeline()

	res := txn.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("Fail to set: %w", err)
	}

	if err := txn.SAdd(ctx, "images", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("Fail to add images key: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("Fail to exec: %w", err)
	}

	return nil
}

var ErrNotExist = errors.New("image do not exist")

func (r *RedisRepo) FindByID(ctx context.Context, sid string) (ImageLink, error) {
	key := imageIDKey(sid)
	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return ImageLink{}, ErrNotExist
	} else if err != nil {
		return ImageLink{}, fmt.Errorf("Fail to get by id: %w", err)
	}

	var image ImageLink
	err = json.Unmarshal([]byte(value), &image)
	if err != nil {
		return ImageLink{}, fmt.Errorf("Fail to decode image json: %w", err)
	}
	return image, nil
}

func (r *RedisRepo) DeleteByID(ctx context.Context, sid string) error {
	key := imageIDKey(sid)

	txn := r.Client.TxPipeline()

	err := txn.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExist
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("Fail to delete by id: %w", err)
	}

	if err := txn.SRem(ctx, "images", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("Fail to add images key: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("Fail to exec: %w", err)
	}

	return nil
}

func (r *RedisRepo) Update(ctx context.Context, image ImageLink) error {
	data, err := json.Marshal(image)
	if err != nil {
		return fmt.Errorf("Fail to encode image: %w", err)
	}

	key := imageIDKey(image.ImageID)

	res := r.Client.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("Fail to update: %w", err)
	}

	return nil
}

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Images []ImageLink
	Cursor uint64
}

func (r *RedisRepo) FindAll(ctx context.Context, page FindAllPage) (FindResult, error) {
	res := r.Client.SScan(ctx, "images", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get image ids: %w", err)
	}

	if len(keys) == 0 {
		return FindResult{
			Images: []ImageLink{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return FindResult{}, fmt.Errorf("Failed to get images: %w", err)
	}

	images := make([]ImageLink, len(xs))

	for i, x := range xs {
		x := x.(string)
		var image ImageLink

		err := json.Unmarshal([]byte(x), &image)
		if err != nil {
			return FindResult{}, fmt.Errorf("Failed to decode images: %w", err)
		}

		images[i] = image
	}

	return FindResult{Images: images, Cursor: cursor}, nil
}
