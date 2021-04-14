package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	profile_handler "projectlist/application/adapters/handler"
	repository_mysql "projectlist/infrastructure/persistence/mysql/repository"
	service_profile "projectlist/usecase/profile/service"
)

func main() {
	dsn := "root:1234lupa@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		println(err)
		panic("failed to connect database")
	}

	profileRepository := repository_mysql.NewRepository(db)

	profileUseCase := service_profile.NewService(profileRepository)

	profileHandler := profile_handler.NewProfileHandler(profileUseCase)

	router := gin.Default()

	api := router.Group("/api")

	api.POST("/profile/create", profileHandler.CreateProfile)

	router.Run(":8800")

}