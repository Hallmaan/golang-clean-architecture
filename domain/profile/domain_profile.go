package profile

import (
	"projectlist/domain"
	profile_input "projectlist/usecase/profile/input"
)

type Profile struct {
	ID int
	Name string
	Email string
	PhoneNumber int
}

func NewProfile(input profile_input.CreateProfileInput) (*Profile, error) {
	profile := &Profile{
		Name: input.Name,
		Email: input.Email,
		PhoneNumber: input.PhoneNumber,
	}

	err := profile.Validate()

	if err != nil {
		return profile, err
	}

	return profile, nil
}

func (p *Profile) Validate() error {
	if(p.Name == "" || p.Email == "" || p.PhoneNumber == 0){
		return domain.ErrInvalidEntity
	}
	return nil
}

func (p *Profile) Update(name string, email string, number int) error {
	p.Name = name
	p.Email = email
	p.PhoneNumber = number

	return p.Validate()
}