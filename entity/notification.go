package entity

import "time"

type Notification interface {
	IsNotification()
}

type BaseNotification struct {
	CreatedAt time.Time `json:"created_at"`
}

func (BaseNotification) IsNotification() {}

type UnreadWorkRequest struct {
	BaseNotification
	WorkID int `json:"workID"`
	Title string `json:"title"`
}

type UnreadMessagesNotification struct {
	BaseNotification
	Count int `json:"count"`
}