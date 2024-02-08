package routes

import (
	taskRoutes "app/src/modules/tasks/infra/http/routes"
	userRoutes "app/src/modules/users/infra/http/routes"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	taskRoutes.TaskRoutes(router)
	userRoutes.UserRoutes(router)
}
