package app

import (
	"github.com/4726/twitch-chat-logger/config"
	"github.com/4726/twitch-chat-logger/storage/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Server struct {
	h    *Handlers
	conf config.Config
}

func (s *Server) router() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET(s.conf.HTTP.SearchRoute, s.h.searchHandler)
	return r
}

func NewServer(conf config.Config) (*Server, error) {
	store := mongodb.New(conf.Mongo.Addr, conf.Mongo.DBName, conf.Mongo.CollectionName)
	s := &Server{
		h:    &Handlers{store},
		conf: conf,
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
	if err := s.h.Close(); err != nil {
		log.Error("shut down server error: ", err)
	} else {
		log.Info("shut down server complete")
	}
}
