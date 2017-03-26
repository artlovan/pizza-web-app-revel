package controllers

import (
	"github.com/revel/revel"
	"revapp/app/models"
	"fmt"
)

type Accounts struct {
	*revel.Controller
}

func (c Accounts) Create() revel.Result {
	return c.Render()
}

func (c Accounts) CreatePost() revel.Result {
	var account models.Account

	account.FirstName = c.Params.Values.Get("firstName")
	account.LastName = c.Params.Values.Get("lastName")
	account.Address1 = c.Params.Values.Get("address1")
	account.Address2 = c.Params.Values.Get("address2")
	account.City = c.Params.Values.Get("city")
	account.State = c.Params.Values.Get("state")
	account.ZipCode = c.Params.Values.Get("zip")
	//c.Params.Bind(&account, "account")

	c.Validation.Required(account.FirstName)
	c.Validation.Required(account.LastName)
	c.Validation.Required(account.Address1)
	c.Validation.Required(account.City)
	c.Validation.Required(account.State)
	c.Validation.Required(account.ZipCode)

	c.Validation.Keep()
	c.FlashParams()

	fmt.Printf("Validation has error: %v\n", c.Validation.HasErrors())
	fmt.Printf("Account info: %#v\n", account)

	return c.RenderTemplate("accounts/create.html")
}
