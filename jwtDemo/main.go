package main

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/router"
)

func main() {
	r := gin.Default()
	r = router.UserRouter(r)
	r.Run()
}
