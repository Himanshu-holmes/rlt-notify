package storage

import (
	"context"
	"sync"

	"github.com/himanshu-holmes/rlt-notify/entity"
)

type userStorage struct {
	mu            *sync.Mutex
	notifications []entity.Notification
}

type memoryWithList struct {
	size    int
	storage *sync.Map
}

func NewMemoryWithList(size int) Storage {
	return &memoryWithList{
		size:    size,
		storage: new(sync.Map),
	}
}
func (m *memoryWithList) Push(ctx context.Context, clientID int, notification entity.Notification) error {
	item := m.get(ctx, clientID)
	item.mu.Lock()
	defer item.mu.Unlock()

	if len(item.notifications) == m.size {
		item.notifications = item.notifications[1:]
	}
	item.notifications = append(item.notifications, notification)
	return nil
}
func (m *memoryWithList)Count(ctx context.Context, clientID int)(int,error){
	item := m.get(ctx, clientID)
	return len(item.notifications),nil
}
func (m *memoryWithList)Pop(ctx context.Context, clientID int)(entity.Notification,error){
	item := m.get(ctx,clientID)
	if len(item.notifications) == 0{
		return nil,ErrEmpty
	}
	item.mu.Lock()
	defer item.mu.Unlock()
	notifcation := item.notifications[0]
	item.notifications = item.notifications[1:]
	return notifcation, nil
}
func(m *memoryWithList)PopAll(ctx context.Context, clientID int)([]entity.Notification,error){
	item := m.get(ctx,clientID)
	if len(item.notifications) == 0{
		return nil,ErrEmpty
	}
	item.mu.Lock()
	defer item.mu.Unlock()
	defer func(){
		item.notifications = nil
	}()

	return item.notifications, nil	
}
func (m *memoryWithList) get(ctx context.Context, clientID int) *userStorage {
	item, _ := m.storage.LoadOrStore(clientID, &userStorage{mu: new(sync.Mutex)})
	return item.(*userStorage)
}
