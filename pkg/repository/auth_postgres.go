package repository

import (
	"fmt"
	"onlineshop/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, phone_number, password_hash) values($1, $2,$3,$4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Email, user.PhoneNumber, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}

func (r *AuthPostgres) IsAdmin(userId int) (bool, error) {
	var isAdmin bool
	query := fmt.Sprintf("SELECT is_admin FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&isAdmin, query, userId)
	return isAdmin, err
}
