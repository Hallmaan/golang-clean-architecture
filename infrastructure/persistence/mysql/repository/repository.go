package repository_mysql

import (
	"gorm.io/gorm"
	"projectlist/domain/profile"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(profile *profile.Profile) (*profile.Profile, error) {
	err := r.db.Create(&profile).Error
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) Update(profile *profile.Profile) (*profile.Profile, error) {
	err := r.db.Updates(&profile).Error
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (r *repository) Get(id int) (*profile.Profile, error) {

	return &profile.Profile{}, nil

}

func (r *repository) GetWithPagination(page int, sort string) ([]*profile.Profile, error) {
	return nil, nil
}
