package framework

import "github.com/yujian0213/self-web/framework/gin"

//IGroup 代表前缀分组
type IGroup interface {
	//实现httpMethod方法
	Get(string,...gin.HandlerFunc)
	Post(string,...gin.HandlerFunc)
	Put(string,...gin.HandlerFunc)
	Delete(string,...gin.HandlerFunc)
}
//Group实现IGroup
type Group struct {
	core *Core
	parent *Group	//指向上一个group（如果有）
	prefix string //这个group的通用前缀
	middlewares []gin.HandlerFunc
}
//初始化group
func NewGroup(core *Core,prefix string) *Group  {
	return &Group{
		core: core,
		parent: nil,
		prefix: prefix,
		middlewares: []gin.HandlerFunc{},
	}
}
//注册中间件
func (g *Group) Use(middlewares ...gin.HandlerFunc)  {
	g.middlewares = append(g.middlewares,middlewares...)
}
//获取group的绝对路径
func (g *Group)  GetAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.GetAbsolutePrefix() + g.prefix
}
//注册分组路由
func (g *Group) Get(uri string,handlers ...gin.HandlerFunc)  {
	uri = g.GetAbsolutePrefix() + uri
	allHandlers := append(g.middlewares,handlers...)
	g.core.Get(uri,allHandlers...)
}
func (g *Group) Post(uri string,handlers ...gin.HandlerFunc)  {
	uri = g.GetAbsolutePrefix() + uri
	allHandlers := append(g.middlewares,handlers...)
	g.core.Post(uri,allHandlers...)
}
func (g *Group) Put(uri string,handlers ...gin.HandlerFunc)  {
	uri = g.GetAbsolutePrefix() + uri
	allHandlers := append(g.middlewares,handlers...)
	g.core.Put(uri,allHandlers...)
}
func (g *Group) Delete(uri string,handlers ...gin.HandlerFunc)  {
	uri = g.GetAbsolutePrefix() + uri
	allHandlers := append(g.middlewares,handlers...)
	g.core.Delete(uri,allHandlers...)
}