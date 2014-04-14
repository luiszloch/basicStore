package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
)

type App struct {
	*revel.Controller
}

type ProductList struct {
	Products []Products
}

type Products struct {
	Id          int
	Img         string
	Name        string
	Description string
}

func (c App) Index() revel.Result {
	b, r := ioutil.ReadFile("public/json/descriptions.json")
	if r != nil {
		fmt.Println(r)
	}

	var m ProductList
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.Products[0].Name)

	greeting1 := "Aloha!"
	greeting2 := " Bien-venido a Mexico!"

	return c.Render(greeting1, greeting2, m)
}
