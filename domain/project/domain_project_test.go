package project_test

import (
	"github.com/stretchr/testify/assert"
	errorDomain "projectlist/domain"
	dProfile "projectlist/domain/profile"
	dProject "projectlist/domain/project"
	"testing"
)

func TestNewProject(t *testing.T){
	domainProfile, err := dProfile.NewProfile("Agung", "Agunghallmanmaliki@gmail.com", 0123)
	project, err := dProject.NewProject("MsightCascade", domainProfile, "project.jpg", "Desc project")
	assert.Nil(t, err)
	assert.Equal(t, "MsightCascade", project.Name)
	assert.Equal(t, domainProfile, project.Profile)
	assert.Equal(t, "project.jpg", project.Image)
	assert.Equal(t, "Desc project", project.Description)
}

func TestProfileValidate(t *testing.T){
	type TestCase struct {
		Name string
		Profile *dProfile.Profile
		Image string
		Description string
		want error
	}

	domainProfile, _ := dProfile.NewProfile("Agung", "Agunghallmanmaliki@gmail.com", 0123)
	print(domainProfile.Name)

	tests := []TestCase{
		{
			Name:  "Agung",
			Profile: domainProfile,
			Image: "project.jpg",
			Description: "Desc project",
			want:  nil,
		},
		{
			Name:  "Agung",
			Profile: nil,
			Image: "",
			Description: "Desc project",
			want:  errorDomain.ErrInvalidEntity,
		},
		{
			Name:  "Agung",
			Profile: domainProfile,
			Image: "",
			Description: "Desc project",
			want:  errorDomain.ErrInvalidEntity,
		},
		{
			Name:  "Agung",
			Profile: domainProfile,
			Image: "project.jpg",
			Description: "",
			want:  errorDomain.ErrInvalidEntity,
		},
	}

	for _, data := range tests {
		_, err := dProject.NewProject(data.Name, data.Profile, data.Image, data.Description)
		assert.Equal(t, data.want, err)
	}

}