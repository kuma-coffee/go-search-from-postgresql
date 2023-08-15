package routerr

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

type Router interface {
	GET(uri string, f func(c *gin.Context))
	POST(uri string, f func(c *gin.Context))
	SERVE(port string)
}

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(c *gin.Context)) {
	router.GET(uri, f)
}

func (*chiRouter) POST(uri string, f func(c *gin.Context)) {
	router.POST(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi Server listening on port: %v", port)
	router.Run(port)
}
