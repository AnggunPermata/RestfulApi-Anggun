package routes

import (
	"project/anggun/constants"
	"project/anggun/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//user routing
	/*
		e.GET("/users", controllers.GetUsersController)
		e.GET("/users/:id", controllers.GetUserController)
	*/
	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)
	/*
		e.DELETE("/users/:id", controllers.DeleteUserController)
		e.PUT("/users/:id", controllers.UpdateUserController)
	*/
	//JWT group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.DELETE("/users:id", controllers.DeleteUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	// r.GET("/users/:id", controllers.GetUserDetailControllers)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	return e
}
