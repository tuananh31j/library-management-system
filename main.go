package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/tuananh31j/library-management-system/config"
	"github.com/tuananh31j/library-management-system/database"
	"github.com/tuananh31j/library-management-system/middleware"
	"github.com/tuananh31j/library-management-system/router"
	"github.com/tuananh31j/library-management-system/utils"
	"gorm.io/gorm"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	app := setupFiber()
	db := database.Connect()
	defer closeDatabase(db)
	setupRouter(app, db)

	serverErrors := make(chan error, 1)
	address := fmt.Sprintf("%s:%d", config.AppHost, config.AppPort)
	go startServer(app, address, serverErrors)
	handleShutdown(ctx, app, serverErrors)
}
func closeDatabase(db *gorm.DB) {
	sqlDB, errDB := db.DB()
	if errDB != nil {
		utils.Log.Errorf("Error getting database instance: %v", errDB)
		return
	}

	if err := sqlDB.Close(); err != nil {
		utils.Log.Errorf("Error closing database connection: %v", err)
	} else {
		utils.Log.Info("Database connection closed successfully")
	}
}

func setupRouter(app *fiber.App, db *gorm.DB) {

	router.InitRouter(app, db)
}

func setupFiber() *fiber.App {
	app := fiber.New(config.FiberConfig())
	app.Use(middleware.LoggerConfig())
	app.Use(middleware.RecoverConfig())
	return app
}

func startServer(app *fiber.App, address string, errs chan<- error) {
	if err := app.Listen(address); err != nil {
		errs <- fmt.Errorf("error starting server: %w", err)
	}
}

func handleShutdown(ctx context.Context, app *fiber.App, serverErrors <-chan error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		utils.Log.Fatalf("Server error: %v", err)
	case <-quit:
		utils.Log.Info("Shutting down server...")
		if err := app.Shutdown(); err != nil {
			utils.Log.Fatalf("Error during server shutdown: %v", err)
		}
	case <-ctx.Done():
		utils.Log.Info("Server exiting due to context cancellation")
	}

	utils.Log.Info("Server exited")
}
