package main

import (
	"assesment/delivery"
	"assesment/repository"
	"assesment/usecase"
)

func main() {
	repo := repository.InitRepo()
	usecase := usecase.InitUsecase(repo)
	develiry := delivery.InitDelivery(usecase)

	develiry.GetTracking()
}
