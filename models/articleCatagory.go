package models

type ArticleCatagory struct {
	Article `xorm:"extends"`
	Catagory `xorm:"extends"`
}

func GetArticleCatagory(page int64) ([]*ArticleCatagory, error) {
	awcs := make([]*ArticleCatagory, 0)
	start := (page - 1) * 10
	//err := db.Limit(1,1).Sql("select a.*,(select title from catagory where id = a.catagory_id) as cata_name from article a order by created desc").Find(&awcs)
	err := db.Table("article").Join("left","catagory","article.catagory_id=catagory.id").Desc("article.Created").Limit(10,int(start)).Find(&awcs)
	return awcs, err
}
