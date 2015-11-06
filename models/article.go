package models

import (
	"strconv"
	"time"
)

type Article struct {
	Id        int64
	Title     string
	Content   string
	ViewCount int64
	Created   time.Time
}

func AddArticle(title,content string) error {
	_, err := db.Insert(&Article{Title: title,Content:content, Created:time.Now()})
	return err
}

func GetArticles() ([]*Article, error) {
	articles := make([]*Article, 0)
	err := db.Find(&articles)
	return articles, err
}

func DelArticle(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = db.Delete(&Article{Id: cid})
	return err
}
