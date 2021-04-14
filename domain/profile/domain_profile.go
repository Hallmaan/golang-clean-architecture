package profile

import (
	"math/rand"
	"projectlist/domain"
)

type Profile struct {
	ID int
	Name string
	Email string
	PhoneNumber int
}

func NewProfile(name string, email string, phoneNumber int) (*Profile, error) {
	profile := &Profile{
		ID: rand.Int(),
		Name: name,
		Email: email,
		PhoneNumber: phoneNumber,
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