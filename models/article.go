package models

import (
	"strconv"
	"time"
)

type Article struct {
	Id         int64
	Title      string
	Content    string
	Author     string
	CatagoryId int64
	ViewCount  int64
	Created    time.Time
}

type ArticleWithCatagory struct {
	Article
	Catagory
}

func AddArticle(title, cid, content string) error {
	ccid, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}
	_, err = db.Insert(&Article{Title: title, CatagoryId: ccid, Content:content, Author:"admin", Created:time.Now()})
	return err
}

func GetArticles() ([]*Article, error) {
	articles := make([]*Article, 0)
	err := db.Desc("Created").Find(&articles)
	return articles, err
}

func GetArticle(id string) (*Article, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	article := &Article{Id:cid}
	_, err = db.Get(article)
	return article, err
}

func DelArticle(id string) error {
	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = db.Delete(&Article{Id: aid})
	return err
}
