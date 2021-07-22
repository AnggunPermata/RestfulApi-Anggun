package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	// users  []User
	users2 = map[int]*User{}
	seq    = 1
)

// ------------------ controller ----------------

//get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users2,
	})
}

//get user by Id
func GetUserController(c echo.Context) error {
	//your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users2[id])
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	//your solution here
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users2, id)
	return c.NoContent(http.StatusNoContent)
}

//update user by id
func UpdateUserController(c echo.Context) error {
	//your solution here
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	users2[id].Name = u.Name
	users2[id].Email = u.Email
	users2[id].Password = u.Password
	return c.JSON(http.StatusOK, users2[id])

}

//create new user
func CreateUserController(c echo.Context) error {
	//binding data
	u := &User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users2[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func main() {
	e := echo.New()
	//routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	e.DELETE("users/:id", DeleteUserController)
	e.GET("/users/:id", GetUserController)
	e.PUT("/users/:id", UpdateUserController)
	//start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
