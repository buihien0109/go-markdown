package router

import(
	c "markdown/controller"
	"github.com/kataras/iris/v12"
)

func PostRouter(app *iris.Application) {
	app.Get("/posts", c.GetPosts)
	app.Get("/post/create", c.GetMarkdown)
	app.Post("/post/create", c.CreateMarkdown)
	app.Get("/post/{id}", c.GetPostById)
	app.Get("/post/update/{id}", c.GetUpdatePostById)
	app.Put("/post/update/{id}", c.UpdatePostById)
}