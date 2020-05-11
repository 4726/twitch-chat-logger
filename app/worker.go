package app

import (
	"github.com/4726/twitch-chat-logger/config"
	"github.com/4726/twitch-chat-logger/storage"
	"github.com/4726/twitch-chat-logger/storage/mongodb"
	twitch "github.com/gempir/go-twitch-irc/v2"
)

//Worker reads users' messages from Twitch and stores it into a database
type Worker struct {
	chatClient *twitch.Client
	store      storage.Storage
}

//NewWorker returns a new worker
func NewWorker(conf config.Config) *Worker {
	chatClient := twitch.NewAnonymousClient()
	chatClient.Join(conf.Channels...)
	return &Worker{
		chatClient: chatClient,
		store:      mongodb.New(),
	}
}

//Init connects to the storage database as well as Twitch's irc server
func (w *Worker) Init() error {
	if err := w.store.Connect(); err != nil {
		return err
	}
	w.chatClient.OnPrivateMessage(w.StoreMessage)
	return w.chatClient.Connect()
}

//StoreMessage is the callback used with OnPrivateMessage()
func (w *Worker) StoreMessage(privmsg twitch.PrivateMessage) {
	if err := w.store.Add(privateMessageToStorageMessage(privmsg)); err != nil {
		log.Error("store error: ", err)
	}
}
