package handler

import (
	"crypto/rand"
	"encoding/base64"
	"strings"

	"github.com/micro/go-micro/errors"
	"github.com/suicidegang/user-srv/db"
	account "github.com/suicidegang/user-srv/proto/account"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

const (
	x = "breakstuff123"
)

var (
	alphanum = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

type Account struct{}

func (a *Account) Create(ctx context.Context, req *account.CreateRequest, rsp *account.CreateResponse) error {

	// Lets make some hash hash (iykwim :P)
	hash, err := bcrypt.GenerateFromPassword([]byte(x+req.Password), 10)
	if err != nil {
		return errors.InternalServerError("sg.micro.srv.user.Create", err.Error())
	}

	pp := base64.StdEncoding.EncodeToString(hash)

	req.User.Username = strings.ToLower(req.User.Username)
	req.User.Email = strings.ToLower(req.User.Email)

	return db.Create(req.User, pp)
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
