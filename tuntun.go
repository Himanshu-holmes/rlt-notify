package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/himanshu-holmes/rlt-notify/entity"
	"github.com/himanshu-holmes/rlt-notify/signal"
	"github.com/himanshu-holmes/rlt-notify/storage"
)

type Tuntun struct {
	Storage        storage.Storage
	Signal         signal.Signal
	defaultTimeout time.Duration
}

func NewTuntun(str storage.Storage, sig signal.Signal) *Tuntun {
	return &Tuntun{Storage: str, Signal: sig,
		defaultTimeout: 2 * time.Minute,
	}
}

func Build()*Tuntun{
	str := storage.NewMemoryWithChannel(100)
	sig := signal.NewSignal()
	return NewTuntun(str,sig)
}

func (b *Tuntun)GetNotifications(ctx context.Context, clientID int)([]entity.Notification,error){
	c, err := b.Storage.Count(ctx,clientID)
	if err!= nil {
		return nil,fmt.Errorf("error while counting notifications: %w",err)
	}
	if c>0 {
		return b.Storage.PopAll(ctx,clientID)
	}
	ch,cancel,err := b.Signal.Subscribe("user#"+strconv.Itoa(clientID))
	defer cancel()
	if err != nil {
		return nil,fmt.Errorf("error while trying to lisen to notification topic: %w",err)
	}
	select {
	case <- ctx.Done():
		return nil,ctx.Err()
	case <- ch:
		return b.Storage.PopAll(ctx,clientID)
	}
}