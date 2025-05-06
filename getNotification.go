package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/himanshu-holmes/rlt-notify/entity"
)

func (b *Tuntun) Notify(ctx context.Context, userID int, notification entity.Notification)error{
	if err := b.Storage.Push(ctx,userID,notification);err !=nil {
		return fmt.Errorf("error while trying to push the new notification: %w",err)
	}
	_ = b.Signal.Publish("user#"+strconv.Itoa(userID))
	return nil
}