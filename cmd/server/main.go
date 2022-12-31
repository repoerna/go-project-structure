package main

import (
	"go-boilerplate-api/internal/app"
	_ "go-boilerplate-api/pkg/logger"
)

func main() {

	app := app.New()
	app.Run()

}
