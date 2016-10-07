package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/errors"
	proto "github.com/suicidegang/user-srv/proto/account"
)

var (
	Url string = "root:root@tcp(127.0.0.1:3306)/user"
	db  *gorm.DB
)

func Init() {
	var d *gorm.DB
	var err error

	d, err = gorm.Open("mysql", Url)
	if err != nil {
		log.Fatal(err)
	}

	d.AutoMigrate(&User{})
	d.LogMode(true)
	db = d
}

func Create(user *proto.User, password string) error {
	u := User{
		Email:    user.Email,
		Username: user.Username,
		Name:     user.Username,
		Password: password,
	}

	db.Create(&u)

	return db.Error
}

func Read(id uint64) (*proto.User, error) {
	var row User
	if err := db.First(&row, uint(id)).Error; err != nil {
		return nil, err
	}
	return transformUser(row), nil
}

func Search(query string, within []uint64) ([]*proto.User, error) {
	var rows []User
	q := db

	if len(query) == 0 && len(within) == 0 {
		return make([]*proto.User, 0), errors.BadRequest("sg.micro.srv.user.Search", "Empty params.")
	}

	if len(query) > 0 {
		like := "%" + query + "%"
		q = q.Where("email LIKE ? OR name LIKE ? OR username LIKE ?", like, like, like)
	}

	if len(within) > 0 {
		q = q.Where("id IN (?)", within)
	}

	if err := q.Find(&rows).Error; err != nil {
		return make([]*proto.User, 0), err
	}

	var users []*proto.User
	for _, usr := range rows {
		users = append(users, transformUser(usr))
	}

	return users, nil
}

func transformUser(user User) *proto.User {
	return &proto.User{
		Id:       uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Created:  user.CreatedAt.String(),
		Updated:  user.UpdatedAt.String(),
	}
}
