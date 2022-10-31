package framework

import (
	"context"
	"fmt"
	"log"
	"time"
)

func TimeoutHandler(fun ControllerHandler,d time.Duration) ControllerHandler {
	return func(c *Context) error {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)
		durationCtx, cancel := context.WithTimeout(c.BaseContext(), d)
		defer cancel()
		c.request.WithContext(durationCtx)
		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()
			//执行具体逻辑
			_ = c.Next()
			finish <- struct{}{}
		}()
		select {
		case p := <- panicChan:
			_ = c.Json(500, "time out")
			log.Println(p)
		case <- durationCtx.Done():
			c.SetHasTimeout()
			_ = c.Json(500, "time out")
			case <- finish:
			fmt.Println("finish")
		}
		return nil
	}
}
