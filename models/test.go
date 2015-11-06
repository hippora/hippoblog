package models

type User struct {
	Id      int64
	Name    string
	Address string
}

func AddUser() error {
	_, err := db.Insert(&User{Name:"name",Address:"address"})
	return err
}

func GetUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := db.Find(&users)
	return users, err
}
