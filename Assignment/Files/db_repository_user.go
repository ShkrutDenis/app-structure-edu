package main

import (
	"database/sql"
	"errors"
)

type UserEntity struct{}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(user *UserEntity) error {
	return errors.New("logic not implemented")
}

func (r *UserRepository) Get(id int64) (*UserEntity, error) {
	return nil, errors.New("logic not implemented")
}

func (r *UserRepository) Update(user *UserEntity) error {
	return errors.New("logic not implemented")
}
