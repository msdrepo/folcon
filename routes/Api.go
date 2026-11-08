package routes

import (
	"github.com/msdrepo/folcon/controllers"

	folcon "github.com/msdrepo/framework"
)

func Route() *folcon.Engine {

	router := folcon.New()

	// api routes ----------------

	public := router.Group("/api")

	public.GET("/user", controllers.UserCall)
	public.GET("/test", controllers.Greet)
	public.POST("/login", controllers.Login)
	public.GET("/user", controllers.GetParamTest)

	admin := router.Group("/admin")

	admin.GET("/test", controllers.TestCall)
	admin.POST("/login", controllers.AdminLogin)
	// admin.GET("/user", controllers.AdminLogin)

	return router

}
