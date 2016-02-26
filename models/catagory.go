package models

import (
	"strconv"
	"time"
)

type Catagory struct {
	Id    int64
	Title string
	ArticleCount int64
	Created time.Time
}

func AddCatagory(title string) error {
	_, err := db.Insert(&Catagory{Title: title,Created:time.Now()})
	return err
}
func GetCatagories() ([]*Catagory, error) {
	catagories := make([]*Catagory, 0)
	err := db.Find(&catagories)
	return catagories, err
}

func DelCatagory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	_, err = db.Delete(&Catagory{Id: cid})
	return err
}