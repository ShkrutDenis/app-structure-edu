package main

import (
	"database/sql"
	"errors"
)

type OrderEntity struct{}

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Insert(Order *OrderEntity) error {
	return errors.New("logic not implemented")
}

func (r *OrderRepository) Get(id int64) (*OrderEntity, error) {
	return nil, errors.New("logic not implemented")
}

func (r *OrderRepository) Update(Order *OrderEntity) error {
	return errors.New("logic not implemented")
}
