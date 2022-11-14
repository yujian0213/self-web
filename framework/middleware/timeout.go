package middleware

import (
	"context"
	"fmt"
	"github.com/yujian0213/self-web/framework/gin"
	"log"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context)  {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//执行具体逻辑
			c.Next()
			finish <- struct{}{}
		}()
		select {
		case p := <- panicChan:
			 c.ISetStatus(500).IJson("timeout")
			log.Println(p)
		case <- durationCtx.Done():
			c.ISetStatus(500).IJson("timeout")
			case <- finish:
			fmt.Println("finish")
		}
	}
}
