package main

import (
	"cse-foodcourt/src/database"
	"github.com/gin-gonic/gin"
)

type App struct {
	r *gin.Engine
}

func (app *App) Init() {
	app.r = gin.Default()
	database.DatabaseConection()
}

func (app *App) Run() {
	if app.r != nil {

	} else {
		panic("Run time error")
	}
}
