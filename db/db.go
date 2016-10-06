package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
