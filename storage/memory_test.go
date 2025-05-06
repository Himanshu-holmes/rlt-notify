package storage_test

import (
	"context"
	"math/rand"
	"testing"

	"github.com/himanshu-holmes/rlt-notify/entity"
	"github.com/himanshu-holmes/rlt-notify/pkg"
	"github.com/himanshu-holmes/rlt-notify/storage"
	"github.com/stretchr/testify/assert"
)

func testNewMemory(m storage.Storage,t *testing.T){
	ctx := context.Background()
	m.Push(ctx,10,entity.UnreadMessagesNotification{Count: 1})
	m.Push(ctx,10,entity.UnreadMessagesNotification{Count: 2})
	m.Push(ctx,10,entity.UnreadMessagesNotification{Count: 3})
	c,_ := m.Count(ctx,10)
	assert.Equal(t,3,c)

	p,err := m.Pop(ctx,10)
	assert.NoError(t,err)
	assert.Equal(t,1,p.(entity.UnreadMessagesNotification).Count)
	
	all,_ := m.PopAll(ctx,10)
	assert.Equal(t,2,len(all))

	for i:=0; i<15; i++ {
		m.Push(ctx,10,entity.UnreadMessagesNotification{Count: i})
	}
	f,err := m.Pop(ctx,10)
	assert.NoError(t,err)
	assert.Equal(t,5,f.(entity.UnreadMessagesNotification).Count)
}

func benchmarkMemory_PushAverage(m storage.Storage,b *testing.B){
	ctx := context.Background()
	for i:=0;i<b.N;i++{
		id := rand.Intn(1000)
		m.Push(ctx, id, entity.UnreadMessagesNotification{
			Count: i,
		})
	}
	b.StopTimer()
    pkg.PrintMemUsage()
}
func benchmarkMemory_PushNewItem(m storage.Storage, b *testing.B){
	ctx := context.Background()
	counter := 0
	for i:=0; i< b.N; i++ {
		m.Push(ctx,i,entity.UnreadMessagesNotification{Count: i})
		counter++
	}
	b.StopTimer()
	b.Log("for ",b.N," notifications: ")
	pkg.PrintMemUsage()
}
func BenchmarkMemoryWithChannel_PushNewItem(b *testing.B){
	benchmarkMemory_PushNewItem(storage.NewMemoryWithChannel(1000),b)
}
func BenchmarkMemoryWithList_PushNewItem(b *testing.B){
	benchmarkMemory_PushNewItem(storage.NewMemoryWithList(1000),b)
}