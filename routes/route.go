package routes

import (
	ha "github.com/berrylradianh/go-jewelry/modules/handler/auth"
	ra "github.com/berrylradianh/go-jewelry/modules/repository/auth"
	ua "github.com/berrylradianh/go-jewelry/modules/usecase/auth"

	hp "github.com/berrylradianh/go-jewelry/modules/handler/products"
	rp "github.com/berrylradianh/go-jewelry/modules/repository/products"
	up "github.com/berrylradianh/go-jewelry/modules/usecase/products"

	hpc "github.com/berrylradianh/go-jewelry/modules/handler/product_categories"
	rpc "github.com/berrylradianh/go-jewelry/modules/repository/product_categories"
	upc "github.com/berrylradianh/go-jewelry/modules/usecase/product_categories"

	db "github.com/berrylradianh/go-jewelry/databases"

	svc "github.com/berrylradianh/go-jewelry/modules/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

var (
	authRepo    ra.Repository
	authHandler ha.Handler
	authUsecase ua.Usecase

	productRepo    rp.Repository
	productHandler hp.Handler
	productUsecase up.Usecase

	productCategoryRepo    rpc.Repository
	productCategoryHandler hpc.Handler
	productCategoryUsecase upc.Usecase
)

func declare() {
	authRepo = ra.Repository{DB: db.DB}
	authUsecase = ua.Usecase{Repository: authRepo}
	authHandler = ha.Handler{Usecase: &authUsecase}

	productRepo = rp.Repository{DB: db.DB}
	productUsecase = up.Usecase{Repository: productRepo}
	productHandler = hp.Handler{Usecase: &productUsecase}

	productCategoryRepo = rpc.Repository{DB: db.DB}
	productCategoryUsecase = upc.Usecase{Repository: productCategoryRepo}
	productCategoryHandler = hpc.Handler{Usecase: &productCategoryUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.Validator = &svc.CustomValidator{Validator: validator.New()}

	account := e.Group("/account")
	account.POST("/login", authHandler.LoginUser())
	account.POST("/register", authHandler.RegisterUser())
	account.POST("/logout", authHandler.LogoutUser())

	product := e.Group("/products")
	product.GET("", productHandler.GetAllProducts())
	product.GET("/:id", productHandler.GetProductById())
	product.POST("", productHandler.CreateProduct())
	product.PUT("/:id", productHandler.UpdateProduct())
	product.DELETE("/:id", productHandler.DeleteProduct())

	productCategory := e.Group("/products/categories")
	productCategory.GET("", productCategoryHandler.GetAllProductCategories())
	productCategory.GET("/:id", productCategoryHandler.GetProductCategoryById())
	productCategory.POST("", productCategoryHandler.CreateProductCategory())
	productCategory.PUT("/:id", productCategoryHandler.UpdateProductCategory())
	productCategory.DELETE("/:id", productCategoryHandler.DeleteProductCategory())

	return e
}
