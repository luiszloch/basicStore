package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting1 := "Aloha!"
	greeting2 := "Bien-venido a Mexico!"
	return c.Render(greeting1, greeting2)
}
