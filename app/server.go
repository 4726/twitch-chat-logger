package app

import (
	"github.com/4726/twitch-chat-logger/config"
	"github.com/4726/twitch-chat-logger/storage/mongodb"
	"github.com/gin-gonic/gin"
)

type Server struct {
	h    *Handlers
	conf config.Config
}

func (s *Server) router() *gin.Engine {
	r := gin.Default()
	r.GET(s.conf.HTTP.SearchRoute, s.h.searchHandler)
	return r
}

func NewServer(conf config.Config) (*Server, error) {
	store := mongodb.New(conf.Mongo.Addr, conf.Mongo.DBName, conf.Mongo.CollectionName)
	s := &Server{
		h:    &Handlers{store},
		conf: conf,
	}
	if err := s.h.store.Connect(); err != nil {
		return nil, err
	}

	worker := NewWorker(conf, store)
	go func() {
		if err := worker.Init(); err != nil {
			log.Error("worker error: ", err)
		}
	}()

	return s, nil
}

func (s *Server) Run() error {
	r := s.router()
	return r.Run(s.conf.HTTP.Addr)
}

func (s *Server) Close() error {
	return s.h.Close()
}
