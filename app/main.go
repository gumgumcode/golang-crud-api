package main

import (

	"github.com/omkarscode/restapi/controller"
	"github.com/omkarscode/restapi/model"
)

func main() {
	model.Init()
	controller.Start()
}