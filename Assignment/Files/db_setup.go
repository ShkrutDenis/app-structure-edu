package main

import "database/sql"

func NewDbConnection(config any) (*sql.DB, error) {
	return sql.Open("driver", "secure_text")
}
