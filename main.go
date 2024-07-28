package main

import (
	"web-crowdfunding/handler"
	"web-crowdfunding/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/web-crowdfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// login := user.LoginInput{
	// 	Email:    "formatter9@example.com",
	// 	Password: "formatter123",
	// }

	// user, err := userService.Login(login)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(user.Nama)

	// userByEmail, err := userRepository.FindByEmail("ryujin@example.com")

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan")
	// } else {
	// 	fmt.Println(userByEmail.Nama)
	// }

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/check_email", userHandler.IsEmailAvailable)

	router.Run()
	// user := user.User{
	// 	Nama: "Test save",
	// }

	// userRepository.Save(user)
}
