package mutc

import "time"

type EventModel struct {
	Name, Uuid         string
	MinPrice, MaxPrice uint16
	NotificationSentAt time.Time
}
