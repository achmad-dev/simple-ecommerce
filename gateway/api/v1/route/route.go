package route

import (
	"database/sql"
	"os"

	customMiddleware "github.com/achmad-dev/simple-ecommerce/gateway/api/v1/middleware"
	handlers "github.com/achmad-dev/simple-ecommerce/gateway/api/v1/route/handler"
	"github.com/achmad-dev/simple-ecommerce/gateway/internal/utils"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/achmad-dev/simple-ecommerce/gateway/service"
	initLog "github.com/achmad-dev/simple-ecommerce/pkg/log"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func ServeRoute() {
	// Initialize logrus logger
	log := initLog.InitLog()

	// Set up database connection
	connStr := "user=postgress password=postgress dbname=simple_ecommerce sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Set up repositories
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	productRepo := repository.NewProductRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Set up services
	cartService := service.NewCartService(cartRepo, userRepo, log)
	orderService := service.NewOrderService(orderRepo, cartRepo, productRepo, log)
	paymentService := service.NewPaymentService(paymentRepo, log)
	productService := service.NewProductService(productRepo, log)
	userService := service.NewUserService(userRepo, log)

	// Set up Echo framework
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	jwtHelper := utils.NewJwtHelper("whatever")

	// Set up handlers
	authHandler := handlers.NewAuthHandler(userService, jwtHelper)
	paymentHandler := handlers.NewPaymentHandler(paymentService, userService)
	productHandler := handlers.NewProductHandler(productService)
	cartHandler := handlers.NewCartHandler(cartService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Set up routes
	e.POST("/register", authHandler.Register)
	e.GET("/login", authHandler.Login)
	protected := e.Group("/simple-ecommerce", customMiddleware.AuthMiddleware(jwtHelper))
	protected.POST("/payments", paymentHandler.CreatePaymentMethod)

	protected.GET("/home", productHandler.FetchAllProducts)
	protected.GET("/home-paginate", productHandler.FetchProductsPaginated)
	protected.POST("/cart/:productID", cartHandler.AddProductToCart)
	protected.DELETE("/cart/:productID", cartHandler.RemoveProductFromCart)
	protected.POST("/orders/:orderID/pay", orderHandler.PayOrder)
	protected.GET("/orders/:userID", orderHandler.GetOrdersByUserID)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
