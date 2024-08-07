package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ngc/handler"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	// Heroes
	router.POST("/heroes", handler.CreateHeroHandler(app.db))
	router.GET("/heroes", handler.GetAllHeroesHandler(app.db))
	router.GET("/heroes/:id", handler.GetHeroByIDHandler(app.db))
	router.PUT("/heroes/:id", handler.UpdateHeroHandler(app.db))
	router.DELETE("/heroes/:id", handler.DeleteHeroHandler(app.db))

	// Villains
	router.POST("/villains", handler.CreateVillainHandler(app.db))
	router.GET("/villains", handler.GetAllVillainsHandler(app.db))
	router.GET("/villains/:id", handler.GetVillainByIDHandler(app.db))
	router.PUT("/villains/:id", handler.UpdateVillainHandler(app.db))
	router.DELETE("/villains/:id", handler.DeleteVillainHandler(app.db))
	return router
}
