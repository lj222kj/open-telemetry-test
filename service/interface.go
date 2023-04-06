package service

import "context"

type Service interface {
	Hello(ctx context.Context) string
}
