package service_profile

import (
	errorDomain "projectlist/domain"
	domain "projectlist/domain/profile"
)

type inmem struct {
	m map[int64]*domain.Profile
}

func newInMem() *inmem {
	var m = map[int64]*domain.Profile{}
	return &inmem{m: m}
}

func (m *inmem) Create(profile *domain.Profile) (*domain.Profile, error) {
	m.m[int64(profile.ID)] = profile
	return profile, nil
}

func (m *inmem) Update(profile *domain.Profile) (*domain.Profile, error) {
	profile, err := m.Get(int(profile.ID))
	if err != nil {
		return profile, err
	}
	m.m[int64(profile.ID)] = profile
	return profile, nil
}

func (m *inmem) Get(id int) (*domain.Profile, error) {
	newid := int64(id)
	if m.m[newid] == nil {
		return nil, errorDomain.ErrInvalidEntity
	}
	return m.m[newid], nil
}

func (m *inmem) GetWithPagination(page int, sort string) ([]*domain.Profile, error) {
	var d []*domain.Profile
	for _, j := range m.m {
		d = append(d, j)
	}
	return d, nil
}