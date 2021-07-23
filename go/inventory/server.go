package inventory

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const (
	HandlerGet    = "GET"
	HandlerPost   = "POST"
	HandlerPut    = "PUT"
	HandlerDelete = "DELETE"
)

var (
	ErrInvalidHandler = errors.New("invlaid handler")
)

type Server struct {
	engine *gin.Engine
}

func NewServerWithContext() *Server {
	e := gin.Default()

	e.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &Server{
		engine: e,
	}
}

func NewServer() *Server {
	e := gin.Default()

	e.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &Server{
		engine: e,
	}
}

func (s *Server) AddHandler(handlerType, path string, handlerFunc func(*gin.Context)) error {
	switch handlerType {
	case "GET":
		s.engine.GET(path, handlerFunc)
		return nil
	case "POST":
		s.engine.POST(path, handlerFunc)
		return nil
	case "PUT":
		s.engine.PUT(path, handlerFunc)
		return nil
	case "DELETE":
		s.engine.PUT(path, handlerFunc)
		return nil
	default:
		return ErrInvalidHandler
	}
}

func (s *Server) Run() {
	// listen and serve on 0.0.0.0:8080
	s.engine.Run()
}
