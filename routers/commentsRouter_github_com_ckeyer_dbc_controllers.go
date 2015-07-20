package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/ckeyer/dbc/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/ckeyer/dbc/controllers:UserController"],
		beego.ControllerComments{
			"Get",
			`/user/:key`,
			[]string{"get"},
			nil})

}
