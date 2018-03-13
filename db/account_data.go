package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Account struct {
	Db    *sql.DB
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewAccount(db *sql.DB) *Account {
	return &Account{
		Db: db,
	}
}

func (account *Account) Fetch(id int) (err error) {
	fmt.Printf("Fetching id: %d", id)
	err = account.Db.QueryRow(
		"select id, name, email from accounts where id = $1",
		id).Scan(&account.Id, &account.Name, &account.Email)
	return
}
