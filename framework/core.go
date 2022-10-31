package framework

import (
	"log"
	"net/http"
	"strings"
)
//框架核心结构
type Core struct {
	//router map[string]map[string]ControllerHandler
	router map[string]*Tree
	middlewares []ControllerHandler
}
//初始化框架核心结构
func NewCore() *Core  {
	//定义路由map
	//getRouter    := map[string]ControllerHandler{}
	//postRouter   := map[string]ControllerHandler{}
	//putRouter    := map[string]ControllerHandler{}
	//deleteRouter := map[string]ControllerHandler{}
	//初始化路由
	router := map[string]*Tree{}
	router["GET"] 	 = NewTree()
	router["POST"]   = NewTree()
	router["PUT"]  	 = NewTree()
	router["DELETE"] = NewTree()
	return &Core{router: router}
}
//批量注册注册通用中间件
func (c *Core) Use(middlewares ...ControllerHandler)  {
	c.middlewares = append(c.middlewares,middlewares...)
}
//初始化Group
func (c *Core) Group(prefix string) IGroup  {
	return NewGroup(c,prefix)
}
//注册静态路由（按方法拆分）
func (c *Core) Get(uri string,handlers ...ControllerHandler)  {
	//将core的middlewares和handlers合并
	allHandlers := append(c.middlewares,handlers...)
	if err := c.router["GET"].AddRouter(uri,allHandlers);err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Post(uri string,handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares,handlers...)
	if err := c.router["POST"].AddRouter(uri,allHandlers);err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Put(uri string,handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares,handlers...)
	if err := c.router["PUT"].AddRouter(uri,allHandlers);err != nil {
		log.Fatal("add router error: ", err)
	}
}
func (c *Core) Delete(uri string,handlers ...ControllerHandler)  {
	allHandlers := append(c.middlewares,handlers...)
	if err := c.router["DELETE"].AddRouter(uri,allHandlers);err != nil {
		log.Fatal("add router error: ", err)
	}
}
// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteNodeByRequest(request *http.Request) *node {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.root.matchNode(uri)
	}
	return nil
}
//框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter,request *http.Request)  {
	log.Println("core.serverHTTP")
	ctx := NewContext(request,response)

	log.Println("core.router")
	//匹配路由，找到路由对应的handler
	node := c.FindRouteNodeByRequest(request)
	if node == nil {
		_ = ctx.SetStatus(404).Json("not found")
		return
	}
	ctx.SetHandlers(node.handlers)
	// 设置路由参数
	params := node.parseParamsFromEndNode(request.URL.Path)
	ctx.SetParams(params)
	if err := ctx.Next();err != nil {
		_ = ctx.SetStatus(500).Json("inner err")
		return
	}
}
