package controllers

import (
	"github.com/ckeyer/dbc/models"
	// "strconv"
)

type UserController struct {
	Controller
}

// @router /user/:key [get]
func (u *UserController) Get() {
	key, e := u.GetInt64(":key")
	if e != nil {
		log.Error(e.Error())
	}
	user := models.FindUser(key)
	log.Debug(user.String())
	u.Ctx.WriteString(user.String())
}
