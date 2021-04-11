package main

import (
	"markdown/router"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()

	// Logging terminal
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	// Serve Static file
	app.HandleDir("/", "./public")

	// Register view
	tmpl := iris.HTML("./views", ".html").Reload(true)
	app.RegisterView(tmpl)

	// Register router
	router.PostRouter(app)

	// Listenning app
	app.Listen(":8080")
}

