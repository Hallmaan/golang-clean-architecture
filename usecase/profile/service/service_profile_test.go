package service_profile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Create(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	_, err := service.Create("Name", "Agunghallmanmaliki@gmail.com", 14045)
	assert.Nil(t, err)
}

func TestService_Get(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	profile, _ := service.Create("Name", "Agunghallmanmaliki@gmail.com", 14045)

	product, err := service.Get(profile.ID)
	assert.Nil(t, err)
	assert.Equal(t, "Name", product.Name)
}

func TestService_Update(t *testing.T) {
	inmem := newInMem()
	service := NewService(inmem)
	profile, _ := service.Create("Name", "Agunghallmanmaliki@gmail.com", 14045)

	profile, err := service.Update(profile.ID, "Name", "Test", 14022)
	assert.Nil(t, err)
}