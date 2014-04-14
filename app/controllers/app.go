package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	//"strings"
)

type App struct {
	*revel.Controller
}

type Product struct {
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

	var productList []Product
	err := json.Unmarshal(b, &productList)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(productList[0].Name)

	greeting1 := "Aloha!"
	greeting2 := " Bien-venido a Mexico!"

	return c.Render(greeting1, greeting2, productList)
}
