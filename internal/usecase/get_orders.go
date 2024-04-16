package usecase

import "github.com/obrunogonzaga/cleanArch/internal/entity"

type FindAllOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewFindAllOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *FindAllOrdersUseCase {
	return &FindAllOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *FindAllOrdersUseCase) Execute() ([]entity.Order, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
