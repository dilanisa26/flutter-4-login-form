package main

import (
	"log"
	"os"
	"task-golang-db/handlers"
	"task-golang-db/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Initialize the database
	db := NewDatabase()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get DB from GORM:", err)
	}
	defer sqlDB.Close()

	// Get signing key from environment
	signingKey := os.Getenv("SIGNING_KEY")
	if signingKey == "" {
		log.Fatal("SIGNING_KEY environment variable is not set")
	}

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:61814"}, // Allowed origin for your Vue app
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Allow "Authorization" header
		AllowCredentials: true,
	}))

	// Homepage route
	router.GET("/homepage", handlers.Homepage(db))

	// Grouping routes under /auth
	authHandler := handlers.NewAuth(db, []byte(signingKey))
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/upsert", authHandler.Upsert)
		// Use the middleware here to secure the password change route
		authRoutes.POST("/changepassword", middleware.AuthMiddleware(signingKey), authHandler.ChangePassword) // New Change Password route
	}

	// Grouping routes under /account
	accountHandler := handlers.NewAccount(db)
	accountRoutes := router.Group("/account")
	{
		accountRoutes.POST("/create", accountHandler.Create)
		accountRoutes.GET("/read/:id", accountHandler.Read)
		accountRoutes.PATCH("/update/:id", accountHandler.Update)
		accountRoutes.DELETE("/delete/:id", accountHandler.Delete)
		accountRoutes.GET("/list", accountHandler.List)
		accountRoutes.POST("/topup", accountHandler.TopUp)
		accountRoutes.GET("/my", middleware.AuthMiddleware(signingKey), accountHandler.My)
		accountRoutes.GET("/balance", middleware.AuthMiddleware(signingKey), accountHandler.Balance)
		accountRoutes.POST("/transfer", middleware.AuthMiddleware(signingKey), accountHandler.Transfer)
		accountRoutes.GET("/mutation", middleware.AuthMiddleware(signingKey), accountHandler.Mutation)
	}

	// Grouping routes under /transaction-category
	transCatHandler := handlers.NewTransCat(db)
	transCatRoutes := router.Group("/transaction-category")
	{
		transCatRoutes.POST("/create", transCatHandler.Create)
		transCatRoutes.GET("/read/:id", transCatHandler.Read)
		transCatRoutes.PATCH("/update/:id", transCatHandler.Update)
		transCatRoutes.DELETE("/delete/:id", transCatHandler.Delete)
		transCatRoutes.GET("/list", transCatHandler.List)
		transCatRoutes.GET("/my", middleware.AuthMiddleware(signingKey), transCatHandler.My)
	}

	// Grouping routes under /transaction
	transactionHandler := handlers.NewTrans(db)
	transactionRoutes := router.Group("/transaction")
	{
		transactionRoutes.POST("/new", transactionHandler.NewTransaction)
		transactionRoutes.GET("/list", transactionHandler.TransactionList)
	}

	// Start the server
	router.Run(":8080")
}

// NewDatabase initializes the database connection
func NewDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE")
	if dsn == "" {
		log.Fatal("DATABASE environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get DB object:", err)
	}

	var currentDB string
	err = sqlDB.QueryRow("SELECT current_database()").Scan(&currentDB)
	if err != nil {
		log.Fatal("Failed to query current database:", err)
	}

	log.Printf("Connected to Database: %s\n", currentDB)

	return db
}
