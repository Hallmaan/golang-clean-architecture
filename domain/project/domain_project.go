package project

import (
	"math/rand"
	"projectlist/domain"
	dProfile "projectlist/domain/profile"
)

type Project struct {
	ID int
	Name string
	Profile *dProfile.Profile
	Image string
	Description string
}

func NewProject(name string, profile *dProfile.Profile, image string, description string) (*Project, error) {
	project := &Project{
		ID: rand.Int(),
		Name: name,
		Profile: profile,
		Image: image,
		Description: description,
	}

	err := project.Validate()

	if err != nil {
		return project, err
	}

	return project, nil
}

func (p *Project) Validate() error {

	if p.Name == "" || p.Image == "" || p.Description == "" || p.Profile == nil{
		return domain.ErrInvalidEntity
	}

	return nil
}