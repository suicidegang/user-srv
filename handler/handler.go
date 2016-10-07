package handler

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/suicidegang/user-srv/db"
	proto "github.com/suicidegang/user-srv/proto/account"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const (
	x        = "breakstuff123"
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

type Account struct{}

func (a *Account) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {

	// Lets make some hash hash (iykwim :P)
	hash, err := bcrypt.GenerateFromPassword([]byte(x+req.Password), 10)
	if err != nil {
		return errors.InternalServerError("sg.micro.srv.user.Create", err.Error())
	}

	pp := base64.StdEncoding.EncodeToString(hash)
	req.User.Username = strings.ToLower(req.User.Username)
	req.User.Email = strings.ToLower(req.User.Email)
	err = db.Create(req.User, pp)

	if err != nil {
		return errors.New("sg.micro.srv.user.Create", err.Error(), 409)
	}

	return nil
}

func (a *Account) Read(ctx context.Context, req *proto.ReadRequest, res *proto.ReadResponse) error {
	user, err := db.Read(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (a *Account) Search(ctx context.Context, req *proto.SearchRequest, res *proto.SearchResponse) error {
	users, err := db.Search(req.Query, req.Id)
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func random(i int) string {
	bytes := make([]byte, i)
	for {
		rand.Read(bytes)
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}
		return string(bytes)
	}
	return "ughwhy?!!!"
}
