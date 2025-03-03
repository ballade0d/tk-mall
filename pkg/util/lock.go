package util

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
	"time"
)

type RedisLock struct {
	client    *redis.Client
	key       string
	value     string
	ttl       time.Duration
	waitQueue string // 等待队列
}

var file, _ = os.ReadFile("scripts/release_lock.lua")
var unlockScriptBLPOP = redis.NewScript(string(file))

// NewRedisLock 创建锁
func NewRedisLock(client *redis.Client, productID int, ttl time.Duration) *RedisLock {
	lockKey := fmt.Sprintf("lock:inventory:%d", productID) // 锁的 key
	return &RedisLock{
		client:    client,
		key:       lockKey,
		value:     fmt.Sprintf("%d", time.Now().UnixNano()),
		ttl:       ttl,
		waitQueue: lockKey + "_queue",
	}
}

// TryLock 尝试获取锁
func (r *RedisLock) TryLock(ctx context.Context) bool {
	success, err := r.client.SetNX(ctx, r.key, r.value, r.ttl).Result()
	if err != nil {
		log.Fatal(err)
		return false
	}
	return success
}

// Unlock 释放锁并通知队列中的下一个等待者
func (r *RedisLock) Unlock(ctx context.Context) {
	_, err := unlockScriptBLPOP.Run(ctx, r.client, []string{r.key, r.waitQueue}, r.value).Result()
	if err != nil {
		log.Fatal(err)
	}
}

// LockWithQueue 等待锁释放
func (r *RedisLock) LockWithQueue(ctx context.Context, maxWait time.Duration) bool {
	if r.TryLock(ctx) {
		return true
	}

	// 添加到等待队列
	uniqueID := r.value
	r.client.RPush(ctx, r.waitQueue, uniqueID)

	// 订阅等待队列
	sub := r.client.Subscribe(ctx, r.waitQueue)
	defer func(sub *redis.PubSub) {
		err := sub.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sub)

	timer := time.NewTimer(maxWait)
	defer timer.Stop()

	for {
		select {
		case msg := <-sub.Channel():
			if msg.Payload == uniqueID {
				if r.TryLock(ctx) {
					return true
				}
			}
		case <-timer.C:
			r.client.LRem(ctx, r.waitQueue, 0, uniqueID) // 超时后移除自己
			return false
		}
	}
}
