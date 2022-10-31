package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	request *http.Request
	responseWriter http.ResponseWriter
	ctx context.Context
	handlers []ControllerHandler
	index int

	//写保护机制
	writeMux *sync.Mutex
	//是否超时标记位
	hasTimeout bool
	params map[string]string // url路由匹配的参数
}

func NewContext(r *http.Request,w http.ResponseWriter) *Context  {
	return &Context{
		request: r,
		responseWriter: w,
		ctx: r.Context(),
		index: -1,
		writeMux: &sync.Mutex{},
	}
}
//为ctx设置handlers
func (ctx *Context) SetHandlers(handlers []ControllerHandler)  {
	ctx.handlers = handlers
}
//实现链条调用handler
func (ctx *Context) Next() error  {
	ctx.index++
	if ctx.index < len(ctx.handlers) {
		if err := ctx.handlers[ctx.index](ctx);err != nil {
			return err
		}
	}
	return nil
}
// 设置参数
func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}
//#region base function
//对外暴露锁
func (ctx *Context) WriteMux() *sync.Mutex  {
	return ctx.writeMux
}
func (ctx *Context) GetRequest() *http.Request {
	return ctx.request
}
func (ctx *Context) GetResponse() http.ResponseWriter  {
	return ctx.responseWriter
}
func (ctx *Context) SetHasTimeout()  {
	ctx.hasTimeout = true
}
func (ctx *Context) HasTimeout() bool  {
	return ctx.hasTimeout
}
//#end region

func (ctx *Context) BaseContext() context.Context  {
	return ctx.request.Context()
}

func (ctx *Context) Deadline() (deadline time.Time,ok bool)  {
	return ctx.BaseContext().Deadline()
}
func (ctx *Context) Done() <-chan struct{}  {
	return ctx.BaseContext().Done()
}
func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}
func (ctx *Context) Value(key interface{}) interface{}  {
	return ctx.BaseContext().Value(key)
}
func (ctx *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}



