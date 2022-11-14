package main

import (
	"github.com/yujian0213/self-web/framework"
	"github.com/yujian0213/self-web/framework/gin"
)

// 注册路由规则
func registerRouter(core *gin.Engine) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.GET("/user/login", UserLoginController)

	// 需求3:批量通用前缀
	subjectApi := core.Group("/subject")
	{
		// 需求4:动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
