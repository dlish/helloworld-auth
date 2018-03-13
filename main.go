package main

import (
	"github.com/dlish/helloworld-auth/db"
	"github.com/dlish/helloworld-auth/server"
)

func main() {
	db, err := db.NewDbConnection()
	if err != nil {
		panic(err)
	}
	server := server.NewServer(db)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
