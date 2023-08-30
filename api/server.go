package api

import (
	db "github.com/UnTea/WindingPacketsOnAPear/db/sqlc"
	"github.com/gin-gonic/gin"
	"log"
)

// Server server HTTP requests for banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	err := server.router.SetTrustedProxies(nil)
	if err != nil {
		log.Println("cannot set trusted proxies parameter: ", err)
	}

	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
