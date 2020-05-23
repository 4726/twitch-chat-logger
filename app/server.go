package app

import (
	"time"

	"github.com/4726/twitch-chat-logger/config"
	"github.com/4726/twitch-chat-logger/storage/mongodb"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	h    *Handlers
	conf config.Config
	done chan struct{}
}

func (s *Server) router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET(s.conf.HTTP.SearchRoute, s.h.searchHandler)
	return r
}

func NewServer(conf config.Config) (*Server, error) {
	store := mongodb.New(conf.Mongo.Addr, conf.Mongo.DBName, conf.Mongo.CollectionName)
	done := make(chan struct{}, 1)
	s := &Server{
		h:    &Handlers{store},
		conf: conf,
		done: done,
	}
	log.Info("connecting to mongo")
	if err := s.h.store.Connect(); err != nil {
		log.Error("could not connect to mongo: ", err)
		return nil, err
	}
	log.Info("connected to mongo")

	worker := NewWorker(conf, store)
	go func() {
		log.Info("starting worker")
		go logMessageCount(worker, done)
		if err := worker.Init(); err != nil {
			log.Error("worker error: ", err)
		}
		log.Info("worker stopped")
	}()

	return s, nil
}

func (s *Server) Run() error {
	r := s.router()
	return r.Run(s.conf.HTTP.Addr)
}

func (s *Server) Close() {
	log.Info("shutting down server")
	s.done <- struct{}{}
	if err := s.h.Close(); err != nil {
		log.Error("shut down server error: ", err)
	} else {
		log.Info("shut down server complete")
	}
}

func logMessageCount(w *Worker, done <-chan struct{}) {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			log.Info("messages saved: ", w.PopMessagesCount())
		}
	}
}
