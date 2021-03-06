package app

import (
	"sync/atomic"

	"github.com/4726/twitch-chat-logger/config"
	"github.com/4726/twitch-chat-logger/storage"
	"github.com/4726/twitch-chat-logger/storage/mongodb"
	twitch "github.com/gempir/go-twitch-irc/v2"
)

//Worker reads users' messages from Twitch and stores it into a database
type Worker struct {
	chatClient *twitch.Client
	store      storage.Storage
	messages   int64
}

//NewWorker returns a new worker
func NewWorker(conf config.Config, store *mongodb.Storage) *Worker {
	w := &Worker{
		store: store,
	}
	w.chatClient = twitch.NewAnonymousClient()
	w.chatClient.OnPrivateMessage(w.StoreMessage)
	w.chatClient.Join(conf.Channels...)
	return w
}

//Init connects to Twitch's irc server and blocks until an error occurs
func (w *Worker) Init() error {
	return w.chatClient.Connect()
}

//StoreMessage is the callback used with OnPrivateMessage()
func (w *Worker) StoreMessage(privmsg twitch.PrivateMessage) {
	atomic.AddInt64(&w.messages, 1)
	if err := w.store.Add(privateMessageToStorageMessage(privmsg)); err != nil {
		log.Error("add store error: ", err)
	}
}

func (w *Worker) PopMessagesCount() int64 {
	return atomic.SwapInt64(&w.messages, 0)
}
