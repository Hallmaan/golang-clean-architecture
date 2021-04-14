package profile_formatter

import "projectlist/domain/profile"

type ProfileFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	PhoneNumber int `json:"phoneNumber"`
}

func FormatProfile(profile *profile.Profile) ProfileFormatter {
	formatter := ProfileFormatter{
		ID:    profile.ID,
		Name:  profile.Name,
		Email: profile.Email,
		PhoneNumber: profile.PhoneNumber,
	}

	return formatter
}
