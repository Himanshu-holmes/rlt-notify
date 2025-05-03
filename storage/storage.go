package storage

import (
	"context"
	"errors"

	"github.com/himanshu-holmes/rlt-notify/entity"
)

var ErrEmpty = errors.New("no notifications found")

type Storage interface {
	Push(ctx context.Context, clientID int, notification entity.Notification)error
	Count(ctx context.Context, clientID int)(int,error)
	Pop(ctx context.Context, clientID int)(entity.Notification,error)
	PopAll(ctx context.Context,clientID int)([]entity.Notification,error)
}
