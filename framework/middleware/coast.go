package middleware

import (
	"github.com/yujian0213/self-web/framework/gin"
	"log"
	"time"
)

func Coast() gin.HandlerFunc {
	return func(c *gin.Context)  {
		start := time.Now()
		c.Next()
		end := time.Now()
		coast := end.Sub(start)
		log.Printf("api uri:%v,coast:%v",c.Request.RequestURI,coast)
	}
}
