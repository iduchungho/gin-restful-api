package main

import (
	"cse-foodcourt/src/route"
	"github.com/gin-gonic/gin"
	"sync"
)

type App struct {
	r *gin.Engine
}

var lock = &sync.Mutex{}
var application *App

func GetApplication() *App {
	// check app is already exist
	if application == nil {
		// Ensure that the instance hasn't yet been
		// initialized by another thread while this one
		// has been waiting for the lock's release.
		lock.Lock()
		defer lock.Unlock()
		if application == nil {
			application = &App{
				r: gin.Default(),
			}
		} else {
			return application
		}
	}
	return application
}

func (app *App) Run() {
	if app.r != nil {
		route.CustomerRoutes(app.r)
		err := app.r.Run("localhost:8080")
		if err != nil {
			panic("Can't run gin engine")
		}

	} else {
		panic("Gin Engine not found")
	}
}
