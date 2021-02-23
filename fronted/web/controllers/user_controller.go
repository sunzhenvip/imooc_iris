package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"imooc/imooc_iris/datamodels"
	"imooc/imooc_iris/services"
	"imooc/imooc_iris/tool"
	"strconv"
)

type UserController struct {
	Ctx     iris.Context
	Service services.IUserService
	Session *sessions.Session
}

func (c *UserController) GetRegister() mvc.View {
	return mvc.View{Name: "user/register.html"}
}

func (c *UserController) PostRegister() {
	var (
		nickName = c.Ctx.FormValue("nickName")
		userName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	// fmt.Println(nickName)
	// fmt.Println(userName)
	// fmt.Println(password)
	// ozzo-validation
	user := &datamodels.User{
		UserName:     userName,
		NickName:     nickName,
		HashPassword: password,
	}
	_, err := c.Service.AddUser(user)
	// fmt.Println(err)
	if err != nil {
		c.Ctx.Redirect("/user/error")
		return
	}
	c.Ctx.Redirect("/user/login")
	return
}

func (c *UserController) GetLogin() mvc.View {
	return mvc.View{
		Name: "/user/login.html",
	}
}

func (c *UserController) PostLogin() mvc.Response {
	var (
		UserName = c.Ctx.FormValue("userName")
		password = c.Ctx.FormValue("password")
	)
	user, isOk := c.Service.IsPwdSuccess(UserName, password)
	if !isOk {
		fmt.Println("没有通过!")
		return mvc.Response{
			Path: "/user/login",
		}
	}
	tool.GlobalCookie(c.Ctx, "uid", strconv.FormatInt(user.ID, 10))
	c.Session.Set("userID", strconv.FormatInt(user.ID, 10))
	return mvc.Response{Path: "/product"}
}
