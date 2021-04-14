package service_profile

import domain "projectlist/domain/profile"

type Writer interface {
	Create(profile *domain.Profile) (*domain.Profile, error)
	Update(profile *domain.Profile) (*domain.Profile, error)
}

type Reader interface {
	Get(id int) (*domain.Profile, error)
	GetWithPagination(page int, sort string) ([]*domain.Profile, error)
}

type Repository interface {
	Writer
	Reader
}

type UseCase interface {
	Get(id int) (*domain.Profile, error)
	Create(name string, email string, phoneNumber int) (*domain.Profile, error)
	Update(id int, name string, email string, phoneNumber int) (*domain.Profile, error)
	GetWithPagination(page int, sort string) ([]*domain.Profile, error)
}
