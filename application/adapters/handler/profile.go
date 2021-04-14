package profile_handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	helper "projectlist/application"
	profile_formatter "projectlist/usecase/profile/formatter"
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

	fmt.Println(input)

	newProfile, err := p.profileService.Create(input)

	if err != nil {
		response := helper.ApiResponse("Add profile failed", http.StatusBadRequest, "Error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := profile_formatter.FormatProfile(newProfile)

	response := helper.ApiResponse("Profile has been added.", http.StatusOK, "Success", formatter)

	c.JSON(http.StatusOK, response)
}
