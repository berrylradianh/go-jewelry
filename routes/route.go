package routes

import (
	ah "github.com/berrylradianh/go-jewelry/modules/handler/auth"
	ar "github.com/berrylradianh/go-jewelry/modules/repository/auth"
	au "github.com/berrylradianh/go-jewelry/modules/usecase/auth"

	db "github.com/berrylradianh/go-jewelry/databases"

	svc "github.com/berrylradianh/go-jewelry/modules/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var (
	// userRepo    user_repo.Repository
	// userHandler user_handler.Handler
	// userUsecase user_usecase.Usecase

	authRepo    ar.Repository
	authHandler ah.Handler
	authUsecase au.Usecase
)

func declare() {
	// userRepo = user_repo.Repository{DB: db.DB}
	// userUsecase = user_usecase.Usecase{Repo: userRepo}
	// userHandler = user_handler.Handler{Usecase: userUsecase}

	authRepo = ar.Repository{DB: db.DB}
	authUsecase = au.Usecase{Repository: authRepo}
	authHandler = ah.Handler{Usecase: &authUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.Validator = &svc.CustomValidator{Validator: validator.New()}

	// user := e.Group("/users")
	// user.GET("", userHandler.GetAllUsers())
	// user.GET("/:id", userHandler.GetUser())
	// user.POST("", userHandler.CreateUser())
	// user.DELETE("/:id", userHandler.DeleteUser())
	// user.PUT("/:id", userHandler.UpdateUser())

	account := e.Group("/account")
	account.POST("/login", authHandler.LoginUser())
	account.POST("/register", authHandler.RegisterUser())
	account.POST("/logout", authHandler.LogoutUser())

	return e
}
