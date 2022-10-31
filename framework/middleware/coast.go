package middleware

import (
	"log"
	"self-web/framework"
	"time"
)

func Coast() framework.ControllerHandler {
	return func(c *framework.Context) error {
		start := time.Now()
		c.Next()
		end := time.Now()
		coast := end.Sub(start)
		log.Printf("api uri:%v,coast:%v",c.GetRequest().RequestURI,coast)
		return nil
	}
}
