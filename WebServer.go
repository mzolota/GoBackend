package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(repo *Repository) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// POST - registracija korisnika
	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Neispravan format podataka"})
			return
		}

		if _, err := repo.CreateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusOK, "Korisnik kreiran")
	})

	// PUT - update korisnika po ID-u
	r.PUT("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Neispravan ID"})
			return
		}

		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Neispravan format podataka"})
			return
		}

		user.ID = id
		if err := repo.UpdateUser(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.String(http.StatusOK, "Korisnik ažuriran")
	})

	// POST - prijava korisnika
	r.POST("/login", func(c *gin.Context) {
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.BindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Neispravan zahtjev"})
			return
		}

		user, err := repo.GetUserByEmail(creds.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Korisnik ne postoji"})
			return
		}

		if user.Password != creds.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Pogrešna lozinka"})
			return
		}

		// Sakrij lozinku prije slanja odgovora
		user.Password = ""

		c.JSON(http.StatusOK, gin.H{
			"message": "Prijava uspješna",
			"user":    user,
		})
	})

	return r
}

func StartWebserver() {
	err := Connect()
	if err != nil {
		fmt.Println("Neuspjela konekcija na bazu:", err)
		return
	}

	repo := &Repository{Conn: Conn}
	router := SetUpRouter(repo)
	router.Run(":8080")
}
