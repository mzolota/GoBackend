package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(repo *Repository) *gin.Engine {

	//POST - dodaje korisnika u bazu
	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		repo.CreateUser(user)
		c.String(http.StatusOK, "korisnik kreiran")
	})

	//PUT - update korisnika po ID-u
	r.PUT("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id") // Dohvati id iz urla
		id, _ := strconv.ParseInt(idParam, 10, 64)
		var user User
		c.BindJSON(&user)
		user.ID = id
		repo.UpdateUser(user)
		c.String(http.StatusOK, "korisnik azuriran")

	})
	return r
}

func StartWebserver() {
	Connect()
	repo := &Repository{Conn: Conn}
	router := SetUpRouter(repo)
	router.Run(":8080")
}
