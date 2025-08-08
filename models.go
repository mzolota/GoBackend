package main

import "time"

// atributi bi trebali poceti velikim slovom --> ako budem radio exporte
// koristen string jer broj pocinje s 0 , ako je +385 nece prihvatiti int zbog plusa
// “ backticks za JSON prijenos
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
