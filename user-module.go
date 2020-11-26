package user_module

import (
	"github.com/gorilla/mux"
	"github.com/verottaa/user-module/controllers"
)

func RegistryControllers(router *mux.Router) {
	var controller = controllers.CreateController()
	controller.InitController(router)
}
