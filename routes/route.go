package routes

import (
	ha "github.com/berrylradianh/go-jewelry/modules/handler/auth"
	ra "github.com/berrylradianh/go-jewelry/modules/repository/auth"
	ua "github.com/berrylradianh/go-jewelry/modules/usecase/auth"

	hp "github.com/berrylradianh/go-jewelry/modules/handler/products"
	rp "github.com/berrylradianh/go-jewelry/modules/repository/products"
	up "github.com/berrylradianh/go-jewelry/modules/usecase/products"

	hu "github.com/berrylradianh/go-jewelry/modules/handler/users"
	ru "github.com/berrylradianh/go-jewelry/modules/repository/users"
	uu "github.com/berrylradianh/go-jewelry/modules/usecase/users"

	hr "github.com/berrylradianh/go-jewelry/modules/handler/roles"
	rr "github.com/berrylradianh/go-jewelry/modules/repository/roles"
	ur "github.com/berrylradianh/go-jewelry/modules/usecase/roles"

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

	productCategoryRepo    rp.Repository
	productCategoryHandler hp.Handler
	productCategoryUsecase up.Usecase

	productMaterialRepo    rp.Repository
	productMaterialHandler hp.Handler
	productMaterialUsecase up.Usecase

	productDescriptionRepo    rp.Repository
	productDescriptionHandler hp.Handler
	productDescriptionUsecase up.Usecase

	userRepo    ru.Repository
	userHandler hu.Handler
	userUsecase uu.Usecase

	userDetailRepo    ru.Repository
	userDetailHandler hu.Handler
	userDetailUsecase uu.Usecase

	roleRepo    rr.Repository
	roleHandler hr.Handler
	roleUsecase ur.Usecase
)

func declare() {
	authRepo = ra.Repository{DB: db.DB}
	authUsecase = ua.Usecase{Repository: authRepo}
	authHandler = ha.Handler{Usecase: &authUsecase}

	productRepo = rp.Repository{DB: db.DB}
	productUsecase = up.Usecase{Repository: productRepo}
	productHandler = hp.Handler{Usecase: &productUsecase}

	productCategoryRepo = rp.Repository{DB: db.DB}
	productCategoryUsecase = up.Usecase{Repository: productCategoryRepo}
	productCategoryHandler = hp.Handler{Usecase: &productCategoryUsecase}

	productMaterialRepo = rp.Repository{DB: db.DB}
	productMaterialUsecase = up.Usecase{Repository: productMaterialRepo}
	productMaterialHandler = hp.Handler{Usecase: &productMaterialUsecase}

	productDescriptionRepo = rp.Repository{DB: db.DB}
	productDescriptionUsecase = up.Usecase{Repository: productDescriptionRepo}
	productDescriptionHandler = hp.Handler{Usecase: &productDescriptionUsecase}

	userRepo = ru.Repository{DB: db.DB}
	userUsecase = uu.Usecase{Repository: userRepo}
	userHandler = hu.Handler{Usecase: &userUsecase}

	userDetailRepo = ru.Repository{DB: db.DB}
	userDetailUsecase = uu.Usecase{Repository: userDetailRepo}
	userDetailHandler = hu.Handler{Usecase: &userDetailUsecase}

	roleRepo = rr.Repository{DB: db.DB}
	roleUsecase = ur.Usecase{Repository: roleRepo}
	roleHandler = hr.Handler{Usecase: &roleUsecase}
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

	productMaterial := e.Group("/products/materials")
	productMaterial.GET("", productMaterialHandler.GetAllProductMaterials())
	productMaterial.GET("/:id", productMaterialHandler.GetProductMaterialById())
	productMaterial.POST("", productMaterialHandler.CreateProductMaterial())
	productMaterial.PUT("/:id", productMaterialHandler.UpdateProductMaterial())
	productMaterial.DELETE("/:id", productMaterialHandler.DeleteProductMaterial())

	productDescription := e.Group("/products/descriptions")
	productDescription.GET("", productDescriptionHandler.GetAllProductDescriptions())
	productDescription.GET("/:id", productDescriptionHandler.GetProductDescriptionById())
	productDescription.POST("", productDescriptionHandler.CreateProductDescription())
	productDescription.PUT("/:id", productDescriptionHandler.UpdateProductDescription())

	user := e.Group("/users")
	user.GET("", userHandler.GetAllUsers())
	user.GET("/:id", userHandler.GetUserById())
	user.POST("", userHandler.CreateUser())
	user.PUT("/:id", userHandler.UpdateUser())
	user.DELETE("/:id", userHandler.DeleteUser())

	userDetail := e.Group("/users/details")
	userDetail.GET("", userDetailHandler.GetAllUserDetails())
	userDetail.GET("/:id", userDetailHandler.GetUserDetailById())
	userDetail.POST("", userDetailHandler.CreateUserDetail())
	userDetail.PUT("/:id", userDetailHandler.UpdateUserDetail())
	userDetail.DELETE("/:id", userDetailHandler.DeleteUserDetail())

	role := e.Group("/roles")
	role.GET("", roleHandler.GetAllRoles())
	role.GET("/:id", roleHandler.GetRoleById())
	role.POST("", roleHandler.CreateRole())
	role.PUT("/:id", roleHandler.UpdateRole())
	role.DELETE("/:id", roleHandler.DeleteRole())

	return e
}
