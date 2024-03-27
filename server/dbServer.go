package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"gram/cmd/controllers"
	"gram/repositories"
	"gram/services"
	"log"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	userController    *controllers.UserController
	socialController  *controllers.SocialController
	photoController   *controllers.PhotoController
	commentController *controllers.CommentController
}

func initHttpServer(config *viper.Viper, db *gorm.DB) *HttpServer {
	userRepository := repositories.NewUserRepository(db)
	socialRepository := repositories.NewSocialRepository(db)
	photoRepository := repositories.NewPhotoRepository(db)
	commentRepository := repositories.NewCommentRepository(db)
	userService := services.NewUserService(userRepository)
	socialService := services.NewSocialService(socialRepository)
	photoService := services.NewPhotoService(photoRepository)
	commentService := services.NewCommentService(commentRepository)
	jwtService := services.NewJWTService()
	userController := controllers.NewUserController(userService)
	socialController := controllers.NewSocialController(socialService)
	photoController := controllers.NewPhotoController(photoService)
	commentController := controllers.NewCommentController(commentService)
	jwtController := controllers.NewJWTController(jwtService)

	router := gin.Default()

	authMiddleware := jwtController.JWTMiddleware()

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(authMiddleware)
	{

		socialGroup := protectedRoutes.Group("/socials")
		{
			socialGroup.GET("", socialController.GetSocials)
			socialGroup.GET("/:id", socialController.GetSocial)
			socialGroup.POST("", socialController.CreateSocial)
			socialGroup.PUT("/:id", socialController.UpdateSocial)
			socialGroup.DELETE("/:id", socialController.DeleteSocial)

			photoGroup := socialGroup.Group("/photos")
			{
				photoGroup.GET("", photoController.GetPhotos)
				photoGroup.GET("/:id", photoController.GetPhoto)
				photoGroup.POST("", photoController.CreatePhoto)
				photoGroup.PUT("/:id", photoController.UpdatePhoto)
				photoGroup.DELETE("/:id", photoController.DeletePhoto)

				commentGroup := photoGroup.Group("/comments")
				{
					commentGroup.GET("", commentController.GetComments)
					commentGroup.GET("/:id", commentController.GetComment)
					commentGroup.POST("", commentController.CreateComment)
					commentGroup.PUT("/:id", commentController.UpdateComment)
					commentGroup.DELETE("/:id", commentController.DeleteComment)
				}
			}
		}
	}

	router.POST("/login", userController.CreateUser)

	return &HttpServer{
		config:            config,
		router:            router,
		userController:    userController,
		socialController:  socialController,
		photoController:   photoController,
		commentController: commentController,
	}
}

func (s *HttpServer) Start() {
	err := s.router.Run(s.config.GetString("server.port"))
	if err != nil {
		log.Fatal("Error while starting server", err)
	}
}
