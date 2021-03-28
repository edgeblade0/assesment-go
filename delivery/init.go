package delivery

import "assesment/usecase"

type deliver struct {
	usecase usecase.Usecase
}

type Delivery interface {
	GetTracking()
}

func InitDelivery(u usecase.Usecase) Delivery {
	return &deliver{
		usecase: u,
	}
}
