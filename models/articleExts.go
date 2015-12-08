package models


type ArticleWithCataName struct {
	Article `xorm:"extends"`
	CataName string
}

func GetArticleWithCataNames() ([]*ArticleWithCataName, error) {
	awcs := make([]*ArticleWithCataName, 0)
	err := db.Sql("select a.*,(select title from catagory where id = a.catagory_id) as cata_name from article a order by created desc").Find(&awcs)
	return awcs, err
}
/*
func GetArticleWithCataName(id string) (*ArticleWithCataName, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	article := &ArticleWithCataName{}
	_, err = db.Sql("select a.*,(select title from catagory where id = a.catagory_id) as cata_name from article a where id=?",cid).Get(&article)
	fmt.Println(article)
	return article, err
}
*/