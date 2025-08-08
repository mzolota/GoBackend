package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// klasa komunikaciju s bazom -> CRUD
type Repository struct {
	Conn *pgx.Conn
}

// dodavanje korisnika u bazu
func (r *Repository) CreateUser(u User) (int64, error) {
	var id int64
	query := `INSERT INTO "Users" (name, lastname, phone, email) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.Conn.QueryRow(context.Background(), query, u.Name, u.Lastname, u.Phone, u.Email).Scan(&id)
	if err != nil {

		return 0, err
	}
	return id, nil
}

// brisanje korisnika u bazi
func (r *Repository) DeleteUser(id int64) error {
	commandTag, err := r.Conn.Exec(context.Background(), `DELETE FROM "Users" WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("nema korisnika s ID %d", id)
	}

	return nil
}

// update korisnika u bazi

func (r *Repository) UpdateUser(u User) error {
	query := `UPDATE "Users" SET name=$1, lastname=$2, phone=$3, email=$4 WHERE id=$5`
	commandTag, err := r.Conn.Exec(context.Background(), query, u.Name, u.Lastname, u.Phone, u.Email, u.ID)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("korisnik s ID %d nije pronađen", u.ID)
	}
	return nil
}
