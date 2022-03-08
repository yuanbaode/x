package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Config Config
}

func NewServer(config Config) *Server {
	return &Server{
		Config: config,

	}}

func (s *Server) Start() (err error) {
	//serveMux := http.NewServeMux()
	//serveMux.Handle("/", http.StripPrefix("/xxx",http.FileServer(http.Dir("./"))))
	//http.ListenAndServe("0.0.0.0:8081",serveMux)

	r := gin.Default()
	r.StaticFS("/", gin.Dir("./",true))
	return r.Run(s.Config.Host)
}
