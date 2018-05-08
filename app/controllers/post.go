package controllers

import (
	"github.com/revel/revel"
	"github.com/koreset/homefnew/app/services"
)

type Post struct {
	*revel.Controller
}

func (p Post) GetPost(postid int) revel.Result {
	post := services.GetPost(postid)
	newsfeed := services.GetNewsFeed()

	return p.Render(post, newsfeed)
}
