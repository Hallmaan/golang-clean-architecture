package service_profile

import domain "projectlist/domain/profile"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Get(id int) (*domain.Profile, error) {
	return s.r.Get(id)
}

func (s *Service) Create(name string, email string, phoneNumber int) (*domain.Profile, error) {
	profile, err := domain.NewProfile(name, email, phoneNumber)
	if err != nil {
		return profile, err
	}

	return s.r.Create(profile)
}

func (s *Service) Update(id int, name string, email string, phoneNumber int) (*domain.Profile, error) {
	profile, err := s.r.Get(id)
	if err != nil {
		return profile, err
	}

	err = profile.Update(name, email, phoneNumber)
	if err != nil {
		return profile, err
	}

	return s.r.Update(profile)
}

func (s *Service) GetWithPagination(page int, sort string) ([]*domain.Profile, error) {
	return s.r.GetWithPagination(page, sort)
}
