package helper

import (
	"context"
	"ddxs-api/internal/svc"
	json "github.com/bytedance/sonic"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisCacheHelper struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedisCacheHelper(ctx context.Context, svcCtx *svc.ServiceContext) *RedisCacheHelper {
	return &RedisCacheHelper{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (r *RedisCacheHelper) GetFromCache(key string, value interface{}) error {
	result, err := r.svcCtx.Redis.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	err = json.Unmarshal([]byte(result), value)
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisCacheHelper) SetToCache(key string, value interface{}, cacheExpiration ...time.Duration) error {
	jsonResp, err := json.Marshal(value)
	if err != nil {
		return err
	}

	var expiration time.Duration
	if len(cacheExpiration) > 0 {
		expiration = cacheExpiration[0]
	} else {
		expiration = time.Duration(r.svcCtx.Config.Redis.CacheExpiration) * time.Second
	}

	err = r.svcCtx.Redis.Set(r.ctx, key, jsonResp, expiration).Err()
	return err
}
