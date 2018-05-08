package services

import (
	gorm "github.com/revel/modules/orm/gorm/app"
	"github.com/koreset/homefnew/app/models"
	"github.com/kennygrant/sanitize"
	"fmt"
)

var defaultTags = []string{"h1", "h2", "h3", "h4", "h5", "h6", "div", "span", "hr", "p", "br", "b", "i", "strong", "em", "ol", "ul", "li", "a", "img", "pre", "code", "blockquote", "article", "section"}

var defaultAttributes = []string{"id", "src", "href", "title", "alt", "name", "rel"}

func GetPosts(start, limit int) []models.Post {
	var posts []models.Post
	gorm.DB.Where("type in (?) and body != ''", []string{"article", "press_release", "news"}).Preload("Images").Order("created desc").Offset(start).Limit(limit).Find(&posts)
	// Lets sanitize the html output and strip off MSOffice tags
	for _, post := range posts {
		post.Body, _ = sanitize.HTMLAllowing(post.Body, defaultTags, defaultAttributes)
	}
	return posts
}

func GetVideos() []models.Post {
	var videos []models.Post
	gorm.DB.LogMode(true)
	gorm.DB.Where("type = 'video'").Preload("Images").Preload("Videos").Order("created desc").Find(&videos)
	return videos
}

func GetPost(postid int) models.Post{
	var post models.Post
	gorm.DB.Where("id = ? ", postid).Preload("Images").Preload("Links").Preload("Videos").First(&post)

	fmt.Println(post)
	return post
}

func GetPublications() []models.Post {
	var publications []models.Post
	gorm.DB.LogMode(true)
	gorm.DB.Where("type = 'publication'").Preload("Images").Preload("Links").Order("created desc").Find(&publications)

	for _, pub := range publications {
		fmt.Println(pub.Images[0].Url, pub.ID)
	}
	return publications
}

func GetPublication(postid int) models.Post  {
	var pub models.Post
	gorm.DB.Where("id = ?", postid).Preload("Images").Preload("Videos").Preload("Links").First(&pub)
	return pub
}
