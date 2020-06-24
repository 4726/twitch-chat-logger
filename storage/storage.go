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

type QueryOptions struct {
	Channel string
	Name    string
	Date    time.Time
}

type Storage interface {
	Connect() error
	Add(ChatMessage) error
	Query(opts QueryOptions) ([]ChatMessage, error)
	Close() error
}
