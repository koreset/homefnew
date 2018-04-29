package controllers

import (
	"github.com/revel/revel"
	"github.com/koreset/homefnew/app/services"
)

type App struct {
	*revel.Controller
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
