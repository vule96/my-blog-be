package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vule96/my-blog-be/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/posts", server.createPost)
	router.GET("/post/:id", server.getPost)
	router.GET("/posts", server.listPosts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"err": err.Error()}
}
