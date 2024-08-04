package main

import (
	"fmt"
	"net/http"
	"strings"
	"web-crowdfunding/auth"
	"web-crowdfunding/campaign"
	"web-crowdfunding/handler"
	"web-crowdfunding/helper"
	"web-crowdfunding/user"

	"github.com/dgrijalva/jwt-go"
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
	authService := auth.NewAuthService()
	userHandler := handler.NewUserHandler(userService, authService)

	campaignRepository := campaign.NewRepository(db)
	campaignService := campaign.NewService(campaignRepository)

	c, _ := campaignService.FindCampaign(1)
	fmt.Println(len(c))

	//============================Campaign==================================//

	// campaigns, _ := campaignRepository.FindById(2)

	// fmt.Println("Debug")
	// fmt.Println(len(campaigns))

	// for _, campaign := range campaigns {
	// 	fmt.Println(campaign.Name)
	// 	if len(campaign.CampaignImages) > 0 {
	// 		fmt.Println(campaign.CampaignImages[0].FileName)
	// 	} else {
	// 		fmt.Println(campaign.Name, "Tidak memiliki gambar")
	// 	}

	// }

	// fmt.Println(authService.GenerateToken(1001))

	// login := user.LoginInput{
	// 	Email:    "formatter9@example.com",
	// 	Password: "formatter123",
	// }

	// user, err := userService.Login(login)//

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
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
	// user := user.User{
	// 	Nama: "Test save",
	// }

	// userRepository.Save(user)
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetID(userID)
		if err != nil {
			response := helper.ResponseAPI("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}

}
