package repository

type repo struct {
}

type Repository interface {
	GetTracking() (string, int, error)
}

func InitRepo() Repository {
	return &repo{}
}
