package signal_test

import (
	"testing"

	"github.com/himanshu-holmes/rlt-notify/signal"
	"github.com/stretchr/testify/assert"
)

func TestNewChannel(t *testing.T){
	s := signal.NewSignal()
	ac, _, err := s.Subscribe("a")

	assert.NoError(t, err)
	
	ac2,_,err:= s.Subscribe("a")
	assert.NoError(t, err)
	
	s.Publish("a")
	select {
	case <- ac:
	default:
		t.Fatal("didn't recieve the signal")
	}
	select {
	case <- ac2:
	default:
		t.Fatal("didn't recieve the signal")
	}
	err = s.Publish("b")
	assert.ErrorIs(t, err,signal.ErrEmpty)
	
	_, cancel,err := s.Subscribe("c")
	assert.NoError(t, err)
	cancel()
	err = s.Publish("c")
	assert.ErrorIs(t, err,signal.ErrEmpty)
	
}