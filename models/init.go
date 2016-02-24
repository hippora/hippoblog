package models
import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"log"
)

var db *xorm.Engine

func init()  {
	var err error
	db,err = xorm.NewEngine("sqlite3","test.db")
	if err != nil {
		log.Fatal(err)
	}
	//engine.SetMaxIdleConns(10)
	//engine.SetMaxOpenConns(10)
	db.ShowSQL(true)
	db.ShowExecTime(true)
	//engine.ShowWarn = true

	db.Sync2(new(User),new(Catagory),new(Article))
}