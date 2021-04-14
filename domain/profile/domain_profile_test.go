package profile_test

import (
	"github.com/stretchr/testify/assert"
	errorDomain "projectlist/domain"
	domain "projectlist/domain/profile"
	profile_input "projectlist/usecase/profile/input"
	"testing"
)

func TestNewProfile(t *testing.T){
	input := profile_input.CreateProfileInput{"Agung", "Agunghallmanmaliki@gmail.com", 1}

	profile, err := domain.NewProfile(input)
	assert.Nil(t, err)
	assert.Equal(t, "Agung", profile.Name)
	assert.Equal(t, "Agunghallmanmaliki@gmail.com", profile.Email)
	assert.Equal(t, 1, profile.PhoneNumber)
}

func TestProfileValidate(t *testing.T){
	type TestCase struct {
		Name string
		Email string
		PhoneNumber int
		want error
	}

	tests := []TestCase{
		{
			Name:  "Agung",
			Email: "Agunghallmanmaliki@gmail.com",
			PhoneNumber: 0123,
			want:  nil,
		},
		{
			Name:  "",
			Email: "Agunghallmanmaliki@gmail.com",
			PhoneNumber: 0123,
			want:  errorDomain.ErrInvalidEntity,
		},
		{
			Name:  "Agung",
			Email: "",
			PhoneNumber: 0123,
			want:  errorDomain.ErrInvalidEntity,
		},
		{
			Name:  "",
			Email: "Agunghallmanmaliki@gmail.com",
			PhoneNumber: 0,
			want:  errorDomain.ErrInvalidEntity,
		},
	}

	for _, data := range tests {
		input := profile_input.CreateProfileInput{data.Name, data.Email, data.PhoneNumber}
		_, err := domain.NewProfile(input)
		assert.Equal(t, data.want, err)
	}

}
