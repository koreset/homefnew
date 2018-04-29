package services

import (
	"fmt"

	"github.com/koreset/homefnew/app/models"
	gorm "github.com/revel/modules/orm/gorm/app"
)

func GetNewsFeed() []models.FeedItem {
	var results []models.FeedItem

	gorm.DB.Limit(10).Find(&results)
	fmt.Println(results)
	return results
}
