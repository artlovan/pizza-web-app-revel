package controllers

import (
	"github.com/revel/revel"
	"revapp/app/models"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var rndGen = rand.New(rand.NewSource(time.Now().Unix()))
const change = 0.5

type Orders struct {
	*revel.Controller
	winner bool
}

func init()  {
	revel.InterceptMethod((*Orders).randomDrawing, revel.BEFORE)
}

func (c Orders) randomDrawing() revel.Result {
	if rndGen.Float32() < change {
		c.winner = true
	}

	return nil
}

func (c Orders) decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func (c Orders) Create() revel.Result {
	return c.Render()
}

func (c Orders) GetPayment(orderId int) revel.Result {
	println("Order ID: ", orderId)
	return c.RenderTemplate("orders/payment.html")
}

func (c Orders) ApiCreate() revel.Result {
	var order models.Order
	//dec := json.NewDecoder(c.Request.Body)
	//dec.Decode(&order)
	c.decode(c.Request.Body, &order)

	if c.winner {
		fmt.Println("You have won a discount!")
	} else {
		fmt.Println("You didn't win a discount!")
	}

	fmt.Printf("The Order data: %#v\n", order)

	return c.RenderText("OK")
	//resp := map[string]int{"status": 200}
	//return c.RenderJson(resp)
}