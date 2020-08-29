package server

import "github.com/gin-gonic/gin"

// GinServer -
type GinServer struct {
	router       *gin.Engine
	publicGroup  *gin.RouterGroup
	privateGroup *gin.RouterGroup
}

// NewGinServer -
func NewGinServer() *GinServer {
	return &GinServer{
		router: gin.New(),
	}
}

// Router -
func (server *GinServer) Router() *gin.Engine {
	return server.router
}

// SetPublicGroup -
func (server *GinServer) SetPublicGroup(path string) {
	if path != "" {
		server.publicGroup = server.Router().Group(path)
	} else {
		panic("Missing path argument to set public group !")
	}
}

// PublicGroup -
func (server *GinServer) PublicGroup() *gin.RouterGroup {
	return server.publicGroup
}

// SetPrivateGroup -
func (server *GinServer) SetPrivateGroup(path string) {
	if path != "" {
		server.privateGroup = server.Router().Group(path)
	} else {
		panic("Missing path argument to set private group !")
	}

}

// PrivateGroup -
func (server *GinServer) PrivateGroup() *gin.RouterGroup {
	return server.privateGroup
}
