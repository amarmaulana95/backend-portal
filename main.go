package main

import (
	"log"
	"net/http"
	"strings"

	article "github.com/amarmaulana95/backend-portal/artikel"
	"github.com/amarmaulana95/backend-portal/auth"
	"github.com/amarmaulana95/backend-portal/handler"
	"github.com/amarmaulana95/backend-portal/helper"
	user "github.com/amarmaulana95/backend-portal/users"
	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/portal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	articleRepository := article.NewRepository(db)

	userService := user.NewService(userRepository)
	articleService := article.NewService(articleRepository)

	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	articleHandler := handler.NewArticleHandler(articleService)

	//api group versioning (untuk dkebutuhan aja)
	router := gin.Default()
	api := router.Group("api/v1") //jika ada pengembangan makan bisa dijadikan v2

	//--------------------------------USER AREA ENDPOINT------------------------------------//

	api.POST("/users", userHandler.RegisterUser) // endpoint register
	api.POST("/sessions", userHandler.Login)     // endpoint login

	//--------------------------------ARTIKEL AREA ENDPOINT------------------------------------//
	//dengan kondisi user roel => user
	api.POST("/article", authMiddleware(authService, userService), articleHandler.CreateArticle)      //post data artikel by token
	api.POST("/article-images", authMiddleware(authService, userService), articleHandler.UploadImage) // post image
	api.GET("/uarticle", authMiddleware(authService, userService), articleHandler.GetArticleUser)     // get data artikel by token
	api.GET("/uarticle/:id", articleHandler.GetArticle)                                               //get data detail by artikel id

	//------------------------------ADMIN AREA-----------------------------------------------//
	//dengan kondisi user roel => admin
	api.GET("/article", authMiddleware(authService, userService), articleHandler.GetArticlesAdmin) // get all artikel + point reward
	api.GET("/users", authMiddleware(authService, userService), userHandler.FetchUser)             //get all user
	api.PUT("/point/:id", authMiddleware(authService, userService), articleHandler.UpdateArticle)  //update point reward artikel

	router.Static("/images", "./images") //get data image static url

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc { //Membuat midlleware

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// jika di dalam string auth header ada kata "Bearer" ?
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			// abourt white status hentikan
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			// abourt white status hentikan
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			// abourt white status hentikan
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			// abourt white status hentikan
			return
		}
		c.Set("currentUser", user)
	}
	//  intinya kita mengambil Authorization : Bearer tokentoken1234
	//  dari header  Authorization, mabil nilai tokennya aja, lalu di validasi, dan ambil user_id lewat service
	//  set konteks isinya user
}
