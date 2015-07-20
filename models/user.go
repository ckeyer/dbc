package models

import (
	"encoding/json"
	"time"
)

type User struct {
	Id      int64
	Name    string
	Profile *Profile
}

func FindUser(id int64) *User {
	return &User{
		Id:   id,
		Name: "ckeyer",
		Profile: &Profile{
			Email: "me@ckeyer.com",
			Birth: time.Now().Unix(),
		},
	}
}
func (u *User) String() string {
	bs, e := json.Marshal(u)
	if e != nil {
		log.Error(e.Error())
		return e.Error()
	}
	return string(bs)
}
