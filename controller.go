package main

import (
	"fmt"
	"self-web/framework"
)

func UserLoginController(c *framework.Context) error {
	fmt.Println("UserLoginController")
	_ = c.SetOkStatus().Json("ok,UserLoginController")
	return nil
}
func SubjectListController(c *framework.Context) error {
	_ = c.SetOkStatus().Json("ok,SubjectListController")
	fmt.Println("SubjectListController")
	return nil
}
func SubjectDelController(c *framework.Context) error {
	_ = c.SetOkStatus().Json("ok,SubjectDelController")
	fmt.Println("SubjectDelController")
	return nil
}
func SubjectUpdateController(c *framework.Context) error {
	_ = c.SetOkStatus().Json("ok,SubjectUpdateController")
	fmt.Println("SubjectUpdateController")
	return nil
}
func SubjectGetController(c *framework.Context) error {
	_ = c.SetOkStatus().Json("ok,SubjectGetController")
	fmt.Println("SubjectGetController")
	return nil
}
func FooControllerHandler(c *framework.Context) error {
	return nil
}
