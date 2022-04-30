package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *solution) newFront() front {
	return front{
		Config: app.Config.Front,
	}
}

func (f *front) setupRoutes() error {
	f.loadTemplates(appAssetsRoot + "templates/*")
	f.Gin.Static("/assets", appAssetsRoot+"assets")

	f.Gin.GET("/", func(c *gin.Context) {
		f.renderTemplate(
			c,
			http.StatusOK,
			"main.html",
			f.getBasicHeaders(),
		)
	})

	f.Gin.NoRoute(func(c *gin.Context) {
		f.renderTemplate(
			c,
			http.StatusNotFound,
			"404.html",
			f.getBasicHeaders(),
		)
	})

	return nil
}

func (f *front) getBasicHeaders() gin.H {
	return gin.H{
		"version": f.Config.Version,
	}
}

func (f *front) runRouter() error {
	go func() {
		err := f.Gin.Run(":" + f.Config.Port)
		if err != nil {
			log.Fatalln(err)
		}
	}()
	return nil
}

func (f *front) loadTemplates(pattern string) {
	if f.Gin == nil {
		log.Fatalln("gin is not initiated")
	}
	f.Gin.LoadHTMLGlob(pattern)
}

func (f *front) renderTemplate(c *gin.Context, code int, name string, obj interface{}) {
	c.HTML(code, name, obj)
}

func (app *solution) openFront() error {
	time.Sleep(time.Millisecond * 150) // wait for gin init
	openBrowser("http://127.0.0.1:" + app.Config.Front.Port)
	return nil
}

func (app *solution) setupFront() error {
	app.Front = app.newFront()

	// setup debug mode
	if app.Config.Front.IsDebug {
		log.Println("debug mode activated")
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// setup gin
	app.Front.Gin = gin.New()

	// setup routes
	return checkErrors(
		app.Front.setupRoutes,
		app.Front.runRouter,
	)
}
