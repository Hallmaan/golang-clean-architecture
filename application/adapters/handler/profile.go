package profile_handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	helper "projectlist/application"
	profile_input "projectlist/usecase/profile/input"
	service_profile "projectlist/usecase/profile/service"
)

type profileHandler struct {
	profileService service_profile.UseCase
}

func NewProfileHandler(profileService service_profile.UseCase) *profileHandler {
	return &profileHandler{profileService}
}

func (p *profileHandler) CreateProfile(c *gin.Context) {
	var input profile_input.CreateProfileInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Add profile failed", http.StatusUnprocessableEntity, "Error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//newProfile, err := p.profileService.Create(input)
}
