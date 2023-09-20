package api

import (
	"fmt"
	"github.com/UnTea/WindingPacketsOnAPear/util"

	db "github.com/UnTea/WindingPacketsOnAPear/db/sqlc"
	"github.com/UnTea/WindingPacketsOnAPear/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server server HTTP requests for banking service
type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

// NewServer creates a new HTTP server and set up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			errorResponse(err)
		}
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authorizationRoute := router.Group("/").Use(authorizationMiddleware(server.tokenMaker))

	authorizationRoute.POST("/accounts", server.createAccount)
	authorizationRoute.GET("/accounts/:id", server.getAccount)
	authorizationRoute.GET("/accounts", server.listAccounts)

	authorizationRoute.POST("/transfers", server.createTransfer)

	server.router = router
	if err := server.router.SetTrustedProxies(nil); err != nil {
		errorResponse(err)
	}
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
