package controller

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gomarkdown/markdown"
	"github.com/kataras/iris/v12"
	"github.com/rs/xid"
)

type ReqCreateMarkdown struct {
	Title string
	Content string
	Description string
}

type Post struct {
	Id string
	Title string
	Content string
	Description string
}

var posts []Post

func GetPosts(ctx iris.Context) {
	ctx.ViewData("posts", posts)
	ctx.View("index.html")
}

func GetMarkdown(ctx iris.Context)  {
	ctx.View("create.html")
}

func CreateMarkdown(ctx iris.Context)  {
	var req ReqCreateMarkdown
	err := ctx.ReadJSON(&req)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
            return
	}

	var post Post

	post.Id = xid.New().String()
	post.Title = req.Title
	post.Content = string(markdown.ToHTML([]byte(req.Content), nil, nil))
	post.Description = req.Description
	log.Println("=========== Create Post ==========")
	log.Println(post)

	posts = append(posts, post)
	log.Println("=========== List Post ==========")
	log.Println(posts)

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Create Post Success")
}

func GetPostById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var post Post

	for i := 0; i < len(posts); i++ {
		if id == posts[i].Id {
			post = posts[i]
		}
	}

	ctx.ViewData("post", post)
	ctx.View("post-detail.html")
}

func GetUpdatePostById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var post Post

	for i := 0; i < len(posts); i++ {
		if id == posts[i].Id {
			post = posts[i]
		}
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(post.Content)

	if err != nil {
		log.Fatal(err)
	}

	post.Content = markdown

	ctx.ViewData("post", post)
	ctx.View("detail.html")
}

func UpdatePostById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var req ReqCreateMarkdown
	err := ctx.ReadJSON(&req)

	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
            return
	}

	for i := 0; i < len(posts); i++ {
		if id == posts[i].Id {
			posts[i].Title = req.Title
			posts[i].Content = req.Content
			posts[i].Description = req.Description
		}
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON("Cập nhật bài viết thành công")
}