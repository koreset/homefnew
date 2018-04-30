package controllers

import (
	"github.com/revel/revel"
	"github.com/koreset/homefnew/app/services"
)

type App struct {
	*revel.Controller
}

type MyResult struct {
	ID int
}

func (c App) Index() revel.Result {
	posts := services.GetPosts(0, 10)
	newsfeed := services.GetNewsFeed()
	videos := services.GetVideos()
	publications := services.GetPublications()
	return c.Render(posts, newsfeed, videos, publications)
}

func (c App) GetContent() revel.Result {
	return c.Render()
}

func (c App) GetPost(postid int) revel.Result {
	post := services.GetPost(postid)
	return c.Render(post)
}

func (c App) CheckHealth() revel.Result {
	return c.RenderJSON("OK")
}
