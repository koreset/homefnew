package services

import (
	gorm "github.com/revel/modules/orm/gorm/app"
)

type MenuItem struct {
	Type  string
	Count uint
}

func GetContentCounts() []MenuItem {
	rows, err := gorm.DB.Table("contents").Select("count(type) as count, type").Group("type").Rows()
	if err != nil {
		panic(err)
	}
	var items []MenuItem
	for rows.Next() {
		var item MenuItem
		gorm.DB.ScanRows(rows, &item)
		items = append(items, item)
	}
	return items
}
