package storage

import "time"

type ChatMessage struct {
	Channel, Message, RoomID, ID string
	Bits                         int
	SubscribeMonths              int
	Admin                        bool
	GlobalMod                    bool
	Moderator                    bool
	Staff                        bool
	Turbo                        bool
	Subscriber                   bool
	DisplayName                  string
	Time                         int64
	UserID                       string
	Name                         string
}

type Storage interface {
	Connect() error
	Add(ChatMessage) error
	Query(channel, term, name string, date time.Time) ([]ChatMessage, error)
	QuerySubscriber(channel, term, name string, date time.Time, subscribeMin int) ([]ChatMessage, error)
	Close() error
}
