package service

import "context"

type ILocalCache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) bool
	SetWithTTL(ctx context.Context, key string, value interface{}) bool // set time để nó xoá cache đi
	Del(ctx context.Context, key string)
}
