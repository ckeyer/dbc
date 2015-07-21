package controllers

type DefaultController struct {
	Controller
}

func (u *DefaultController) Get() {
	u.Ctx.WriteString("Hello world!!!")
}
