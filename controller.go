package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Repo *Repository
}

// POST /user
func (uc *UserController) CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	// dodaj u bazu
	uc.Repo.CreateUser(user)
	c.String(http.StatusOK, "Korisnik kreiran")
}

// UPDATE /user
func (uc *UserController) UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id") // uzima id iz url-a
	c.BindJSON(&user)

	// pretvara id iz str u int
	intId, _ := strconv.ParseInt(id, 10, 64)
	user.ID = intId

	// azurira u bazi
	uc.Repo.UpdateUser(user)
	c.String(http.StatusOK, "Korisnik updatean")
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	uc.Repo.DeleteUser(id)
	c.String(http.StatusOK, "Korisnik obrisan")

}
