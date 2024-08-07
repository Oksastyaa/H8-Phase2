package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ngc_release2/handler"
	"ngc_release2/service"
)

func (app *application) routes() http.Handler {
	invService := service.NewInventoryService()
	invHandler := handler.NewInventoryHandler(invService, app.db)

	router := httprouter.New()
	router.GET("/inventories", invHandler.GetInventories)
	router.GET("/inventories/:id", invHandler.GetInventory)
	router.POST("/inventories", invHandler.CreateInventory)
	router.PUT("/inventories/:id", invHandler.UpdateInventory)
	router.DELETE("/inventories/:id", invHandler.DeleteInventory)

	return router
}
