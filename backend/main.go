package main

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"imooc/imooc_iris/backend/web/controllers"
	"imooc/imooc_iris/common"
	"imooc/imooc_iris/repositories"
	"imooc/imooc_iris/services"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 注册模板
	tmplate := iris.HTML("./web/views",
		".html").Layout(
		"shared/layout.html").Reload(true)
	app.RegisterView(tmplate)
	// 4.设置模板目标
	// app.StaticWeb("/assets","./backend/web/assets")  此方法没有了  HandleDir 代替了
	app.HandleDir("/assets", "./web/assets")
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.View("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})
	// 连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		fmt.Println("MySQL连接错误:", err)
		return
		// log.Error(err)
	}
	// os.Exit(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	productRepository := repositories.NewProductManager("product", db)
	productSerivce := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productSerivce)
	product.Handle(new(controllers.ProductController))

	app.Run(
		iris.Addr("localhost:8080"),
		// iris.WithoutVersionChecker,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
