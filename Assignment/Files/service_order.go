package main

// list of functions to manipulate with Order entity

type Order struct{}

type OrderService struct {
	dependencies any
}

func NewOrderService(deps any) *OrderService {
	return &OrderService{dependencies: deps}
}

func (s *OrderService) OpenOrder(data any) Order {
	return Order{}
}

func (s *OrderService) CloseOrder(order Order) {}

func (s *OrderService) UpdateOrder(data any, order Order) {}

func (s *OrderService) GetOrderStatus(id int64) string {
	return "done"
}
