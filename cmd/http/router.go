package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	Handler httpHandler
}

func NewRouter(handler httpHandler) *Router {
	return &Router{
		router:  gin.Default(),
		Handler: handler,
	}
}

func (r *Router) routes() {
	r.router.GET("/ping", r.Handler.Ping)
	r.router.POST("/login", r.Handler.Login)
	r.router.GET("/merchant/omzet", r.Handler.MerchantOmzetReport)
	r.router.GET("/outlet/omzet", r.Handler.OutletOmzetReport)
}
