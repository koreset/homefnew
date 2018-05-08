package controllers

import (
	"github.com/revel/revel"
	"github.com/koreset/homefnew/app/services"
)

type Publication struct {
	*revel.Controller
}


func (p Publication) GetPublications() revel.Result{
	newsfeed := services.GetNewsFeed()
	publications := services.GetPublications()
	return p.Render(newsfeed, publications)
}

func (p Publication) GetPublication(postid int) revel.Result {
	publication := services.GetPublication(postid)
	return p.Render(publication)
}
