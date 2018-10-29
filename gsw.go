package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"time"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)

//----------
// Handlers
//----------

func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

//curl -H "SLEEPTIME: 3"  http://localhost:8080/sleep  -v
func toSleep(c echo.Context) error {
	//sleeptime := c.Request().Header.Get("SLEEPTIME")
	sleeptime := c.Request().Header.Get("SLEEPTIME")
	ist, _ := strconv.Atoi(sleeptime)
	fmt.Println(sleeptime)
	time.Sleep(time.Duration(ist) * time.Second)
	return c.JSON(http.StatusOK, sleeptime)
}

//curl -H "SLEEPTIME: 3" -H "KEY:UUU" http://localhost:8080/sleep  -v
func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		fmt.Println("inside Process1")
		getKey := c.Request().Header.Get("KEY")
		if getKey != "UUU" {
			fmt.Println("direct return")
			return c.JSON(http.StatusForbidden, getKey)
		}
		fmt.Println("next")
		return next(c)
	}
}

/*
func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		fmt.Println("inside Process")
		return nil
	}
}
*/

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(Process)

	// Routes
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.GET("sleep", toSleep)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
