package session

import (
	"github.com/jmoiron/sqlx"
	"github.com/srinathgs/mysqlstore"
)

func NewSession(db *sqlx.DB) (*mysqlstore.MySQLStore, error) {
	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 60*60*24*14, []byte("secret-token"))
	if err != nil {
		panic(err)
	}

	return store, err
}
